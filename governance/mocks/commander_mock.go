// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/governance (interfaces: Commander)

// Package mocks is a generated GoMock package.
package mocks

import (
	blockchain "code.vegaprotocol.io/vega/blockchain"
	nodewallet "code.vegaprotocol.io/vega/nodewallet"
	gomock "github.com/golang/mock/gomock"
	proto "github.com/golang/protobuf/proto"
	reflect "reflect"
)

// MockCommander is a mock of Commander interface
type MockCommander struct {
	ctrl     *gomock.Controller
	recorder *MockCommanderMockRecorder
}

// MockCommanderMockRecorder is the mock recorder for MockCommander
type MockCommanderMockRecorder struct {
	mock *MockCommander
}

// NewMockCommander creates a new mock instance
func NewMockCommander(ctrl *gomock.Controller) *MockCommander {
	mock := &MockCommander{ctrl: ctrl}
	mock.recorder = &MockCommanderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCommander) EXPECT() *MockCommanderMockRecorder {
	return m.recorder
}

// Command mocks base method
func (m *MockCommander) Command(arg0 nodewallet.Wallet, arg1 blockchain.Command, arg2 proto.Message) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Command", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Command indicates an expected call of Command
func (mr *MockCommanderMockRecorder) Command(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Command", reflect.TypeOf((*MockCommander)(nil).Command), arg0, arg1, arg2)
}
