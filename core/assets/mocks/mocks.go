// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/core/assets (interfaces: ERC20BridgeView,Notary)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	types "code.vegaprotocol.io/vega/core/types"
	v1 "code.vegaprotocol.io/vega/protos/vega/commands/v1"
	gomock "github.com/golang/mock/gomock"
)

// MockERC20BridgeView is a mock of ERC20BridgeView interface.
type MockERC20BridgeView struct {
	ctrl     *gomock.Controller
	recorder *MockERC20BridgeViewMockRecorder
}

// MockERC20BridgeViewMockRecorder is the mock recorder for MockERC20BridgeView.
type MockERC20BridgeViewMockRecorder struct {
	mock *MockERC20BridgeView
}

// NewMockERC20BridgeView creates a new mock instance.
func NewMockERC20BridgeView(ctrl *gomock.Controller) *MockERC20BridgeView {
	mock := &MockERC20BridgeView{ctrl: ctrl}
	mock.recorder = &MockERC20BridgeViewMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockERC20BridgeView) EXPECT() *MockERC20BridgeViewMockRecorder {
	return m.recorder
}

// FindAsset mocks base method.
func (m *MockERC20BridgeView) FindAsset(arg0 *types.AssetDetails) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAsset", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// FindAsset indicates an expected call of FindAsset.
func (mr *MockERC20BridgeViewMockRecorder) FindAsset(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAsset", reflect.TypeOf((*MockERC20BridgeView)(nil).FindAsset), arg0)
}

// MockNotary is a mock of Notary interface.
type MockNotary struct {
	ctrl     *gomock.Controller
	recorder *MockNotaryMockRecorder
}

// MockNotaryMockRecorder is the mock recorder for MockNotary.
type MockNotaryMockRecorder struct {
	mock *MockNotary
}

// NewMockNotary creates a new mock instance.
func NewMockNotary(ctrl *gomock.Controller) *MockNotary {
	mock := &MockNotary{ctrl: ctrl}
	mock.recorder = &MockNotaryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNotary) EXPECT() *MockNotaryMockRecorder {
	return m.recorder
}

// OfferSignatures mocks base method.
func (m *MockNotary) OfferSignatures(arg0 v1.NodeSignatureKind, arg1 func(string) []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "OfferSignatures", arg0, arg1)
}

// OfferSignatures indicates an expected call of OfferSignatures.
func (mr *MockNotaryMockRecorder) OfferSignatures(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "OfferSignatures", reflect.TypeOf((*MockNotary)(nil).OfferSignatures), arg0, arg1)
}

// StartAggregate mocks base method.
func (m *MockNotary) StartAggregate(arg0 string, arg1 v1.NodeSignatureKind, arg2 []byte) {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "StartAggregate", arg0, arg1, arg2)
}

// StartAggregate indicates an expected call of StartAggregate.
func (mr *MockNotaryMockRecorder) StartAggregate(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "StartAggregate", reflect.TypeOf((*MockNotary)(nil).StartAggregate), arg0, arg1, arg2)
}
