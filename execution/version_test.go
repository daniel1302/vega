package execution_test

import (
	"context"
	"testing"
	"time"

	types "code.vegaprotocol.io/vega/proto"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

func TestVersioning(t *testing.T) {
	party1 := "party1"
	now := time.Unix(10, 0)
	closingAt := time.Unix(10000000000, 0)
	tm := getTestMarket(t, now, closingAt)
	price := uint64(100)
	size := uint64(100)

	addAccount(tm, party1)
	tm.orderStore.EXPECT().Add(gomock.Any()).Times(1)
	tm.accountBuf.EXPECT().Add(gomock.Any()).AnyTimes()

	orderBuy := &types.Order{
		Status:      types.Order_STATUS_ACTIVE,
		Type:        types.Order_LIMIT,
		TimeInForce: types.Order_TIF_GTC,
		Id:          "someid",
		Side:        types.Side_SIDE_BUY,
		PartyID:     party1,
		MarketID:    tm.market.GetID(),
		Size:        size,
		Price:       price,
		Remaining:   100,
		CreatedAt:   now.UnixNano(),
		Reference:   "party1-buy-order",
	}
	// Create an order and check version is set to 1
	confirmation, err := tm.market.SubmitOrder(context.TODO(), orderBuy)
	assert.NotNil(t, confirmation)
	assert.NoError(t, err)
	assert.EqualValues(t, confirmation.GetOrder().Version, uint64(1))

	orderID := confirmation.GetOrder().Id

	// Amend price up, check version moves to 2
	amend := &types.OrderAmendment{
		OrderID:  orderID,
		MarketID: tm.market.GetID(),
		PartyID:  party1,
		Price:    &types.Price{Value: price + 1},
	}

	tm.orderStore.EXPECT().Add(gomock.Any()).Times(1).Do(func(order types.Order) {
		assert.EqualValues(t, order.Id, orderID)
		assert.EqualValues(t, order.Price, price+1)
		assert.EqualValues(t, order.Size, size)
		assert.EqualValues(t, order.Version, uint64(2))
	})
	amendment, err := tm.market.AmendOrder(context.TODO(), amend)
	assert.NotNil(t, amendment)
	assert.NoError(t, err)

	// Amend price down, check version moves to 3
	amend.Price = &types.Price{Value: price - 1}
	tm.orderStore.EXPECT().Add(gomock.Any()).Times(1).Do(func(order types.Order) {
		assert.EqualValues(t, order.Id, orderID)
		assert.EqualValues(t, order.Price, price-1)
		assert.EqualValues(t, order.Size, size)
		assert.EqualValues(t, order.Version, uint64(3))
	})
	amendment, err = tm.market.AmendOrder(context.TODO(), amend)
	assert.NotNil(t, amendment)
	assert.NoError(t, err)

	// Amend quantity up, check version moves to 4
	amend.Price = nil
	amend.SizeDelta = 1
	tm.orderStore.EXPECT().Add(gomock.Any()).Times(1).Do(func(order types.Order) {
		assert.EqualValues(t, order.Id, orderID)
		assert.EqualValues(t, order.Price, price-1)
		assert.EqualValues(t, order.Size, size+1)
		assert.EqualValues(t, order.Remaining, size+1)
		assert.EqualValues(t, order.Version, uint64(4))
	})
	amendment, err = tm.market.AmendOrder(context.TODO(), amend)
	assert.NotNil(t, amendment)
	assert.NoError(t, err)

	// Amend quantity down, check version moves to 5
	amend.Price = nil
	amend.SizeDelta = -2
	tm.orderStore.EXPECT().Add(gomock.Any()).Times(1).Do(func(order types.Order) {
		assert.EqualValues(t, order.Id, orderID)
		assert.EqualValues(t, order.Price, price-1)
		assert.EqualValues(t, order.Size, size-1)
		assert.EqualValues(t, order.Remaining, size-1)
		assert.EqualValues(t, order.Version, uint64(5))
	})
	amendment, err = tm.market.AmendOrder(context.TODO(), amend)
	assert.NotNil(t, amendment)
	assert.NoError(t, err)

	// Flip to GTT, check version moves to 6
	amend.TimeInForce = types.Order_TIF_GTT
	amend.ExpiresAt = &types.Timestamp{Value: now.UnixNano() + 100000000000}
	amend.SizeDelta = 0
	tm.orderStore.EXPECT().Add(gomock.Any()).Times(1).Do(func(order types.Order) {
		assert.EqualValues(t, order.Id, orderID)
		assert.EqualValues(t, order.Price, price-1)
		assert.EqualValues(t, order.Size, size-1)
		assert.EqualValues(t, order.Remaining, size-1)
		assert.EqualValues(t, order.TimeInForce, types.Order_TIF_GTT)
		assert.EqualValues(t, order.Version, uint64(6))
	})
	amendment, err = tm.market.AmendOrder(context.TODO(), amend)
	assert.NotNil(t, amendment)
	assert.NoError(t, err)

	// Update expiry time, check version moves to 7
	amend.ExpiresAt = &types.Timestamp{Value: now.UnixNano() + 100000000000}
	tm.orderStore.EXPECT().Add(gomock.Any()).Times(1).Do(func(order types.Order) {
		assert.EqualValues(t, order.Id, orderID)
		assert.EqualValues(t, order.Price, price-1)
		assert.EqualValues(t, order.Size, size-1)
		assert.EqualValues(t, order.Remaining, size-1)
		assert.EqualValues(t, order.TimeInForce, types.Order_TIF_GTT)
		assert.EqualValues(t, order.Version, uint64(7))
	})
	amendment, err = tm.market.AmendOrder(context.TODO(), amend)
	assert.NotNil(t, amendment)
	assert.NoError(t, err)

	// Flip back GTC, check version moves to 8
	amend.TimeInForce = types.Order_TIF_GTC
	amend.ExpiresAt = nil
	tm.orderStore.EXPECT().Add(gomock.Any()).Times(1).Do(func(order types.Order) {
		assert.EqualValues(t, order.Id, orderID)
		assert.EqualValues(t, order.Price, price-1)
		assert.EqualValues(t, order.Size, size-1)
		assert.EqualValues(t, order.Remaining, size-1)
		assert.EqualValues(t, order.TimeInForce, types.Order_TIF_GTC)
		assert.EqualValues(t, order.Version, uint64(8))
	})
	amendment, err = tm.market.AmendOrder(context.TODO(), amend)
	assert.NotNil(t, amendment)
	assert.NoError(t, err)

	// Cancel the order and check the version does not increase from 8
	tm.orderStore.EXPECT().Add(gomock.Any()).Times(1).Do(func(order types.Order) {
		assert.EqualValues(t, order.Id, orderID)
		assert.EqualValues(t, order.Price, price-1)
		assert.EqualValues(t, order.Size, size-1)
		assert.EqualValues(t, order.Remaining, size-1)
		assert.EqualValues(t, order.Version, uint64(8))
	})
	cancelled, err := tm.market.CancelOrderByID(orderID)
	assert.NoError(t, err)
	if assert.NotNil(t, cancelled, "cancelled freshly submitted order") {
		assert.EqualValues(t, confirmation.Order.Id, cancelled.Order.Id)
	}
}
