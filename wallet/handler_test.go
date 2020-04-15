package wallet_test

import (
	"encoding/base64"
	"errors"
	"os"
	"testing"

	"code.vegaprotocol.io/vega/fsutil"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/wallet"
	"code.vegaprotocol.io/vega/wallet/crypto"
	"code.vegaprotocol.io/vega/wallet/mocks"

	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

type testHandler struct {
	*wallet.Handler
	ctrl    *gomock.Controller
	auth    *mocks.MockAuth
	rootDir string
}

func getTestHandler(t *testing.T) *testHandler {
	ctrl := gomock.NewController(t)
	auth := mocks.NewMockAuth(ctrl)
	rootPath := rootDir()
	fsutil.EnsureDir(rootPath)
	wallet.EnsureBaseFolder(rootPath)

	h := wallet.NewHandler(logging.NewTestLogger(), auth, rootPath)
	return &testHandler{
		Handler: h,
		ctrl:    ctrl,
		auth:    auth,
		rootDir: rootPath,
	}
}

func TestHandler(t *testing.T) {
	t.Run("create a wallet success then login", testHandlerCreateWalletThenLogin)
	t.Run("create a wallet failure - already exists", testHandlerCreateWalletFailureAlreadyExists)
	t.Run("login failure on non wallet", testHandlerLoginFailureOnNonCreatedWallet)
	t.Run("revoke token success", testHandlerRevokeTokenSuccess)
	t.Run("revoke token failure", testHandlerRevokeTokenFailure)
	t.Run("generate keypair success and list public keys", testVerifyTokenSuccess)
	t.Run("generate keypair failure - invalid token", testVerifyTokenInvalidToken)
	t.Run("generate keypair failure - wallet not found", testVerifyTokenWalletNotFound)
	t.Run("list public key failure - invalid token", testListPubInvalidToken)
	t.Run("list public key failure - wallet not found", testListPubWalletNotFound)
	t.Run("sign tx - success", testSignTxSuccess)
	t.Run("sign tx - failure key tainted", testSignTxFailure)
	t.Run("taint key - success", testTaintKeySuccess)
	t.Run("taint key failure - invalid token", testTaintKeyInvalidToken)
	t.Run("taint key failure - wallet not found", testTaintKeyWalletNotFound)
	t.Run("taint key failure - already tainted", testTaintKeyAlreadyFailAlreadyTainted)
	t.Run("update meta failure - pub key does not exists", testTaintKeyPubKeyDoesNotExists)
	t.Run("update meta - success", testUpdateMetaSuccess)
	t.Run("update meta failure - invalid token", testUpdateMetaInvalidToken)
	t.Run("update meta taint key failure - wallet not found", testUpdateMetaWalletNotFound)
	t.Run("update meta failure - pub key does not exists", testUpdateMetaPubKeyDoesNotExists)
}

func testHandlerCreateWalletThenLogin(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	h.auth.EXPECT().NewSession(gomock.Any()).Times(2).
		Return("some fake token", nil)

	tok, err := h.CreateWallet("jeremy", "thisisasecurepassphraseinnit")
	assert.NoError(t, err)
	assert.NotEmpty(t, tok)

	tok, err = h.LoginWallet("jeremy", "thisisasecurepassphraseinnit")
	assert.NoError(t, err)
	assert.NotEmpty(t, tok)

	assert.NoError(t, os.RemoveAll(h.rootDir))
}

func testHandlerCreateWalletFailureAlreadyExists(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	h.auth.EXPECT().NewSession(gomock.Any()).Times(1).
		Return("some fake token", nil)

	// create the wallet once.
	tok, err := h.CreateWallet("jeremy", "thisisasecurepassphraseinnit")
	assert.NoError(t, err)
	assert.NotEmpty(t, tok)

	// try to create it again
	tok, err = h.CreateWallet("jeremy", "we can use a different passphrase yo!")
	assert.EqualError(t, err, wallet.ErrWalletAlreadyExists.Error())
	assert.Empty(t, tok)

	assert.NoError(t, os.RemoveAll(h.rootDir))
}

func testHandlerLoginFailureOnNonCreatedWallet(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	tok, err := h.LoginWallet("jeremy", "thisisasecurepassphraseinnit")
	assert.EqualError(t, err, wallet.ErrWalletDoesNotExists.Error())
	assert.Empty(t, tok)

	assert.NoError(t, os.RemoveAll(h.rootDir))
}

func testHandlerRevokeTokenSuccess(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	h.auth.EXPECT().NewSession(gomock.Any()).Times(1).
		Return("some fake token", nil)

	tok, err := h.CreateWallet("jeremy", "thisisasecurepassphraseinnit")
	assert.NoError(t, err)
	assert.NotEmpty(t, tok)

	h.auth.EXPECT().Revoke(gomock.Any()).Times(1).
		Return(nil)
	err = h.RevokeToken(tok)
	assert.NoError(t, err)

	assert.NoError(t, os.RemoveAll(h.rootDir))
}

func testHandlerRevokeTokenFailure(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	h.auth.EXPECT().NewSession(gomock.Any()).Times(1).
		Return("some fake token", nil)

	tok, err := h.CreateWallet("jeremy", "thisisasecurepassphraseinnit")
	assert.NoError(t, err)
	assert.NotEmpty(t, tok)

	h.auth.EXPECT().Revoke(gomock.Any()).Times(1).
		Return(errors.New("bad token"))
	err = h.RevokeToken(tok)
	assert.EqualError(t, err, "bad token")

	assert.NoError(t, os.RemoveAll(h.rootDir))
}

func testVerifyTokenSuccess(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// first create the wallet
	h.auth.EXPECT().NewSession(gomock.Any()).Times(1).
		Return("some fake token", nil)

	tok, err := h.CreateWallet("jeremy", "thisisasecurepassphraseinnit")
	assert.NoError(t, err)
	assert.NotEmpty(t, tok)

	// then start the test
	h.auth.EXPECT().VerifyToken(gomock.Any()).Times(2).
		Return("jeremy", nil)

	key, err := h.GenerateKeypair(tok, "thisisasecurepassphraseinnit")
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	// now make sure we have the new key saved
	keys, err := h.ListPublicKeys(tok)
	assert.NoError(t, err)
	assert.Len(t, keys, 1)
	assert.Equal(t, key, keys[0].Pub)

	assert.NoError(t, os.RemoveAll(h.rootDir))
}

func testVerifyTokenInvalidToken(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// then start the test
	h.auth.EXPECT().VerifyToken(gomock.Any()).Times(1).
		Return("", errors.New("bad token"))

	key, err := h.GenerateKeypair("yolo token", "whatever")
	assert.EqualError(t, err, "bad token")
	assert.Empty(t, key)

	assert.NoError(t, os.RemoveAll(h.rootDir))

}

// this should never happend but beeeh....
func testVerifyTokenWalletNotFound(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// then start the test
	h.auth.EXPECT().VerifyToken(gomock.Any()).Times(1).
		Return("jeremy", nil)

	key, err := h.GenerateKeypair("yolo token", "whatever")
	assert.EqualError(t, err, wallet.ErrWalletDoesNotExists.Error())
	assert.Empty(t, key)

	assert.NoError(t, os.RemoveAll(h.rootDir))
}

func testListPubInvalidToken(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// then start the test
	h.auth.EXPECT().VerifyToken(gomock.Any()).Times(1).
		Return("", errors.New("bad token"))

	key, err := h.ListPublicKeys("yolo token")
	assert.EqualError(t, err, "bad token")
	assert.Empty(t, key)

	assert.NoError(t, os.RemoveAll(h.rootDir))

}

// this should never happend but beeeh....
func testListPubWalletNotFound(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// then start the test
	h.auth.EXPECT().VerifyToken(gomock.Any()).Times(1).
		Return("jeremy", nil)

	key, err := h.ListPublicKeys("yolo token")
	assert.EqualError(t, err, wallet.ErrWalletDoesNotExists.Error())
	assert.Empty(t, key)

	assert.NoError(t, os.RemoveAll(h.rootDir))
}

func testSignTxSuccess(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// then start the test
	h.auth.EXPECT().VerifyToken(gomock.Any()).AnyTimes().
		Return("jeremy", nil)

	// first create the wallet
	h.auth.EXPECT().NewSession(gomock.Any()).Times(1).
		Return("some fake token", nil)

	tok, err := h.CreateWallet("jeremy", "thisisasecurepassphraseinnit")
	assert.NoError(t, err)
	assert.NotEmpty(t, tok)

	key, err := h.GenerateKeypair(tok, "thisisasecurepassphraseinnit")
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	message := "hello world."

	signedBundle, err := h.SignTx(tok, base64.StdEncoding.EncodeToString([]byte(message)), key)
	assert.NoError(t, err)

	// verify signature then
	alg, err := crypto.NewSignatureAlgorithm(crypto.Ed25519)
	assert.NoError(t, err)

	v, err := alg.Verify(signedBundle.PubKey, []byte(message), signedBundle.Sig)
	assert.NoError(t, err)
	assert.True(t, v)

	assert.NoError(t, os.RemoveAll(h.rootDir))
}

func testSignTxFailure(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// then start the test
	h.auth.EXPECT().VerifyToken(gomock.Any()).AnyTimes().
		Return("jeremy", nil)

	// first create the wallet
	h.auth.EXPECT().NewSession(gomock.Any()).Times(1).
		Return("some fake token", nil)

	tok, err := h.CreateWallet("jeremy", "thisisasecurepassphraseinnit")
	assert.NoError(t, err)
	assert.NotEmpty(t, tok)

	key, err := h.GenerateKeypair(tok, "thisisasecurepassphraseinnit")
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	// taint the key
	err = h.TaintKey(tok, key, "thisisasecurepassphraseinnit")
	assert.NoError(t, err)

	message := "hello world."
	_, err = h.SignTx(tok, base64.StdEncoding.EncodeToString([]byte(message)), key)
	assert.EqualError(t, err, wallet.ErrPubKeyIsTainted.Error())

	assert.NoError(t, os.RemoveAll(h.rootDir))
}

func testTaintKeySuccess(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// first create the wallet
	h.auth.EXPECT().NewSession(gomock.Any()).Times(1).
		Return("some fake token", nil)

	tok, err := h.CreateWallet("jeremy", "thisisasecurepassphraseinnit")
	assert.NoError(t, err)
	assert.NotEmpty(t, tok)

	// then start the test
	h.auth.EXPECT().VerifyToken(gomock.Any()).AnyTimes().
		Return("jeremy", nil)

	key, err := h.GenerateKeypair(tok, "thisisasecurepassphraseinnit")
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	// taint the key
	err = h.TaintKey(tok, key, "thisisasecurepassphraseinnit")
	assert.NoError(t, err)

	// now make sure we have the new key saved
	keys, err := h.ListPublicKeys(tok)
	assert.NoError(t, err)
	assert.Len(t, keys, 1)
	assert.Equal(t, key, keys[0].Pub)
	assert.True(t, keys[0].Tainted)

	assert.NoError(t, os.RemoveAll(h.rootDir))
}

func testTaintKeyInvalidToken(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// then the test
	h.auth.EXPECT().VerifyToken(gomock.Any()).AnyTimes().
		Return("", errors.New("invalid token"))

	// taint the key
	err := h.TaintKey("some token", "some key", "thisisasecurepassphraseinnit")
	assert.EqualError(t, err, "invalid token")

	assert.NoError(t, os.RemoveAll(h.rootDir))

}
func testTaintKeyPubKeyDoesNotExists(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// first create the wallet
	h.auth.EXPECT().NewSession(gomock.Any()).Times(1).
		Return("some fake token", nil)

	tok, err := h.CreateWallet("jeremy", "thisisasecurepassphraseinnit")
	assert.NoError(t, err)
	assert.NotEmpty(t, tok)

	// then start the test
	h.auth.EXPECT().VerifyToken(gomock.Any()).AnyTimes().
		Return("jeremy", nil)

	// taint the key
	err = h.TaintKey(tok, "some key", "thisisasecurepassphraseinnit")
	assert.EqualError(t, err, wallet.ErrPubKeyDoesNotExists.Error())

	assert.NoError(t, os.RemoveAll(h.rootDir))
}

func testTaintKeyWalletNotFound(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// then start the test
	h.auth.EXPECT().VerifyToken(gomock.Any()).AnyTimes().
		Return("jeremy", nil)

	// taint the key
	err := h.TaintKey("some token", "some key", "thisisasecurepassphraseinnit")
	assert.EqualError(t, err, wallet.ErrWalletDoesNotExists.Error())

	assert.NoError(t, os.RemoveAll(h.rootDir))
}

func testTaintKeyAlreadyFailAlreadyTainted(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// first create the wallet
	h.auth.EXPECT().NewSession(gomock.Any()).Times(1).
		Return("some fake token", nil)

	tok, err := h.CreateWallet("jeremy", "thisisasecurepassphraseinnit")
	assert.NoError(t, err)
	assert.NotEmpty(t, tok)

	// then start the test
	h.auth.EXPECT().VerifyToken(gomock.Any()).AnyTimes().
		Return("jeremy", nil)

	key, err := h.GenerateKeypair(tok, "thisisasecurepassphraseinnit")
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	// taint the key
	err = h.TaintKey(tok, key, "thisisasecurepassphraseinnit")
	assert.NoError(t, err)

	// taint the key again which produce an error
	err = h.TaintKey(tok, key, "thisisasecurepassphraseinnit")
	assert.Error(t, err, wallet.ErrPubKeyAlreadyTainted)

	assert.NoError(t, os.RemoveAll(h.rootDir))
}

func testUpdateMetaSuccess(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// first create the wallet
	h.auth.EXPECT().NewSession(gomock.Any()).Times(1).
		Return("some fake token", nil)

	tok, err := h.CreateWallet("jeremy", "thisisasecurepassphraseinnit")
	assert.NoError(t, err)
	assert.NotEmpty(t, tok)

	// then start the test
	h.auth.EXPECT().VerifyToken(gomock.Any()).AnyTimes().
		Return("jeremy", nil)

	key, err := h.GenerateKeypair(tok, "thisisasecurepassphraseinnit")
	assert.NoError(t, err)
	assert.NotEmpty(t, key)

	// add meta
	err = h.UpdateMeta(tok, key, "thisisasecurepassphraseinnit", []wallet.Meta{wallet.Meta{Key: "primary", Value: "yes"}})
	assert.NoError(t, err)

	// now make sure we have the new key saved
	keys, err := h.ListPublicKeys(tok)
	assert.NoError(t, err)
	assert.Len(t, keys, 1)
	assert.Equal(t, key, keys[0].Pub)
	assert.Len(t, keys[0].Meta, 1)
	assert.Equal(t, keys[0].Meta[0].Key, "primary")
	assert.Equal(t, keys[0].Meta[0].Value, "yes")

	assert.NoError(t, os.RemoveAll(h.rootDir))
}

func testUpdateMetaInvalidToken(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// then the test
	h.auth.EXPECT().VerifyToken(gomock.Any()).AnyTimes().
		Return("", errors.New("invalid token"))

	// taint the key
	err := h.UpdateMeta("some token", "some key", "thisisasecurepassphraseinnit", []wallet.Meta{})
	assert.EqualError(t, err, "invalid token")

	assert.NoError(t, os.RemoveAll(h.rootDir))

}

func testUpdateMetaPubKeyDoesNotExists(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// first create the wallet
	h.auth.EXPECT().NewSession(gomock.Any()).Times(1).
		Return("some fake token", nil)

	tok, err := h.CreateWallet("jeremy", "thisisasecurepassphraseinnit")
	assert.NoError(t, err)
	assert.NotEmpty(t, tok)

	// then start the test
	h.auth.EXPECT().VerifyToken(gomock.Any()).AnyTimes().
		Return("jeremy", nil)

	// update meta
	err = h.UpdateMeta(tok, "some key", "thisisasecurepassphraseinnit", []wallet.Meta{})
	assert.EqualError(t, err, wallet.ErrPubKeyDoesNotExists.Error())

	assert.NoError(t, os.RemoveAll(h.rootDir))
}

func testUpdateMetaWalletNotFound(t *testing.T) {
	h := getTestHandler(t)
	defer h.ctrl.Finish()

	// then start the test
	h.auth.EXPECT().VerifyToken(gomock.Any()).AnyTimes().
		Return("jeremy", nil)

	// taint the key
	err := h.UpdateMeta("some token", "some key", "thisisasecurepassphraseinnit", []wallet.Meta{})
	assert.EqualError(t, err, wallet.ErrWalletDoesNotExists.Error())

	assert.NoError(t, os.RemoveAll(h.rootDir))
}