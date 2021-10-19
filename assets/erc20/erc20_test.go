package erc20_test

import (
	"context"
	"encoding/hex"
	"math/big"
	"testing"

	"code.vegaprotocol.io/vega/assets/erc20"
	vcrypto "code.vegaprotocol.io/vega/crypto"
	"code.vegaprotocol.io/vega/types"
	"code.vegaprotocol.io/vega/types/num"

	ethnw "code.vegaprotocol.io/vega/nodewallets/eth"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	ethcommon "github.com/ethereum/go-ethereum/common"
	ethtypes "github.com/ethereum/go-ethereum/core/types"
	ethcrypto "github.com/ethereum/go-ethereum/crypto"
	"github.com/stretchr/testify/assert"
	"golang.org/x/crypto/ed25519"
)

const (
	privKey       = "9feb9cbee69c1eeb30db084544ff8bf92166bf3fddefa6a021b458b4de04c66758a127387b1dff15b71fd7d0a9fd104ed75da4aac549efd5d149051ea57cefaf"
	pubKey        = "58a127387b1dff15b71fd7d0a9fd104ed75da4aac549efd5d149051ea57cefaf"
	tokenID       = "af077ace8cbf3179f826f2d3485b812f6efef07d913f2ed02f295360dd78b30e"
	ethPartyAddr  = "0x1ebe188952ab6035adad21ea1c4f64fd2eac60e1"
	bridgeAddress = "0xcB84d72e61e383767C4DFEb2d8ff7f4FB89abc6e"
)

var token = &types.AssetDetails{
	Name:        "VEGA",
	Symbol:      "VEGA",
	TotalSupply: num.NewUint(10000),
	Decimals:    18,
	MinLpStake:  num.NewUint(1),
	Source: &types.AssetDetailsErc20{
		Erc20: &types.ERC20{
			ContractAddress: "0x1FaA74E181092A97Fecc923015293ce57eE1208A",
		},
	},
}

type testERC20 struct {
	*erc20.ERC20
	wallet    *ethnw.Wallet
	ethClient testEthClient
}

func newTestERC20(t *testing.T) *testERC20 {
	t.Helper()
	wallet := ethnw.NewWallet(testWallet{})
	ethClient := testEthClient{}
	erc20Token, err := erc20.New(tokenID, token, wallet, ethClient)
	assert.NoError(t, err)

	return &testERC20{
		ERC20:     erc20Token,
		wallet:    wallet,
		ethClient: ethClient,
	}
}

func TestERC20Signatures(t *testing.T) {
	t.Run("withdraw_asset", testWithdrawAsset)
	t.Run("list_asset", testListAsset)
}

func testWithdrawAsset(t *testing.T) {
	token := newTestERC20(t)
	msg, sig, err := token.SignWithdrawal(
		num.NewUint(42),
		1000,
		ethPartyAddr,
		big.NewInt(84),
	)

	assert.NoError(t, err)
	assert.NotNil(t, msg)
	assert.NotNil(t, sig)
	assert.True(t, verifySignature(msg, sig))
	assert.Equal(t,
		"8185a0025d157b822e9a1077febc45038abef27fb2007855b75e3207536baba2d3ed69f9292f80c049c35a2ae7036bca884cdf4784f685f685af619d4326c70b",
		hex.EncodeToString(sig),
	)
}

func testListAsset(t *testing.T) {
	token := newTestERC20(t)
	msg, sig, err := token.SignBridgeListing()

	assert.NoError(t, err)
	assert.NotNil(t, msg)
	assert.NotNil(t, sig)
	assert.True(t, verifySignature(msg, sig))
	assert.Equal(t,
		"f754629dd9489307abf772831957f1da5f686e7c78ea55c71fb718062fd718fe09a217b1939c4f34ad32f214256d79c7c85dfa461efdd22d3a0a24c61e821e03",
		hex.EncodeToString(sig),
	)
}

type testEthClient struct {
	bind.ContractBackend
}

func (testEthClient) HeaderByNumber(context.Context, *big.Int) (*ethtypes.Header, error) {
	return nil, nil
}
func (testEthClient) BridgeAddress() ethcommon.Address              { return ethcommon.HexToAddress(bridgeAddress) }
func (testEthClient) CurrentHeight(context.Context) (uint64, error) { return 100, nil }
func (testEthClient) ConfirmationsRequired() uint32                 { return 1 }

type testWallet struct{}

func (testWallet) Cleanup() error { return nil }
func (testWallet) Name() string   { return "eth" }
func (testWallet) Chain() string  { return "eth" }
func (testWallet) Sign(data []byte) ([]byte, error) {
	priv, _ := hex.DecodeString(privKey)
	return ed25519.Sign(ed25519.PrivateKey(priv), data), nil
}
func (testWallet) Algo() string              { return "eth" }
func (testWallet) Version() (string, error)  { return "1", nil }
func (testWallet) PubKey() vcrypto.PublicKey { return vcrypto.PublicKey{} }

func verifySignature(msg, sig []byte) bool {
	pub, _ := hex.DecodeString(pubKey)
	hash := ethcrypto.Keccak256(msg)
	return ed25519.Verify(ed25519.PublicKey(pub), hash, sig)
}