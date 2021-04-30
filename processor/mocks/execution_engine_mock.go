// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/processor (interfaces: ExecutionEngine)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockExecutionEngine is a mock of ExecutionEngine interface
type MockExecutionEngine struct {
	ctrl     *gomock.Controller
	recorder *MockExecutionEngineMockRecorder
}

// MockExecutionEngineMockRecorder is the mock recorder for MockExecutionEngine
type MockExecutionEngineMockRecorder struct {
	mock *MockExecutionEngine
}

// NewMockExecutionEngine creates a new mock instance
func NewMockExecutionEngine(ctrl *gomock.Controller) *MockExecutionEngine {
	mock := &MockExecutionEngine{ctrl: ctrl}
	mock.recorder = &MockExecutionEngineMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockExecutionEngine) EXPECT() *MockExecutionEngineMockRecorder {
	return m.recorder
}

// AmendOrder mocks base method
func (m *MockExecutionEngine) AmendOrder(arg0 context.Context, arg1 *proto.OrderAmendment) (*proto.OrderConfirmation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AmendOrder", arg0, arg1)
	ret0, _ := ret[0].(*proto.OrderConfirmation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AmendOrder indicates an expected call of AmendOrder
func (mr *MockExecutionEngineMockRecorder) AmendOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AmendOrder", reflect.TypeOf((*MockExecutionEngine)(nil).AmendOrder), arg0, arg1)
}

// CancelOrder mocks base method
func (m *MockExecutionEngine) CancelOrder(arg0 context.Context, arg1 *proto.OrderCancellation, arg2 string) ([]*proto.OrderCancellationConfirmation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelOrder", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*proto.OrderCancellationConfirmation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelOrder indicates an expected call of CancelOrder
func (mr *MockExecutionEngineMockRecorder) CancelOrder(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOrder", reflect.TypeOf((*MockExecutionEngine)(nil).CancelOrder), arg0, arg1, arg2)
}

// Hash mocks base method
func (m *MockExecutionEngine) Hash() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Hash")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// Hash indicates an expected call of Hash
func (mr *MockExecutionEngineMockRecorder) Hash() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Hash", reflect.TypeOf((*MockExecutionEngine)(nil).Hash))
}

// RejectMarket mocks base method
func (m *MockExecutionEngine) RejectMarket(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RejectMarket", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// RejectMarket indicates an expected call of RejectMarket
func (mr *MockExecutionEngineMockRecorder) RejectMarket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RejectMarket", reflect.TypeOf((*MockExecutionEngine)(nil).RejectMarket), arg0, arg1)
}

// StartOpeningAuction mocks base method
func (m *MockExecutionEngine) StartOpeningAuction(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "StartOpeningAuction", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// StartOpeningAuction indicates an expected call of StartOpeningAuction
func (mr *MockExecutionEngineMockRecorder) StartOpeningAuction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartOpeningAuction", reflect.TypeOf((*MockExecutionEngine)(nil).StartOpeningAuction), arg0, arg1)
}

// SubmitLiquidityProvision mocks base method
func (m *MockExecutionEngine) SubmitLiquidityProvision(arg0 context.Context, arg1 *proto.LiquidityProvisionSubmission, arg2, arg3 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitLiquidityProvision", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitLiquidityProvision indicates an expected call of SubmitLiquidityProvision
func (mr *MockExecutionEngineMockRecorder) SubmitLiquidityProvision(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitLiquidityProvision", reflect.TypeOf((*MockExecutionEngine)(nil).SubmitLiquidityProvision), arg0, arg1, arg2, arg3)
}

// SubmitMarket mocks base method
func (m *MockExecutionEngine) SubmitMarket(arg0 context.Context, arg1 *proto.Market) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitMarket", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitMarket indicates an expected call of SubmitMarket
func (mr *MockExecutionEngineMockRecorder) SubmitMarket(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitMarket", reflect.TypeOf((*MockExecutionEngine)(nil).SubmitMarket), arg0, arg1)
}

// SubmitMarketWithLiquidityProvision mocks base method
func (m *MockExecutionEngine) SubmitMarketWithLiquidityProvision(arg0 context.Context, arg1 *proto.Market, arg2 *proto.LiquidityProvisionSubmission, arg3, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitMarketWithLiquidityProvision", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// SubmitMarketWithLiquidityProvision indicates an expected call of SubmitMarketWithLiquidityProvision
func (mr *MockExecutionEngineMockRecorder) SubmitMarketWithLiquidityProvision(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitMarketWithLiquidityProvision", reflect.TypeOf((*MockExecutionEngine)(nil).SubmitMarketWithLiquidityProvision), arg0, arg1, arg2, arg3, arg4)
}

// SubmitOrder mocks base method
func (m *MockExecutionEngine) SubmitOrder(arg0 context.Context, arg1 *proto.Order) (*proto.OrderConfirmation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitOrder", arg0, arg1)
	ret0, _ := ret[0].(*proto.OrderConfirmation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitOrder indicates an expected call of SubmitOrder
func (mr *MockExecutionEngineMockRecorder) SubmitOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitOrder", reflect.TypeOf((*MockExecutionEngine)(nil).SubmitOrder), arg0, arg1)
}
