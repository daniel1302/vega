// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/processor (interfaces: OracleAdaptors)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	v1 "code.vegaprotocol.io/protos/vega/commands/v1"
	crypto "code.vegaprotocol.io/vega/crypto"
	oracles "code.vegaprotocol.io/vega/oracles"
	gomock "github.com/golang/mock/gomock"
)

// MockOracleAdaptors is a mock of OracleAdaptors interface
type MockOracleAdaptors struct {
	ctrl     *gomock.Controller
	recorder *MockOracleAdaptorsMockRecorder
}

// MockOracleAdaptorsMockRecorder is the mock recorder for MockOracleAdaptors
type MockOracleAdaptorsMockRecorder struct {
	mock *MockOracleAdaptors
}

// NewMockOracleAdaptors creates a new mock instance
func NewMockOracleAdaptors(ctrl *gomock.Controller) *MockOracleAdaptors {
	mock := &MockOracleAdaptors{ctrl: ctrl}
	mock.recorder = &MockOracleAdaptorsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockOracleAdaptors) EXPECT() *MockOracleAdaptorsMockRecorder {
	return m.recorder
}

// Normalise mocks base method
func (m *MockOracleAdaptors) Normalise(arg0 crypto.PublicKey, arg1 v1.OracleDataSubmission) (*oracles.OracleData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Normalise", arg0, arg1)
	ret0, _ := ret[0].(*oracles.OracleData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Normalise indicates an expected call of Normalise
func (mr *MockOracleAdaptorsMockRecorder) Normalise(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Normalise", reflect.TypeOf((*MockOracleAdaptors)(nil).Normalise), arg0, arg1)
}
