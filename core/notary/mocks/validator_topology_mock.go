// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/notary (interfaces: ValidatorTopology)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockValidatorTopology is a mock of ValidatorTopology interface.
type MockValidatorTopology struct {
	ctrl     *gomock.Controller
	recorder *MockValidatorTopologyMockRecorder
}

// MockValidatorTopologyMockRecorder is the mock recorder for MockValidatorTopology.
type MockValidatorTopologyMockRecorder struct {
	mock *MockValidatorTopology
}

// NewMockValidatorTopology creates a new mock instance.
func NewMockValidatorTopology(ctrl *gomock.Controller) *MockValidatorTopology {
	mock := &MockValidatorTopology{ctrl: ctrl}
	mock.recorder = &MockValidatorTopologyMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockValidatorTopology) EXPECT() *MockValidatorTopologyMockRecorder {
	return m.recorder
}

// IsTendermintValidator mocks base method.
func (m *MockValidatorTopology) IsTendermintValidator(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsTendermintValidator", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsTendermintValidator indicates an expected call of IsTendermintValidator.
func (mr *MockValidatorTopologyMockRecorder) IsTendermintValidator(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsTendermintValidator", reflect.TypeOf((*MockValidatorTopology)(nil).IsTendermintValidator), arg0)
}

// IsValidator mocks base method.
func (m *MockValidatorTopology) IsValidator() bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidator")
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidator indicates an expected call of IsValidator.
func (mr *MockValidatorTopologyMockRecorder) IsValidator() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidator", reflect.TypeOf((*MockValidatorTopology)(nil).IsValidator))
}

// IsValidatorVegaPubKey mocks base method.
func (m *MockValidatorTopology) IsValidatorVegaPubKey(arg0 string) bool {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "IsValidatorVegaPubKey", arg0)
	ret0, _ := ret[0].(bool)
	return ret0
}

// IsValidatorVegaPubKey indicates an expected call of IsValidatorVegaPubKey.
func (mr *MockValidatorTopologyMockRecorder) IsValidatorVegaPubKey(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "IsValidatorVegaPubKey", reflect.TypeOf((*MockValidatorTopology)(nil).IsValidatorVegaPubKey), arg0)
}

// Len mocks base method.
func (m *MockValidatorTopology) Len() int {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Len")
	ret0, _ := ret[0].(int)
	return ret0
}

// Len indicates an expected call of Len.
func (mr *MockValidatorTopologyMockRecorder) Len() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Len", reflect.TypeOf((*MockValidatorTopology)(nil).Len))
}

// SelfVegaPubKey mocks base method.
func (m *MockValidatorTopology) SelfVegaPubKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SelfVegaPubKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// SelfVegaPubKey indicates an expected call of SelfVegaPubKey.
func (mr *MockValidatorTopologyMockRecorder) SelfVegaPubKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SelfVegaPubKey", reflect.TypeOf((*MockValidatorTopology)(nil).SelfVegaPubKey))
}