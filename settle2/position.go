package settle2

import (
	"code.vegaprotocol.io/vega/events"
	types "code.vegaprotocol.io/vega/proto"

	"github.com/pkg/errors"
)

var (
	ErrPartyDoesNotMatch = errors.New("event party and position party do not match")
)

// See positions.MarketPosition
type pos struct {
	// embed the type, we will copy the three main fields because those should be immutable
	// which we can't guarantee through an embedded interface
	events.MarketPosition
	party string
	size  int64
	price uint64
}

type mtmTransfer struct {
	events.MarketPosition
	transfer *types.Transfer
}

func newPos(evt events.MarketPosition) *pos {
	return &pos{
		MarketPosition: evt,
		party:          evt.Party(),
		size:           evt.Size(),
		price:          evt.Price(),
	}
}

// update - set the size/price of an event accordingly
func (p *pos) update(evt events.MarketPosition) error {
	// this check, in theory, should not be needed...
	if p.party != evt.Party() {
		return ErrPartyDoesNotMatch
	}
	// embed updated event
	p.MarketPosition = evt
	p.size = evt.Size()
	p.price = evt.Price()
	return nil
}

// Party - part of the MarketPosition interface, used to update position after SettlePreTrade
func (p pos) Party() string {
	return p.party
}

// Size - part of the MarketPosition interface, used to update position after SettlePreTrade
func (p pos) Size() int64 {
	return p.size
}

// Price - part of the MarketPosition interface, used to update position after SettlePreTrade
func (p pos) Price() uint64 {
	return p.price
}

// Transfer - part of the Transfer interface
func (m mtmTransfer) Transfer() *types.Transfer {
	return m.transfer
}
