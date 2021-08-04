package validators

import (
	"context"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"sync"

	"code.vegaprotocol.io/vega/crypto"
	"code.vegaprotocol.io/vega/logging"
	commandspb "code.vegaprotocol.io/protos/vega/commands/v1"
)

var (
	ErrVegaNodeAlreadyRegisterForChain = errors.New("a vega node is already registered with the blockchain node")
	ErrInvalidChainPubKey              = errors.New("invalid blockchain public key")
)

//go:generate go run github.com/golang/mock/mockgen -destination mocks/wallet_mock.go -package mocks code.vegaprotocol.io/vega/validators Wallet
type Wallet interface {
	PubKeyOrAddress() crypto.PublicKeyOrAddress
}

// ValidatorMapping maps a tendermint pubkey with a vega pubkey
type ValidatorMapping map[string]string

type Topology struct {
	log    *logging.Logger
	cfg    Config
	wallet Wallet

	// tendermint validator pubkey to vega pubkey
	validators ValidatorMapping
	// just pubkeys of vega node for easy lookup
	vegaValidatorRefs map[string]struct{}
	chainValidators   [][]byte

	isValidator bool

	mu sync.Mutex
}

func NewTopology(log *logging.Logger, cfg Config, wallet Wallet) *Topology {
	log = log.Named(namedLogger)
	log.SetLevel(cfg.Level.Get())

	t := &Topology{
		log:               log,
		cfg:               cfg,
		wallet:            wallet,
		validators:        ValidatorMapping{},
		chainValidators:   [][]byte{},
		vegaValidatorRefs: map[string]struct{}{},
	}

	return t
}

// ReloadConf updates the internal configuration
func (t *Topology) ReloadConf(cfg Config) {
	t.log.Info("reloading configuration")
	if t.log.GetLevel() != cfg.Level.Get() {
		t.log.Info("updating log level",
			logging.String("old", t.log.GetLevel().String()),
			logging.String("new", cfg.Level.String()),
		)
		t.log.SetLevel(cfg.Level.Get())
	}

	t.cfg = cfg
}

func (t *Topology) IsValidator() bool {
	return t.isValidator
}

func (t *Topology) Len() int {
	return len(t.vegaValidatorRefs)
}

// Exists check if a vega public key is part of the validator set
func (t *Topology) Exists(key []byte) bool {
	_, ok := t.vegaValidatorRefs[string(key)]
	return ok
}

func (t *Topology) AllPubKeys() [][]byte {
	t.mu.Lock()
	defer t.mu.Unlock()
	keys := make([][]byte, 0, len(t.validators))
	for _, key := range t.validators {
		keys = append(keys, []byte(key))
	}
	return keys
}

func (t *Topology) SelfVegaPubKey() []byte {
	return t.wallet.PubKeyOrAddress().Bytes()
}

// UpdateValidatorSet updates the chain validator set
// It overwrites the previous set.
func (t *Topology) UpdateValidatorSet(keys [][]byte) {
	t.mu.Lock()
	defer t.mu.Unlock()
	t.chainValidators = keys
}

//IsValidatorNode takes a nodeID and returns true if the node is a validator node
func (t *Topology) IsValidatorNode(nodeID string) bool {
	_, ok := t.validators[nodeID]
	return ok
}

func (t *Topology) AddNodeRegistration(nr *commandspb.NodeRegistration) error {
	t.mu.Lock()
	defer t.mu.Unlock()

	key := hex.EncodeToString(nr.ChainPubKey)
	if _, ok := t.validators[key]; ok {
		return ErrVegaNodeAlreadyRegisterForChain
	}
	// check if this tm pubkey exists in the network
	var ok bool
	for _, k := range t.chainValidators {
		if string(k) == string(nr.ChainPubKey) {
			ok = true
			break
		}
	}
	if !ok {
		return ErrInvalidChainPubKey
	}

	// then add it to the topology
	t.validators[key] = string(nr.PubKey)
	t.vegaValidatorRefs[string(nr.PubKey)] = struct{}{}
	t.log.Info("new node registration successful",
		logging.String("node-key", hex.EncodeToString(nr.PubKey)),
		logging.String("tm-key", hex.EncodeToString(nr.ChainPubKey)))
	return nil
}

func (t *Topology) LoadValidatorsOnGenesis(_ context.Context, rawstate []byte) error {
	state, err := LoadGenesisState(rawstate)
	if err != nil {
		return err
	}

	pubKey := t.wallet.PubKeyOrAddress().Hex()

	// vals is a map of tm pubkey -> vega pubkey
	// tm is base64 encoded, vega is hex
	for tm, vega := range state {
		tmBytes, err := base64.StdEncoding.DecodeString(tm)
		if err != nil {
			return err
		}

		vegaBytes, err := hex.DecodeString(vega)
		if err != nil {
			return err
		}

		if pubKey == vega {
			t.isValidator = true
		}

		nr := &commandspb.NodeRegistration{
			PubKey:      vegaBytes,
			ChainPubKey: tmBytes,
		}
		if err := t.AddNodeRegistration(nr); err != nil {
			return err
		}
	}

	return nil
}
