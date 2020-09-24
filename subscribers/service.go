package subscribers

import (
	"context"

	"code.vegaprotocol.io/vega/broker"
	"code.vegaprotocol.io/vega/events"
	types "code.vegaprotocol.io/vega/proto"
)

//go:generate go run github.com/golang/mock/mockgen -destination mocks/event_bus_mock.go -package mocks code.vegaprotocol.io/vega/events Broker
type Broker interface {
	Subscribe(s broker.Subscriber) int
	Unsubscribe(id int)
}

type Service struct {
	broker Broker
}

func NewService(broker Broker) *Service {
	return &Service{
		broker: broker,
	}
}

func (s *Service) ObserveEvents(ctx context.Context, retries int, eTypes []events.Type, filters ...EventFilter) <-chan []*types.BusEvent {
	out := make(chan []*types.BusEvent)
	ctx, cfunc := context.WithCancel(ctx)
	// use stream subscriber
	sub := NewStreamSub(ctx, eTypes, filters...)
	id := s.broker.Subscribe(sub)
	go func() {
		defer func() {
			s.broker.Unsubscribe(id)
			close(out)
			cfunc()
		}()
		ret := retries
		for {
			// wait for actual changes
			data := sub.GetData()
			select {
			case <-ctx.Done():
				return
			case out <- data:
				ret = retries
			default:
				if ret == 0 {
					return
				}
				ret--
			}
		}
	}()
	return out
}
