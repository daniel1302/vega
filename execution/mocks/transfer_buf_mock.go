// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/execution (interfaces: TransferBuf)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTransferBuf is a mock of TransferBuf interface
type MockTransferBuf struct {
	ctrl     *gomock.Controller
	recorder *MockTransferBufMockRecorder
}

// MockTransferBufMockRecorder is the mock recorder for MockTransferBuf
type MockTransferBufMockRecorder struct {
	mock *MockTransferBuf
}

// NewMockTransferBuf creates a new mock instance
func NewMockTransferBuf(ctrl *gomock.Controller) *MockTransferBuf {
	mock := &MockTransferBuf{ctrl: ctrl}
	mock.recorder = &MockTransferBufMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTransferBuf) EXPECT() *MockTransferBufMockRecorder {
	return m.recorder
}

// Add mocks base method
func (m *MockTransferBuf) Add(arg0 []*proto.TransferResponse) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Add", arg0)
}

// Add indicates an expected call of Add
func (mr *MockTransferBufMockRecorder) Add(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Add", reflect.TypeOf((*MockTransferBuf)(nil).Add), arg0)
}

// Flush mocks base method
func (m *MockTransferBuf) Flush() error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Flush")
	ret0, _ := ret[0].(error)
	return ret0
}

// Flush indicates an expected call of Flush
func (mr *MockTransferBufMockRecorder) Flush() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Flush", reflect.TypeOf((*MockTransferBuf)(nil).Flush))
}