// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/api (interfaces: TradeOrderService)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTradeOrderService is a mock of TradeOrderService interface
type MockTradeOrderService struct {
	ctrl     *gomock.Controller
	recorder *MockTradeOrderServiceMockRecorder
}

// MockTradeOrderServiceMockRecorder is the mock recorder for MockTradeOrderService
type MockTradeOrderServiceMockRecorder struct {
	mock *MockTradeOrderService
}

// NewMockTradeOrderService creates a new mock instance
func NewMockTradeOrderService(ctrl *gomock.Controller) *MockTradeOrderService {
	mock := &MockTradeOrderService{ctrl: ctrl}
	mock.recorder = &MockTradeOrderServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTradeOrderService) EXPECT() *MockTradeOrderServiceMockRecorder {
	return m.recorder
}

// AmendOrder mocks base method
func (m *MockTradeOrderService) AmendOrder(arg0 context.Context, arg1 *proto.OrderAmendment) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AmendOrder", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AmendOrder indicates an expected call of AmendOrder
func (mr *MockTradeOrderServiceMockRecorder) AmendOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AmendOrder", reflect.TypeOf((*MockTradeOrderService)(nil).AmendOrder), arg0, arg1)
}

// CancelOrder mocks base method
func (m *MockTradeOrderService) CancelOrder(arg0 context.Context, arg1 *proto.OrderCancellation) (*proto.PendingOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelOrder", arg0, arg1)
	ret0, _ := ret[0].(*proto.PendingOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelOrder indicates an expected call of CancelOrder
func (mr *MockTradeOrderServiceMockRecorder) CancelOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOrder", reflect.TypeOf((*MockTradeOrderService)(nil).CancelOrder), arg0, arg1)
}

// CreateOrder mocks base method
func (m *MockTradeOrderService) CreateOrder(arg0 context.Context, arg1 *proto.OrderSubmission) (*proto.PendingOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", arg0, arg1)
	ret0, _ := ret[0].(*proto.PendingOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder
func (mr *MockTradeOrderServiceMockRecorder) CreateOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockTradeOrderService)(nil).CreateOrder), arg0, arg1)
}
