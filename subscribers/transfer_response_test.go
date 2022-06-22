// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package subscribers_test

import (
	"context"
	"sync"
	"testing"
	"time"

	"code.vegaprotocol.io/data-node/logging"
	"code.vegaprotocol.io/data-node/subscribers"
	"code.vegaprotocol.io/data-node/subscribers/mocks"
	types "code.vegaprotocol.io/protos/vega"
	"code.vegaprotocol.io/vega/events"

	eventspb "code.vegaprotocol.io/protos/vega/events/v1"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type timeStub struct {
	t time.Time
}

type trStub struct {
	r []*types.TransferResponse
}

type trTst struct {
	*subscribers.TransferResponse
	ctrl  *gomock.Controller
	ctx   context.Context
	cfunc context.CancelFunc
	store *mocks.MockTransferResponseStore
}

func getTestSub(t *testing.T, ack bool) *trTst {
	ctrl := gomock.NewController(t)
	ctx, cfunc := context.WithCancel(context.Background())
	store := mocks.NewMockTransferResponseStore(ctrl)
	tr := subscribers.NewTransferResponse(ctx, store, logging.NewTestLogger(), ack)
	return &trTst{
		TransferResponse: tr,
		ctrl:             ctrl,
		ctx:              ctx,
		cfunc:            cfunc,
		store:            store,
	}
}

func TestTypes(t *testing.T) {
	tr := getTestSub(t, true)
	defer tr.Finish()
	types := tr.Types()
	assert.Equal(t, 2, len(types))
	assert.Contains(t, types, events.TimeUpdate)
	assert.Contains(t, types, events.TransferResponses)
}

func TestPushME(t *testing.T) {
	t.Run("Push several transfer batches, then push time event", testPushSeveralSuccess)
}

func TestChannelPushOptional(t *testing.T) {
	t.Run("Events are sent through channel", testChannelOptionalSuccess)
	t.Run("No events are sent when the subscriber is paused", testChannelOptionalSkip)
}

func testPushSeveralSuccess(t *testing.T) {
	tr := getTestSub(t, true)
	defer tr.Finish()
	time := timeStub{
		t: time.Now(),
	}
	trs := []*types.TransferResponse{
		{},
		{},
	}
	stubs := []*trStub{
		{
			r: []*types.TransferResponse{trs[0]},
		},
		{
			r: []*types.TransferResponse{trs[1]},
		},
	}
	for _, e := range stubs {
		tr.Push(e)
	}
	// only now do we expect a call to the store:
	tr.store.EXPECT().SaveBatch(gomock.Any()).Times(1).Return(nil).Do(func(_ []*types.TransferResponse) {
	})
	tr.Push(time)
}

func testChannelOptionalSuccess(t *testing.T) {
	tr := getTestSub(t, false)
	defer tr.Finish()
	resps := []*types.TransferResponse{
		{},
		{},
	}
	evt := &trStub{
		r: resps,
	}
	skipped := false
	select {
	case <-tr.Skip():
		skipped = true
	case <-tr.Closed():
		t.FailNow()
	case tr.C() <- []events.Event{evt}:
		skipped = false
	}
	assert.False(t, skipped)
	time := timeStub{
		t: time.Now(),
	}
	wg := sync.WaitGroup{}
	wg.Add(1)
	tr.store.EXPECT().SaveBatch(gomock.Any()).Times(1).Return(nil).Do(func(_ []*types.TransferResponse) {
		wg.Done()
	})
	// push time event
	tr.C() <- []events.Event{time}
	// Make sure the time event triggers the call to the storage
	wg.Wait()
}

func testChannelOptionalSkip(t *testing.T) {
	tr := getTestSub(t, false)
	defer tr.Finish()
	resps := []*types.TransferResponse{
		{},
		{},
	}
	tr.Pause()
	evt := &trStub{
		r: resps,
	}
	skipped := false
	select {
	case <-tr.Skip():
		skipped = true
	case tr.C() <- []events.Event{evt}:
		skipped = false
	}
	assert.True(t, skipped)
	time := timeStub{
		t: time.Now(),
	}
	// no expected calls to storage
	tr.Resume()
	tr.Push(time)
}

func (t *trTst) Finish() {
	t.cfunc()
	t.ctrl.Finish()
}

func (t timeStub) Context() context.Context {
	return context.TODO()
}

func (t timeStub) Time() time.Time {
	return t.t
}

func (t timeStub) TraceID() string {
	return "test-trace-id"
}

func (t timeStub) TxHash() string {
	return "test-tx-hash"
}

func (t timeStub) Type() events.Type {
	return events.TimeUpdate
}

func (t trStub) Context() context.Context {
	return context.TODO()
}

func (t trStub) Type() events.Type {
	return events.TransferResponses
}

func (t trStub) TraceID() string {
	return "test-trace-id"
}

func (t trStub) TxHash() string {
	return "test-tx-hash"
}

func (t trStub) TransferResponses() []*types.TransferResponse {
	return t.r
}

func (t trStub) StreamMessage() *eventspb.BusEvent {
	return &eventspb.BusEvent{}
}

func (t trStub) SetSequenceID(s uint64) {}
func (t trStub) Sequence() uint64       { return 0 }
func (t trStub) ChainID() string        { return "testchain" }
func (t trStub) BlockNr() int64         { return 0 }

func (t timeStub) SetSequenceID(s uint64) {}
func (t timeStub) Sequence() uint64       { return 0 }
func (t timeStub) ChainID() string        { return "testchain" }
func (t timeStub) BlockNr() int64         { return 0 }

func (t timeStub) StreamMessage() *eventspb.BusEvent {
	return &eventspb.BusEvent{}
}
