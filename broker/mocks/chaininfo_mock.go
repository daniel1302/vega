// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/data-node/broker (interfaces: ChainInfoI)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockChainInfoI is a mock of ChainInfoI interface.
type MockChainInfoI struct {
	ctrl     *gomock.Controller
	recorder *MockChainInfoIMockRecorder
}

// MockChainInfoIMockRecorder is the mock recorder for MockChainInfoI.
type MockChainInfoIMockRecorder struct {
	mock *MockChainInfoI
}

// NewMockChainInfoI creates a new mock instance.
func NewMockChainInfoI(ctrl *gomock.Controller) *MockChainInfoI {
	mock := &MockChainInfoI{ctrl: ctrl}
	mock.recorder = &MockChainInfoIMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockChainInfoI) EXPECT() *MockChainInfoIMockRecorder {
	return m.recorder
}

// GetChainID mocks base method.
func (m *MockChainInfoI) GetChainID() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetChainID")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetChainID indicates an expected call of GetChainID.
func (mr *MockChainInfoIMockRecorder) GetChainID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetChainID", reflect.TypeOf((*MockChainInfoI)(nil).GetChainID))
}

// SetChainID mocks base method.
func (m *MockChainInfoI) SetChainID(arg0 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SetChainID", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// SetChainID indicates an expected call of SetChainID.
func (mr *MockChainInfoIMockRecorder) SetChainID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SetChainID", reflect.TypeOf((*MockChainInfoI)(nil).SetChainID), arg0)
}
