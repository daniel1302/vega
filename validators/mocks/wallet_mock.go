// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/validators (interfaces: Wallet)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	crypto "code.vegaprotocol.io/vega/crypto"
	gomock "github.com/golang/mock/gomock"
)

// MockWallet is a mock of Wallet interface.
type MockWallet struct {
	ctrl     *gomock.Controller
	recorder *MockWalletMockRecorder
}

// MockWalletMockRecorder is the mock recorder for MockWallet.
type MockWalletMockRecorder struct {
	mock *MockWallet
}

// NewMockWallet creates a new mock instance.
func NewMockWallet(ctrl *gomock.Controller) *MockWallet {
	mock := &MockWallet{ctrl: ctrl}
	mock.recorder = &MockWalletMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWallet) EXPECT() *MockWalletMockRecorder {
	return m.recorder
}

// ID mocks base method.
func (m *MockWallet) ID() crypto.PublicKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ID")
	ret0, _ := ret[0].(crypto.PublicKey)
	return ret0
}

// ID indicates an expected call of ID.
func (mr *MockWalletMockRecorder) ID() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ID", reflect.TypeOf((*MockWallet)(nil).ID))
}

// PubKey mocks base method.
func (m *MockWallet) PubKey() crypto.PublicKey {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PubKey")
	ret0, _ := ret[0].(crypto.PublicKey)
	return ret0
}

// PubKey indicates an expected call of PubKey.
func (mr *MockWalletMockRecorder) PubKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PubKey", reflect.TypeOf((*MockWallet)(nil).PubKey))
}
