package nullchain

import (
	"context"
	"encoding/json"
	"errors"
	"sync"
	"time"

	vgfs "code.vegaprotocol.io/shared/libs/fs"
	vgrand "code.vegaprotocol.io/shared/libs/rand"
	"code.vegaprotocol.io/vega/logging"

	abci "github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/proto/tendermint/types"
	tmctypes "github.com/tendermint/tendermint/rpc/core/types"
	tmtypes "github.com/tendermint/tendermint/types"
)

var (
	ErrNotImplemented      = errors.New("not implemented for nullblockchain")
	ErrGenesisFileRequired = errors.New("--blockchain.nullchain.genesis-file is required")
)

//go:generate go run github.com/golang/mock/mockgen -destination mocks/application_service_mock.go -package mocks code.vegaprotocol.io/vega/blockchain/nullchain ApplicationService
type ApplicationService interface {
	InitChain(res abci.RequestInitChain) (resp abci.ResponseInitChain)
	BeginBlock(req abci.RequestBeginBlock) (resp abci.ResponseBeginBlock)
	EndBlock(req abci.RequestEndBlock) (resp abci.ResponseEndBlock)
	Commit() (resp abci.ResponseCommit)
	DeliverTx(req abci.RequestDeliverTx) (resp abci.ResponseDeliverTx)
}

type NullBlockchain struct {
	log                  *logging.Logger
	service              ApplicationService
	chainID              string
	genesisFile          string
	genesisTime          time.Time
	transactionsPerBlock uint64
	blockDuration        time.Duration

	now         time.Time
	pending     []*abci.RequestDeliverTx
	blockHeight int64

	mu sync.Mutex
}

func NewClient(
	log *logging.Logger,
	cfg Config,
	service ApplicationService,
) *NullBlockchain {
	// setup logger
	log = log.Named(namedLogger)
	log.SetLevel(cfg.Level.Get())
	log.Info("starting nullblockchain")

	now := time.Now()
	n := &NullBlockchain{
		log:                  log,
		blockHeight:          0,
		service:              service,
		chainID:              vgrand.RandomStr(12),
		transactionsPerBlock: cfg.TransactionsPerBlock,
		blockDuration:        cfg.BlockDuration.Duration,
		genesisFile:          cfg.GenesisFile,
		genesisTime:          now,
		now:                  now,
		pending:              make([]*abci.RequestDeliverTx, 0),
	}

	return n
}

func (n *NullBlockchain) Start() error {
	err := n.InitChain(n.genesisFile)
	if err != nil {
		return err
	}

	// Start the first block
	n.BeginBlock()
	return nil
}

func (n *NullBlockchain) processBlock() {
	n.log.Debugf("processing block %d with %d transactions", n.blockHeight, len(n.pending))
	for _, tx := range n.pending {
		n.service.DeliverTx(*tx)
	}
	n.pending = n.pending[:0]

	n.blockHeight++
	n.EndBlock()
	n.service.Commit()

	// Increment time, start and start new block
	n.now = n.now.Add(n.blockDuration)
	n.BeginBlock()
}

func (n *NullBlockchain) handleTransaction(tx []byte) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.pending = append(n.pending, &abci.RequestDeliverTx{Tx: tx})

	n.log.Debugf("transaction added to block: %d of %d", len(n.pending), n.transactionsPerBlock)
	if len(n.pending) == int(n.transactionsPerBlock) {
		n.processBlock()
	}
}

// ForwardTime moves the chain time forward by the given duration, delivering any pending
// transaction and creating any extra empty blocks if time is stepped forward by more than
// a block duration.
func (n *NullBlockchain) ForwardTime(d time.Duration) {
	n.log.Debugf("time-forwarding by %s", d)

	nBlocks := d / n.blockDuration
	if nBlocks == 0 {
		n.log.Debugf("not a full block-duration, not moving time: %s < %s", d, n.blockDuration)
		return
	}

	n.mu.Lock()
	defer n.mu.Unlock()
	for i := 0; i < int(nBlocks); i++ {
		n.processBlock()
	}
}

// InitChain processes the given genesis file setting the chain's time, and passing the
// appstate through to the processors InitChain.
func (n *NullBlockchain) InitChain(genesisFile string) error {
	exists, err := vgfs.FileExists(genesisFile)
	if !exists || err != nil {
		return ErrGenesisFileRequired
	}

	b, err := vgfs.ReadFile(genesisFile)
	if err != nil {
		return err
	}

	// Parse the appstate of the genesis file, same layout as a TM genesis-file
	genesis := struct {
		GenesisTime *time.Time      `json:"genesis_time"`
		Appstate    json.RawMessage `json:"app_state"`
	}{}

	err = json.Unmarshal(b, &genesis)
	if err != nil {
		return err
	}

	// Set chain time from genesis file
	if genesis.GenesisTime != nil {
		n.genesisTime = *genesis.GenesisTime
		n.now = *genesis.GenesisTime
	}

	n.service.InitChain(
		abci.RequestInitChain{
			Time:          n.now,
			ChainId:       n.chainID,
			InitialHeight: n.blockHeight,
			AppStateBytes: genesis.Appstate,
		})
	return nil
}

func (n *NullBlockchain) BeginBlock() *NullBlockchain {
	r := abci.RequestBeginBlock{
		Header: types.Header{
			Time: n.now,
		},
	}
	n.service.BeginBlock(r)
	return n
}

func (n *NullBlockchain) EndBlock() *NullBlockchain {
	r := abci.RequestEndBlock{
		Height: n.blockHeight,
	}
	n.service.EndBlock(r)
	return n
}

func (n *NullBlockchain) GetGenesisTime(context.Context) (time.Time, error) {
	return n.genesisTime, nil
}

func (n *NullBlockchain) GetChainID(context.Context) (string, error) {
	return n.chainID, nil
}

func (n *NullBlockchain) GetStatus(context.Context) (*tmctypes.ResultStatus, error) {
	return &tmctypes.ResultStatus{
		NodeInfo: p2p.DefaultNodeInfo{
			Version: "0.34.12",
		},
		SyncInfo: tmctypes.SyncInfo{
			CatchingUp: false,
		},
	}, nil
}

func (n *NullBlockchain) GetNetworkInfo(context.Context) (*tmctypes.ResultNetInfo, error) {
	return &tmctypes.ResultNetInfo{
		Listening: true,
		Listeners: []string{},
		NPeers:    0,
	}, nil
}

func (n *NullBlockchain) GetUnconfirmedTxCount(context.Context) (int, error) {
	n.mu.Lock()
	defer n.mu.Unlock()
	return len(n.pending), nil
}

func (n *NullBlockchain) Health(_ context.Context) (*tmctypes.ResultHealth, error) {
	return &tmctypes.ResultHealth{}, nil
}

func (n *NullBlockchain) SendTransactionAsync(ctx context.Context, tx []byte) (string, error) {
	go func() {
		n.handleTransaction(tx)
	}()
	return vgrand.RandomStr(64), nil
}

func (n *NullBlockchain) SendTransactionSync(ctx context.Context, tx []byte) (string, error) {
	n.handleTransaction(tx)
	return vgrand.RandomStr(64), nil
}

func (n *NullBlockchain) SendTransactionCommit(ctx context.Context, tx []byte) (string, error) {
	// I think its worth only implementing this if needed. With time-forwarding we already have
	// control over when a block ends and gets committed, so I don't think its worth adding the
	// the complexity of trying to keep track of tx deliveries here.
	n.log.Error("not implemented")
	return "", ErrNotImplemented
}

func (n *NullBlockchain) Validators(_ context.Context) ([]*tmtypes.Validator, error) {
	n.log.Error("not implemented")
	return nil, ErrNotImplemented
}

func (n *NullBlockchain) GenesisValidators(_ context.Context) ([]*tmtypes.Validator, error) {
	n.log.Error("not implemented")
	return nil, ErrNotImplemented
}

func (n *NullBlockchain) Subscribe(context.Context, func(tmctypes.ResultEvent) error, ...string) error {
	n.log.Error("not implemented")
	return ErrNotImplemented
}
