// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/processor (interfaces: Banking)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	types "code.vegaprotocol.io/vega/core/types"
	num "code.vegaprotocol.io/vega/libs/num"
	gomock "github.com/golang/mock/gomock"
)

// MockBanking is a mock of Banking interface.
type MockBanking struct {
	ctrl     *gomock.Controller
	recorder *MockBankingMockRecorder
}

// MockBankingMockRecorder is the mock recorder for MockBanking.
type MockBankingMockRecorder struct {
	mock *MockBanking
}

// NewMockBanking creates a new mock instance.
func NewMockBanking(ctrl *gomock.Controller) *MockBanking {
	mock := &MockBanking{ctrl: ctrl}
	mock.recorder = &MockBankingMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockBanking) EXPECT() *MockBankingMockRecorder {
	return m.recorder
}

// CancelTransferFunds mocks base method.
func (m *MockBanking) CancelTransferFunds(arg0 context.Context, arg1 *types.CancelTransferFunds) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CancelTransferFunds", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CancelTransferFunds indicates an expected call of CancelTransferFunds.
func (mr *MockBankingMockRecorder) CancelTransferFunds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CancelTransferFunds", reflect.TypeOf((*MockBanking)(nil).CancelTransferFunds), arg0, arg1)
}

// DepositBuiltinAsset mocks base method.
func (m *MockBanking) DepositBuiltinAsset(arg0 context.Context, arg1 *types.BuiltinAssetDeposit, arg2 string, arg3 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DepositBuiltinAsset", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// DepositBuiltinAsset indicates an expected call of DepositBuiltinAsset.
func (mr *MockBankingMockRecorder) DepositBuiltinAsset(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DepositBuiltinAsset", reflect.TypeOf((*MockBanking)(nil).DepositBuiltinAsset), arg0, arg1, arg2, arg3)
}

// DepositERC20 mocks base method.
func (m *MockBanking) DepositERC20(arg0 context.Context, arg1 *types.ERC20Deposit, arg2 string, arg3, arg4 uint64, arg5 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DepositERC20", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// DepositERC20 indicates an expected call of DepositERC20.
func (mr *MockBankingMockRecorder) DepositERC20(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DepositERC20", reflect.TypeOf((*MockBanking)(nil).DepositERC20), arg0, arg1, arg2, arg3, arg4, arg5)
}

// ERC20WithdrawalEvent mocks base method.
func (m *MockBanking) ERC20WithdrawalEvent(arg0 context.Context, arg1 *types.ERC20Withdrawal, arg2, arg3 uint64, arg4 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ERC20WithdrawalEvent", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// ERC20WithdrawalEvent indicates an expected call of ERC20WithdrawalEvent.
func (mr *MockBankingMockRecorder) ERC20WithdrawalEvent(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ERC20WithdrawalEvent", reflect.TypeOf((*MockBanking)(nil).ERC20WithdrawalEvent), arg0, arg1, arg2, arg3, arg4)
}

// EnableBuiltinAsset mocks base method.
func (m *MockBanking) EnableBuiltinAsset(arg0 context.Context, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableBuiltinAsset", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableBuiltinAsset indicates an expected call of EnableBuiltinAsset.
func (mr *MockBankingMockRecorder) EnableBuiltinAsset(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableBuiltinAsset", reflect.TypeOf((*MockBanking)(nil).EnableBuiltinAsset), arg0, arg1)
}

// EnableERC20 mocks base method.
func (m *MockBanking) EnableERC20(arg0 context.Context, arg1 *types.ERC20AssetList, arg2 string, arg3, arg4 uint64, arg5 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EnableERC20", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// EnableERC20 indicates an expected call of EnableERC20.
func (mr *MockBankingMockRecorder) EnableERC20(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EnableERC20", reflect.TypeOf((*MockBanking)(nil).EnableERC20), arg0, arg1, arg2, arg3, arg4, arg5)
}

// TransferFunds mocks base method.
func (m *MockBanking) TransferFunds(arg0 context.Context, arg1 *types.TransferFunds) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TransferFunds", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// TransferFunds indicates an expected call of TransferFunds.
func (mr *MockBankingMockRecorder) TransferFunds(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TransferFunds", reflect.TypeOf((*MockBanking)(nil).TransferFunds), arg0, arg1)
}

// UpdateERC20 mocks base method.
func (m *MockBanking) UpdateERC20(arg0 context.Context, arg1 *types.ERC20AssetLimitsUpdated, arg2 string, arg3, arg4 uint64, arg5 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateERC20", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateERC20 indicates an expected call of UpdateERC20.
func (mr *MockBankingMockRecorder) UpdateERC20(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateERC20", reflect.TypeOf((*MockBanking)(nil).UpdateERC20), arg0, arg1, arg2, arg3, arg4, arg5)
}

// WithdrawBuiltinAsset mocks base method.
func (m *MockBanking) WithdrawBuiltinAsset(arg0 context.Context, arg1, arg2, arg3 string, arg4 *num.Uint) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithdrawBuiltinAsset", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithdrawBuiltinAsset indicates an expected call of WithdrawBuiltinAsset.
func (mr *MockBankingMockRecorder) WithdrawBuiltinAsset(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawBuiltinAsset", reflect.TypeOf((*MockBanking)(nil).WithdrawBuiltinAsset), arg0, arg1, arg2, arg3, arg4)
}

// WithdrawERC20 mocks base method.
func (m *MockBanking) WithdrawERC20(arg0 context.Context, arg1, arg2, arg3 string, arg4 *num.Uint, arg5 *types.Erc20WithdrawExt) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "WithdrawERC20", arg0, arg1, arg2, arg3, arg4, arg5)
	ret0, _ := ret[0].(error)
	return ret0
}

// WithdrawERC20 indicates an expected call of WithdrawERC20.
func (mr *MockBankingMockRecorder) WithdrawERC20(arg0, arg1, arg2, arg3, arg4, arg5 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "WithdrawERC20", reflect.TypeOf((*MockBanking)(nil).WithdrawERC20), arg0, arg1, arg2, arg3, arg4, arg5)
}
