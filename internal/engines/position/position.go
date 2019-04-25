package position

import (
	"fmt"
	"sync"

	"code.vegaprotocol.io/vega/internal/logging"
	types "code.vegaprotocol.io/vega/proto"
)

type MarketPosition struct {
	size    int64
	margins map[string]uint64
	partyID string
	price   uint64
}

func (m MarketPosition) String() string {
	return fmt.Sprintf("size: %v, margins: %v, partyID: %v", m.size, m.margins, m.partyID)
}

// Margins returns a copy of the current margins map
func (m *MarketPosition) Margins() map[string]uint64 {
	out := make(map[string]uint64, 0)
	for k, v := range m.margins {
		out[k] = v
	}
	return out
}

// UpdateMargin updates the margin value for a single asset
func (m *MarketPosition) UpdateMargin(assetID string, margin uint64) {
	m.margins[assetID] = margin
}

func (m MarketPosition) Size() int64 {
	return m.size
}

func (m MarketPosition) Party() string {
	return m.partyID
}

func (m MarketPosition) Price() uint64 {
	return m.price
}

type Engine struct {
	log *logging.Logger
	Config

	cfgMu sync.Mutex
	mu    *sync.RWMutex
	// partyID -> MarketPosition
	positions map[string]*MarketPosition
}

func New(log *logging.Logger, config Config) *Engine {
	// setup logger
	log = log.Named(namedLogger)
	log.SetLevel(config.Level.Get())

	return &Engine{
		Config:    config,
		log:       log,
		mu:        &sync.RWMutex{},
		positions: map[string]*MarketPosition{},
	}
}

func (e *Engine) ReloadConf(cfg Config) {
	e.log.Info("reloading configuration")
	if e.log.GetLevel() != cfg.Level.Get() {
		e.log.Info("updating log level",
			logging.String("old", e.log.GetLevel().String()),
			logging.String("new", cfg.Level.String()),
		)
		e.log.SetLevel(cfg.Level.Get())
	}

	e.cfgMu.Lock()
	e.Config = cfg
	e.cfgMu.Unlock()
}

func (e *Engine) Update(trade *types.Trade) {

	// Not using defer e.mu.Unlock(), because defer calls add some overhead
	// and this is called for each transaction, so we want to optimise as much as possible
	// there aren't multiple returns here anyway, so just unlock as and when it's needed
	e.mu.Lock()
	// todo(cdm): overflow should be managed at the trade/order creation point. We shouldn't accept an order onto
	// your book that would overflow your position. Order validation requires position store/state lookup.

	buyer, ok := e.positions[trade.Buyer]
	if !ok {
		e.positions[trade.Buyer] = &MarketPosition{
			margins: map[string]uint64{},
			partyID: trade.Buyer,
		}
		buyer = e.positions[trade.Buyer]
	}

	seller, ok := e.positions[trade.Seller]
	if !ok {
		e.positions[trade.Seller] = &MarketPosition{
			margins: map[string]uint64{},
			partyID: trade.Seller,
		}
		seller = e.positions[trade.Seller]
	}

	// get net value of trade, add that to total price running for buyer/seller
	price := trade.Size * trade.Price

	// Buyer INCREASED position size buy trade.Size
	buyer.size += int64(trade.Size)
	buyer.price += price

	// Seller DECREASED position size buy trade.Size
	seller.size -= int64(trade.Size)
	// add price, still. this is keeping a running total of the sell price
	seller.price += price

	e.log.Debug("Positions Updated for trade",
		logging.Trade(*trade),
		logging.String("buyer-position", fmt.Sprintf("%+v", buyer)),
		logging.String("seller-position", fmt.Sprintf("%+v", seller)))

	// we've set all the values now, unlock after logging
	// because we're working on MarketPosition pointers
	e.mu.Unlock()
}

func (e *Engine) Positions() []MarketPosition {
	e.mu.RLock()
	out := make([]MarketPosition, 0, len(e.positions))
	for _, value := range e.positions {
		out = append(out, *value)
	}
	e.mu.RUnlock()
	return out
}