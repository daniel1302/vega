// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/monitor/liquidity (interfaces: AuctionState)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockAuctionState is a mock of AuctionState interface
type MockAuctionState struct {
	ctrl     *gomock.Controller
	recorder *MockAuctionStateMockRecorder
}

// MockAuctionStateMockRecorder is the mock recorder for MockAuctionState
type MockAuctionStateMockRecorder struct {
	mock *MockAuctionState
}

// NewMockAuctionState creates a new mock instance
func NewMockAuctionState(ctrl *gomock.Controller) *MockAuctionState {
	mock := &MockAuctionState{ctrl: ctrl}
	mock.recorder = &MockAuctionStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAuctionState) EXPECT() *MockAuctionStateMockRecorder {
	return m.recorder
}

// EndAuction mocks base method
func (m *MockAuctionState) EndAuction() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "EndAuction")
}

// EndAuction indicates an expected call of EndAuction
func (mr *MockAuctionStateMockRecorder) EndAuction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndAuction", reflect.TypeOf((*MockAuctionState)(nil).EndAuction))
}

// ExtendAuction mocks base method
func (m *MockAuctionState) ExtendAuction(arg0 proto.AuctionDuration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "ExtendAuction", arg0)
}

// ExtendAuction indicates an expected call of ExtendAuction
func (mr *MockAuctionStateMockRecorder) ExtendAuction(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ExtendAuction", reflect.TypeOf((*MockAuctionState)(nil).ExtendAuction), arg0)
}

// InAuction mocks base method
func (m *MockAuctionState) InAuction() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InAuction")
	ret0, _ := ret[0].(bool)
	return ret0
}

// InAuction indicates an expected call of InAuction
func (mr *MockAuctionStateMockRecorder) InAuction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InAuction", reflect.TypeOf((*MockAuctionState)(nil).InAuction))
}

// IsLiquidityAuction mocks base method
func (m *MockAuctionState) IsLiquidityAuction() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsLiquidityAuction")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsLiquidityAuction indicates an expected call of IsLiquidityAuction
func (mr *MockAuctionStateMockRecorder) IsLiquidityAuction() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsLiquidityAuction", reflect.TypeOf((*MockAuctionState)(nil).IsLiquidityAuction))
}

// StartLiquidityAuction mocks base method
func (m *MockAuctionState) StartLiquidityAuction(arg0 time.Time, arg1 *proto.AuctionDuration) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartLiquidityAuction", arg0, arg1)
}

// StartLiquidityAuction indicates an expected call of StartLiquidityAuction
func (mr *MockAuctionStateMockRecorder) StartLiquidityAuction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartLiquidityAuction", reflect.TypeOf((*MockAuctionState)(nil).StartLiquidityAuction), arg0, arg1)
}
