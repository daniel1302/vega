// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/api/endpoints/grpc (interfaces: OrderService)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockOrderService is a mock of OrderService interface
type MockOrderService struct {
	ctrl     *gomock.Controller
	recorder *MockOrderServiceMockRecorder
}

// MockOrderServiceMockRecorder is the mock recorder for MockOrderService
type MockOrderServiceMockRecorder struct {
	mock *MockOrderService
}

// NewMockOrderService creates a new mock instance
func NewMockOrderService(ctrl *gomock.Controller) *MockOrderService {
	mock := &MockOrderService{ctrl: ctrl}
	mock.recorder = &MockOrderServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOrderService) EXPECT() *MockOrderServiceMockRecorder {
	return m.recorder
}

// AmendOrder mocks base method
func (m *MockOrderService) AmendOrder(arg0 context.Context, arg1 *proto.OrderAmendment) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AmendOrder", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AmendOrder indicates an expected call of AmendOrder
func (mr *MockOrderServiceMockRecorder) AmendOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AmendOrder", reflect.TypeOf((*MockOrderService)(nil).AmendOrder), arg0, arg1)
}

// CancelOrder mocks base method
func (m *MockOrderService) CancelOrder(arg0 context.Context, arg1 *proto.OrderCancellation) (*proto.PendingOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelOrder", arg0, arg1)
	ret0, _ := ret[0].(*proto.PendingOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelOrder indicates an expected call of CancelOrder
func (mr *MockOrderServiceMockRecorder) CancelOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOrder", reflect.TypeOf((*MockOrderService)(nil).CancelOrder), arg0, arg1)
}

// CreateOrder mocks base method
func (m *MockOrderService) CreateOrder(arg0 context.Context, arg1 *proto.OrderSubmission) (*proto.PendingOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", arg0, arg1)
	ret0, _ := ret[0].(*proto.PendingOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder
func (mr *MockOrderServiceMockRecorder) CreateOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockOrderService)(nil).CreateOrder), arg0, arg1)
}

// GetByMarket mocks base method
func (m *MockOrderService) GetByMarket(arg0 context.Context, arg1 string, arg2, arg3 uint64, arg4 bool, arg5 *bool) ([]*proto.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByMarket", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]*proto.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByMarket indicates an expected call of GetByMarket
func (mr *MockOrderServiceMockRecorder) GetByMarket(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByMarket", reflect.TypeOf((*MockOrderService)(nil).GetByMarket), arg0, arg1, arg2, arg3, arg4, arg5)
}

// GetByMarketAndId mocks base method
func (m *MockOrderService) GetByMarketAndId(arg0 context.Context, arg1, arg2 string) (*proto.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByMarketAndId", arg0, arg1, arg2)
	ret0, _ := ret[0].(*proto.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByMarketAndId indicates an expected call of GetByMarketAndId
func (mr *MockOrderServiceMockRecorder) GetByMarketAndId(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByMarketAndId", reflect.TypeOf((*MockOrderService)(nil).GetByMarketAndId), arg0, arg1, arg2)
}

// GetByParty mocks base method
func (m *MockOrderService) GetByParty(arg0 context.Context, arg1 string, arg2, arg3 uint64, arg4 bool, arg5 *bool) ([]*proto.Order, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByParty", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].([]*proto.Order)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByParty indicates an expected call of GetByParty
func (mr *MockOrderServiceMockRecorder) GetByParty(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByParty", reflect.TypeOf((*MockOrderService)(nil).GetByParty), arg0, arg1, arg2, arg3, arg4, arg5)
}
