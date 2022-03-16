// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/nodewallets (interfaces: Chain)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1 "code.vegaprotocol.io/protos/vega/commands/v1"
	gomock "github.com/golang/mock/gomock"
	coretypes "github.com/tendermint/tendermint/rpc/core/types"
)

// MockChain is a mock of Chain interface.
type MockChain struct {
	ctrl     *gomock.Controller
	recorder *MockChainMockRecorder
}

// MockChainMockRecorder is the mock recorder for MockChain.
type MockChainMockRecorder struct {
	mock *MockChain
}

// NewMockChain creates a new mock instance.
func NewMockChain(ctrl *gomock.Controller) *MockChain {
	mock := &MockChain{ctrl: ctrl}
	mock.recorder = &MockChainMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChain) EXPECT() *MockChainMockRecorder {
	return m.recorder
}

// SubmitTransactionAsync mocks base method.
func (m *MockChain) SubmitTransactionAsync(arg0 context.Context, arg1 *v1.Transaction) (*coretypes.ResultBroadcastTx, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SubmitTransactionAsync", arg0, arg1)
	ret0, _ := ret[0].(*coretypes.ResultBroadcastTx)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SubmitTransactionAsync indicates an expected call of SubmitTransactionAsync.
func (mr *MockChainMockRecorder) SubmitTransactionAsync(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SubmitTransactionAsync", reflect.TypeOf((*MockChain)(nil).SubmitTransactionAsync), arg0, arg1)
}
