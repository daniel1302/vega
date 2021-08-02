package fee

import (
	"context"
	"math"
	"strconv"

	"code.vegaprotocol.io/vega/logging"
	types "code.vegaprotocol.io/protos/vega"
)

// MarketStore ...
//go:generate go run github.com/golang/mock/mockgen -destination mocks/market_store_mock.go -package mocks code.vegaprotocol.io/vega/fee MarketStore
type MarketStore interface {
	GetByID(name string) (*types.Market, error)
}

// MarketDataStore ...
//go:generate go run github.com/golang/mock/mockgen -destination mocks/market_data_store_mock.go -package mocks code.vegaprotocol.io/vega/fee MarketDataStore
type MarketDataStore interface {
	GetByID(marketID string) (types.MarketData, error)
}

type Svc struct {
	cfg          Config
	log          *logging.Logger
	mktStore     MarketStore
	mktDataStore MarketDataStore
}

func NewService(log *logging.Logger, cfg Config, mktStore MarketStore, mktDataStore MarketDataStore) *Svc {
	return &Svc{
		cfg:          cfg,
		log:          log,
		mktStore:     mktStore,
		mktDataStore: mktDataStore,
	}
}

// ReloadConf is used in order to reload the internal configuration of
// the of the fee engine
func (s *Svc) ReloadConf(cfg Config) {
	s.log.Info("reloading configuration")
	if s.log.GetLevel() != cfg.Level.Get() {
		s.log.Info("updating log level",
			logging.String("old", s.log.GetLevel().String()),
			logging.String("new", cfg.Level.String()),
		)
		s.log.SetLevel(cfg.Level.Get())
	}

	s.cfg = cfg
}

func (s *Svc) EstimateFee(ctx context.Context, o *types.Order) (*types.Fee, error) {
	mkt, err := s.mktStore.GetByID(o.MarketId)
	if err != nil {
		return nil, err
	}
	price := o.Price
	if o.PeggedOrder != nil {
		return &types.Fee{
			MakerFee:          0,
			InfrastructureFee: 0,
			LiquidityFee:      0,
		}, nil
	}

	base := float64(price * o.Size)
	maker, infra, liquidity, err := s.feeFactors(mkt)
	if err != nil {
		return nil, err
	}

	fee := &types.Fee{
		MakerFee:          uint64(math.Ceil(base * maker)),
		InfrastructureFee: uint64(math.Ceil(base * infra)),
		LiquidityFee:      uint64(math.Ceil(base * liquidity)),
	}

	// if mkt.State == types.MarketState_MARKET_STATE_OPENING_AUCTION {
	// 	// half price paid by both partis
	// 	fee.MakerFee = fee.MakerFee / 2
	// 	fee.InfrastructureFee = fee.InfrastructureFee / 2
	// 	fee.LiquidityFee = fee.LiquidityFee / 2
	// }

	return fee, nil
}

func (s *Svc) feeFactors(mkt *types.Market) (maker, infra, liquidity float64, err error) {
	maker, err = strconv.ParseFloat(mkt.Fees.Factors.MakerFee, 64)
	if err != nil {
		return
	}
	infra, err = strconv.ParseFloat(mkt.Fees.Factors.InfrastructureFee, 64)
	if err != nil {
		return
	}
	liquidity, err = strconv.ParseFloat(mkt.Fees.Factors.LiquidityFee, 64)
	if err != nil {
		return
	}
	return
}
