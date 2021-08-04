package events

import (
	"context"

	eventspb "code.vegaprotocol.io/protos/vega/events/v1"
	"code.vegaprotocol.io/vega/types/num"
)

type DelegationBalance struct {
	*Base
	party  string
	nodeID string
	amount *num.Uint
}

func NewDelegationBalance(ctx context.Context, party, nodeID string, amount *num.Uint) *DelegationBalance {
	return &DelegationBalance{
		Base:   newBase(ctx, DelegationBalanceEvent),
		party:  party,
		nodeID: nodeID,
		amount: amount,
	}
}

func (db DelegationBalance) Proto() eventspb.DelegationBalanceEvent {
	return eventspb.DelegationBalanceEvent{
		Party:  db.party,
		NodeId: db.nodeID,
		Amount: db.amount.Uint64(),
	}
}

func (db DelegationBalance) StreamMessage() *eventspb.BusEvent {
	p := db.Proto()
	return &eventspb.BusEvent{
		Id:    db.eventID(),
		Block: db.TraceID(),
		Type:  db.et.ToProto(),
		Event: &eventspb.BusEvent_DelegationBalance{
			DelegationBalance: &p,
		},
	}
}

type PendingDelegationBalance struct {
	*Base
	party              string
	nodeID             string
	delegationAmount   *num.Uint
	undelegationAmount *num.Uint
}

func NewPendingDelegationBalance(ctx context.Context, party, nodeID string, delegationAmount *num.Uint, undelegationAmount *num.Uint) *PendingDelegationBalance {
	return &PendingDelegationBalance{
		Base:               newBase(ctx, PendingDelegationBalanceEvent),
		party:              party,
		nodeID:             nodeID,
		delegationAmount:   delegationAmount,
		undelegationAmount: undelegationAmount,
	}
}

func (pdb PendingDelegationBalance) Proto() eventspb.PendingDelegationBalanceEvent {
	return eventspb.PendingDelegationBalanceEvent{
		Party:              pdb.party,
		NodeId:             pdb.nodeID,
		DelegationAmount:   pdb.delegationAmount.Uint64(),
		UndelegationAmount: pdb.undelegationAmount.Uint64(),
	}
}

func (pdb PendingDelegationBalance) StreamMessage() *eventspb.BusEvent {
	p := pdb.Proto()
	return &eventspb.BusEvent{
		Id:    pdb.eventID(),
		Block: pdb.TraceID(),
		Type:  pdb.et.ToProto(),
		Event: &eventspb.BusEvent_PendingDelegationBalance{
			PendingDelegationBalance: &p,
		},
	}
}
