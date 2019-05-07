// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/gateway/graphql (interfaces: TradingClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	api "code.vegaprotocol.io/vega/proto/api"
	context "context"
	gomock "github.com/golang/mock/gomock"
	grpc "google.golang.org/grpc"
	reflect "reflect"
)

// MockTradingClient is a mock of TradingClient interface
type MockTradingClient struct {
	ctrl     *gomock.Controller
	recorder *MockTradingClientMockRecorder
}

// MockTradingClientMockRecorder is the mock recorder for MockTradingClient
type MockTradingClientMockRecorder struct {
	mock *MockTradingClient
}

// NewMockTradingClient creates a new mock instance
func NewMockTradingClient(ctrl *gomock.Controller) *MockTradingClient {
	mock := &MockTradingClient{ctrl: ctrl}
	mock.recorder = &MockTradingClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTradingClient) EXPECT() *MockTradingClientMockRecorder {
	return m.recorder
}

// AmendOrder mocks base method
func (m *MockTradingClient) AmendOrder(arg0 context.Context, arg1 *proto.OrderAmendment, arg2 ...grpc.CallOption) (*api.OrderResponse, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "AmendOrder", varargs...)
	ret0, _ := ret[0].(*api.OrderResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AmendOrder indicates an expected call of AmendOrder
func (mr *MockTradingClientMockRecorder) AmendOrder(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AmendOrder", reflect.TypeOf((*MockTradingClient)(nil).AmendOrder), varargs...)
}

// CancelOrder mocks base method
func (m *MockTradingClient) CancelOrder(arg0 context.Context, arg1 *proto.OrderCancellation, arg2 ...grpc.CallOption) (*proto.PendingOrder, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "CancelOrder", varargs...)
	ret0, _ := ret[0].(*proto.PendingOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelOrder indicates an expected call of CancelOrder
func (mr *MockTradingClientMockRecorder) CancelOrder(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOrder", reflect.TypeOf((*MockTradingClient)(nil).CancelOrder), varargs...)
}

// SubmitOrder mocks base method
func (m *MockTradingClient) SubmitOrder(arg0 context.Context, arg1 *proto.OrderSubmission, arg2 ...grpc.CallOption) (*proto.PendingOrder, error) {
	m.ctrl.T.Helper()
	varargs := []interface{}{arg0, arg1}
	for _, a := range arg2 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "SubmitOrder", varargs...)
	ret0, _ := ret[0].(*proto.PendingOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitOrder indicates an expected call of SubmitOrder
func (mr *MockTradingClientMockRecorder) SubmitOrder(arg0, arg1 interface{}, arg2 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	varargs := append([]interface{}{arg0, arg1}, arg2...)
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitOrder", reflect.TypeOf((*MockTradingClient)(nil).SubmitOrder), varargs...)
}
