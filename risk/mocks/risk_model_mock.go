// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/risk (interfaces: Model)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
	time "time"
)

// MockModel is a mock of Model interface
type MockModel struct {
	ctrl     *gomock.Controller
	recorder *MockModelMockRecorder
}

// MockModelMockRecorder is the mock recorder for MockModel
type MockModelMockRecorder struct {
	mock *MockModel
}

// NewMockModel creates a new mock instance
func NewMockModel(ctrl *gomock.Controller) *MockModel {
	mock := &MockModel{ctrl: ctrl}
	mock.recorder = &MockModelMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockModel) EXPECT() *MockModelMockRecorder {
	return m.recorder
}

// CalculateRiskFactors mocks base method
func (m *MockModel) CalculateRiskFactors(arg0 *proto.RiskResult) (bool, *proto.RiskResult) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculateRiskFactors", arg0)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(*proto.RiskResult)
	return ret0, ret1
}

// CalculateRiskFactors indicates an expected call of CalculateRiskFactors
func (mr *MockModelMockRecorder) CalculateRiskFactors(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculateRiskFactors", reflect.TypeOf((*MockModel)(nil).CalculateRiskFactors), arg0)
}

// CalculationInterval mocks base method
func (m *MockModel) CalculationInterval() time.Duration {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CalculationInterval")
	ret0, _ := ret[0].(time.Duration)
	return ret0
}

// CalculationInterval indicates an expected call of CalculationInterval
func (mr *MockModelMockRecorder) CalculationInterval() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CalculationInterval", reflect.TypeOf((*MockModel)(nil).CalculationInterval))
}

// PriceRange mocks base method
func (m *MockModel) PriceRange(arg0, arg1, arg2 float64) (float64, float64) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "PriceRange", arg0, arg1, arg2)
	ret0, _ := ret[0].(float64)
	ret1, _ := ret[1].(float64)
	return ret0, ret1
}

// PriceRange indicates an expected call of PriceRange
func (mr *MockModelMockRecorder) PriceRange(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "PriceRange", reflect.TypeOf((*MockModel)(nil).PriceRange), arg0, arg1, arg2)
}
