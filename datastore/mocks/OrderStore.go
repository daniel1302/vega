// Code generated by mockery v1.0.0. DO NOT EDIT.
package mocks

import datastore "vega/datastore"
import mock "github.com/stretchr/testify/mock"
import msg "vega/msg"

// OrderStore is an autogenerated mock type for the OrderStore type
type OrderStore struct {
	mock.Mock
}

// Delete provides a mock function with given fields: r
func (_m *OrderStore) Delete(r datastore.Order) error {
	ret := _m.Called(r)

	var r0 error
	if rf, ok := ret.Get(0).(func(datastore.Order) error); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// GetByMarket provides a mock function with given fields: market, params
func (_m *OrderStore) GetByMarket(market string, params datastore.GetParams) ([]datastore.Order, error) {
	ret := _m.Called(market, params)

	var r0 []datastore.Order
	if rf, ok := ret.Get(0).(func(string, datastore.GetParams) []datastore.Order); ok {
		r0 = rf(market, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]datastore.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, datastore.GetParams) error); ok {
		r1 = rf(market, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByMarketAndId provides a mock function with given fields: market, id
func (_m *OrderStore) GetByMarketAndId(market string, id string) (datastore.Order, error) {
	ret := _m.Called(market, id)

	var r0 datastore.Order
	if rf, ok := ret.Get(0).(func(string, string) datastore.Order); ok {
		r0 = rf(market, id)
	} else {
		r0 = ret.Get(0).(datastore.Order)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(market, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByParty provides a mock function with given fields: party, params
func (_m *OrderStore) GetByParty(party string, params datastore.GetParams) ([]datastore.Order, error) {
	ret := _m.Called(party, params)

	var r0 []datastore.Order
	if rf, ok := ret.Get(0).(func(string, datastore.GetParams) []datastore.Order); ok {
		r0 = rf(party, params)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]datastore.Order)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, datastore.GetParams) error); ok {
		r1 = rf(party, params)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetByPartyAndId provides a mock function with given fields: party, id
func (_m *OrderStore) GetByPartyAndId(party string, id string) (datastore.Order, error) {
	ret := _m.Called(party, id)

	var r0 datastore.Order
	if rf, ok := ret.Get(0).(func(string, string) datastore.Order); ok {
		r0 = rf(party, id)
	} else {
		r0 = ret.Get(0).(datastore.Order)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string, string) error); ok {
		r1 = rf(party, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetMarkets provides a mock function with given fields:
func (_m *OrderStore) GetMarkets() ([]string, error) {
	ret := _m.Called()

	var r0 []string
	if rf, ok := ret.Get(0).(func() []string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]string)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetOrderBookDepth provides a mock function with given fields: market
func (_m *OrderStore) GetOrderBookDepth(market string) (*msg.OrderBookDepth, error) {
	ret := _m.Called(market)

	var r0 *msg.OrderBookDepth
	if rf, ok := ret.Get(0).(func(string) *msg.OrderBookDepth); ok {
		r0 = rf(market)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*msg.OrderBookDepth)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(string) error); ok {
		r1 = rf(market)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Post provides a mock function with given fields: r
func (_m *OrderStore) Post(r datastore.Order) error {
	ret := _m.Called(r)

	var r0 error
	if rf, ok := ret.Get(0).(func(datastore.Order) error); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Put provides a mock function with given fields: r
func (_m *OrderStore) Put(r datastore.Order) error {
	ret := _m.Called(r)

	var r0 error
	if rf, ok := ret.Get(0).(func(datastore.Order) error); ok {
		r0 = rf(r)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}
