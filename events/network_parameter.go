package events

import (
	"context"

	proto "code.vegaprotocol.io/protos/vega"
	eventspb "code.vegaprotocol.io/protos/vega/events/v1"
)

type NetworkParameter struct {
	*Base
	np proto.NetworkParameter
}

func NewNetworkParameterEvent(ctx context.Context, key, value string) *NetworkParameter {
	return &NetworkParameter{
		Base: newBase(ctx, NetworkParameterEvent),
		np:   proto.NetworkParameter{Key: key, Value: value},
	}
}

func (n *NetworkParameter) NetworkParameter() proto.NetworkParameter {
	return n.np
}

func (n NetworkParameter) Proto() proto.NetworkParameter {
	return n.np
}

func (n NetworkParameter) StreamMessage() *eventspb.BusEvent {
	return &eventspb.BusEvent{
		Version: eventspb.Version,
		Id:      n.eventID(),
		Block:   n.TraceID(),
		Type:    n.et.ToProto(),
		Event: &eventspb.BusEvent_NetworkParameter{
			NetworkParameter: &n.np,
		},
	}
}

func NetworkParameterEventFromStream(ctx context.Context, be *eventspb.BusEvent) *NetworkParameter {
	return &NetworkParameter{
		Base: newBaseFromStream(ctx, NetworkParameterEvent, be),
		np:   proto.NetworkParameter{Key: be.GetNetworkParameter().Key, Value: be.GetNetworkParameter().Value},
	}
}
