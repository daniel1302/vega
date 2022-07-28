// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/datanode/sqlsubscribers (interfaces: AccountSource)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entities "code.vegaprotocol.io/vega/datanode/entities"
	gomock "github.com/golang/mock/gomock"
)

// MockAccountSource is a mock of AccountSource interface.
type MockAccountSource struct {
	ctrl     *gomock.Controller
	recorder *MockAccountSourceMockRecorder
}

// MockAccountSourceMockRecorder is the mock recorder for MockAccountSource.
type MockAccountSourceMockRecorder struct {
	mock *MockAccountSource
}

// NewMockAccountSource creates a new mock instance.
func NewMockAccountSource(ctrl *gomock.Controller) *MockAccountSource {
	mock := &MockAccountSource{ctrl: ctrl}
	mock.recorder = &MockAccountSourceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAccountSource) EXPECT() *MockAccountSourceMockRecorder {
	return m.recorder
}

// GetByID mocks base method.
func (m *MockAccountSource) GetByID(arg0 int64) (entities.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0)
	ret0, _ := ret[0].(entities.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockAccountSourceMockRecorder) GetByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockAccountSource)(nil).GetByID), arg0)
}

// Obtain mocks base method.
func (m *MockAccountSource) Obtain(arg0 context.Context, arg1 *entities.Account) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Obtain", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Obtain indicates an expected call of Obtain.
func (mr *MockAccountSourceMockRecorder) Obtain(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Obtain", reflect.TypeOf((*MockAccountSource)(nil).Obtain), arg0, arg1)
}
