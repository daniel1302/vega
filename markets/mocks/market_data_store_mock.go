// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/markets (interfaces: MarketDataStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockMarketDataStore is a mock of MarketDataStore interface
type MockMarketDataStore struct {
	ctrl     *gomock.Controller
	recorder *MockMarketDataStoreMockRecorder
}

// MockMarketDataStoreMockRecorder is the mock recorder for MockMarketDataStore
type MockMarketDataStoreMockRecorder struct {
	mock *MockMarketDataStore
}

// NewMockMarketDataStore creates a new mock instance
func NewMockMarketDataStore(ctrl *gomock.Controller) *MockMarketDataStore {
	mock := &MockMarketDataStore{ctrl: ctrl}
	mock.recorder = &MockMarketDataStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockMarketDataStore) EXPECT() *MockMarketDataStoreMockRecorder {
	return m.recorder
}

// GetAll mocks base method
func (m *MockMarketDataStore) GetAll() []proto.MarketData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]proto.MarketData)
	return ret0
}

// GetAll indicates an expected call of GetAll
func (mr *MockMarketDataStoreMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockMarketDataStore)(nil).GetAll))
}

// GetByID mocks base method
func (m *MockMarketDataStore) GetByID(arg0 string) (proto.MarketData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(proto.MarketData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID
func (mr *MockMarketDataStoreMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockMarketDataStore)(nil).GetByID), arg0)
}

// Subscribe mocks base method
func (m *MockMarketDataStore) Subscribe(arg0 chan<- []proto.MarketData) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Subscribe indicates an expected call of Subscribe
func (mr *MockMarketDataStoreMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockMarketDataStore)(nil).Subscribe), arg0)
}

// Unsubscribe mocks base method
func (m *MockMarketDataStore) Unsubscribe(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe
func (mr *MockMarketDataStoreMockRecorder) Unsubscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockMarketDataStore)(nil).Unsubscribe), arg0)
}