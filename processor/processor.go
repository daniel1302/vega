package processor

import (
	"encoding/hex"
	"fmt"
	"time"

	"code.vegaprotocol.io/vega/blockchain"
	"code.vegaprotocol.io/vega/logging"
	types "code.vegaprotocol.io/vega/proto"
	"code.vegaprotocol.io/vega/vegatime"

	"github.com/golang/protobuf/proto"
	"github.com/pkg/errors"
)

var (
	ErrUnknownCommand                               = errors.New("unknown command when validating payload")
	ErrInvalidSignature                             = errors.New("invalid signature")
	ErrOrderSubmissionPartyAndPubKeyDoesNotMatch    = errors.New("order submission party and pubkey does not match")
	ErrOrderCancellationPartyAndPubKeyDoesNotMatch  = errors.New("order cancellation party and pubkey does not match")
	ErrOrderAmendmentPartyAndPubKeyDoesNotMatch     = errors.New("order amendment party and pubkey does not match")
	ErrProposalSubmissionPartyAndPubKeyDoesNotMatch = errors.New("proposal submission party and pubkey does not match")
	ErrVoteSubmissionPartyAndPubKeyDoesNotMatch     = errors.New("vote submission party and pubkey does not match")
	ErrWithdrawPartyAndPublKeyDoesNotMatch          = errors.New("withdraw party and pubkey does not match")
	ErrCommandKindUnknown                           = errors.New("unknown command kind when validating payload")
	ErrUnknownNodeKey                               = errors.New("node pubkey unknown")
	ErrUnknownProposal                              = errors.New("proposal unknown")
	ErrNotAnAssetProposal                           = errors.New("proposal is not a new asset proposal")
)

//go:generate go run github.com/golang/mock/mockgen -destination mocks/time_service_mock.go -package mocks code.vegaprotocol.io/vega/processor TimeService
type TimeService interface {
	GetTimeNow() (time.Time, error)
	GetTimeLastBatch() (time.Time, error)
	NotifyOnTick(f func(time.Time))
}

//go:generate go run github.com/golang/mock/mockgen -destination mocks/execution_engine_mock.go -package mocks code.vegaprotocol.io/vega/processor ExecutionEngine
type ExecutionEngine interface {
	SubmitOrder(order *types.Order) (*types.OrderConfirmation, error)
	CancelOrder(order *types.OrderCancellation) (*types.OrderCancellationConfirmation, error)
	AmendOrder(order *types.OrderAmendment) (*types.OrderConfirmation, error)
	NotifyTraderAccount(notif *types.NotifyTraderAccount) error
	Withdraw(*types.Withdraw) error
	Generate() error
	SubmitProposal(proposal *types.Proposal) error
	VoteOnProposal(vote *types.Vote) error
}

//go:generate go run github.com/golang/mock/mockgen -destination mocks/stats_mock.go -package mocks code.vegaprotocol.io/vega/processor Stats
type Stats interface {
	IncTotalCreateOrder()
	AddCurrentTradesInBatch(i uint64)
	AddTotalTrades(i uint64) uint64
	IncTotalOrders()
	IncCurrentOrdersInBatch()
	IncTotalCancelOrder()
	IncTotalAmendOrder()
	// batch stats
	IncTotalBatches()
	NewBatch()
	TotalOrders() uint64
	TotalBatches() uint64
	SetAverageOrdersPerBatch(i uint64)
	SetBlockDuration(uint64)
	CurrentOrdersInBatch() uint64
	CurrentTradesInBatch() uint64
	SetOrdersPerSecond(i uint64)
	SetTradesPerSecond(i uint64)
}

type nodeProposal struct {
	*types.Proposal
	votes     map[string]struct{}
	validTime time.Time
}

// Processor handle processing of all transaction sent through the node
type Processor struct {
	log *logging.Logger
	Config
	stat              Stats
	exec              ExecutionEngine
	time              TimeService
	nodes             map[string]struct{} // all other nodes in the network
	nodeProposals     map[string]*nodeProposal
	pendingValidation []*types.Proposal
	blockCommands     [][]byte
	currentTimestamp  time.Time
	previousTimestamp time.Time
}

// NewProcessor instantiates a new transactions processor
func New(log *logging.Logger, config Config, exec ExecutionEngine, ts TimeService, stat Stats) *Processor {
	// setup logger
	log = log.Named(namedLogger)
	log.SetLevel(config.Level.Get())

	p := &Processor{
		log:               log,
		stat:              stat,
		Config:            config,
		exec:              exec,
		time:              ts,
		nodes:             map[string]struct{}{},
		nodeProposals:     map[string]*nodeProposal{},
		pendingValidation: []*types.Proposal{},
		blockCommands:     [][]byte{},
	}
	ts.NotifyOnTick(p.onTick)
	return p
}

func (p *Processor) Begin2() ([]byte, error) {
	if err := p.Begin(); err != nil {
		return nil, err
	}
	// command byte + payload
	reg := &types.NodeRegistration{
		PubKey: "this value comes from config?", // @TODO
	}
	raw, err := proto.Marshal(reg)
	if err != nil {
		return nil, err
	}
	return append([]byte{byte(blockchain.RegisterNodeCommand)}, raw...), nil
}

// Begin update timestamps
func (p *Processor) Begin() error {
	if p.log.GetLevel() == logging.DebugLevel {
		p.log.Debug("Processor service BEGIN starting")
	}
	var err error
	// Load the latest consensus block time
	if p.currentTimestamp, err = p.time.GetTimeNow(); err != nil {
		return err
	}

	if p.previousTimestamp, err = p.time.GetTimeLastBatch(); err != nil {
		return err
	}

	if p.log.GetLevel() == logging.DebugLevel {
		p.log.Debug("ABCI service BEGIN completed",
			logging.Int64("current-timestamp", p.currentTimestamp.UnixNano()),
			logging.Int64("previous-timestamp", p.previousTimestamp.UnixNano()),
			logging.String("current-datetime", vegatime.Format(p.currentTimestamp)),
			logging.String("previous-datetime", vegatime.Format(p.previousTimestamp)),
		)
	}
	return nil
}

func (p *Processor) Commit() error {
	if p.log.GetLevel() == logging.DebugLevel {
		p.log.Debug("Processor COMMIT starting")
	}
	p.stats()
	if err := p.exec.Generate(); err != nil {
		return errors.Wrap(err, "failure generating data in execution engine (commit)")
	}
	if p.log.GetLevel() == logging.DebugLevel {
		p.log.Debug("Processor COMMIT completed")
	}
	return nil
}

func (p *Processor) stats() {
	p.stat.IncTotalBatches()
	avg := p.stat.TotalOrders() / p.stat.TotalBatches()
	p.stat.SetAverageOrdersPerBatch(avg)
	duration := time.Duration(p.currentTimestamp.UnixNano() - p.previousTimestamp.UnixNano()).Seconds()
	var (
		currentOrders, currentTrades uint64
	)
	p.stat.SetBlockDuration(uint64(duration * float64(time.Second.Nanoseconds())))
	if duration > 0 {
		currentOrders, currentTrades = uint64(float64(p.stat.CurrentOrdersInBatch())/duration),
			uint64(float64(p.stat.CurrentTradesInBatch())/duration)
	}
	p.stat.SetOrdersPerSecond(currentOrders)
	p.stat.SetTradesPerSecond(currentTrades)
	// log stats
	p.log.Debug("Processor batch stats",
		logging.Int64("previousTimestamp", p.previousTimestamp.UnixNano()),
		logging.Int64("currentTimestamp", p.currentTimestamp.UnixNano()),
		logging.Float64("duration", duration),
		logging.Uint64("currentOrdersInBatch", p.stat.CurrentOrdersInBatch()),
		logging.Uint64("currentTradesInBatch", p.stat.CurrentTradesInBatch()),
		logging.Uint64("total-batches", p.stat.TotalBatches()),
		logging.Uint64("avg-orders-batch", avg),
		logging.Uint64("orders-per-sec", currentOrders),
		logging.Uint64("trades-per-sec", currentTrades),
	)
	p.stat.NewBatch() // sets previous batch orders/trades to current, zeroes current tally
}

func (p *Processor) SetTime(now time.Time) {
	p.previousTimestamp = p.currentTimestamp
	p.currentTimestamp = now
}

// ReloadConf update the internal configuration of the processor
func (p *Processor) ReloadConf(cfg Config) {
	p.log.Info("reloading configuration")
	if p.log.GetLevel() != cfg.Level.Get() {
		p.log.Info("updating log level",
			logging.String("old", p.log.GetLevel().String()),
			logging.String("new", cfg.Level.String()),
		)
		p.log.SetLevel(cfg.Level.Get())
	}

	p.Config = cfg
}

func (p *Processor) getOrder(payload []byte) (*types.Order, error) {
	order := &types.Order{}
	if err := proto.Unmarshal(payload, order); err != nil {
		return nil, err
	}
	return order, nil
}

func (p *Processor) getOrderSubmission(payload []byte) (*types.Order, error) {
	orderSubmission := &types.OrderSubmission{}
	err := proto.Unmarshal(payload, orderSubmission)
	if err != nil {
		return nil, err
	}

	order := types.Order{
		Id:          orderSubmission.Id,
		MarketID:    orderSubmission.MarketID,
		PartyID:     orderSubmission.PartyID,
		Price:       orderSubmission.Price,
		Size:        orderSubmission.Size,
		Side:        orderSubmission.Side,
		TimeInForce: orderSubmission.TimeInForce,
		Type:        orderSubmission.Type,
		ExpiresAt:   orderSubmission.ExpiresAt,
		Reference:   orderSubmission.Reference,
		Status:      types.Order_Active,
		CreatedAt:   0,
		Remaining:   orderSubmission.Size,
	}

	return &order, nil
}

func (p *Processor) getOrderCancellation(payload []byte) (*types.OrderCancellation, error) {
	order := &types.OrderCancellation{}
	err := proto.Unmarshal(payload, order)
	if err != nil {
		return nil, err
	}
	return order, nil
}

func (p *Processor) getOrderAmendment(payload []byte) (*types.OrderAmendment, error) {
	amendment := &types.OrderAmendment{}
	err := proto.Unmarshal(payload, amendment)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding order to proto")
	}
	return amendment, nil
}

func (p *Processor) getNotifyTraderAccount(payload []byte) (*types.NotifyTraderAccount, error) {
	notif := &types.NotifyTraderAccount{}
	err := proto.Unmarshal(payload, notif)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding order to proto")
	}
	return notif, nil
}

func (p *Processor) getWithdraw(payload []byte) (*types.Withdraw, error) {
	w := &types.Withdraw{}
	err := proto.Unmarshal(payload, w)
	if err != nil {
		return nil, errors.Wrap(err, "error decoding order to proto")
	}
	return w, nil
}

func (p *Processor) getProposalSubmission(payload []byte) (*types.Proposal, error) {
	proposalSubmission := &types.Proposal{}
	err := proto.Unmarshal(payload, proposalSubmission)
	if err != nil {
		return nil, err
	}
	return proposalSubmission, nil
}

func (p *Processor) getVoteSubmission(payload []byte) (*types.Vote, error) {
	voteSubmission := &types.Vote{}
	err := proto.Unmarshal(payload, voteSubmission)
	if err != nil {
		return nil, err
	}
	return voteSubmission, nil
}

// ValidateSigned - validates a signed transaction. This sits here because it's actual data processing
// related. We need to unmarshal the payload to validate the partyID
func (p *Processor) ValidateSigned(key, data []byte, cmd blockchain.Command) error {
	switch cmd {
	case blockchain.SubmitOrderCommand:
		order, err := p.getOrderSubmission(data)
		if err != nil {
			return err
		}
		// partyID is hex encoded pubkey
		if order.PartyID != hex.EncodeToString(key) {
			return ErrOrderSubmissionPartyAndPubKeyDoesNotMatch
		}
		return nil
	case blockchain.CancelOrderCommand:
		order, err := p.getOrderCancellation(data)
		if err != nil {
			return err
		}
		// partyID is hex encoded pubkey
		if order.PartyID != hex.EncodeToString(key) {
			return ErrOrderCancellationPartyAndPubKeyDoesNotMatch
		}
		return nil
	case blockchain.AmendOrderCommand:
		order, err := p.getOrderAmendment(data)
		if err != nil {
			return err
		}
		// partyID is hex encoded pubkey
		if order.PartyID != hex.EncodeToString(key) {
			return ErrOrderAmendmentPartyAndPubKeyDoesNotMatch
		}
		return nil
	case blockchain.ProposeCommand:
		proposal, err := p.getProposalSubmission(data)
		if err != nil {
			return err
		}
		// partyID is hex encoded pubkey
		if proposal.PartyID != hex.EncodeToString(key) {
			return ErrProposalSubmissionPartyAndPubKeyDoesNotMatch
		}
		return nil
	case blockchain.VoteCommand:
		vote, err := p.getVoteSubmission(data)
		if err != nil {
			return err
		}
		// partyID is hex encoded pubkey
		if vote.PartyID != hex.EncodeToString(key) {
			return ErrVoteSubmissionPartyAndPubKeyDoesNotMatch
		}
		return nil
	case blockchain.WithdrawCommand:
		withdraw, err := p.getWithdraw(data)
		if err != nil {
			return err
		}
		if withdraw.PartyID != hex.EncodeToString(key) {
			return ErrWithdrawPartyAndPublKeyDoesNotMatch
		}
		return nil
	}
	return errors.New("unknown command when validating payload")
}

// Process performs validation and then sends the command and data to
// the underlying blockchain service handlers e.g. submit order, etc.
func (p *Processor) Process(data []byte, cmd blockchain.Command) error {
	// first is that a signed or unsigned command?
	switch cmd {
	case blockchain.SubmitOrderCommand:
		order, err := p.getOrderSubmission(data)
		if err != nil {
			return err
		}
		err = p.submitOrder(order)
	case blockchain.CancelOrderCommand:
		order, err := p.getOrderCancellation(data)
		if err != nil {
			return err
		}
		return p.cancelOrder(order)
	case blockchain.AmendOrderCommand:
		order, err := p.getOrderAmendment(data)
		if err != nil {
			return err
		}
		return p.amendOrder(order)
	case blockchain.WithdrawCommand:
		withdraw, err := p.getWithdraw(data)
		if err != nil {
			return err
		}
		return p.exec.Withdraw(withdraw)
	case blockchain.ProposeCommand:
		proposal, err := p.getProposalSubmission(data)
		if err != nil {
			return err
		}
		// proposal is a new asset proposal?
		if na := proposal.Terms.GetNewAsset(); na != nil {
			p.nodeProposals[proposal.Reference] = &nodeProposal{
				Proposal: proposal,
				votes:    map[string]struct{}{},
			}
			if _, err := p.validateAsset(proposal); err != nil {
				return err
			}
			//@TODO validate proposal here + cast vote
			return nil
		}
		return p.exec.SubmitProposal(proposal)
	case blockchain.VoteCommand:
		vote, err := p.getVoteSubmission(data)
		if err != nil {
			return err
		}
		return p.exec.VoteOnProposal(vote)
	case blockchain.RegisterNodeCommand:
		node, err := p.getNodeRegistration(data)
		if err != nil {
			return err
		}
		p.nodes[node.PubKey] = struct{}{}
	case blockchain.NodeVoteCommand:
		vote, err := p.getNodeVote(data)
		if err != nil {
			return err
		}
		if _, ok := p.nodes[vote.PubKey]; !ok {
			return ErrUnknownNodeKey
		}
		prop, ok := p.nodeProposals[vote.Reference]
		if !ok {
			return ErrUnknownProposal
		}
		prop.votes[vote.PubKey] = struct{}{}
	case blockchain.NotifyTraderAccountCommand:
		notify, err := p.getNotifyTraderAccount(data)
		if err != nil {
			return err
		}
		return p.exec.NotifyTraderAccount(notify)
	default:
		p.log.Warn("Unknown command received", logging.String("command", cmd.String()))
		return fmt.Errorf("unknown command received: %s", cmd)
	}
	return nil
}

func (p *Processor) getNodeVote(payload []byte) (*types.NodeVote, error) {
	vote := &types.NodeVote{}
	if err := proto.Unmarshal(payload, vote); err != nil {
		return nil, err
	}
	return vote, nil
}

func (p *Processor) getNodeRegistration(payload []byte) (*types.NodeRegistration, error) {
	cmd := &types.NodeRegistration{}
	err := proto.Unmarshal(payload, cmd)
	if err != nil {
		return nil, err
	}
	return cmd, nil
}

func (p *Processor) submitOrder(o *types.Order) error {
	p.stat.IncTotalCreateOrder()
	if p.log.GetLevel() == logging.DebugLevel {
		p.log.Debug("Processor received a SUBMIT ORDER request", logging.Order(*o))
	}

	o.CreatedAt = p.currentTimestamp.UnixNano()

	// Submit the create order request to the execution engine
	conf, err := p.exec.SubmitOrder(o)
	if conf != nil {

		if p.log.GetLevel() == logging.DebugLevel {
			p.log.Debug("Order confirmed",
				logging.Order(*o),
				logging.OrderWithTag(*conf.Order, "aggressive-order"),
				logging.String("passive-trades", fmt.Sprintf("%+v", conf.Trades)),
				logging.String("passive-orders", fmt.Sprintf("%+v", conf.PassiveOrdersAffected)))
		}
		p.stat.AddCurrentTradesInBatch(uint64(len(conf.Trades)))
		p.stat.AddTotalTrades(uint64(len(conf.Trades)))
		p.stat.IncCurrentOrdersInBatch()
	}

	// increment total orders, even for failures so current ID strategy is valid.
	p.stat.IncTotalOrders()

	if err != nil {
		p.log.Error("error message on creating order",
			logging.Order(*o),
			logging.Error(err))
	}

	return err
}

func (p *Processor) cancelOrder(order *types.OrderCancellation) error {
	p.stat.IncTotalCancelOrder()
	if p.log.GetLevel() == logging.DebugLevel {
		p.log.Debug("Blockchain service received a CANCEL ORDER request", logging.String("order-id", order.OrderID))
	}

	// Submit the cancel new order request to the Vega trading core
	msg, err := p.exec.CancelOrder(order)
	if err != nil {
		p.log.Error("error on cancelling order",
			logging.String("order-id", order.OrderID),
			logging.Error(err),
		)
		return err
	}
	if p.LogOrderCancelDebug {
		p.log.Debug("Order cancelled", logging.Order(*msg.Order))
	}

	return nil
}

func (p *Processor) amendOrder(order *types.OrderAmendment) error {
	p.stat.IncTotalAmendOrder()
	if p.log.GetLevel() == logging.DebugLevel {
		p.log.Debug("Blockchain service received a AMEND ORDER request",
			logging.String("order", order.String()))
	}

	// Submit the Amendment new order request to the Vega trading core
	_, err := p.exec.AmendOrder(order)
	if err != nil {
		p.log.Error("Error amending order",
			logging.String("order", order.String()),
			logging.Error(err),
		)
		return err
	}
	if p.LogOrderAmendDebug {
		p.log.Debug("Order amended", logging.String("order", order.String()))
	}
	return nil
}

func (p *Processor) validateAsset(prop *types.Proposal) (bool, error) {
	asset := prop.Terms.GetNewAsset()
	if asset == nil {
		return false, ErrNotAnAssetProposal
	}
	p.log.Debug("Validating asset",
		logging.String("asset-id", asset.ID),
	)
	// need dep to validate this proposal
	// if validation failed, add to pendingPropopsals to retry every so often
	return true, nil
}

// check the asset proposals on tick
func (p *Processor) onTick(t time.Time) {
	for k, prop := range p.nodeProposals {
		// this proposal has passed the node-voting period
		if prop.validTime.Before(t) {
			// if not all nodes have approved, just remove
			if len(prop.votes) < len(p.nodes) {
				p.log.Warn("proposal was not accepted by all nodes",
					logging.String("proposal", prop.Proposal.String()),
					logging.Int("vote-count", len(prop.votes)),
					logging.Int("node-count", len(p.nodes)),
				)
			} else if err := p.exec.SubmitProposal(prop.Proposal); err != nil {
				p.log.Error("Failed to submit node-approved proposal",
					logging.String("proposal", prop.Proposal.String()),
				)
				continue // try again next block
			}
			// either proposal wasn't accepted, or it's been passed on to governance
			delete(p.nodeProposals, k)
		}
	}
}
