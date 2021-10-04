package clef

import (
	"context"
	"fmt"
	"time"

	"code.vegaprotocol.io/vega/crypto"
	"github.com/ethereum/go-ethereum/accounts"
	ethcommon "github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const requestTimeout = time.Second * 5

//go:generate go run github.com/golang/mock/mockgen -destination mocks/rpc_client_mock.go -package mocks code.vegaprotocol.io/vega/nodewallets/eth/clef Client
type Client interface {
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
	Close()
}

type wallet struct {
	client   Client
	endpoint string
	name     string
	account  *accounts.Account
}

func newAccount(accountAddr ethcommon.Address, endpoint string) *accounts.Account {
	return &accounts.Account{
		URL: accounts.URL{
			Scheme: "clef",
			Path:   endpoint,
		},
		Address: accountAddr,
	}
}

func NewWallet(client Client, endpoint string, accountAddr ethcommon.Address) (*wallet, error) {
	w := &wallet{
		name:     fmt.Sprintf("clef-%s", endpoint),
		client:   client,
		endpoint: endpoint,
	}

	if !w.contains(accountAddr) {
		return nil, fmt.Errorf("account with address %q not found", accountAddr)
	}

	w.account = newAccount(accountAddr, w.endpoint)

	return w, nil
}

// GenerateNewWallet new wallet will create new account in Clef and returns wallet.
// Caveat: generating new wallet in Clef has to be manually approved and only key store backend is supported.
func GenerateNewWallet(client Client, endpoint string) (*wallet, error) {
	w := &wallet{
		name:     fmt.Sprintf("clef-%s", endpoint),
		client:   client,
		endpoint: endpoint,
	}

	acc, err := w.generateAccount()
	if err != nil {
		return nil, err
	}

	w.account = acc

	return w, nil
}

func (w *wallet) generateAccount() (*accounts.Account, error) {
	// increase timeout here as generating new account has to be manually approved in Clef
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout*20)
	defer cancel()

	var res string
	if err := w.client.CallContext(ctx, &res, "account_new"); err != nil {
		return nil, err
	}

	return newAccount(ethcommon.HexToAddress(res), w.endpoint), nil
}

func (w *wallet) contains(testAddr ethcommon.Address) bool {
	addresses, err := w.listAccounts()
	if err != nil {
		// TODO log the error here
		return false
	}

	for _, addr := range addresses {
		if testAddr == addr {
			return true
		}
	}

	return false
}

func (w *wallet) listAccounts() ([]ethcommon.Address, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	var res []ethcommon.Address
	if err := w.client.CallContext(ctx, &res, "account_list"); err != nil {
		return nil, err
	}
	return res, nil
}

// Cleanup is noop
func (w *wallet) Cleanup() error {
	w.client.Close()
	return nil
}

func (w *wallet) Name() string {
	return w.name
}

func (w *wallet) Chain() string {
	return "ethereum"
}

func (w *wallet) Sign(data []byte) ([]byte, error) {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	var res hexutil.Bytes
	signAddress := ethcommon.NewMixedcaseAddress(w.account.Address)

	if err := w.client.CallContext(
		ctx,
		&res,
		"account_signData",
		accounts.MimetypeTypedData,
		&signAddress, // Need to use the pointer here, because of how MarshalJSON is defined
		hexutil.Encode(data),
	); err != nil {
		return nil, err
	}

	return res, nil
}

func (w *wallet) Algo() string {
	return "eth"
}

func (w *wallet) Version() string {
	ctx, cancel := context.WithTimeout(context.Background(), requestTimeout)
	defer cancel()

	var v string
	if err := w.client.CallContext(ctx, &v, "account_version"); err != nil {
		return ""
	}

	return v
}

func (w *wallet) PubKeyOrAddress() crypto.PublicKeyOrAddress {
	return crypto.NewPublicKeyOrAddress(w.account.Address.Hex(), w.account.Address.Bytes())
}
