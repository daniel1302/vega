// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/liquidity/target (interfaces: OpenInterestCalculator)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	types "code.vegaprotocol.io/vega/core/types"
	gomock "github.com/golang/mock/gomock"
)

// MockOpenInterestCalculator is a mock of OpenInterestCalculator interface.
type MockOpenInterestCalculator struct {
	ctrl     *gomock.Controller
	recorder *MockOpenInterestCalculatorMockRecorder
}

// MockOpenInterestCalculatorMockRecorder is the mock recorder for MockOpenInterestCalculator.
type MockOpenInterestCalculatorMockRecorder struct {
	mock *MockOpenInterestCalculator
}

// NewMockOpenInterestCalculator creates a new mock instance.
func NewMockOpenInterestCalculator(ctrl *gomock.Controller) *MockOpenInterestCalculator {
	mock := &MockOpenInterestCalculator{ctrl: ctrl}
	mock.recorder = &MockOpenInterestCalculatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockOpenInterestCalculator) EXPECT() *MockOpenInterestCalculatorMockRecorder {
	return m.recorder
}

// GetOpenInterestGivenTrades mocks base method.
func (m *MockOpenInterestCalculator) GetOpenInterestGivenTrades(arg0 []*types.Trade) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetOpenInterestGivenTrades", arg0)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// GetOpenInterestGivenTrades indicates an expected call of GetOpenInterestGivenTrades.
func (mr *MockOpenInterestCalculatorMockRecorder) GetOpenInterestGivenTrades(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetOpenInterestGivenTrades", reflect.TypeOf((*MockOpenInterestCalculator)(nil).GetOpenInterestGivenTrades), arg0)
}
