// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/checkpoint (interfaces: CollateralState)

// Package mocks is a generated GoMock package.
package mocks

import (
	types "code.vegaprotocol.io/vega/types"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockCollateralState is a mock of CollateralState interface
type MockCollateralState struct {
	ctrl     *gomock.Controller
	recorder *MockCollateralStateMockRecorder
}

// MockCollateralStateMockRecorder is the mock recorder for MockCollateralState
type MockCollateralStateMockRecorder struct {
	mock *MockCollateralState
}

// NewMockCollateralState creates a new mock instance
func NewMockCollateralState(ctrl *gomock.Controller) *MockCollateralState {
	mock := &MockCollateralState{ctrl: ctrl}
	mock.recorder = &MockCollateralStateMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockCollateralState) EXPECT() *MockCollateralStateMockRecorder {
	return m.recorder
}

// Checkpoint mocks base method
func (m *MockCollateralState) Checkpoint() ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Checkpoint")
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Checkpoint indicates an expected call of Checkpoint
func (mr *MockCollateralStateMockRecorder) Checkpoint() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Checkpoint", reflect.TypeOf((*MockCollateralState)(nil).Checkpoint))
}

// EnableAsset mocks base method
func (m *MockCollateralState) EnableAsset(arg0 context.Context, arg1 types.Asset) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableAsset", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableAsset indicates an expected call of EnableAsset
func (mr *MockCollateralStateMockRecorder) EnableAsset(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableAsset", reflect.TypeOf((*MockCollateralState)(nil).EnableAsset), arg0, arg1)
}

// Load mocks base method
func (m *MockCollateralState) Load(arg0 []byte) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Load", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Load indicates an expected call of Load
func (mr *MockCollateralStateMockRecorder) Load(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Load", reflect.TypeOf((*MockCollateralState)(nil).Load), arg0)
}

// Name mocks base method
func (m *MockCollateralState) Name() types.CheckpointName {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Name")
	ret0, _ := ret[0].(types.CheckpointName)
	return ret0
}

// Name indicates an expected call of Name
func (mr *MockCollateralStateMockRecorder) Name() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Name", reflect.TypeOf((*MockCollateralState)(nil).Name))
}
