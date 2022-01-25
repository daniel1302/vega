// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/processor (interfaces: NetworkParameters)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockNetworkParameters is a mock of NetworkParameters interface.
type MockNetworkParameters struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkParametersMockRecorder
}

// MockNetworkParametersMockRecorder is the mock recorder for MockNetworkParameters.
type MockNetworkParametersMockRecorder struct {
	mock *MockNetworkParameters
}

// NewMockNetworkParameters creates a new mock instance.
func NewMockNetworkParameters(ctrl *gomock.Controller) *MockNetworkParameters {
	mock := &MockNetworkParameters{ctrl: ctrl}
	mock.recorder = &MockNetworkParametersMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkParameters) EXPECT() *MockNetworkParametersMockRecorder {
	return m.recorder
}

// DispatchChanges mocks base method.
func (m *MockNetworkParameters) DispatchChanges(arg0 context.Context) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "DispatchChanges", arg0)
}

// DispatchChanges indicates an expected call of DispatchChanges.
func (mr *MockNetworkParametersMockRecorder) DispatchChanges(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DispatchChanges", reflect.TypeOf((*MockNetworkParameters)(nil).DispatchChanges), arg0)
}

// Update mocks base method.
func (m *MockNetworkParameters) Update(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// Update indicates an expected call of Update.
func (mr *MockNetworkParametersMockRecorder) Update(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockNetworkParameters)(nil).Update), arg0, arg1, arg2)
}
