// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/data-node/api (interfaces: GovernanceDataService)

// Package mocks is a generated GoMock package.
package mocks

import (
	vega "code.vegaprotocol.io/data-node/proto/vega"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockGovernanceDataService is a mock of GovernanceDataService interface
type MockGovernanceDataService struct {
	ctrl     *gomock.Controller
	recorder *MockGovernanceDataServiceMockRecorder
}

// MockGovernanceDataServiceMockRecorder is the mock recorder for MockGovernanceDataService
type MockGovernanceDataServiceMockRecorder struct {
	mock *MockGovernanceDataService
}

// NewMockGovernanceDataService creates a new mock instance
func NewMockGovernanceDataService(ctrl *gomock.Controller) *MockGovernanceDataService {
	mock := &MockGovernanceDataService{ctrl: ctrl}
	mock.recorder = &MockGovernanceDataServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockGovernanceDataService) EXPECT() *MockGovernanceDataServiceMockRecorder {
	return m.recorder
}

// GetNetworkParametersProposals mocks base method
func (m *MockGovernanceDataService) GetNetworkParametersProposals(arg0 *vega.Proposal_State) []*vega.GovernanceData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNetworkParametersProposals", arg0)
	ret0, _ := ret[0].([]*vega.GovernanceData)
	return ret0
}

// GetNetworkParametersProposals indicates an expected call of GetNetworkParametersProposals
func (mr *MockGovernanceDataServiceMockRecorder) GetNetworkParametersProposals(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNetworkParametersProposals", reflect.TypeOf((*MockGovernanceDataService)(nil).GetNetworkParametersProposals), arg0)
}

// GetNewAssetProposals mocks base method
func (m *MockGovernanceDataService) GetNewAssetProposals(arg0 *vega.Proposal_State) []*vega.GovernanceData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNewAssetProposals", arg0)
	ret0, _ := ret[0].([]*vega.GovernanceData)
	return ret0
}

// GetNewAssetProposals indicates an expected call of GetNewAssetProposals
func (mr *MockGovernanceDataServiceMockRecorder) GetNewAssetProposals(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNewAssetProposals", reflect.TypeOf((*MockGovernanceDataService)(nil).GetNewAssetProposals), arg0)
}

// GetNewMarketProposals mocks base method
func (m *MockGovernanceDataService) GetNewMarketProposals(arg0 *vega.Proposal_State) []*vega.GovernanceData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetNewMarketProposals", arg0)
	ret0, _ := ret[0].([]*vega.GovernanceData)
	return ret0
}

// GetNewMarketProposals indicates an expected call of GetNewMarketProposals
func (mr *MockGovernanceDataServiceMockRecorder) GetNewMarketProposals(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetNewMarketProposals", reflect.TypeOf((*MockGovernanceDataService)(nil).GetNewMarketProposals), arg0)
}

// GetProposalByID mocks base method
func (m *MockGovernanceDataService) GetProposalByID(arg0 string) (*vega.GovernanceData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProposalByID", arg0)
	ret0, _ := ret[0].(*vega.GovernanceData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProposalByID indicates an expected call of GetProposalByID
func (mr *MockGovernanceDataServiceMockRecorder) GetProposalByID(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProposalByID", reflect.TypeOf((*MockGovernanceDataService)(nil).GetProposalByID), arg0)
}

// GetProposalByReference mocks base method
func (m *MockGovernanceDataService) GetProposalByReference(arg0 string) (*vega.GovernanceData, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProposalByReference", arg0)
	ret0, _ := ret[0].(*vega.GovernanceData)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetProposalByReference indicates an expected call of GetProposalByReference
func (mr *MockGovernanceDataServiceMockRecorder) GetProposalByReference(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProposalByReference", reflect.TypeOf((*MockGovernanceDataService)(nil).GetProposalByReference), arg0)
}

// GetProposals mocks base method
func (m *MockGovernanceDataService) GetProposals(arg0 *vega.Proposal_State) []*vega.GovernanceData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProposals", arg0)
	ret0, _ := ret[0].([]*vega.GovernanceData)
	return ret0
}

// GetProposals indicates an expected call of GetProposals
func (mr *MockGovernanceDataServiceMockRecorder) GetProposals(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProposals", reflect.TypeOf((*MockGovernanceDataService)(nil).GetProposals), arg0)
}

// GetProposalsByParty mocks base method
func (m *MockGovernanceDataService) GetProposalsByParty(arg0 string, arg1 *vega.Proposal_State) []*vega.GovernanceData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetProposalsByParty", arg0, arg1)
	ret0, _ := ret[0].([]*vega.GovernanceData)
	return ret0
}

// GetProposalsByParty indicates an expected call of GetProposalsByParty
func (mr *MockGovernanceDataServiceMockRecorder) GetProposalsByParty(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetProposalsByParty", reflect.TypeOf((*MockGovernanceDataService)(nil).GetProposalsByParty), arg0, arg1)
}

// GetUpdateMarketProposals mocks base method
func (m *MockGovernanceDataService) GetUpdateMarketProposals(arg0 string, arg1 *vega.Proposal_State) []*vega.GovernanceData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetUpdateMarketProposals", arg0, arg1)
	ret0, _ := ret[0].([]*vega.GovernanceData)
	return ret0
}

// GetUpdateMarketProposals indicates an expected call of GetUpdateMarketProposals
func (mr *MockGovernanceDataServiceMockRecorder) GetUpdateMarketProposals(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetUpdateMarketProposals", reflect.TypeOf((*MockGovernanceDataService)(nil).GetUpdateMarketProposals), arg0, arg1)
}

// GetVotesByParty mocks base method
func (m *MockGovernanceDataService) GetVotesByParty(arg0 string) []*vega.Vote {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetVotesByParty", arg0)
	ret0, _ := ret[0].([]*vega.Vote)
	return ret0
}

// GetVotesByParty indicates an expected call of GetVotesByParty
func (mr *MockGovernanceDataServiceMockRecorder) GetVotesByParty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetVotesByParty", reflect.TypeOf((*MockGovernanceDataService)(nil).GetVotesByParty), arg0)
}

// ObserveGovernance mocks base method
func (m *MockGovernanceDataService) ObserveGovernance(arg0 context.Context, arg1 int) <-chan []vega.GovernanceData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObserveGovernance", arg0, arg1)
	ret0, _ := ret[0].(<-chan []vega.GovernanceData)
	return ret0
}

// ObserveGovernance indicates an expected call of ObserveGovernance
func (mr *MockGovernanceDataServiceMockRecorder) ObserveGovernance(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObserveGovernance", reflect.TypeOf((*MockGovernanceDataService)(nil).ObserveGovernance), arg0, arg1)
}

// ObservePartyProposals mocks base method
func (m *MockGovernanceDataService) ObservePartyProposals(arg0 context.Context, arg1 int, arg2 string) <-chan []vega.GovernanceData {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObservePartyProposals", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan []vega.GovernanceData)
	return ret0
}

// ObservePartyProposals indicates an expected call of ObservePartyProposals
func (mr *MockGovernanceDataServiceMockRecorder) ObservePartyProposals(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObservePartyProposals", reflect.TypeOf((*MockGovernanceDataService)(nil).ObservePartyProposals), arg0, arg1, arg2)
}

// ObservePartyVotes mocks base method
func (m *MockGovernanceDataService) ObservePartyVotes(arg0 context.Context, arg1 int, arg2 string) <-chan []vega.Vote {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObservePartyVotes", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan []vega.Vote)
	return ret0
}

// ObservePartyVotes indicates an expected call of ObservePartyVotes
func (mr *MockGovernanceDataServiceMockRecorder) ObservePartyVotes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObservePartyVotes", reflect.TypeOf((*MockGovernanceDataService)(nil).ObservePartyVotes), arg0, arg1, arg2)
}

// ObserveProposalVotes mocks base method
func (m *MockGovernanceDataService) ObserveProposalVotes(arg0 context.Context, arg1 int, arg2 string) <-chan []vega.Vote {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObserveProposalVotes", arg0, arg1, arg2)
	ret0, _ := ret[0].(<-chan []vega.Vote)
	return ret0
}

// ObserveProposalVotes indicates an expected call of ObserveProposalVotes
func (mr *MockGovernanceDataServiceMockRecorder) ObserveProposalVotes(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObserveProposalVotes", reflect.TypeOf((*MockGovernanceDataService)(nil).ObserveProposalVotes), arg0, arg1, arg2)
}
