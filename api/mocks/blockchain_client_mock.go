// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/api (interfaces: BlockchainClient)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	context "context"
	gomock "github.com/golang/mock/gomock"
	types "github.com/tendermint/tendermint/rpc/core/types"
	reflect "reflect"
	time "time"
)

// MockBlockchainClient is a mock of BlockchainClient interface
type MockBlockchainClient struct {
	ctrl     *gomock.Controller
	recorder *MockBlockchainClientMockRecorder
}

// MockBlockchainClientMockRecorder is the mock recorder for MockBlockchainClient
type MockBlockchainClientMockRecorder struct {
	mock *MockBlockchainClient
}

// NewMockBlockchainClient creates a new mock instance
func NewMockBlockchainClient(ctrl *gomock.Controller) *MockBlockchainClient {
	mock := &MockBlockchainClient{ctrl: ctrl}
	mock.recorder = &MockBlockchainClientMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBlockchainClient) EXPECT() *MockBlockchainClientMockRecorder {
	return m.recorder
}

// AmendOrder mocks base method
func (m *MockBlockchainClient) AmendOrder(arg0 context.Context, arg1 *proto.OrderAmendment) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AmendOrder", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// AmendOrder indicates an expected call of AmendOrder
func (mr *MockBlockchainClientMockRecorder) AmendOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AmendOrder", reflect.TypeOf((*MockBlockchainClient)(nil).AmendOrder), arg0, arg1)
}

// CancelOrder mocks base method
func (m *MockBlockchainClient) CancelOrder(arg0 context.Context, arg1 *proto.OrderCancellation) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelOrder", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CancelOrder indicates an expected call of CancelOrder
func (mr *MockBlockchainClientMockRecorder) CancelOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelOrder", reflect.TypeOf((*MockBlockchainClient)(nil).CancelOrder), arg0, arg1)
}

// CreateOrder mocks base method
func (m *MockBlockchainClient) CreateOrder(arg0 context.Context, arg1 *proto.Order) (*proto.PendingOrder, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateOrder", arg0, arg1)
	ret0, _ := ret[0].(*proto.PendingOrder)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateOrder indicates an expected call of CreateOrder
func (mr *MockBlockchainClientMockRecorder) CreateOrder(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateOrder", reflect.TypeOf((*MockBlockchainClient)(nil).CreateOrder), arg0, arg1)
}

// GetGenesisTime mocks base method
func (m *MockBlockchainClient) GetGenesisTime(arg0 context.Context) (time.Time, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetGenesisTime", arg0)
	ret0, _ := ret[0].(time.Time)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetGenesisTime indicates an expected call of GetGenesisTime
func (mr *MockBlockchainClientMockRecorder) GetGenesisTime(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetGenesisTime", reflect.TypeOf((*MockBlockchainClient)(nil).GetGenesisTime), arg0)
}

// GetNetworkInfo mocks base method
func (m *MockBlockchainClient) GetNetworkInfo(arg0 context.Context) (*types.ResultNetInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkInfo", arg0)
	ret0, _ := ret[0].(*types.ResultNetInfo)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNetworkInfo indicates an expected call of GetNetworkInfo
func (mr *MockBlockchainClientMockRecorder) GetNetworkInfo(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkInfo", reflect.TypeOf((*MockBlockchainClient)(nil).GetNetworkInfo), arg0)
}

// GetStatus mocks base method
func (m *MockBlockchainClient) GetStatus(arg0 context.Context) (*types.ResultStatus, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetStatus", arg0)
	ret0, _ := ret[0].(*types.ResultStatus)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetStatus indicates an expected call of GetStatus
func (mr *MockBlockchainClientMockRecorder) GetStatus(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetStatus", reflect.TypeOf((*MockBlockchainClient)(nil).GetStatus), arg0)
}

// GetUnconfirmedTxCount mocks base method
func (m *MockBlockchainClient) GetUnconfirmedTxCount(arg0 context.Context) (int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUnconfirmedTxCount", arg0)
	ret0, _ := ret[0].(int)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetUnconfirmedTxCount indicates an expected call of GetUnconfirmedTxCount
func (mr *MockBlockchainClientMockRecorder) GetUnconfirmedTxCount(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUnconfirmedTxCount", reflect.TypeOf((*MockBlockchainClient)(nil).GetUnconfirmedTxCount), arg0)
}

// Health mocks base method
func (m *MockBlockchainClient) Health() (*types.ResultHealth, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Health")
	ret0, _ := ret[0].(*types.ResultHealth)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Health indicates an expected call of Health
func (mr *MockBlockchainClientMockRecorder) Health() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Health", reflect.TypeOf((*MockBlockchainClient)(nil).Health))
}

// NotifyTraderAccount mocks base method
func (m *MockBlockchainClient) NotifyTraderAccount(arg0 context.Context, arg1 *proto.NotifyTraderAccount) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyTraderAccount", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotifyTraderAccount indicates an expected call of NotifyTraderAccount
func (mr *MockBlockchainClientMockRecorder) NotifyTraderAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyTraderAccount", reflect.TypeOf((*MockBlockchainClient)(nil).NotifyTraderAccount), arg0, arg1)
}

// SubmitTransaction mocks base method
func (m *MockBlockchainClient) SubmitTransaction(arg0 context.Context, arg1 *proto.SignedBundle) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitTransaction", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitTransaction indicates an expected call of SubmitTransaction
func (mr *MockBlockchainClientMockRecorder) SubmitTransaction(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitTransaction", reflect.TypeOf((*MockBlockchainClient)(nil).SubmitTransaction), arg0, arg1)
}

// Withdraw mocks base method
func (m *MockBlockchainClient) Withdraw(arg0 context.Context, arg1 *proto.Withdraw) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Withdraw", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Withdraw indicates an expected call of Withdraw
func (mr *MockBlockchainClientMockRecorder) Withdraw(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Withdraw", reflect.TypeOf((*MockBlockchainClient)(nil).Withdraw), arg0, arg1)
}
