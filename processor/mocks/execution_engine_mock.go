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
func (m *MockExecutionEngine) CancelOrder(arg0 context.Context, arg1 *proto.OrderCancellation) (*proto.OrderCancellationConfirmation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelOrder", arg0, arg1)
	ret0, _ := ret[0].(*proto.OrderCancellationConfirmation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelOrder indicates an expected call of CancelOrder
func (mr *MockExecutionEngineMockRecorder) CancelOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOrder", reflect.TypeOf((*MockExecutionEngine)(nil).CancelOrder), arg0, arg1)
}

// Generate mocks base method
func (m *MockExecutionEngine) Generate() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Generate")
	ret0, _ := ret[0].(error)
	return ret0
}

// Generate indicates an expected call of Generate
func (mr *MockExecutionEngineMockRecorder) Generate() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Generate", reflect.TypeOf((*MockExecutionEngine)(nil).Generate))
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
