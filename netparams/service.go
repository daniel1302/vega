package netparams

import (
	"context"
	"sync"

	"code.vegaprotocol.io/vega/events"
	types "code.vegaprotocol.io/vega/proto"
	"code.vegaprotocol.io/vega/subscribers"
)

type NetParamEvent interface {
	events.Event
	NetParam() types.NetworkParameter
}

type Service struct {
	*subscribers.Base

	params map[string]types.NetworkParameter
	mu     sync.RWMutex
	ch     chan types.NetworkParameter
}

func NewService(ctx context.Context) *Service {
	return &Service{
		Base:   subscribers.NewBase(ctx, 10, true),
		params: map[string]types.NetworkParameter{},
	}
}

// GetAll return the list of all current network parameters
func (s *Service) GetAll() []types.NetworkParameter {
	s.mu.RLock()
	defer s.mu.RUnlock()

	out := make([]types.NetworkParameter, 0, len(s.params))
	for _, v := range s.params {
		out = append(out, v)
	}
	return out
}

func (s *Service) Push(evts ...events.Event) {
	for _, e := range evts {
		if nse, ok := e.(NetParamEvent); ok {
			s.ch <- nse.NetParam()
		}
	}
}

func (s *Service) consume() {
	defer func() { close(s.ch) }()
	for {
		select {
		case <-s.Closed():
			return
		case np, ok := <-s.ch:
			if !ok {
				// cleanup base
				s.Halt()
				// channel is closed
				return
			}
			s.mu.Lock()
			s.params[np.Key] = np
			s.mu.Unlock()
		}
	}
}

func (n *Service) Types() []events.Type {
	return []events.Type{
		events.NetworkParameterEvent,
	}
}
