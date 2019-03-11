// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import context "context"
import mock "github.com/stretchr/testify/mock"
import proto "code.vegaprotocol.io/vega/proto"
import storage "code.vegaprotocol.io/vega/internal/storage"

// CandleStore is an autogenerated mock type for the CandleStore type
type CandleStore struct {
	mock.Mock
}

// AddTradeToBuffer provides a mock function with given fields: trade
func (_m *CandleStore) AddTradeToBuffer(trade proto.Trade) error {
	ret := _m.Called(trade)

	var r0 error
	if rf, ok := ret.Get(0).(func(proto.Trade) error); ok {
		r0 = rf(trade)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Close provides a mock function with given fields:
func (_m *CandleStore) Close() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GenerateCandlesFromBuffer provides a mock function with given fields: market
func (_m *CandleStore) GenerateCandlesFromBuffer(market string) error {
	ret := _m.Called(market)

	var r0 error
	if rf, ok := ret.Get(0).(func(string) error); ok {
		r0 = rf(market)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetCandles provides a mock function with given fields: ctx, market, sinceTimestamp, interval
func (_m *CandleStore) GetCandles(ctx context.Context, market string, sinceTimestamp uint64, interval proto.Interval) ([]*proto.Candle, error) {
	ret := _m.Called(ctx, market, sinceTimestamp, interval)

	var r0 []*proto.Candle
	if rf, ok := ret.Get(0).(func(context.Context, string, uint64, proto.Interval) []*proto.Candle); ok {
		r0 = rf(ctx, market, sinceTimestamp, interval)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]*proto.Candle)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, uint64, proto.Interval) error); ok {
		r1 = rf(ctx, market, sinceTimestamp, interval)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// StartNewBuffer provides a mock function with given fields: marketId, timestamp
func (_m *CandleStore) StartNewBuffer(marketId string, timestamp uint64) error {
	ret := _m.Called(marketId, timestamp)

	var r0 error
	if rf, ok := ret.Get(0).(func(string, uint64) error); ok {
		r0 = rf(marketId, timestamp)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Subscribe provides a mock function with given fields: iT
func (_m *CandleStore) Subscribe(iT *storage.InternalTransport) uint64 {
	ret := _m.Called(iT)

	var r0 uint64
	if rf, ok := ret.Get(0).(func(*storage.InternalTransport) uint64); ok {
		r0 = rf(iT)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	return r0
}

// Unsubscribe provides a mock function with given fields: id
func (_m *CandleStore) Unsubscribe(id uint64) error {
	ret := _m.Called(id)

	var r0 error
	if rf, ok := ret.Get(0).(func(uint64) error); ok {
		r0 = rf(id)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
