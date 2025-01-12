package crypto_test

import (
	"crypto"
	"testing"

	wcrypto "code.vegaprotocol.io/vega/wallet/crypto"
	"github.com/oasisprotocol/curve25519-voi/primitives/ed25519"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSignature(t *testing.T) {
	t.Run("create signature ed25519 success", testCreateEd25519SignatureOK)
	t.Run("create signature ed25519 fail", testCreateSignatureFailureNotAnAlgo)
	t.Run("verify success", testVerifyOK)
	t.Run("verify fail wrong message", testVerifyFailWrongMessage)
	t.Run("verify fail wrong pubkey", testVerifyFailWrongPubKey)
	t.Run("sign fail bad key length", testSignBadKeyLength)
	t.Run("verify fail bad key length", testVerifyBadKeyLength)
}

func testCreateEd25519SignatureOK(t *testing.T) {
	_, err := wcrypto.NewSignatureAlgorithm(wcrypto.Ed25519, 1)
	assert.NoError(t, err)
}

func testCreateSignatureFailureNotAnAlgo(t *testing.T) {
	_, err := wcrypto.NewSignatureAlgorithm("not an algo", 1)
	assert.ErrorIs(t, err, wcrypto.ErrUnsupportedSignatureAlgorithm)
}

func testVerifyOK(t *testing.T) {
	s, err := wcrypto.NewSignatureAlgorithm(wcrypto.Ed25519, 1)
	assert.NoError(t, err)
	pub, priv := generateKey(t)
	assert.NoError(t, err)

	message := []byte("hello world")

	sig, err := s.Sign(priv, message)
	assert.NoError(t, err)
	assert.NotEmpty(t, sig)

	ok, err := s.Verify(pub, message, sig)
	assert.NoError(t, err)
	assert.True(t, ok)
}

func testSignBadKeyLength(t *testing.T) {
	s, err := wcrypto.NewSignatureAlgorithm(wcrypto.Ed25519, 1)
	assert.NoError(t, err)
	_, priv := generateKey(t)

	assert.NoError(t, err)

	message := []byte("hello world")

	// Chop one byte off the key
	priv2, ok := priv.([]byte)
	require.True(t, ok)
	priv3 := priv2[0 : len(priv2)-1]
	sig, err := s.Sign(crypto.PrivateKey(priv3), message)
	assert.Error(t, err)
	assert.Nil(t, sig)
}

func testVerifyBadKeyLength(t *testing.T) {
	s, err := wcrypto.NewSignatureAlgorithm(wcrypto.Ed25519, 1)
	assert.NoError(t, err)
	pub, priv := generateKey(t)

	assert.NoError(t, err)

	message := []byte("hello world")

	sig, err := s.Sign(priv, message)
	assert.NoError(t, err)
	assert.NotEmpty(t, sig)

	// Chop one byte off the key
	pub2, ok := pub.([]byte)
	require.True(t, ok)
	pub3 := pub2[0 : len(pub2)-1]
	ok, err = s.Verify(crypto.PublicKey(pub3), message, sig)
	assert.Error(t, err)
	assert.False(t, ok)
}

func testVerifyFailWrongMessage(t *testing.T) {
	s, err := wcrypto.NewSignatureAlgorithm(wcrypto.Ed25519, 1)
	assert.NoError(t, err)
	pub, priv := generateKey(t)
	assert.NoError(t, err)

	message := []byte("hello world")
	wrongmessage := []byte("yolo")

	sig, err := s.Sign(priv, message)
	assert.NoError(t, err)
	assert.NotEmpty(t, sig)

	ok, err := s.Verify(pub, wrongmessage, sig)
	assert.NoError(t, err)
	assert.False(t, ok)
}

func testVerifyFailWrongPubKey(t *testing.T) {
	s, err := wcrypto.NewSignatureAlgorithm(wcrypto.Ed25519, 1)
	assert.NoError(t, err)
	// gen 2 sets of  keys
	_, priv := generateKey(t)
	assert.NoError(t, err)
	pub, _ := generateKey(t)
	assert.NoError(t, err)

	message := []byte("hello world")

	sig, err := s.Sign(priv, message)
	assert.NoError(t, err)
	assert.NotEmpty(t, sig)

	ok, err := s.Verify(pub, message, sig)
	assert.NoError(t, err)
	assert.False(t, ok)
}

func generateKey(t *testing.T) (crypto.PublicKey, crypto.PrivateKey) {
	t.Helper()
	pub, priv, err := ed25519.GenerateKey(nil)
	if err != nil {
		t.Fatalf("couldn't generate key: %v", err)
	}

	return []byte(pub), []byte(priv)
}
