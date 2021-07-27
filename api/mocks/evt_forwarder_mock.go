// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/api (interfaces: EvtForwarder)

// Package mocks is a generated GoMock package.
package mocks

import (
	v1 "code.vegaprotocol.io/protos/vega/commands/v1"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockEvtForwarder is a mock of EvtForwarder interface
type MockEvtForwarder struct {
	ctrl     *gomock.Controller
	recorder *MockEvtForwarderMockRecorder
}

// MockEvtForwarderMockRecorder is the mock recorder for MockEvtForwarder
type MockEvtForwarderMockRecorder struct {
	mock *MockEvtForwarder
}

// NewMockEvtForwarder creates a new mock instance
func NewMockEvtForwarder(ctrl *gomock.Controller) *MockEvtForwarder {
	mock := &MockEvtForwarder{ctrl: ctrl}
	mock.recorder = &MockEvtForwarderMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockEvtForwarder) EXPECT() *MockEvtForwarderMockRecorder {
	return m.recorder
}

// Forward mocks base method
func (m *MockEvtForwarder) Forward(arg0 context.Context, arg1 *v1.ChainEvent, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Forward", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Forward indicates an expected call of Forward
func (mr *MockEvtForwarderMockRecorder) Forward(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Forward", reflect.TypeOf((*MockEvtForwarder)(nil).Forward), arg0, arg1, arg2)
}
