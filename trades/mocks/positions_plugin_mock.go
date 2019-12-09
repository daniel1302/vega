// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/trades (interfaces: PositionsPlugin)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockPositionsPlugin is a mock of PositionsPlugin interface
type MockPositionsPlugin struct {
	ctrl     *gomock.Controller
	recorder *MockPositionsPluginMockRecorder
}

// MockPositionsPluginMockRecorder is the mock recorder for MockPositionsPlugin
type MockPositionsPluginMockRecorder struct {
	mock *MockPositionsPlugin
}

// NewMockPositionsPlugin creates a new mock instance
func NewMockPositionsPlugin(ctrl *gomock.Controller) *MockPositionsPlugin {
	mock := &MockPositionsPlugin{ctrl: ctrl}
	mock.recorder = &MockPositionsPluginMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockPositionsPlugin) EXPECT() *MockPositionsPluginMockRecorder {
	return m.recorder
}

// GetPositionsByMarket mocks base method
func (m *MockPositionsPlugin) GetPositionsByMarket(arg0 string) ([]*proto.Position, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPositionsByMarket", arg0)
	ret0, _ := ret[0].([]*proto.Position)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPositionsByMarket indicates an expected call of GetPositionsByMarket
func (mr *MockPositionsPluginMockRecorder) GetPositionsByMarket(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPositionsByMarket", reflect.TypeOf((*MockPositionsPlugin)(nil).GetPositionsByMarket), arg0)
}

// GetPositionsByMarketAndParty mocks base method
func (m *MockPositionsPlugin) GetPositionsByMarketAndParty(arg0, arg1 string) (*proto.Position, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPositionsByMarketAndParty", arg0, arg1)
	ret0, _ := ret[0].(*proto.Position)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPositionsByMarketAndParty indicates an expected call of GetPositionsByMarketAndParty
func (mr *MockPositionsPluginMockRecorder) GetPositionsByMarketAndParty(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPositionsByMarketAndParty", reflect.TypeOf((*MockPositionsPlugin)(nil).GetPositionsByMarketAndParty), arg0, arg1)
}

// GetPositionsByParty mocks base method
func (m *MockPositionsPlugin) GetPositionsByParty(arg0 string) ([]*proto.Position, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPositionsByParty", arg0)
	ret0, _ := ret[0].([]*proto.Position)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPositionsByParty indicates an expected call of GetPositionsByParty
func (mr *MockPositionsPluginMockRecorder) GetPositionsByParty(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPositionsByParty", reflect.TypeOf((*MockPositionsPlugin)(nil).GetPositionsByParty), arg0)
}
