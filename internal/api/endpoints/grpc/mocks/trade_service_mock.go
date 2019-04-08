// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/internal/api/endpoints/grpc (interfaces: TradeService)

// Package mocks is a generated GoMock package.
package mocks

import (
	proto "code.vegaprotocol.io/vega/proto"
	context "context"
	gomock "github.com/golang/mock/gomock"
	reflect "reflect"
)

// MockTradeService is a mock of TradeService interface
type MockTradeService struct {
	ctrl     *gomock.Controller
	recorder *MockTradeServiceMockRecorder
}

// MockTradeServiceMockRecorder is the mock recorder for MockTradeService
type MockTradeServiceMockRecorder struct {
	mock *MockTradeService
}

// NewMockTradeService creates a new mock instance
func NewMockTradeService(ctrl *gomock.Controller) *MockTradeService {
	mock := &MockTradeService{ctrl: ctrl}
	mock.recorder = &MockTradeServiceMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use
func (m *MockTradeService) EXPECT() *MockTradeServiceMockRecorder {
	return m.recorder
}

// GetByMarket mocks base method
func (m *MockTradeService) GetByMarket(arg0 context.Context, arg1 string, arg2, arg3 uint64, arg4 bool) ([]*proto.Trade, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByMarket", arg0, arg1, arg2, arg3, arg4)
	ret0, _ := ret[0].([]*proto.Trade)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByMarket indicates an expected call of GetByMarket
func (mr *MockTradeServiceMockRecorder) GetByMarket(arg0, arg1, arg2, arg3, arg4 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByMarket", reflect.TypeOf((*MockTradeService)(nil).GetByMarket), arg0, arg1, arg2, arg3, arg4)
}

// GetPositionsByParty mocks base method
func (m *MockTradeService) GetPositionsByParty(arg0 context.Context, arg1 string) ([]*proto.MarketPosition, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPositionsByParty", arg0, arg1)
	ret0, _ := ret[0].([]*proto.MarketPosition)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPositionsByParty indicates an expected call of GetPositionsByParty
func (mr *MockTradeServiceMockRecorder) GetPositionsByParty(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPositionsByParty", reflect.TypeOf((*MockTradeService)(nil).GetPositionsByParty), arg0, arg1)
}