// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/governance (interfaces: Accounts)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockAccounts is a mock of Accounts interface
type MockAccounts struct {
	ctrl     *gomock.Controller
	recorder *MockAccountsMockRecorder
}

// MockAccountsMockRecorder is the mock recorder for MockAccounts
type MockAccountsMockRecorder struct {
	mock *MockAccounts
}

// NewMockAccounts creates a new mock instance
func NewMockAccounts(ctrl *gomock.Controller) *MockAccounts {
	mock := &MockAccounts{ctrl: ctrl}
	mock.recorder = &MockAccountsMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockAccounts) EXPECT() *MockAccountsMockRecorder {
	return m.recorder
}

// GetAssetTotalSupply mocks base method
func (m *MockAccounts) GetAssetTotalSupply(arg0 string) (uint64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAssetTotalSupply", arg0)
	ret0, _ := ret[0].(uint64)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAssetTotalSupply indicates an expected call of GetAssetTotalSupply
func (mr *MockAccountsMockRecorder) GetAssetTotalSupply(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAssetTotalSupply", reflect.TypeOf((*MockAccounts)(nil).GetAssetTotalSupply), arg0)
}

// GetPartyGeneralAccount mocks base method
func (m *MockAccounts) GetPartyGeneralAccount(arg0, arg1 string) (*proto.Account, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartyGeneralAccount", arg0, arg1)
	ret0, _ := ret[0].(*proto.Account)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartyGeneralAccount indicates an expected call of GetPartyGeneralAccount
func (mr *MockAccountsMockRecorder) GetPartyGeneralAccount(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartyGeneralAccount", reflect.TypeOf((*MockAccounts)(nil).GetPartyGeneralAccount), arg0, arg1)
}