// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/datanode/api (interfaces: NetParamsService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	vega "code.vegaprotocol.io/protos/vega"
	gomock "github.com/golang/mock/gomock"
)

// MockNetParamsService is a mock of NetParamsService interface.
type MockNetParamsService struct {
	ctrl     *gomock.Controller
	recorder *MockNetParamsServiceMockRecorder
}

// MockNetParamsServiceMockRecorder is the mock recorder for MockNetParamsService.
type MockNetParamsServiceMockRecorder struct {
	mock *MockNetParamsService
}

// NewMockNetParamsService creates a new mock instance.
func NewMockNetParamsService(ctrl *gomock.Controller) *MockNetParamsService {
	mock := &MockNetParamsService{ctrl: ctrl}
	mock.recorder = &MockNetParamsServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetParamsService) EXPECT() *MockNetParamsServiceMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockNetParamsService) GetAll() []vega.NetworkParameter {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll")
	ret0, _ := ret[0].([]vega.NetworkParameter)
	return ret0
}

// GetAll indicates an expected call of GetAll.
func (mr *MockNetParamsServiceMockRecorder) GetAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockNetParamsService)(nil).GetAll))
}
