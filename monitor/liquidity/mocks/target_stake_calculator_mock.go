// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/monitor/liquidity (interfaces: TargetStakeCalculator)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	types "code.vegaprotocol.io/vega/types"
	num "code.vegaprotocol.io/vega/types/num"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockTargetStakeCalculator is a mock of TargetStakeCalculator interface
type MockTargetStakeCalculator struct {
	ctrl     *gomock.Controller
	recorder *MockTargetStakeCalculatorMockRecorder
}

// MockTargetStakeCalculatorMockRecorder is the mock recorder for MockTargetStakeCalculator
type MockTargetStakeCalculatorMockRecorder struct {
	mock *MockTargetStakeCalculator
}

// NewMockTargetStakeCalculator creates a new mock instance
func NewMockTargetStakeCalculator(ctrl *gomock.Controller) *MockTargetStakeCalculator {
	mock := &MockTargetStakeCalculator{ctrl: ctrl}
	mock.recorder = &MockTargetStakeCalculatorMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTargetStakeCalculator) EXPECT() *MockTargetStakeCalculatorMockRecorder {
	return m.recorder
}

// GetTheoreticalTargetStake mocks base method
func (m *MockTargetStakeCalculator) GetTheoreticalTargetStake(arg0 proto.RiskFactor, arg1 time.Time, arg2 *num.Uint, arg3 []*types.Trade) float64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetTheoreticalTargetStake", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(float64)
	return ret0
}

// GetTheoreticalTargetStake indicates an expected call of GetTheoreticalTargetStake
func (mr *MockTargetStakeCalculatorMockRecorder) GetTheoreticalTargetStake(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetTheoreticalTargetStake", reflect.TypeOf((*MockTargetStakeCalculator)(nil).GetTheoreticalTargetStake), arg0, arg1, arg2, arg3)
}
