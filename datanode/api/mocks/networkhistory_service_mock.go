// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/datanode/api (interfaces: NetworkHistoryService)

// Package mocks is a generated GoMock package.
package mocks

import (
	segment "code.vegaprotocol.io/vega/datanode/networkhistory/segment"
	context "context"
	gomock "github.com/golang/mock/gomock"
	io "io"
	reflect "reflect"
)

// MockNetworkHistoryService is a mock of NetworkHistoryService interface.
type MockNetworkHistoryService struct {
	ctrl     *gomock.Controller
	recorder *MockNetworkHistoryServiceMockRecorder
}

// MockNetworkHistoryServiceMockRecorder is the mock recorder for MockNetworkHistoryService.
type MockNetworkHistoryServiceMockRecorder struct {
	mock *MockNetworkHistoryService
}

// NewMockNetworkHistoryService creates a new mock instance.
func NewMockNetworkHistoryService(ctrl *gomock.Controller) *MockNetworkHistoryService {
	mock := &MockNetworkHistoryService{ctrl: ctrl}
	mock.recorder = &MockNetworkHistoryServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNetworkHistoryService) EXPECT() *MockNetworkHistoryServiceMockRecorder {
	return m.recorder
}

// CopyHistorySegmentToFile mocks base method.
func (m *MockNetworkHistoryService) CopyHistorySegmentToFile(arg0 context.Context, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CopyHistorySegmentToFile", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// CopyHistorySegmentToFile indicates an expected call of CopyHistorySegmentToFile.
func (mr *MockNetworkHistoryServiceMockRecorder) CopyHistorySegmentToFile(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CopyHistorySegmentToFile", reflect.TypeOf((*MockNetworkHistoryService)(nil).CopyHistorySegmentToFile), arg0, arg1, arg2)
}

// FetchHistorySegment mocks base method.
func (m *MockNetworkHistoryService) FetchHistorySegment(arg0 context.Context, arg1 string) (segment.Full, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FetchHistorySegment", arg0, arg1)
	ret0, _ := ret[0].(segment.Full)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FetchHistorySegment indicates an expected call of FetchHistorySegment.
func (mr *MockNetworkHistoryServiceMockRecorder) FetchHistorySegment(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FetchHistorySegment", reflect.TypeOf((*MockNetworkHistoryService)(nil).FetchHistorySegment), arg0, arg1)
}

// GetActivePeerIPAddresses mocks base method.
func (m *MockNetworkHistoryService) GetActivePeerIPAddresses() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetActivePeerIPAddresses")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetActivePeerIPAddresses indicates an expected call of GetActivePeerIPAddresses.
func (mr *MockNetworkHistoryServiceMockRecorder) GetActivePeerIPAddresses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetActivePeerIPAddresses", reflect.TypeOf((*MockNetworkHistoryService)(nil).GetActivePeerIPAddresses))
}

// GetBootstrapPeers mocks base method.
func (m *MockNetworkHistoryService) GetBootstrapPeers() []string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetBootstrapPeers")
	ret0, _ := ret[0].([]string)
	return ret0
}

// GetBootstrapPeers indicates an expected call of GetBootstrapPeers.
func (mr *MockNetworkHistoryServiceMockRecorder) GetBootstrapPeers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetBootstrapPeers", reflect.TypeOf((*MockNetworkHistoryService)(nil).GetBootstrapPeers))
}

// GetConnectedPeerAddresses mocks base method.
func (m *MockNetworkHistoryService) GetConnectedPeerAddresses() ([]string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetConnectedPeerAddresses")
	ret0, _ := ret[0].([]string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetConnectedPeerAddresses indicates an expected call of GetConnectedPeerAddresses.
func (mr *MockNetworkHistoryServiceMockRecorder) GetConnectedPeerAddresses() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetConnectedPeerAddresses", reflect.TypeOf((*MockNetworkHistoryService)(nil).GetConnectedPeerAddresses))
}

// GetHighestBlockHeightHistorySegment mocks base method.
func (m *MockNetworkHistoryService) GetHighestBlockHeightHistorySegment() (segment.Full, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHighestBlockHeightHistorySegment")
	ret0, _ := ret[0].(segment.Full)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetHighestBlockHeightHistorySegment indicates an expected call of GetHighestBlockHeightHistorySegment.
func (mr *MockNetworkHistoryServiceMockRecorder) GetHighestBlockHeightHistorySegment() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHighestBlockHeightHistorySegment", reflect.TypeOf((*MockNetworkHistoryService)(nil).GetHighestBlockHeightHistorySegment))
}

// GetHistorySegmentReader mocks base method.
func (m *MockNetworkHistoryService) GetHistorySegmentReader(arg0 context.Context, arg1 string) (io.ReadSeekCloser, int64, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetHistorySegmentReader", arg0, arg1)
	ret0, _ := ret[0].(io.ReadSeekCloser)
	ret1, _ := ret[1].(int64)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// GetHistorySegmentReader indicates an expected call of GetHistorySegmentReader.
func (mr *MockNetworkHistoryServiceMockRecorder) GetHistorySegmentReader(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetHistorySegmentReader", reflect.TypeOf((*MockNetworkHistoryService)(nil).GetHistorySegmentReader), arg0, arg1)
}

// GetIpfsAddress mocks base method.
func (m *MockNetworkHistoryService) GetIpfsAddress() (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetIpfsAddress")
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetIpfsAddress indicates an expected call of GetIpfsAddress.
func (mr *MockNetworkHistoryServiceMockRecorder) GetIpfsAddress() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetIpfsAddress", reflect.TypeOf((*MockNetworkHistoryService)(nil).GetIpfsAddress))
}

// GetSwarmKey mocks base method.
func (m *MockNetworkHistoryService) GetSwarmKey() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSwarmKey")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSwarmKey indicates an expected call of GetSwarmKey.
func (mr *MockNetworkHistoryServiceMockRecorder) GetSwarmKey() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSwarmKey", reflect.TypeOf((*MockNetworkHistoryService)(nil).GetSwarmKey))
}

// GetSwarmKeySeed mocks base method.
func (m *MockNetworkHistoryService) GetSwarmKeySeed() string {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetSwarmKeySeed")
	ret0, _ := ret[0].(string)
	return ret0
}

// GetSwarmKeySeed indicates an expected call of GetSwarmKeySeed.
func (mr *MockNetworkHistoryServiceMockRecorder) GetSwarmKeySeed() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetSwarmKeySeed", reflect.TypeOf((*MockNetworkHistoryService)(nil).GetSwarmKeySeed))
}

// ListAllHistorySegments mocks base method.
func (m *MockNetworkHistoryService) ListAllHistorySegments() (segment.Segments[segment.Full], error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListAllHistorySegments")
	ret0, _ := ret[0].(segment.Segments[segment.Full])
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListAllHistorySegments indicates an expected call of ListAllHistorySegments.
func (mr *MockNetworkHistoryServiceMockRecorder) ListAllHistorySegments() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListAllHistorySegments", reflect.TypeOf((*MockNetworkHistoryService)(nil).ListAllHistorySegments))
}
