package orders

import (
	"context"
	"sync"
	"testing"

	"code.vegaprotocol.io/vega/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestObserveOrders(t *testing.T) {
	t.Run("Observe orders - all markets/parties success", testObserveAllOrdersSuccess)
}

func testObserveAllOrdersSuccess(t *testing.T) {
	svc := getTestService(t)
	defer svc.ctrl.Finish()
	ctx, cfunc := context.WithCancel(context.Background())
	wg := sync.WaitGroup{}
	done := make(chan struct{})
	subRef := uint64(1)
	orders := []proto.Order{
		{
			Id:     "order_id1",
			Market: "market1",
			Party:  "party1",
		},
		{
			Id:     "order_id2",
			Market: "market2",
			Party:  "party2",
		},
	}

	wg.Add(1)
	subscriber := func(ch chan<- []proto.Order) {
		defer wg.Done()
		ch <- orders
	}
	svc.orderStore.EXPECT().Subscribe(gomock.Any()).Times(1).Return(subRef).Do(func(ch chan<- []proto.Order) {
		go subscriber(ch)
	})
	svc.orderStore.EXPECT().Unsubscribe(subRef).Times(1).Return(nil).Do(func(_ uint64) {
		done <- struct{}{}
	})
	// all orders
	ch, ref := svc.svc.ObserveOrders(ctx, 0, nil, nil)
	assert.Equal(t, subRef, ref)
	gotOrders := <-ch

	wg.Wait()
	cfunc() // cancel context
	<-done
	assert.Equal(t, len(orders), len(gotOrders))
	for i := range orders {
		assert.Equal(t, orders[i], gotOrders[i])
	}
}
