// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/data-node/delegations (interfaces: DelegationStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	vega "code.vegaprotocol.io/protos/vega"
	gomock "github.com/golang/mock/gomock"
)

// MockDelegationStore is a mock of DelegationStore interface.
type MockDelegationStore struct {
	ctrl     *gomock.Controller
	recorder *MockDelegationStoreMockRecorder
}

// MockDelegationStoreMockRecorder is the mock recorder for MockDelegationStore.
type MockDelegationStoreMockRecorder struct {
	mock *MockDelegationStore
}

// NewMockDelegationStore creates a new mock instance.
func NewMockDelegationStore(ctrl *gomock.Controller) *MockDelegationStore {
	mock := &MockDelegationStore{ctrl: ctrl}
	mock.recorder = &MockDelegationStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDelegationStore) EXPECT() *MockDelegationStoreMockRecorder {
	return m.recorder
}

// GetAllDelegations mocks base method.
func (m *MockDelegationStore) GetAllDelegations() ([]*vega.Delegation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDelegations")
	ret0, _ := ret[0].([]*vega.Delegation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDelegations indicates an expected call of GetAllDelegations.
func (mr *MockDelegationStoreMockRecorder) GetAllDelegations() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDelegations", reflect.TypeOf((*MockDelegationStore)(nil).GetAllDelegations))
}

// GetAllDelegationsOnEpoch mocks base method.
func (m *MockDelegationStore) GetAllDelegationsOnEpoch(arg0 string) ([]*vega.Delegation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllDelegationsOnEpoch", arg0)
	ret0, _ := ret[0].([]*vega.Delegation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllDelegationsOnEpoch indicates an expected call of GetAllDelegationsOnEpoch.
func (mr *MockDelegationStoreMockRecorder) GetAllDelegationsOnEpoch(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllDelegationsOnEpoch", reflect.TypeOf((*MockDelegationStore)(nil).GetAllDelegationsOnEpoch), arg0)
}

// GetNodeDelegations mocks base method.
func (m *MockDelegationStore) GetNodeDelegations(arg0 string) ([]*vega.Delegation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeDelegations", arg0)
	ret0, _ := ret[0].([]*vega.Delegation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeDelegations indicates an expected call of GetNodeDelegations.
func (mr *MockDelegationStoreMockRecorder) GetNodeDelegations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeDelegations", reflect.TypeOf((*MockDelegationStore)(nil).GetNodeDelegations), arg0)
}

// GetNodeDelegationsOnEpoch mocks base method.
func (m *MockDelegationStore) GetNodeDelegationsOnEpoch(arg0, arg1 string) ([]*vega.Delegation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNodeDelegationsOnEpoch", arg0, arg1)
	ret0, _ := ret[0].([]*vega.Delegation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetNodeDelegationsOnEpoch indicates an expected call of GetNodeDelegationsOnEpoch.
func (mr *MockDelegationStoreMockRecorder) GetNodeDelegationsOnEpoch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNodeDelegationsOnEpoch", reflect.TypeOf((*MockDelegationStore)(nil).GetNodeDelegationsOnEpoch), arg0, arg1)
}

// GetPartyDelegations mocks base method.
func (m *MockDelegationStore) GetPartyDelegations(arg0 string) ([]*vega.Delegation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartyDelegations", arg0)
	ret0, _ := ret[0].([]*vega.Delegation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartyDelegations indicates an expected call of GetPartyDelegations.
func (mr *MockDelegationStoreMockRecorder) GetPartyDelegations(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartyDelegations", reflect.TypeOf((*MockDelegationStore)(nil).GetPartyDelegations), arg0)
}

// GetPartyDelegationsOnEpoch mocks base method.
func (m *MockDelegationStore) GetPartyDelegationsOnEpoch(arg0, arg1 string) ([]*vega.Delegation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartyDelegationsOnEpoch", arg0, arg1)
	ret0, _ := ret[0].([]*vega.Delegation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartyDelegationsOnEpoch indicates an expected call of GetPartyDelegationsOnEpoch.
func (mr *MockDelegationStoreMockRecorder) GetPartyDelegationsOnEpoch(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartyDelegationsOnEpoch", reflect.TypeOf((*MockDelegationStore)(nil).GetPartyDelegationsOnEpoch), arg0, arg1)
}

// GetPartyNodeDelegations mocks base method.
func (m *MockDelegationStore) GetPartyNodeDelegations(arg0, arg1 string) ([]*vega.Delegation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartyNodeDelegations", arg0, arg1)
	ret0, _ := ret[0].([]*vega.Delegation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartyNodeDelegations indicates an expected call of GetPartyNodeDelegations.
func (mr *MockDelegationStoreMockRecorder) GetPartyNodeDelegations(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartyNodeDelegations", reflect.TypeOf((*MockDelegationStore)(nil).GetPartyNodeDelegations), arg0, arg1)
}

// GetPartyNodeDelegationsOnEpoch mocks base method.
func (m *MockDelegationStore) GetPartyNodeDelegationsOnEpoch(arg0, arg1, arg2 string) ([]*vega.Delegation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPartyNodeDelegationsOnEpoch", arg0, arg1, arg2)
	ret0, _ := ret[0].([]*vega.Delegation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPartyNodeDelegationsOnEpoch indicates an expected call of GetPartyNodeDelegationsOnEpoch.
func (mr *MockDelegationStoreMockRecorder) GetPartyNodeDelegationsOnEpoch(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPartyNodeDelegationsOnEpoch", reflect.TypeOf((*MockDelegationStore)(nil).GetPartyNodeDelegationsOnEpoch), arg0, arg1, arg2)
}

// Subscribe mocks base method.
func (m *MockDelegationStore) Subscribe(arg0 chan vega.Delegation) uint64 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Subscribe", arg0)
	ret0, _ := ret[0].(uint64)
	return ret0
}

// Subscribe indicates an expected call of Subscribe.
func (mr *MockDelegationStoreMockRecorder) Subscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Subscribe", reflect.TypeOf((*MockDelegationStore)(nil).Subscribe), arg0)
}

// Unsubscribe mocks base method.
func (m *MockDelegationStore) Unsubscribe(arg0 uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Unsubscribe", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Unsubscribe indicates an expected call of Unsubscribe.
func (mr *MockDelegationStoreMockRecorder) Unsubscribe(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Unsubscribe", reflect.TypeOf((*MockDelegationStore)(nil).Unsubscribe), arg0)
}
