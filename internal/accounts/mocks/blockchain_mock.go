// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/accounts (interfaces: Blockchain)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockBlockchain is a mock of Blockchain interface
type MockBlockchain struct {
	ctrl     *gomock.Controller
	recorder *MockBlockchainMockRecorder
}

// MockBlockchainMockRecorder is the mock recorder for MockBlockchain
type MockBlockchainMockRecorder struct {
	mock *MockBlockchain
}

// NewMockBlockchain creates a new mock instance
func NewMockBlockchain(ctrl *gomock.Controller) *MockBlockchain {
	mock := &MockBlockchain{ctrl: ctrl}
	mock.recorder = &MockBlockchainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockBlockchain) EXPECT() *MockBlockchainMockRecorder {
	return m.recorder
}

// NotifyTraderAccount mocks base method
func (m *MockBlockchain) NotifyTraderAccount(arg0 context.Context, arg1 *proto.NotifyTraderAccount) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NotifyTraderAccount", arg0, arg1)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NotifyTraderAccount indicates an expected call of NotifyTraderAccount
func (mr *MockBlockchainMockRecorder) NotifyTraderAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NotifyTraderAccount", reflect.TypeOf((*MockBlockchain)(nil).NotifyTraderAccount), arg0, arg1)
}