package statevar_test

import (
	"bytes"
	"context"
	"testing"

	snapshotpb "code.vegaprotocol.io/protos/vega/snapshot/v1"
	"code.vegaprotocol.io/vega/libs/proto"
	gtypes "code.vegaprotocol.io/vega/types"
	types "code.vegaprotocol.io/vega/types/statevar"
	"github.com/stretchr/testify/require"
)

func TestSnapshot(t *testing.T) {
	engine1 := getTestEngine(t, now).engine
	engine1.RegisterStateVariable("asset1", "market1", "var1", converter{}, defaultStartCalc(), []types.StateVarEventType{types.StateVarEventTypeMarketEnactment, types.StateVarEventTypeTimeTrigger}, defaultResultBack())
	engine1.RegisterStateVariable("asset1", "market1", "var2", converter{}, defaultStartCalc(), []types.StateVarEventType{types.StateVarEventTypeMarketEnactment, types.StateVarEventTypeTimeTrigger}, defaultResultBack())
	engine1.RegisterStateVariable("asset1", "market2", "var1", converter{}, defaultStartCalc(), []types.StateVarEventType{types.StateVarEventTypeMarketEnactment, types.StateVarEventTypeTimeTrigger}, defaultResultBack())
	engine1.RegisterStateVariable("asset1", "market2", "var2", converter{}, defaultStartCalc(), []types.StateVarEventType{types.StateVarEventTypeMarketEnactment, types.StateVarEventTypeTimeTrigger}, defaultResultBack())
	engine1.RegisterStateVariable("asset2", "market1", "var1", converter{}, defaultStartCalc(), []types.StateVarEventType{types.StateVarEventTypeMarketEnactment, types.StateVarEventTypeTimeTrigger}, defaultResultBack())
	engine1.RegisterStateVariable("asset2", "market1", "var1", converter{}, defaultStartCalc(), []types.StateVarEventType{types.StateVarEventTypeMarketEnactment, types.StateVarEventTypeTimeTrigger}, defaultResultBack())

	engine1.ReadyForTimeTrigger("asset1", "market1")
	engine1.ReadyForTimeTrigger("asset1", "market2")

	key := (&gtypes.PayloadFloatingPointConsensus{}).Key()
	state1, _, err := engine1.GetState(key)
	require.NoError(t, err)

	engine2 := getTestEngine(t, now).engine
	engine2.RegisterStateVariable("asset1", "market1", "var1", converter{}, defaultStartCalc(), []types.StateVarEventType{types.StateVarEventTypeMarketEnactment, types.StateVarEventTypeTimeTrigger}, defaultResultBack())
	engine2.RegisterStateVariable("asset1", "market1", "var2", converter{}, defaultStartCalc(), []types.StateVarEventType{types.StateVarEventTypeMarketEnactment, types.StateVarEventTypeTimeTrigger}, defaultResultBack())
	engine2.RegisterStateVariable("asset1", "market2", "var1", converter{}, defaultStartCalc(), []types.StateVarEventType{types.StateVarEventTypeMarketEnactment, types.StateVarEventTypeTimeTrigger}, defaultResultBack())
	engine2.RegisterStateVariable("asset1", "market2", "var2", converter{}, defaultStartCalc(), []types.StateVarEventType{types.StateVarEventTypeMarketEnactment, types.StateVarEventTypeTimeTrigger}, defaultResultBack())
	engine2.RegisterStateVariable("asset2", "market1", "var1", converter{}, defaultStartCalc(), []types.StateVarEventType{types.StateVarEventTypeMarketEnactment, types.StateVarEventTypeTimeTrigger}, defaultResultBack())
	engine2.RegisterStateVariable("asset2", "market1", "var1", converter{}, defaultStartCalc(), []types.StateVarEventType{types.StateVarEventTypeMarketEnactment, types.StateVarEventTypeTimeTrigger}, defaultResultBack())

	pl := snapshotpb.Payload{}
	require.NoError(t, proto.Unmarshal(state1, &pl))
	engine2.LoadState(context.Background(), gtypes.PayloadFromProto(&pl))

	state2, _, err := engine2.GetState(key)
	require.NoError(t, err)
	require.True(t, bytes.Equal(state1, state2))
}

func TestSnapshotChangeFlagSet(t *testing.T) {
	key := (&gtypes.PayloadFloatingPointConsensus{}).Key()
	engine1 := getTestEngine(t, now).engine

	engine1.RegisterStateVariable("asset1", "market1", "var1", converter{}, defaultStartCalc(), []types.StateVarEventType{types.StateVarEventTypeMarketEnactment, types.StateVarEventTypeTimeTrigger}, defaultResultBack())

	state1, _, err := engine1.GetState(key)
	require.NoError(t, err)

	// this should hit the change flag causing us to reserialise at the next hash
	engine1.ReadyForTimeTrigger("asset1", "market1")

	state2, _, err := engine1.GetState(key)
	require.NoError(t, err)
	require.False(t, bytes.Equal(state1, state2))
}
