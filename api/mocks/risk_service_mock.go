// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/data-node/api (interfaces: RiskService)

// Package mocks is a generated GoMock package.
package mocks

import (
	vega "code.vegaprotocol.io/protos/vega"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockRiskService is a mock of RiskService interface
type MockRiskService struct {
	ctrl     *gomock.Controller
	recorder *MockRiskServiceMockRecorder
}

// MockRiskServiceMockRecorder is the mock recorder for MockRiskService
type MockRiskServiceMockRecorder struct {
	mock *MockRiskService
}

// NewMockRiskService creates a new mock instance
func NewMockRiskService(ctrl *gomock.Controller) *MockRiskService {
	mock := &MockRiskService{ctrl: ctrl}
	mock.recorder = &MockRiskServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockRiskService) EXPECT() *MockRiskServiceMockRecorder {
	return m.recorder
}

// EstimateMargin mocks base method
func (m *MockRiskService) EstimateMargin(arg0 context.Context, arg1 *vega.Order) (*vega.MarginLevels, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "EstimateMargin", arg0, arg1)
	ret0, _ := ret[0].(*vega.MarginLevels)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// EstimateMargin indicates an expected call of EstimateMargin
func (mr *MockRiskServiceMockRecorder) EstimateMargin(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "EstimateMargin", reflect.TypeOf((*MockRiskService)(nil).EstimateMargin), arg0, arg1)
}

// GetMarginLevelsByID mocks base method
func (m *MockRiskService) GetMarginLevelsByID(arg0, arg1 string) ([]vega.MarginLevels, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarginLevelsByID", arg0, arg1)
	ret0, _ := ret[0].([]vega.MarginLevels)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetMarginLevelsByID indicates an expected call of GetMarginLevelsByID
func (mr *MockRiskServiceMockRecorder) GetMarginLevelsByID(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarginLevelsByID", reflect.TypeOf((*MockRiskService)(nil).GetMarginLevelsByID), arg0, arg1)
}

// GetMarginLevelsSubscribersCount mocks base method
func (m *MockRiskService) GetMarginLevelsSubscribersCount() int32 {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetMarginLevelsSubscribersCount")
	ret0, _ := ret[0].(int32)
	return ret0
}

// GetMarginLevelsSubscribersCount indicates an expected call of GetMarginLevelsSubscribersCount
func (mr *MockRiskServiceMockRecorder) GetMarginLevelsSubscribersCount() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetMarginLevelsSubscribersCount", reflect.TypeOf((*MockRiskService)(nil).GetMarginLevelsSubscribersCount))
}

// ObserveMarginLevels mocks base method
func (m *MockRiskService) ObserveMarginLevels(arg0 context.Context, arg1 int, arg2, arg3 string) (<-chan []vega.MarginLevels, uint64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ObserveMarginLevels", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(<-chan []vega.MarginLevels)
	ret1, _ := ret[1].(uint64)
	return ret0, ret1
}

// ObserveMarginLevels indicates an expected call of ObserveMarginLevels
func (mr *MockRiskServiceMockRecorder) ObserveMarginLevels(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ObserveMarginLevels", reflect.TypeOf((*MockRiskService)(nil).ObserveMarginLevels), arg0, arg1, arg2, arg3)
}
