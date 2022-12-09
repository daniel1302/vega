// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/datanode/api (interfaces: DeHistoryService)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	store "code.vegaprotocol.io/vega/datanode/dehistory/store"
	gomock "github.com/golang/mock/gomock"
)

// MockDeHistoryService is a mock of DeHistoryService interface.
type MockDeHistoryService struct {
	ctrl     *gomock.Controller
	recorder *MockDeHistoryServiceMockRecorder
}

// MockDeHistoryServiceMockRecorder is the mock recorder for MockDeHistoryService.
type MockDeHistoryServiceMockRecorder struct {
	mock *MockDeHistoryService
}

// NewMockDeHistoryService creates a new mock instance.
func NewMockDeHistoryService(ctrl *gomock.Controller) *MockDeHistoryService {
	mock := &MockDeHistoryService{ctrl: ctrl}
	mock.recorder = &MockDeHistoryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDeHistoryService) EXPECT() *MockDeHistoryServiceMockRecorder {
	return m.recorder
}

// CopyHistorySegmentToFile mocks base method.
func (m *MockDeHistoryService) CopyHistorySegmentToFile(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyHistorySegmentToFile", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CopyHistorySegmentToFile indicates an expected call of CopyHistorySegmentToFile.
func (mr *MockDeHistoryServiceMockRecorder) CopyHistorySegmentToFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyHistorySegmentToFile", reflect.TypeOf((*MockDeHistoryService)(nil).CopyHistorySegmentToFile), arg0, arg1, arg2)
}

// FetchHistorySegment mocks base method.
func (m *MockDeHistoryService) FetchHistorySegment(arg0 context.Context, arg1 string) (store.SegmentIndexEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchHistorySegment", arg0, arg1)
	ret0, _ := ret[0].(store.SegmentIndexEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchHistorySegment indicates an expected call of FetchHistorySegment.
func (mr *MockDeHistoryServiceMockRecorder) FetchHistorySegment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchHistorySegment", reflect.TypeOf((*MockDeHistoryService)(nil).FetchHistorySegment), arg0, arg1)
}

// GetActivePeerAddresses mocks base method.
func (m *MockDeHistoryService) GetActivePeerAddresses() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActivePeerAddresses")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetActivePeerAddresses indicates an expected call of GetActivePeerAddresses.
func (mr *MockDeHistoryServiceMockRecorder) GetActivePeerAddresses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActivePeerAddresses", reflect.TypeOf((*MockDeHistoryService)(nil).GetActivePeerAddresses))
}

// GetHighestBlockHeightHistorySegment mocks base method.
func (m *MockDeHistoryService) GetHighestBlockHeightHistorySegment() (store.SegmentIndexEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHighestBlockHeightHistorySegment")
	ret0, _ := ret[0].(store.SegmentIndexEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHighestBlockHeightHistorySegment indicates an expected call of GetHighestBlockHeightHistorySegment.
func (mr *MockDeHistoryServiceMockRecorder) GetHighestBlockHeightHistorySegment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHighestBlockHeightHistorySegment", reflect.TypeOf((*MockDeHistoryService)(nil).GetHighestBlockHeightHistorySegment))
}

// GetSwarmKey mocks base method.
func (m *MockDeHistoryService) GetSwarmKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSwarmKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSwarmKey indicates an expected call of GetSwarmKey.
func (mr *MockDeHistoryServiceMockRecorder) GetSwarmKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSwarmKey", reflect.TypeOf((*MockDeHistoryService)(nil).GetSwarmKey))
}

// ListAllHistorySegments mocks base method.
func (m *MockDeHistoryService) ListAllHistorySegments() ([]store.SegmentIndexEntry, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllHistorySegments")
	ret0, _ := ret[0].([]store.SegmentIndexEntry)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllHistorySegments indicates an expected call of ListAllHistorySegments.
func (mr *MockDeHistoryServiceMockRecorder) ListAllHistorySegments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllHistorySegments", reflect.TypeOf((*MockDeHistoryService)(nil).ListAllHistorySegments))
}
