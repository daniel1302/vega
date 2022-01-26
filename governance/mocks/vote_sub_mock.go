// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/data-node/governance (interfaces: VoteSub)

// Package mocks is a generated GoMock package.
package mocks

import (
	reflect "reflect"

	subscribers "code.vegaprotocol.io/data-node/subscribers"
	vega "code.vegaprotocol.io/protos/vega"
	gomock "github.com/golang/mock/gomock"
)

// MockVoteSub is a mock of VoteSub interface.
type MockVoteSub struct {
	ctrl     *gomock.Controller
	recorder *MockVoteSubMockRecorder
}

// MockVoteSubMockRecorder is the mock recorder for MockVoteSub.
type MockVoteSubMockRecorder struct {
	mock *MockVoteSub
}

// NewMockVoteSub creates a new mock instance.
func NewMockVoteSub(ctrl *gomock.Controller) *MockVoteSub {
	mock := &MockVoteSub{ctrl: ctrl}
	mock.recorder = &MockVoteSubMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockVoteSub) EXPECT() *MockVoteSubMockRecorder {
	return m.recorder
}

// Filter mocks base method.
func (m *MockVoteSub) Filter(arg0 ...subscribers.VoteFilter) []*vega.Vote {
	m.ctrl.T.Helper()
	varargs := []interface{}{}
	for _, a := range arg0 {
		varargs = append(varargs, a)
	}
	ret := m.ctrl.Call(m, "Filter", varargs...)
	ret0, _ := ret[0].([]*vega.Vote)
	return ret0
}

// Filter indicates an expected call of Filter.
func (mr *MockVoteSubMockRecorder) Filter(arg0 ...interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Filter", reflect.TypeOf((*MockVoteSub)(nil).Filter), arg0...)
}
