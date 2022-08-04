// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE.VEGA file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package execution

import (
	"context"
	"fmt"
	"sort"

	"code.vegaprotocol.io/vega/core/events"
	"code.vegaprotocol.io/vega/core/types"
	"code.vegaprotocol.io/vega/libs/num"
	"code.vegaprotocol.io/vega/logging"
)

type PeggedOrders struct {
	log         *logging.Logger
	timeService TimeService
	// mapping orderid -> partyid
	orders map[string]string

	// parked list
	parked   []*types.Order
	isParked map[string]struct{}
}

func NewPeggedOrders(log *logging.Logger, ts TimeService) *PeggedOrders {
	return &PeggedOrders{
		log:         log,
		timeService: ts,
		orders:      map[string]string{},
		parked:      []*types.Order{},
		isParked:    map[string]struct{}{},
	}
}

func NewPeggedOrdersFromSnapshot(
	log *logging.Logger,
	ts TimeService,
	state *types.PeggedOrdersState,
) *PeggedOrders {
	p := NewPeggedOrders(log, ts)
	p.orders = state.Orders
	p.parked = state.Parked
	for _, v := range p.parked {
		p.isParked[v.ID] = struct{}{}
	}
	return p
}

func (p *PeggedOrders) Changed() bool {
	return true
}

func (p *PeggedOrders) GetState() *types.PeggedOrdersState {
	ordersCopy := make(map[string]string, len(p.orders))
	for k, v := range p.orders {
		ordersCopy[k] = v
	}

	parkedCopy := make([]*types.Order, 0, len(p.parked))
	for _, v := range p.parked {
		parkedCopy = append(parkedCopy, v.Clone())
	}

	return &types.PeggedOrdersState{
		Orders: ordersCopy,
		Parked: parkedCopy,
	}
}

func (p *PeggedOrders) IsParked(id string) bool {
	_, parked := p.isParked[id]
	return parked
}

func (p *PeggedOrders) Park(o *types.Order) {
	fmt.Printf("PARKING: %v\n", o.ID)
	appendOrder("park", o)
	o.UpdatedAt = p.timeService.GetTimeNow().UnixNano()
	o.Status = types.OrderStatusParked
	o.Price = num.UintZero()
	o.OriginalPrice = num.UintZero()
	p.park(o)
}

func (p *PeggedOrders) park(o *types.Order) {
	p.parked = append(p.parked, o)
	p.isParked[o.ID] = struct{}{}
	p.hasDups()
}

func (p *PeggedOrders) Unpark(o *types.Order) {
	fmt.Printf("UNPARKING: %v\n", o.ID)
	for i, po := range p.parked {
		if po.ID == o.ID {
			appendOrder("rm pegf", po)
			// Remove item from slice
			copy(p.parked[i:], p.parked[i+1:])
			p.parked[len(p.parked)-1] = nil
			p.parked = p.parked[:len(p.parked)-1]
			delete(p.isParked, o.ID)
			return
		}
	}
}

func (p *PeggedOrders) GetParkedByID(id string) *types.Order {
	for _, o := range p.parked {
		if o.ID == id {
			return o
		}
	}
	return nil
}

func (p *PeggedOrders) hasDups() {
	orders := map[string]struct{}{}
	for _, v := range p.parked {
		if _, ok := orders[v.ID]; ok {
			panic(fmt.Sprintf("duplicate order: %v", v.ID))
		}
		orders[v.ID] = struct{}{}
	}

}

func (p *PeggedOrders) Add(o *types.Order) {
	appendOrder("add peg", o)
	p.orders[o.ID] = o.Party
	p.hasDups()
}

// Remove from the parked list AND the list of pegged orders
func (p *PeggedOrders) Remove(o *types.Order) {
	defer func() { p.hasDups() }()
	appendOrder("rm peg", o)
	// delete from the list
	delete(p.orders, o.ID)
	// remove if parked
	p.Unpark(o)
}

func (p *PeggedOrders) AmendParked(amended *types.Order) {
	defer func() { p.hasDups() }()
	appendOrder("amend peg", amended)
	for i, o := range p.parked {
		if o.ID == amended.ID {
			appendOrder("amend pegf", o)
			p.parked[i] = amended
			return
		}
	}

	p.log.Panic("tried to amend a non parked order from the parked list", logging.Order(amended))
}

func (p *PeggedOrders) RemoveAllForParty(
	ctx context.Context, party string, status types.OrderStatus,
) (orders []*types.Order, evts []events.Event) {
	defer func() { p.hasDups() }()
	n := 0
	now := p.timeService.GetTimeNow().UnixNano()

	// first all pegged IDs
	for oid, pid := range p.orders {
		if pid == party {
			delete(p.orders, oid)
		}
	}

	// then we look at the parked and delete + create events
	for _, o := range p.parked {
		if o.Party == party /* && o.Status == types.Order_STATUS_PARKED */ {
			o.UpdatedAt = now
			o.Status = status
			orders = append(orders, o)
			evts = append(evts, events.NewOrderEvent(ctx, o))
			delete(p.isParked, o.ID)
			continue
		}
		// here we insert back in the slice
		p.parked[n] = o
		n++
	}
	p.parked = p.parked[:n]

	return
}

func (p *PeggedOrders) RemoveAllParkedForParty(
	ctx context.Context, party string, status types.OrderStatus,
) (orders []*types.Order, evts []events.Event) {
	defer func() { p.hasDups() }()
	n := 0
	now := p.timeService.GetTimeNow().UnixNano()

	for _, o := range p.parked {
		if o.Party == party {
			o.UpdatedAt = now
			o.Status = status
			orders = append(orders, o)
			evts = append(evts, events.NewOrderEvent(ctx, o))
			delete(p.isParked, o.ID)
			delete(p.orders, o.ID)
			continue
		}
		// here we insert back in the slice
		p.parked[n] = o
		n++
	}
	p.parked = p.parked[:n]
	return
}

func (p *PeggedOrders) GetAllActiveOrders() (orders []string) {
	defer func() { p.hasDups() }()
	for k, _ := range p.orders {
		fmt.Printf("order: %v\n", k)
		if _, parked := p.isParked[k]; !parked {
			orders = append(orders, k)
		}
	}

	fmt.Printf("PARK: %#v\n", p.isParked)
	for _, v := range p.parked {
		fmt.Printf("%v\n", v.String())
	}
	return
}

func (p *PeggedOrders) GetIDs() []string {
	ids := make([]string, 0, len(p.orders))
	for k, _ := range p.orders {
		ids = append(ids, k)
	}

	sort.Strings(ids)
	return ids
}

func (p *PeggedOrders) GetAllParkedForParty(party string) (orders []*types.Order) {
	defer func() { p.hasDups() }()
	for _, order := range p.parked {
		if order.Party == party {
			orders = append(orders, order)
		}
	}
	return
}

func (p *PeggedOrders) GetAllForParty(
	party string,
) (live []string, parked []*types.Order) {
	parkedIDs := map[string]struct{}{}
	for _, order := range p.parked {
		if order.Party == party {
			parked = append(parked, order)
			parkedIDs[order.ID] = struct{}{}
		}
	}

	// now iterate over the whole list and get non parked
	for k, _ := range p.orders {
		if _, ok := parkedIDs[k]; !ok {
			live = append(live, k)
		}
	}

	sort.Strings(live)

	return live, parked
}

func (p *PeggedOrders) Settled() []*types.Order {
	defer func() { p.hasDups() }()
	// now we can remove the pegged orders too
	peggedOrders := make([]*types.Order, 0, len(p.parked))
	for _, v := range p.parked {
		order := v.Clone()
		order.Status = types.OrderStatusStopped
		peggedOrders = append(peggedOrders, order)
	}
	sort.Slice(peggedOrders, func(i, j int) bool {
		return peggedOrders[i].ID < peggedOrders[j].ID
	})

	p.parked = nil

	return peggedOrders
}
