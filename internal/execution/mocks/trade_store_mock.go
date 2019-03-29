// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/execution (interfaces: TradeStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTradeStore is a mock of TradeStore interface
type MockTradeStore struct {
	ctrl     *gomock.Controller
	recorder *MockTradeStoreMockRecorder
}

// MockTradeStoreMockRecorder is the mock recorder for MockTradeStore
type MockTradeStoreMockRecorder struct {
	mock *MockTradeStore
}

// NewMockTradeStore creates a new mock instance
func NewMockTradeStore(ctrl *gomock.Controller) *MockTradeStore {
	mock := &MockTradeStore{ctrl: ctrl}
	mock.recorder = &MockTradeStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTradeStore) EXPECT() *MockTradeStoreMockRecorder {
	return m.recorder
}

// Commit mocks base method
func (m *MockTradeStore) Commit() error {
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(error)
	return ret0
}

// Commit indicates an expected call of Commit
func (mr *MockTradeStoreMockRecorder) Commit() *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockTradeStore)(nil).Commit))
}

// Post mocks base method
func (m *MockTradeStore) Post(arg0 *proto.Trade) error {
	ret := m.ctrl.Call(m, "Post", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Post indicates an expected call of Post
func (mr *MockTradeStoreMockRecorder) Post(arg0 interface{}) *gomock.Call {
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Post", reflect.TypeOf((*MockTradeStore)(nil).Post), arg0)
}
