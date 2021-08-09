package delegation

import (
	"context"
	"errors"
	"sort"
	"time"

	"code.vegaprotocol.io/vega/events"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/netparams"
	"code.vegaprotocol.io/vega/types"
	"code.vegaprotocol.io/vega/types/num"
)

var (
	// ErrPartyHasNoStakingAccount is returned when the staking account for the party cannot be found
	ErrPartyHasNoStakingAccount = errors.New("cannot find staking account for the party")
	// ErrInvalidNodeID is returned when the node id passed for delegation/undelegation is not a validator node identifier
	ErrInvalidNodeID = errors.New("invalid node ID")
	// ErrInsufficientBalanceForDelegation is returned when the balance in the staking account is insufficient to cover all committed and pending delegations
	ErrInsufficientBalanceForDelegation = errors.New("insufficient balance for delegation")
	// ErrIncorrectTokenAmountForUndelegation is returned when the amount to undelegation doesn't match the delegation balance (pending + committed) for the party and validator
	ErrIncorrectTokenAmountForUndelegation = errors.New("incorrect token amount for undelegation")
	// ErrAmountLTMinAmountForDelegation is returned when the amount to delegate to a node is lower than the minimum allowed amount from network params
	ErrAmountLTMinAmountForDelegation = errors.New("delegation amount is lower than the minimum amount for delegation for a validator")
)

// ValidatorTopology represents the topology of validators and can check if a given node is a validator
type ValidatorTopology interface {
	IsValidatorNode(nodeID string) bool
}

// Broker send events
// we no longer need to generate this mock here, we can use the broker/mocks package instead
type Broker interface {
	Send(event events.Event)
	SendBatch(events []events.Event)
}

// StakingAccounts provides access to the staking balance of a given party now and within a duration of an epoch
type StakingAccounts interface {
	GetBalanceNow(party string) *num.Uint
	GetBalanceForEpoch(party string, from, to time.Time) *num.Uint
}

// validator delegation state - updated at the end of each epoch
type validatorDelegation struct {
	nodeID         string               // node id
	partyToAmount  map[string]*num.Uint // party -> delegated amount
	totalDelegated *num.Uint            // the total amount delegates by parties
}

// party delegation state - how much is delegated by the party to each validator and in total
type partyDelegation struct {
	party          string               // party ID
	nodeToAmount   map[string]*num.Uint // nodeID -> delegated amount
	totalDelegated *num.Uint            // total amount delegated by party
}

// party delegation state
type pendingPartyDelegation struct {
	party                  string
	nodeToDelegateAmount   map[string]*num.Uint
	nodeToUndelegateAmount map[string]*num.Uint
	totalDelegation        *num.Uint
	totalUndelegation      *num.Uint
}

// Engine is handling the delegations balances from parties to validators
// The delegation engine is designed in the following way with the following assumptions:
// 1. during epoch it is called with delegation requests that are added to a pending data structure and only applied at the end of the epoch
// 2. At the end of the epoch the engine is called and does the following:
// 2.1 updates the delegated balances to match the epoch's staking account balance for each party such that if a party withdrew from their
//     staking account during the epoch it will not count for them for rewarding
// 2.2 capture the state after 2.1 to be returned to the rewarding engine
// 2.3 process all pending delegations
type Engine struct {
	log                  *logging.Logger
	config               Config
	broker               Broker
	topology             ValidatorTopology                  // an interface to the topoology to interact with validator nodes if needed
	stakingAccounts      StakingAccounts                    // an interface to the staking account for getting party balances
	nodeDelegationState  map[string]*validatorDelegation    // validator to active delegations
	partyDelegationState map[string]*partyDelegation        // party to active delegations
	pendingState         map[string]*pendingPartyDelegation // pending delegations/undelegations by party
	netp                 NetParams                          // network parameter interface for reading needed network parameters
}

//NetParams provides access to network parameters
//use mock from mocks/netparams_mock.go
type NetParams interface {
	Get(string) (string, error)
}

//New instantiate a new delegation engine
func New(log *logging.Logger, config Config, broker Broker, topology ValidatorTopology, stakingAccounts StakingAccounts, netp NetParams) *Engine {
	e := &Engine{
		config:               config,
		log:                  log.Named(namedLogger),
		broker:               broker,
		topology:             topology,
		stakingAccounts:      stakingAccounts,
		nodeDelegationState:  map[string]*validatorDelegation{},
		partyDelegationState: map[string]*partyDelegation{},
		pendingState:         map[string]*pendingPartyDelegation{},
		netp:                 netp,
	}
	return e
}

//OnEpochEnd updates the delegation engine state at the end of an epoch and returns the last epoch's validation-delegation data for rewarding
// step 1: process delegation data for the epoch - undelegate if the balance of the staking account doesn't cover all delegations
// step 2: capture validator delegation data to be returned
// step 3: apply pending undelegations
// step 4: apply pending delegations
func (e *Engine) OnEpochEnd(ctx context.Context, start, end time.Time) []*types.ValidatorData {
	if e.log.IsDebug() {
		e.log.Debug("on epoch end:", logging.Time("start", start), logging.Time("end", end))
	}
	e.preprocessEpochForRewarding(ctx, start, end)
	stateForRewards := e.getValidatorData()
	e.processPending(ctx)
	return stateForRewards
}

//Delegate increases the pending delegation balance and potentially decreases the pending undelegation balance for a given validator node
func (e *Engine) Delegate(ctx context.Context, party string, nodeID string, amount *num.Uint) error {
	amt := amount.Clone()

	// check if the node is a validator node
	if !e.topology.IsValidatorNode(nodeID) {
		return ErrInvalidNodeID
	}

	// check if the delegator has a staking account
	partyBalance := e.stakingAccounts.GetBalanceNow(party)
	if partyBalance == nil {
		return ErrPartyHasNoStakingAccount
	}

	// read the delegation min amount network param
	validatorsDelegationMinAmount, err := e.netp.Get(netparams.DelegationMinAmount)
	if err != nil {
		return err
	}
	minAmount, ok := num.UintFromString(validatorsDelegationMinAmount, 10)
	if ok {
		e.log.Panic("unable to read", logging.String(netparams.DelegationMinAmount, validatorsDelegationMinAmount))
	}

	if amt.LT(minAmount) {
		return ErrAmountLTMinAmountForDelegation
	}

	// check if the delegator has sufficient balance in their staking account including all pending and committed delegations and undelegations
	// this is basically just fail fast - the delegation may still fail
	currentPendingPartyDelegation, ok := e.pendingState[party]
	if !ok {
		e.pendingState[party] = &pendingPartyDelegation{
			party:                  party,
			totalDelegation:        num.Zero(),
			totalUndelegation:      num.Zero(),
			nodeToUndelegateAmount: map[string]*num.Uint{},
			nodeToDelegateAmount:   map[string]*num.Uint{},
		}
		currentPendingPartyDelegation = e.pendingState[party]
	}
	partyDelegation, ok := e.partyDelegationState[party]
	partyDelegationBalance := num.Zero()
	if ok {
		partyDelegationBalance = partyDelegation.totalDelegated
	}

	// if the party withdrew from their account and now don't have sufficient cover for their current delegation, prevent them from further delgations
	// no need to immediately undelegate because this will be handled at epoch end
	if partyBalance.LTE(partyDelegationBalance) {
		return ErrInsufficientBalanceForDelegation
	}

	// subrtact the committed delegation balance and apply pending if any

	balanceAvailableForDelegation := num.Zero().Sub(partyBalance, partyDelegationBalance)
	partyPendingDelegation := currentPendingPartyDelegation.totalDelegation
	partyPendingUndelegation := currentPendingPartyDelegation.totalUndelegation

	// add pending undelegations to available balance
	if !partyPendingUndelegation.IsZero() {
		balanceAvailableForDelegation.AddSum(partyPendingUndelegation)
	}
	// subtract pending delegations from available balance
	if !partyPendingDelegation.IsZero() {
		// if there's somehow more pending than available for delegation due to withdrawls return error
		if partyPendingDelegation.GT(balanceAvailableForDelegation) {
			return ErrInsufficientBalanceForDelegation
		}
		balanceAvailableForDelegation = num.Zero().Sub(balanceAvailableForDelegation, partyPendingDelegation)
	}

	// if the balance with committed and pending delegations/undelegations is insufficient to satisfy the delegation return error
	if balanceAvailableForDelegation.LT(amt) {
		return ErrInsufficientBalanceForDelegation
	}

	// all good add to pending delegation
	remainingBalanceForDelegate := amt

	partyPendingUndelegationForNode, udok := currentPendingPartyDelegation.nodeToUndelegateAmount[nodeID]
	partyPendingDelegationForNode, dok := currentPendingPartyDelegation.nodeToDelegateAmount[nodeID]

	if udok { // we have undelegates which we can counter
		if remainingBalanceForDelegate.GTE(partyPendingUndelegationForNode) {
			// the delegation amount is greater than or equal to the undelegated amount, we can clear the whole undelegation and leave the remaining delegation
			remainingBalanceForDelegate = num.Zero().Sub(remainingBalanceForDelegate, partyPendingUndelegationForNode)
			delete(currentPendingPartyDelegation.nodeToUndelegateAmount, nodeID)
			currentPendingPartyDelegation.totalUndelegation = num.Zero().Sub(currentPendingPartyDelegation.totalUndelegation, partyPendingUndelegationForNode)
			currentPendingPartyDelegation.totalDelegation = num.Zero().Add(currentPendingPartyDelegation.totalDelegation, remainingBalanceForDelegate)

			if !remainingBalanceForDelegate.IsZero() {
				currentPendingPartyDelegation.nodeToDelegateAmount[nodeID] = remainingBalanceForDelegate
			} else {
				delete(currentPendingPartyDelegation.nodeToDelegateAmount, nodeID)
				if currentPendingPartyDelegation.totalUndelegation.EQ(num.Zero()) && currentPendingPartyDelegation.totalDelegation.EQ(num.Zero()) {
					delete(e.pendingState, party)
				}
			}
		} else {
			// the delegation amount is lower than the pending undelegate amount - we can just adjust the undelegate amount
			updatedUndelegateAmout := num.Zero().Sub(partyPendingUndelegationForNode, remainingBalanceForDelegate)
			currentPendingPartyDelegation.nodeToUndelegateAmount[nodeID] = updatedUndelegateAmout
			currentPendingPartyDelegation.totalUndelegation = num.Zero().Sub(currentPendingPartyDelegation.totalUndelegation, remainingBalanceForDelegate)

		}
	} else {
		// there are no pending undelegations we can just update the pending delegation
		if !dok {
			partyPendingDelegationForNode = num.Zero()
		}
		currentPendingPartyDelegation.nodeToDelegateAmount[nodeID] = num.Zero().Add(partyPendingDelegationForNode, remainingBalanceForDelegate)
		currentPendingPartyDelegation.totalDelegation = num.Zero().Add(currentPendingPartyDelegation.totalDelegation, remainingBalanceForDelegate)
	}

	e.sendPendingBalanceEvent(ctx, party, nodeID)

	return nil
}

//UndelegateAtEndOfEpoch increases the pending undelegation balance and potentially decreases the pending delegation balance for a given validator node and party
func (e *Engine) UndelegateAtEndOfEpoch(ctx context.Context, party string, nodeID string, amount *num.Uint) error {
	amt := amount.Clone()

	if amt.IsZero() {
		// calculate how much we have available for undelegation including pending and committed
		availableForUndelegationInPending := num.Zero()
		if pendingState, ok := e.pendingState[party]; ok {
			if nodeDelegation, ok := pendingState.nodeToDelegateAmount[nodeID]; ok {
				availableForUndelegationInPending = num.Sum(availableForUndelegationInPending, nodeDelegation)
			}
		}
		availableForUndelegationInActive := num.Zero()
		if partyDelegation, ok := e.partyDelegationState[party]; ok {
			if nodeDelegation, ok := partyDelegation.nodeToAmount[nodeID]; ok {
				availableForUndelegationInActive = num.Sum(availableForUndelegationInActive, nodeDelegation)
			}
		}
		amt = amt.AddSum(availableForUndelegationInPending, availableForUndelegationInActive)
	}

	// check if the node is a validator node
	if e.topology == nil || !e.topology.IsValidatorNode(nodeID) {
		return ErrInvalidNodeID
	}

	// get the delegated balance for the given node
	validatorState, ok := e.nodeDelegationState[nodeID]
	partyDelegatedToNodeAmount := num.Zero()
	if ok {
		partyDelegatedToNodeAmount, ok = validatorState.partyToAmount[party]
		if !ok {
			partyDelegatedToNodeAmount = num.Zero()
		}
	}

	pendingDelegateToNodeAmount := num.Zero()
	pendingUndelegateToNodeAmount := num.Zero()

	// check if there is anything pending
	currentPendingPartyDelegation, ok := e.pendingState[party]
	if ok {
		pendingDelegateToNodeAmount, ok = currentPendingPartyDelegation.nodeToDelegateAmount[nodeID]
		if !ok {
			pendingDelegateToNodeAmount = num.Zero()
		}
		pendingUndelegateToNodeAmount, ok = currentPendingPartyDelegation.nodeToUndelegateAmount[nodeID]
		if !ok {
			pendingUndelegateToNodeAmount = num.Zero()
		}
	} else {
		// if there isn't yet a pending state, construct it here
		currentPendingPartyDelegation = &pendingPartyDelegation{
			party:                  party,
			totalDelegation:        num.Zero(),
			totalUndelegation:      num.Zero(),
			nodeToUndelegateAmount: map[string]*num.Uint{},
			nodeToDelegateAmount:   map[string]*num.Uint{},
		}
	}

	totalDelegationBalance := num.Zero().Add(partyDelegatedToNodeAmount, pendingDelegateToNodeAmount)
	totalDelegationBalance = num.Zero().Sub(totalDelegationBalance, pendingUndelegateToNodeAmount)

	// if the amount is greater than the available balance to undelegate return error
	if amt.GT(totalDelegationBalance) {
		return ErrIncorrectTokenAmountForUndelegation
	}

	remainingBalanceForUndelegate := amt

	if !pendingDelegateToNodeAmount.IsZero() { // we have delegates which we can counter
		if remainingBalanceForUndelegate.GTE(pendingDelegateToNodeAmount) {
			// the undelegation amount is greater than or equal to the delegated amount, we can clear the whole delegation and leave the remaining undelegation
			remainingBalanceForUndelegate = num.Zero().Sub(remainingBalanceForUndelegate, pendingDelegateToNodeAmount)
			currentPendingPartyDelegation.totalDelegation = num.Zero().Sub(currentPendingPartyDelegation.totalDelegation, pendingDelegateToNodeAmount)
			currentPendingPartyDelegation.totalUndelegation = num.Zero().Add(currentPendingPartyDelegation.totalUndelegation, remainingBalanceForUndelegate)

			delete(currentPendingPartyDelegation.nodeToDelegateAmount, nodeID)
			if !remainingBalanceForUndelegate.IsZero() {
				currentPendingPartyDelegation.nodeToUndelegateAmount[nodeID] = remainingBalanceForUndelegate
			} else {
				delete(currentPendingPartyDelegation.nodeToUndelegateAmount, nodeID)
				if currentPendingPartyDelegation.totalUndelegation.EQ(num.Zero()) && currentPendingPartyDelegation.totalDelegation.EQ(num.Zero()) {
					delete(e.pendingState, party)
				}
			}
		} else {
			// the undelegation amount is lower than the pending delegate amount - we can just adjust the delegate amount
			updatedDelegateAmount := num.Zero().Sub(pendingDelegateToNodeAmount, remainingBalanceForUndelegate)
			currentPendingPartyDelegation.totalDelegation = num.Zero().Sub(currentPendingPartyDelegation.totalDelegation, remainingBalanceForUndelegate)
			currentPendingPartyDelegation.nodeToDelegateAmount[nodeID] = updatedDelegateAmount
		}
	} else {
		// there are no pending delegations we can just update the pending undelegation
		currentPendingPartyDelegation.nodeToUndelegateAmount[nodeID] = num.Zero().Add(pendingUndelegateToNodeAmount, remainingBalanceForUndelegate)
		currentPendingPartyDelegation.totalUndelegation = num.Zero().Add(currentPendingPartyDelegation.totalUndelegation, remainingBalanceForUndelegate)
	}

	_, ok = e.pendingState[party]
	// if there was no previous undelegation and we ended up undelegating, add to state
	if !ok && !currentPendingPartyDelegation.totalUndelegation.IsZero() {
		e.pendingState[party] = currentPendingPartyDelegation
	}

	e.sendPendingBalanceEvent(ctx, party, nodeID)
	return nil
}

//UndelegateNow changes the balance of delegation immediately without waiting for the end of the epoch
// if possible it removed balance from pending delegated, if not enough it removes balance from the current epoch delegated amount
func (e *Engine) UndelegateNow(ctx context.Context, party string, nodeID string, amount *num.Uint) error {
	// first check available balance for undelegation and error if the requested amount is greater than
	availableForUndelegationInPending := num.Zero()
	if pendingState, ok := e.pendingState[party]; ok {
		if nodeDelegation, ok := pendingState.nodeToDelegateAmount[nodeID]; ok {
			availableForUndelegationInPending = num.Zero().Add(availableForUndelegationInPending, nodeDelegation)
		}
	}
	availableForUndelegationInActive := num.Zero()
	if partyDelegation, ok := e.partyDelegationState[party]; ok {
		if nodeDelegation, ok := partyDelegation.nodeToAmount[nodeID]; ok {
			availableForUndelegationInActive = num.Zero().Add(availableForUndelegationInActive, nodeDelegation)
		}
	}

	totalAvailableForUndelegation := num.Zero().Add(availableForUndelegationInPending, availableForUndelegationInActive)

	// if the party passes 0 they want to undelegate all
	var amt = amount.Clone()
	if amt.IsZero() {
		amt = totalAvailableForUndelegation.Clone()
	}

	if amt.GT(totalAvailableForUndelegation) {
		return ErrIncorrectTokenAmountForUndelegation
	}

	// strart with undelegating from pending, if not enough go to active
	if !availableForUndelegationInPending.IsZero() {
		pendingState := e.pendingState[party]
		if amt.LTE(availableForUndelegationInPending) {
			pendingState.nodeToDelegateAmount[nodeID] = num.Zero().Sub(availableForUndelegationInPending, amt)
			pendingState.totalDelegation = num.Zero().Sub(pendingState.totalDelegation, amt)
			if pendingState.nodeToDelegateAmount[nodeID].EQ(num.Zero()) {
				delete(pendingState.nodeToDelegateAmount, nodeID)
			}
			amt = num.Zero()
		} else {
			// we don't have enough delegation to cover for the undelegate request
			pendingState.totalDelegation = num.Zero().Sub(pendingState.totalDelegation, availableForUndelegationInPending)
			delete(pendingState.nodeToDelegateAmount, nodeID)
			amt = num.Zero().Sub(amt, availableForUndelegationInPending)
		}

		if pendingState.totalDelegation.EQ(num.Zero()) && pendingState.totalUndelegation.EQ(num.Zero()) {
			delete(e.pendingState, party)
		}

		e.sendPendingBalanceEvent(ctx, party, nodeID)
	}
	// if there's still some balance to undelegate we go to the delegated state
	if !amt.IsZero() {
		partyDelegation := e.partyDelegationState[party]
		partyDelegation.totalDelegated = num.Zero().Sub(partyDelegation.totalDelegated, amt)
		partyDelegation.nodeToAmount[nodeID] = num.Zero().Sub(partyDelegation.nodeToAmount[nodeID], amt)
		if partyDelegation.nodeToAmount[nodeID].EQ(num.Zero()) {
			delete(partyDelegation.nodeToAmount, nodeID)
		}
		if partyDelegation.totalDelegated.EQ(num.Zero()) {
			delete(e.partyDelegationState, party)
		}
		nodeDelegation, ok := e.nodeDelegationState[nodeID]
		if !ok {
			e.log.Panic("party and node delegation state disagree")
		}
		nodeDelegation.totalDelegated = num.Zero().Sub(nodeDelegation.totalDelegated, amt)
		nodeDelegation.partyToAmount[party] = num.Zero().Sub(nodeDelegation.partyToAmount[party], amt)
		if nodeDelegation.partyToAmount[party].EQ(num.Zero()) {
			delete(nodeDelegation.partyToAmount, party)
		}
		if nodeDelegation.totalDelegated.EQ(num.Zero()) {
			delete(e.nodeDelegationState, nodeID)
		}
		e.sendDelegatedBalanceEvent(ctx, party, nodeID)
	}
	return nil
}

func (e *Engine) sendPendingBalanceEvent(ctx context.Context, party, nodeID string) {
	pendingState, ok := e.pendingState[party]

	if ok {
		pendingDelegated, dok := pendingState.nodeToDelegateAmount[nodeID]
		if !dok {
			pendingDelegated = num.Zero()
		}
		pendingUndelegated, udok := pendingState.nodeToUndelegateAmount[nodeID]
		if !udok {
			pendingUndelegated = num.Zero()
		}
		e.broker.Send(events.NewPendingDelegationBalance(ctx, party, nodeID, pendingDelegated, pendingUndelegated))
	} else {
		e.broker.Send(events.NewPendingDelegationBalance(ctx, party, nodeID, num.Zero(), num.Zero()))
	}
}

func (e *Engine) sendDelegatedBalanceEvent(ctx context.Context, party, nodeID string) {
	delegated, ok := e.partyDelegationState[party]

	if ok {
		amt, ok := delegated.nodeToAmount[nodeID]
		if !ok {
			amt = num.Zero()
		}
		e.broker.Send(events.NewDelegationBalance(ctx, party, nodeID, amt))
	} else {
		e.broker.Send(events.NewDelegationBalance(ctx, party, nodeID, num.Zero()))
	}
}

func (e *Engine) decreaseDelegationAmountBy(party, nodeID string, amt *num.Uint) {
	partyDelegation := e.partyDelegationState[party]
	nodeDelegation := e.nodeDelegationState[nodeID]

	// update the balance for the validator for the party
	partyDelegation.nodeToAmount[nodeID] = num.Zero().Sub(partyDelegation.nodeToAmount[nodeID], amt)
	partyDelegation.totalDelegated = num.Zero().Sub(partyDelegation.totalDelegated, amt)

	// if there's no more delegations, remove the entry for the nodeID
	if partyDelegation.nodeToAmount[nodeID].EQ(num.Zero()) {
		delete(partyDelegation.nodeToAmount, nodeID)
	}
	if partyDelegation.totalDelegated.EQ(num.Zero()) {
		delete(e.partyDelegationState, party)
	}

	// update the balance for the party for the validator
	nodeDelegation.partyToAmount[party] = num.Zero().Sub(nodeDelegation.partyToAmount[party], amt)
	nodeDelegation.totalDelegated = num.Zero().Sub(nodeDelegation.totalDelegated, amt)

	// if there's no more delegations, remove the entry for the nodeID
	if nodeDelegation.partyToAmount[party].EQ(num.Zero()) {
		delete(nodeDelegation.partyToAmount, party)
	}
	if nodeDelegation.totalDelegated.EQ(num.Zero()) {
		delete(e.nodeDelegationState, nodeID)
	}

}

// sort node IDs for deterministic processing
func (e *Engine) sortNodes(nodes map[string]*num.Uint) []string {
	nodeIDs := make([]string, 0, len(nodes))
	for nodeID := range nodes {
		nodeIDs = append(nodeIDs, nodeID)
	}

	// sort the parties for deterministic handling
	sort.Strings(nodeIDs)
	return nodeIDs
}

// preprocessEpoch is called at the end of an epoch and updates the state to be returned for rewarding calculation
// check balance for the epoch duration and undelegate if delegations don't have sufficient cover
// the state of the engine by the end of this method reflects the state to be used for reward engine
func (e *Engine) preprocessEpochForRewarding(ctx context.Context, epochStart, epochEnd time.Time) {
	parties := make([]string, 0, len(e.partyDelegationState))
	for party := range e.partyDelegationState {
		parties = append(parties, party)
	}

	// sort the parties for deterministic handling
	sort.Strings(parties)

	// for all parties with delegations in the ended epoch
	for _, party := range parties {
		partyDelegation := e.partyDelegationState[party]

		// get the party stake balance for the epoch
		stakeBalance := e.stakingAccounts.GetBalanceForEpoch(party, epochStart, epochEnd)

		// if the stake covers the total delegated balance nothing to do further for the party
		if stakeBalance.GTE(partyDelegation.totalDelegated) {
			continue
		}

		// if the stake account balance for the epoch is less than the delegated balance - we need to undelegate the difference
		// this will be done evenly as much as possible between all validators with delegation from the party
		remainingBalanceToUndelegate := num.Zero().Sub(partyDelegation.totalDelegated, stakeBalance)

		totalTaken := num.Zero()

		nodeIDs := e.sortNodes(partyDelegation.nodeToAmount)

		// undelegate proportionally across delegated validator nodes
		totalDeletation := partyDelegation.totalDelegated.Clone()
		for _, nodeID := range nodeIDs {
			balance := partyDelegation.nodeToAmount[nodeID]
			balanceToTake := num.Zero().Mul(balance, remainingBalanceToUndelegate)
			balanceToTake = num.Zero().Div(balanceToTake, totalDeletation)

			if balanceToTake.EQ(num.Zero()) {
				continue
			}

			e.decreaseDelegationAmountBy(party, nodeID, balanceToTake)
			totalTaken = num.Zero().Add(totalTaken, balanceToTake)
		}

		// if there was a remainder, the maximum that we need to take more from each node is 1,
		if totalTaken.LT(remainingBalanceToUndelegate) {
			for _, nodeID := range nodeIDs {
				balance, ok := partyDelegation.nodeToAmount[nodeID]
				if !ok {
					continue
				}
				if totalTaken.EQ(remainingBalanceToUndelegate) {
					break
				}
				if !balance.IsZero() {
					e.decreaseDelegationAmountBy(party, nodeID, num.NewUint(1))
					totalTaken = num.Zero().Add(totalTaken, num.NewUint(1))
				}
			}
		}

		if len(partyDelegation.nodeToAmount) == 0 {
			delete(e.partyDelegationState, party)
		}

		for _, nodeID := range nodeIDs {
			e.sendDelegatedBalanceEvent(ctx, party, nodeID)
		}
	}
}

// process pending delegations and undelegations at the end of the epoch and clear the delegation/undelegation maps at the end
func (e *Engine) processPending(ctx context.Context) {
	parties := make([]string, 0, len(e.pendingState))
	partyNodes := map[string][]string{}
	for party, state := range e.pendingState {
		parties = append(parties, party)
		nodes := map[string]bool{}
		for node := range state.nodeToDelegateAmount {
			nodes[node] = true
		}
		for node := range state.nodeToUndelegateAmount {
			nodes[node] = true
		}
		var nodesSlice []string
		for node := range nodes {
			nodesSlice = append(nodesSlice, node)
		}
		sort.Strings(nodesSlice)
		partyNodes[party] = nodesSlice
	}

	// sort the parties for deterministic handling
	sort.Strings(parties)

	// read the delegation min amount network param
	maxStakePerValidatorStr, err := e.netp.Get(netparams.DelegationMaxStakePerValidator)
	if err != nil {
		e.log.Panic("Cannot find validators.delegation.maxStakePerValidator")
	}
	maxStakePerValidator, ok := num.UintFromString(maxStakePerValidatorStr, 10)
	if ok {
		e.log.Panic("unable to read", logging.String(netparams.DelegationMaxStakePerValidator, maxStakePerValidatorStr))
	}

	e.processPendingUndelegations(parties)
	e.processPendingDelegations(parties, maxStakePerValidator)

	e.pendingState = map[string]*pendingPartyDelegation{}
	for _, party := range parties {
		nodes := partyNodes[party]
		for _, node := range nodes {
			e.sendDelegatedBalanceEvent(ctx, party, node)
			e.sendPendingBalanceEvent(ctx, party, node)
		}
	}
}

// process pending undelegations for all parties
func (e *Engine) processPendingUndelegations(parties []string) {
	for _, party := range parties {
		pending, ok := e.pendingState[party]
		if !ok {
			continue
		}

		// get committed delegations for the party
		committedDelegations, ok := e.partyDelegationState[party]
		if !ok {
			committedDelegations = &partyDelegation{
				party:          party,
				totalDelegated: num.Zero(),
				nodeToAmount:   map[string]*num.Uint{},
			}
		}

		// apply undelegations deterministically
		nodeIDs := e.sortNodes(pending.nodeToUndelegateAmount)

		for _, nodeID := range nodeIDs {
			amount, ok := pending.nodeToUndelegateAmount[nodeID]
			if !ok {
				continue
			}
			committedForNode, delegationFoundForParty := committedDelegations.nodeToAmount[nodeID]
			if !delegationFoundForParty {
				// there is nothing to undelegate for this node, log and continue
				e.log.Debug("no committed delegation found for pending undelegation for", logging.String("party", party), logging.String("nodeID", nodeID))
				continue
			}

			validatorDelegation, ok := e.nodeDelegationState[nodeID]
			if !ok {
				// this should never happen
				e.log.Panic("trying to undelegate from an unknown node", logging.String("nodeID", nodeID))
			}

			validatorPartyDelegationAmount, ok := validatorDelegation.partyToAmount[party]
			if !ok == delegationFoundForParty {
				e.log.Panic("party and validator state disagree", logging.String("nodeID", nodeID), logging.String("party", party))
			}

			amountForUndelegate := amount
			if committedForNode.LT(amount) {
				amountForUndelegate = committedForNode
			}

			// undelegate
			// update validator mapping for the party
			validatorDelegation.partyToAmount[party] = num.Zero().Sub(validatorPartyDelegationAmount, amountForUndelegate)

			// if no more delegations for the party for the node, remove the mapping
			if validatorDelegation.partyToAmount[party].EQ(num.Zero()) {
				delete(validatorDelegation.partyToAmount, party)
			}
			validatorDelegation.totalDelegated = num.Zero().Sub(validatorDelegation.totalDelegated, amountForUndelegate)
			// if no more delegations for the node, clear it from the state
			if validatorDelegation.totalDelegated.EQ(num.Zero()) {
				delete(e.nodeDelegationState, nodeID)
			}

			// update undelegation for party
			committedDelegations.totalDelegated = num.Zero().Sub(committedDelegations.totalDelegated, amountForUndelegate)
			committedDelegations.nodeToAmount[nodeID] = num.Zero().Sub(committedForNode, amountForUndelegate)
			if committedDelegations.nodeToAmount[nodeID].EQ(num.Zero()) {
				delete(committedDelegations.nodeToAmount, nodeID)
			}

			if !committedDelegations.totalDelegated.IsZero() {
				e.partyDelegationState[party] = committedDelegations
			} else {
				_, ok := e.partyDelegationState[party]
				if ok {
					delete(e.partyDelegationState, party)
				}
			}
		}
	}
}

// process pending delegations for all parties
func (e *Engine) processPendingDelegations(parties []string, maxStakePerValidator *num.Uint) {
	// process undelegations for all parties first
	for _, party := range parties {
		pending, ok := e.pendingState[party]
		if !ok {
			continue
		}
		// get account balance
		partyBalance := e.stakingAccounts.GetBalanceNow(party)

		// get committed delegations for the party
		committedDelegations, ok := e.partyDelegationState[party]
		if !ok {
			committedDelegations = &partyDelegation{
				party:          party,
				totalDelegated: num.Zero(),
				nodeToAmount:   map[string]*num.Uint{},
			}
		}
		availableForDelegation := num.Zero().Sub(partyBalance, committedDelegations.totalDelegated)

		// apply delegation deterministically
		nodeIDs := e.sortNodes(pending.nodeToDelegateAmount)
		for _, nodeID := range nodeIDs {
			_, ok := pending.nodeToDelegateAmount[nodeID]
			if !ok {
				continue
			}

			// get the amount for delegation and adjust it if needed to the available balance for delegation in the validator
			amount := pending.nodeToDelegateAmount[nodeID].Clone()
			currentNodeDelegationBalance := num.Zero()
			currentNodeDelegation, ok := e.nodeDelegationState[nodeID]
			if ok {
				currentNodeDelegationBalance = currentNodeDelegation.totalDelegated
			}
			if !maxStakePerValidator.IsZero() {
				availableBalanceOnNode := num.Zero().Sub(maxStakePerValidator, currentNodeDelegationBalance)
				if amount.GT(availableBalanceOnNode) {
					amount = availableBalanceOnNode
				}
			}

			// check that the amount is not greater than the available for delegation
			if amount.GT(availableForDelegation) || amount.EQ(num.Zero()) {
				e.log.Debug("the amount requested for delegation is greater than available for delegation at end of epoch", logging.String("party", party), logging.String("nodeID", nodeID), logging.Uint64("amt", amount.Uint64()))
				continue
			}

			// update the validator delegation balance
			currentValidatorDelegation, ok := e.nodeDelegationState[nodeID]
			if !ok {
				currentValidatorDelegation = &validatorDelegation{
					nodeID:         nodeID,
					totalDelegated: num.Zero(),
					partyToAmount:  map[string]*num.Uint{},
				}
			}
			currentDelegationAmtForParty, ok := currentValidatorDelegation.partyToAmount[party]
			if !ok {
				currentDelegationAmtForParty = num.Zero()
			}
			currentValidatorDelegation.partyToAmount[party] = num.Zero().Add(currentDelegationAmtForParty, amount)
			currentValidatorDelegation.totalDelegated = num.Zero().Add(currentValidatorDelegation.totalDelegated, amount)
			e.nodeDelegationState[nodeID] = currentValidatorDelegation

			// update undelegation for party
			committedForNode, ok := committedDelegations.nodeToAmount[nodeID]
			if !ok {
				committedForNode = num.Zero()
			}
			committedDelegations.totalDelegated = num.Zero().Add(committedDelegations.totalDelegated, amount)
			committedDelegations.nodeToAmount[nodeID] = num.Zero().Add(committedForNode, amount)
			e.partyDelegationState[party] = committedDelegations
		}
	}
}

//returns the current state of the delegation per node
func (e *Engine) getValidatorData() []*types.ValidatorData {
	validators := make([]*types.ValidatorData, 0, len(e.nodeDelegationState))

	nodeIDs := make([]string, 0, len(e.nodeDelegationState))
	for nodeID := range e.nodeDelegationState {
		nodeIDs = append(nodeIDs, nodeID)
	}

	// sort the parties for deterministic handling
	sort.Strings(nodeIDs)

	for _, nodeID := range nodeIDs {
		validatorState := e.nodeDelegationState[nodeID]
		validator := &types.ValidatorData{
			NodeID:            nodeID,
			StakeByDelegators: validatorState.totalDelegated.Clone(),
			Delegators:        map[string]*num.Uint{},
		}
		for delegatingParties, amt := range validatorState.partyToAmount {
			validator.Delegators[delegatingParties] = amt.Clone()
		}
		validators = append(validators, validator)
	}

	return validators

}
