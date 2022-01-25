// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/blockchain/nullchain (interfaces: ApplicationService)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	types "github.com/tendermint/tendermint/abci/types"
)

// MockApplicationService is a mock of ApplicationService interface.
type MockApplicationService struct {
	ctrl     *gomock.Controller
	recorder *MockApplicationServiceMockRecorder
}

// MockApplicationServiceMockRecorder is the mock recorder for MockApplicationService.
type MockApplicationServiceMockRecorder struct {
	mock *MockApplicationService
}

// NewMockApplicationService creates a new mock instance.
func NewMockApplicationService(ctrl *gomock.Controller) *MockApplicationService {
	mock := &MockApplicationService{ctrl: ctrl}
	mock.recorder = &MockApplicationServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockApplicationService) EXPECT() *MockApplicationServiceMockRecorder {
	return m.recorder
}

// BeginBlock mocks base method.
func (m *MockApplicationService) BeginBlock(arg0 types.RequestBeginBlock) types.ResponseBeginBlock {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "BeginBlock", arg0)
	ret0, _ := ret[0].(types.ResponseBeginBlock)
	return ret0
}

// BeginBlock indicates an expected call of BeginBlock.
func (mr *MockApplicationServiceMockRecorder) BeginBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "BeginBlock", reflect.TypeOf((*MockApplicationService)(nil).BeginBlock), arg0)
}

// Commit mocks base method.
func (m *MockApplicationService) Commit() types.ResponseCommit {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Commit")
	ret0, _ := ret[0].(types.ResponseCommit)
	return ret0
}

// Commit indicates an expected call of Commit.
func (mr *MockApplicationServiceMockRecorder) Commit() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Commit", reflect.TypeOf((*MockApplicationService)(nil).Commit))
}

// DeliverTx mocks base method.
func (m *MockApplicationService) DeliverTx(arg0 types.RequestDeliverTx) types.ResponseDeliverTx {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DeliverTx", arg0)
	ret0, _ := ret[0].(types.ResponseDeliverTx)
	return ret0
}

// DeliverTx indicates an expected call of DeliverTx.
func (mr *MockApplicationServiceMockRecorder) DeliverTx(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeliverTx", reflect.TypeOf((*MockApplicationService)(nil).DeliverTx), arg0)
}

// EndBlock mocks base method.
func (m *MockApplicationService) EndBlock(arg0 types.RequestEndBlock) types.ResponseEndBlock {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EndBlock", arg0)
	ret0, _ := ret[0].(types.ResponseEndBlock)
	return ret0
}

// EndBlock indicates an expected call of EndBlock.
func (mr *MockApplicationServiceMockRecorder) EndBlock(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EndBlock", reflect.TypeOf((*MockApplicationService)(nil).EndBlock), arg0)
}

// InitChain mocks base method.
func (m *MockApplicationService) InitChain(arg0 types.RequestInitChain) types.ResponseInitChain {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitChain", arg0)
	ret0, _ := ret[0].(types.ResponseInitChain)
	return ret0
}

// InitChain indicates an expected call of InitChain.
func (mr *MockApplicationServiceMockRecorder) InitChain(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitChain", reflect.TypeOf((*MockApplicationService)(nil).InitChain), arg0)
}
