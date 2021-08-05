package limits

import (
	"context"
	"fmt"
	"time"

	"code.vegaprotocol.io/vega/logging"
)

type Engine struct {
	log *logging.Logger
	cfg Config

	// are these action possible?
	canProposeMarket, canProposeAsset bool

	// Settings from the genesis state
	proposeMarketEnabled, proposeAssetEnabled         bool
	proposeMarketEnabledFrom, proposeAssetEnabledFrom time.Time
}

func New(cfg Config, log *logging.Logger) *Engine {
	log = log.Named(namedLogger)
	log.SetLevel(cfg.Level.Get())
	return &Engine{
		log: log,
		cfg: cfg,
	}
}

// UponGenesis load the limits from the genersis state
func (e *Engine) UponGenesis(ctx context.Context, rawState []byte) error {
	e.log.Debug("loading genesis configuration")
	state, err := LoadGenesisState(rawState)
	if err != nil && err != ErrNoLimitsGenesisState {
		e.log.Error("unable to load genesis state",
			logging.Error(err))
		return err
	}

	if err == ErrNoLimitsGenesisState {
		defaultState := DefaultGenesisState()
		state = &defaultState
	}

	// set enabled by default if not genesis state
	if state == nil {
		e.proposeAssetEnabled = true
		e.proposeMarketEnabled = true
		return nil
	}

	e.proposeAssetEnabled = state.ProposeAssetEnabled
	e.proposeMarketEnabled = state.ProposeMarketEnabled
	e.proposeAssetEnabledFrom = timeFromPtr(state.ProposeAssetEnabledFrom)
	e.proposeMarketEnabledFrom = timeFromPtr(state.ProposeMarketEnabledFrom)

	e.log.Info("loaded limits genesis state",
		logging.String("state", fmt.Sprintf("%#v", *state)))

	return nil
}

func (e *Engine) OnTick(_ context.Context, t time.Time) {
	if e.canProposeAsset && e.canProposeMarket {
		return
	}

	if !e.canProposeMarket && e.proposeMarketEnabled && t.After(e.proposeMarketEnabledFrom) {
		e.canProposeMarket = true
	}
	if !e.canProposeAsset && e.proposeAssetEnabled && t.After(e.proposeAssetEnabledFrom) {
		e.canProposeAsset = true
	}

}

func (e *Engine) CanProposeMarket() bool {
	return e.canProposeMarket
}

func (e *Engine) CanProposeAsset() bool {
	return e.canProposeAsset
}

func (e *Engine) CanTrade() bool {
	return e.canProposeAsset && e.canProposeMarket
}

func timeFromPtr(tptr *time.Time) time.Time {
	var t time.Time
	if tptr != nil {
		t = *tptr
	}
	return t
}
