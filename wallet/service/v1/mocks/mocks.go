// Code generated by MockGen. DO NOT EDIT.
// Source: code.vegaprotocol.io/vega/wallet/service/v1 (interfaces: WalletHandler,Auth,NodeForward,RSAStore,SpamHandler)

// Package mocks is a generated GoMock package.
package mocks

import (
	context "context"
	reflect "reflect"

	v1 "code.vegaprotocol.io/vega/protos/vega/api/v1"
	v10 "code.vegaprotocol.io/vega/protos/vega/commands/v1"
	v11 "code.vegaprotocol.io/vega/protos/vega/wallet/v1"
	types "code.vegaprotocol.io/vega/wallet/api/node/types"
	v12 "code.vegaprotocol.io/vega/wallet/service/v1"
	wallet "code.vegaprotocol.io/vega/wallet/wallet"
	gomock "github.com/golang/mock/gomock"
)

// MockWalletHandler is a mock of WalletHandler interface.
type MockWalletHandler struct {
	ctrl     *gomock.Controller
	recorder *MockWalletHandlerMockRecorder
}

// MockWalletHandlerMockRecorder is the mock recorder for MockWalletHandler.
type MockWalletHandlerMockRecorder struct {
	mock *MockWalletHandler
}

// NewMockWalletHandler creates a new mock instance.
func NewMockWalletHandler(ctrl *gomock.Controller) *MockWalletHandler {
	mock := &MockWalletHandler{ctrl: ctrl}
	mock.recorder = &MockWalletHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockWalletHandler) EXPECT() *MockWalletHandlerMockRecorder {
	return m.recorder
}

// CreateWallet mocks base method.
func (m *MockWalletHandler) CreateWallet(arg0, arg1 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CreateWallet", arg0, arg1)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CreateWallet indicates an expected call of CreateWallet.
func (mr *MockWalletHandlerMockRecorder) CreateWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateWallet", reflect.TypeOf((*MockWalletHandler)(nil).CreateWallet), arg0, arg1)
}

// GetPublicKey mocks base method.
func (m *MockWalletHandler) GetPublicKey(arg0, arg1 string) (wallet.PublicKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetPublicKey", arg0, arg1)
	ret0, _ := ret[0].(wallet.PublicKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetPublicKey indicates an expected call of GetPublicKey.
func (mr *MockWalletHandlerMockRecorder) GetPublicKey(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetPublicKey", reflect.TypeOf((*MockWalletHandler)(nil).GetPublicKey), arg0, arg1)
}

// ImportWallet mocks base method.
func (m *MockWalletHandler) ImportWallet(arg0, arg1, arg2 string, arg3 uint32) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ImportWallet", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// ImportWallet indicates an expected call of ImportWallet.
func (mr *MockWalletHandlerMockRecorder) ImportWallet(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ImportWallet", reflect.TypeOf((*MockWalletHandler)(nil).ImportWallet), arg0, arg1, arg2, arg3)
}

// ListPublicKeys mocks base method.
func (m *MockWalletHandler) ListPublicKeys(arg0 string) ([]wallet.PublicKey, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListPublicKeys", arg0)
	ret0, _ := ret[0].([]wallet.PublicKey)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListPublicKeys indicates an expected call of ListPublicKeys.
func (mr *MockWalletHandlerMockRecorder) ListPublicKeys(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListPublicKeys", reflect.TypeOf((*MockWalletHandler)(nil).ListPublicKeys), arg0)
}

// LoginWallet mocks base method.
func (m *MockWalletHandler) LoginWallet(arg0, arg1 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LoginWallet", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// LoginWallet indicates an expected call of LoginWallet.
func (mr *MockWalletHandlerMockRecorder) LoginWallet(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginWallet", reflect.TypeOf((*MockWalletHandler)(nil).LoginWallet), arg0, arg1)
}

// SecureGenerateKeyPair mocks base method.
func (m *MockWalletHandler) SecureGenerateKeyPair(arg0, arg1 string, arg2 []wallet.Metadata) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SecureGenerateKeyPair", arg0, arg1, arg2)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SecureGenerateKeyPair indicates an expected call of SecureGenerateKeyPair.
func (mr *MockWalletHandlerMockRecorder) SecureGenerateKeyPair(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SecureGenerateKeyPair", reflect.TypeOf((*MockWalletHandler)(nil).SecureGenerateKeyPair), arg0, arg1, arg2)
}

// SignAny mocks base method.
func (m *MockWalletHandler) SignAny(arg0 string, arg1 []byte, arg2 string) ([]byte, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignAny", arg0, arg1, arg2)
	ret0, _ := ret[0].([]byte)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignAny indicates an expected call of SignAny.
func (mr *MockWalletHandlerMockRecorder) SignAny(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignAny", reflect.TypeOf((*MockWalletHandler)(nil).SignAny), arg0, arg1, arg2)
}

// SignTx mocks base method.
func (m *MockWalletHandler) SignTx(arg0 string, arg1 *v11.SubmitTransactionRequest, arg2 uint64, arg3 string) (*v10.Transaction, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SignTx", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v10.Transaction)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SignTx indicates an expected call of SignTx.
func (mr *MockWalletHandlerMockRecorder) SignTx(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SignTx", reflect.TypeOf((*MockWalletHandler)(nil).SignTx), arg0, arg1, arg2, arg3)
}

// TaintKey mocks base method.
func (m *MockWalletHandler) TaintKey(arg0, arg1, arg2 string) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "TaintKey", arg0, arg1, arg2)
	ret0, _ := ret[0].(error)
	return ret0
}

// TaintKey indicates an expected call of TaintKey.
func (mr *MockWalletHandlerMockRecorder) TaintKey(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "TaintKey", reflect.TypeOf((*MockWalletHandler)(nil).TaintKey), arg0, arg1, arg2)
}

// UpdateMeta mocks base method.
func (m *MockWalletHandler) UpdateMeta(arg0, arg1, arg2 string, arg3 []wallet.Metadata) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "UpdateMeta", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(error)
	return ret0
}

// UpdateMeta indicates an expected call of UpdateMeta.
func (mr *MockWalletHandlerMockRecorder) UpdateMeta(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateMeta", reflect.TypeOf((*MockWalletHandler)(nil).UpdateMeta), arg0, arg1, arg2, arg3)
}

// VerifyAny mocks base method.
func (m *MockWalletHandler) VerifyAny(arg0, arg1 []byte, arg2 string) (bool, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyAny", arg0, arg1, arg2)
	ret0, _ := ret[0].(bool)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyAny indicates an expected call of VerifyAny.
func (mr *MockWalletHandlerMockRecorder) VerifyAny(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyAny", reflect.TypeOf((*MockWalletHandler)(nil).VerifyAny), arg0, arg1, arg2)
}

// MockAuth is a mock of Auth interface.
type MockAuth struct {
	ctrl     *gomock.Controller
	recorder *MockAuthMockRecorder
}

// MockAuthMockRecorder is the mock recorder for MockAuth.
type MockAuthMockRecorder struct {
	mock *MockAuth
}

// NewMockAuth creates a new mock instance.
func NewMockAuth(ctrl *gomock.Controller) *MockAuth {
	mock := &MockAuth{ctrl: ctrl}
	mock.recorder = &MockAuthMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAuth) EXPECT() *MockAuthMockRecorder {
	return m.recorder
}

// NewSession mocks base method.
func (m *MockAuth) NewSession(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "NewSession", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// NewSession indicates an expected call of NewSession.
func (mr *MockAuthMockRecorder) NewSession(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "NewSession", reflect.TypeOf((*MockAuth)(nil).NewSession), arg0)
}

// Revoke mocks base method.
func (m *MockAuth) Revoke(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Revoke", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Revoke indicates an expected call of Revoke.
func (mr *MockAuthMockRecorder) Revoke(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Revoke", reflect.TypeOf((*MockAuth)(nil).Revoke), arg0)
}

// RevokeAllToken mocks base method.
func (m *MockAuth) RevokeAllToken() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "RevokeAllToken")
}

// RevokeAllToken indicates an expected call of RevokeAllToken.
func (mr *MockAuthMockRecorder) RevokeAllToken() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RevokeAllToken", reflect.TypeOf((*MockAuth)(nil).RevokeAllToken))
}

// VerifyToken mocks base method.
func (m *MockAuth) VerifyToken(arg0 string) (string, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "VerifyToken", arg0)
	ret0, _ := ret[0].(string)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// VerifyToken indicates an expected call of VerifyToken.
func (mr *MockAuthMockRecorder) VerifyToken(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "VerifyToken", reflect.TypeOf((*MockAuth)(nil).VerifyToken), arg0)
}

// MockNodeForward is a mock of NodeForward interface.
type MockNodeForward struct {
	ctrl     *gomock.Controller
	recorder *MockNodeForwardMockRecorder
}

// MockNodeForwardMockRecorder is the mock recorder for MockNodeForward.
type MockNodeForwardMockRecorder struct {
	mock *MockNodeForward
}

// NewMockNodeForward creates a new mock instance.
func NewMockNodeForward(ctrl *gomock.Controller) *MockNodeForward {
	mock := &MockNodeForward{ctrl: ctrl}
	mock.recorder = &MockNodeForwardMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockNodeForward) EXPECT() *MockNodeForwardMockRecorder {
	return m.recorder
}

// CheckTx mocks base method.
func (m *MockNodeForward) CheckTx(arg0 context.Context, arg1 *v10.Transaction, arg2 int) (*v1.CheckTransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckTx", arg0, arg1, arg2)
	ret0, _ := ret[0].(*v1.CheckTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// CheckTx indicates an expected call of CheckTx.
func (mr *MockNodeForwardMockRecorder) CheckTx(arg0, arg1, arg2 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckTx", reflect.TypeOf((*MockNodeForward)(nil).CheckTx), arg0, arg1, arg2)
}

// HealthCheck mocks base method.
func (m *MockNodeForward) HealthCheck(arg0 context.Context) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "HealthCheck", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// HealthCheck indicates an expected call of HealthCheck.
func (mr *MockNodeForwardMockRecorder) HealthCheck(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "HealthCheck", reflect.TypeOf((*MockNodeForward)(nil).HealthCheck), arg0)
}

// LastBlockHeightAndHash mocks base method.
func (m *MockNodeForward) LastBlockHeightAndHash(arg0 context.Context) (*v1.LastBlockHeightResponse, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "LastBlockHeightAndHash", arg0)
	ret0, _ := ret[0].(*v1.LastBlockHeightResponse)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// LastBlockHeightAndHash indicates an expected call of LastBlockHeightAndHash.
func (mr *MockNodeForwardMockRecorder) LastBlockHeightAndHash(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LastBlockHeightAndHash", reflect.TypeOf((*MockNodeForward)(nil).LastBlockHeightAndHash), arg0)
}

// SendTx mocks base method.
func (m *MockNodeForward) SendTx(arg0 context.Context, arg1 *v10.Transaction, arg2 v1.SubmitTransactionRequest_Type, arg3 int) (*v1.SubmitTransactionResponse, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SendTx", arg0, arg1, arg2, arg3)
	ret0, _ := ret[0].(*v1.SubmitTransactionResponse)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// SendTx indicates an expected call of SendTx.
func (mr *MockNodeForwardMockRecorder) SendTx(arg0, arg1, arg2, arg3 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SendTx", reflect.TypeOf((*MockNodeForward)(nil).SendTx), arg0, arg1, arg2, arg3)
}

// SpamStatistics mocks base method.
func (m *MockNodeForward) SpamStatistics(arg0 context.Context, arg1 string) (*v1.GetSpamStatisticsResponse, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "SpamStatistics", arg0, arg1)
	ret0, _ := ret[0].(*v1.GetSpamStatisticsResponse)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// SpamStatistics indicates an expected call of SpamStatistics.
func (mr *MockNodeForwardMockRecorder) SpamStatistics(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "SpamStatistics", reflect.TypeOf((*MockNodeForward)(nil).SpamStatistics), arg0, arg1)
}

// Stop mocks base method.
func (m *MockNodeForward) Stop() {
	m.ctrl.T.Helper()
	m.ctrl.Call(m, "Stop")
}

// Stop indicates an expected call of Stop.
func (mr *MockNodeForwardMockRecorder) Stop() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Stop", reflect.TypeOf((*MockNodeForward)(nil).Stop))
}

// MockRSAStore is a mock of RSAStore interface.
type MockRSAStore struct {
	ctrl     *gomock.Controller
	recorder *MockRSAStoreMockRecorder
}

// MockRSAStoreMockRecorder is the mock recorder for MockRSAStore.
type MockRSAStoreMockRecorder struct {
	mock *MockRSAStore
}

// NewMockRSAStore creates a new mock instance.
func NewMockRSAStore(ctrl *gomock.Controller) *MockRSAStore {
	mock := &MockRSAStore{ctrl: ctrl}
	mock.recorder = &MockRSAStoreMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRSAStore) EXPECT() *MockRSAStoreMockRecorder {
	return m.recorder
}

// GetRsaKeys mocks base method.
func (m *MockRSAStore) GetRsaKeys() (*v12.RSAKeys, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetRsaKeys")
	ret0, _ := ret[0].(*v12.RSAKeys)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetRsaKeys indicates an expected call of GetRsaKeys.
func (mr *MockRSAStoreMockRecorder) GetRsaKeys() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetRsaKeys", reflect.TypeOf((*MockRSAStore)(nil).GetRsaKeys))
}

// MockSpamHandler is a mock of SpamHandler interface.
type MockSpamHandler struct {
	ctrl     *gomock.Controller
	recorder *MockSpamHandlerMockRecorder
}

// MockSpamHandlerMockRecorder is the mock recorder for MockSpamHandler.
type MockSpamHandlerMockRecorder struct {
	mock *MockSpamHandler
}

// NewMockSpamHandler creates a new mock instance.
func NewMockSpamHandler(ctrl *gomock.Controller) *MockSpamHandler {
	mock := &MockSpamHandler{ctrl: ctrl}
	mock.recorder = &MockSpamHandlerMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockSpamHandler) EXPECT() *MockSpamHandlerMockRecorder {
	return m.recorder
}

// CheckSubmission mocks base method.
func (m *MockSpamHandler) CheckSubmission(arg0 *v11.SubmitTransactionRequest, arg1 *types.SpamStatistics) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "CheckSubmission", arg0, arg1)
	ret0, _ := ret[0].(error)
	return ret0
}

// CheckSubmission indicates an expected call of CheckSubmission.
func (mr *MockSpamHandlerMockRecorder) CheckSubmission(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CheckSubmission", reflect.TypeOf((*MockSpamHandler)(nil).CheckSubmission), arg0, arg1)
}

// GenerateProofOfWork mocks base method.
func (m *MockSpamHandler) GenerateProofOfWork(arg0 string, arg1 *types.SpamStatistics) (*v10.ProofOfWork, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GenerateProofOfWork", arg0, arg1)
	ret0, _ := ret[0].(*v10.ProofOfWork)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GenerateProofOfWork indicates an expected call of GenerateProofOfWork.
func (mr *MockSpamHandlerMockRecorder) GenerateProofOfWork(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GenerateProofOfWork", reflect.TypeOf((*MockSpamHandler)(nil).GenerateProofOfWork), arg0, arg1)
}