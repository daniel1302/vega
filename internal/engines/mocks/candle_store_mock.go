// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/execution (interfaces: CandleStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCandleStore is a mock of CandleStore interface
type MockCandleStore struct {
	ctrl     *gomock.Controller
	recorder *MockCandleStoreMockRecorder
}

// MockCandleStoreMockRecorder is the mock recorder for MockCandleStore
type MockCandleStoreMockRecorder struct {
	mock *MockCandleStore
}

// NewMockCandleStore creates a new mock instance
func NewMockCandleStore(ctrl *gomock.Controller) *MockCandleStore {
	mock := &MockCandleStore{ctrl: ctrl}
	mock.recorder = &MockCandleStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCandleStore) EXPECT() *MockCandleStoreMockRecorder {
	return m.recorder
}

// FetchMostRecentCandle mocks base method
func (m *MockCandleStore) FetchMostRecentCandle(arg0 string, arg1 proto.Interval, arg2 bool) (*proto.Candle, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchMostRecentCandle", arg0, arg1, arg2)
	ret0, _ := ret[0].(*proto.Candle)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchMostRecentCandle indicates an expected call of FetchMostRecentCandle
func (mr *MockCandleStoreMockRecorder) FetchMostRecentCandle(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchMostRecentCandle", reflect.TypeOf((*MockCandleStore)(nil).FetchMostRecentCandle), arg0, arg1, arg2)
}

// GenerateCandlesFromBuffer mocks base method
func (m *MockCandleStore) GenerateCandlesFromBuffer(arg0 string, arg1 map[string]proto.Candle) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateCandlesFromBuffer", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// GenerateCandlesFromBuffer indicates an expected call of GenerateCandlesFromBuffer
func (mr *MockCandleStoreMockRecorder) GenerateCandlesFromBuffer(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateCandlesFromBuffer", reflect.TypeOf((*MockCandleStore)(nil).GenerateCandlesFromBuffer), arg0, arg1)
}