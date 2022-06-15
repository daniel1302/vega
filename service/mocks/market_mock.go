// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/data-node/service (interfaces: MarketStore)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	entities "code.vegaprotocol.io/data-node/entities"
	gomock "github.com/golang/mock/gomock"
)

// MockMarketStore is a mock of MarketStore interface.
type MockMarketStore struct {
	ctrl     *gomock.Controller
	recorder *MockMarketStoreMockRecorder
}

// MockMarketStoreMockRecorder is the mock recorder for MockMarketStore.
type MockMarketStoreMockRecorder struct {
	mock *MockMarketStore
}

// NewMockMarketStore creates a new mock instance.
func NewMockMarketStore(ctrl *gomock.Controller) *MockMarketStore {
	mock := &MockMarketStore{ctrl: ctrl}
	mock.recorder = &MockMarketStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockMarketStore) EXPECT() *MockMarketStoreMockRecorder {
	return m.recorder
}

// GetAll mocks base method.
func (m *MockMarketStore) GetAll(arg0 context.Context, arg1 entities.OffsetPagination) ([]entities.Market, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAll", arg0, arg1)
	ret0, _ := ret[0].([]entities.Market)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAll indicates an expected call of GetAll.
func (mr *MockMarketStoreMockRecorder) GetAll(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAll", reflect.TypeOf((*MockMarketStore)(nil).GetAll), arg0, arg1)
}

// GetAllPaged mocks base method.
func (m *MockMarketStore) GetAllPaged(arg0 context.Context, arg1 string, arg2 entities.CursorPagination) ([]entities.Market, entities.PageInfo, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllPaged", arg0, arg1, arg2)
	ret0, _ := ret[0].([]entities.Market)
	ret1, _ := ret[1].(entities.PageInfo)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetAllPaged indicates an expected call of GetAllPaged.
func (mr *MockMarketStoreMockRecorder) GetAllPaged(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllPaged", reflect.TypeOf((*MockMarketStore)(nil).GetAllPaged), arg0, arg1, arg2)
}

// GetByID mocks base method.
func (m *MockMarketStore) GetByID(arg0 context.Context, arg1 string) (entities.Market, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", arg0, arg1)
	ret0, _ := ret[0].(entities.Market)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockMarketStoreMockRecorder) GetByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockMarketStore)(nil).GetByID), arg0, arg1)
}

// Upsert mocks base method.
func (m *MockMarketStore) Upsert(arg0 context.Context, arg1 *entities.Market) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Upsert", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// Upsert indicates an expected call of Upsert.
func (mr *MockMarketStoreMockRecorder) Upsert(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Upsert", reflect.TypeOf((*MockMarketStore)(nil).Upsert), arg0, arg1)
}
