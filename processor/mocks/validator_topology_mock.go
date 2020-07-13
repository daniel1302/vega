// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/processor (interfaces: ValidatorTopology)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockValidatorTopology is a mock of ValidatorTopology interface
type MockValidatorTopology struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorTopologyMockRecorder
}

// MockValidatorTopologyMockRecorder is the mock recorder for MockValidatorTopology
type MockValidatorTopologyMockRecorder struct {
	mock *MockValidatorTopology
}

// NewMockValidatorTopology creates a new mock instance
func NewMockValidatorTopology(ctrl *gomock.Controller) *MockValidatorTopology {
	mock := &MockValidatorTopology{ctrl: ctrl}
	mock.recorder = &MockValidatorTopologyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockValidatorTopology) EXPECT() *MockValidatorTopologyMockRecorder {
	return m.recorder
}

// AddNodeRegistration mocks base method
func (m *MockValidatorTopology) AddNodeRegistration(arg0 *proto.NodeRegistration) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddNodeRegistration", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddNodeRegistration indicates an expected call of AddNodeRegistration
func (mr *MockValidatorTopologyMockRecorder) AddNodeRegistration(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddNodeRegistration", reflect.TypeOf((*MockValidatorTopology)(nil).AddNodeRegistration), arg0)
}

// AllPubKeys mocks base method
func (m *MockValidatorTopology) AllPubKeys() [][]byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AllPubKeys")
	ret0, _ := ret[0].([][]byte)
	return ret0
}

// AllPubKeys indicates an expected call of AllPubKeys
func (mr *MockValidatorTopologyMockRecorder) AllPubKeys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AllPubKeys", reflect.TypeOf((*MockValidatorTopology)(nil).AllPubKeys))
}

// Exists mocks base method
func (m *MockValidatorTopology) Exists(arg0 []byte) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Exists", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// Exists indicates an expected call of Exists
func (mr *MockValidatorTopologyMockRecorder) Exists(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Exists", reflect.TypeOf((*MockValidatorTopology)(nil).Exists), arg0)
}

// IsValidator mocks base method
func (m *MockValidatorTopology) IsValidator() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidator")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidator indicates an expected call of IsValidator
func (mr *MockValidatorTopologyMockRecorder) IsValidator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidator", reflect.TypeOf((*MockValidatorTopology)(nil).IsValidator))
}

// Len mocks base method
func (m *MockValidatorTopology) Len() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

// Len indicates an expected call of Len
func (mr *MockValidatorTopologyMockRecorder) Len() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Len", reflect.TypeOf((*MockValidatorTopology)(nil).Len))
}

// Ready mocks base method
func (m *MockValidatorTopology) Ready() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Ready")
	ret0, _ := ret[0].(bool)
	return ret0
}

// Ready indicates an expected call of Ready
func (mr *MockValidatorTopologyMockRecorder) Ready() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Ready", reflect.TypeOf((*MockValidatorTopology)(nil).Ready))
}

// SelfChainPubKey mocks base method
func (m *MockValidatorTopology) SelfChainPubKey() []byte {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelfChainPubKey")
	ret0, _ := ret[0].([]byte)
	return ret0
}

// SelfChainPubKey indicates an expected call of SelfChainPubKey
func (mr *MockValidatorTopologyMockRecorder) SelfChainPubKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelfChainPubKey", reflect.TypeOf((*MockValidatorTopology)(nil).SelfChainPubKey))
}
