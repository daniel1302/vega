package subscribers

import (
	"context"

	"code.vegaprotocol.io/vega/events"
	types "code.vegaprotocol.io/vega/proto"
)

type GovernanceEvent interface {
	events.Event
	ProposalID() string
	PartyID() string
}

type VoteE interface {
	GovernanceEvent
	Vote() types.Vote
	Value() types.Vote_Value
}

// ProposalFilter - a callback to be applied to the proposals we're interested in
// some common filter calls will be provided
type ProposalFilter func(p types.Proposal) bool

// GovernanceFilter - callback to filter out proposal and vote events we're after
type GovernanceFilter func(e GovernanceEvent) bool

// VoteFilter - callbacks to filter out only the vote events we're after
type VoteFilter func(v types.Vote) bool

// Filter - variadic argument for constructor so we can set different types of filters
// as a single vararg
type Filter func(g *GovernanceSub)

type GovernanceSub struct {
	*Base
	gfilters []GovernanceFilter
	pfilters []ProposalFilter
	vfilters []VoteFilter
	combined []*types.GovernanceData
	byPID    map[string]*types.GovernanceData
}

// Governance - vararg to set governance filters
func Governance(f ...GovernanceFilter) Filter {
	return func(g *GovernanceSub) {
		g.gfilters = f
	}
}

// Proposals - varargs setting filters for proposals
func Proposals(f ...ProposalFilter) Filter {
	return func(g *GovernanceSub) {
		g.pfilters = f
	}
}

// Votes - vararg setting filters on votes
func Votes(f ...VoteFilter) Filter {
	return func(g *GovernanceSub) {
		g.vfilters = f
	}
}

func NewGovernanceSub(ctx context.Context, filters ...Filter) *GovernanceSub {
	g := GovernanceSub{
		Base:     newBase(ctx, 10),
		gfilters: []GovernanceFilter{},
		pfilters: []ProposalFilter{},
		vfilters: []VoteFilter{},
		combined: []*types.GovernanceData{},
		byPID:    map[string]*types.GovernanceData{},
	}
	for _, f := range filters {
		f(&g)
	}
	g.running = true
	go g.loop(g.ctx)
	return &g
}

func (g *GovernanceSub) loop(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			g.Halt()
			return
		case e := <-g.ch:
			if g.isRunning() {
				g.Push(e)
			}
		}
	}
}

func (g *GovernanceSub) filter(e GovernanceEvent) bool {
	for _, f := range g.gfilters {
		if !f(e) {
			return false
		}
	}
	switch et := e.(type) {
	case PropE:
		p := et.Proposal()
		for _, f := range g.pfilters {
			if !f(p) {
				return false
			}
		}
	case VoteE:
		v := et.Vote()
		for _, f := range g.vfilters {
			if !f(v) {
				return false
			}
		}
	}
	return true
}

func (g *GovernanceSub) Push(e events.Event) {
	// if this is a governance event, apply filters to only get events we need/want
	if ge, ok := e.(GovernanceEvent); ok {
		if !g.filter(ge) {
			return
		}
	}
	switch et := e.(type) {
	case PropE:
		prop := et.Proposal()
		gd := g.getData(prop.ID)
		gd.Proposal = &prop
	case VoteE:
		vote := et.Vote()
		gd = g.getData(vote.ProposalID)
		if vote.Value == types.Vote_VALUE_YES {
			delete(gd.NoParty, vote.PartyID)
			gd.YesParty[vote.PartyID] = &vote
		} else {
			delete(gd.YesParty, vote.PartyID)
			gd.NoParty[vote.PartyID] = &vote
		}
	}
}

func (g *GovernanceSub) getData(id string) *types.GovernanceData {
	gd, ok := g.byPID[id]
	if !ok {
		gd = &types.GovernanceData{
			Yes:      []*types.Vote{},
			No:       []*types.Vote{},
			YesParty: map[string]*types.Vote{},
			NoParty:  map[string]*types.Vote{},
		}
		g.byPID[id] = gd
		g.combined = append(g.combined, gd)
	}
	return gd
}

func (g *GovernanceSub) Types() []events.Type {
	return []events.Type{
		events.ProposalEvent,
		events.VoteEvent,
	}
}

// GetGovernanceData - returns current data, this is a VALUE RECEIVER for a reason
// pointer recevers would cause data races
func (g GovernanceSub) GetGovernanceData() []*types.GovernanceData {
	data := g.combined
	// create a copy
	ret := make([]*types.GovernanceData, 0, len(data))
	// copy the votes
	for _, d := range data {
		cpy := *d
		cpy.Yes = make([]*types.Vote, 0, len(cpy.YesParty))
		for _, v := range cpy.YesParty {
			vc := *v
			cpy.Yes = append(cpy.Yes, &vc)
		}
		cpy.No = make([]*types.Vote, 0, len(cpy.NoParty))
		for _, v := range cpy.NoParty {
			vc := *v
			cpy.No = append(cpy.No, &vc)
		}
		ret = append(ret, &cpy)
	}
	return ret
}
