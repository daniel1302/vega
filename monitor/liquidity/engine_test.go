package liquidity_test

import (
	"testing"
	"time"

	"code.vegaprotocol.io/vega/monitor/liquidity"
	"code.vegaprotocol.io/vega/monitor/liquidity/mocks"
	"code.vegaprotocol.io/vega/types"
	"code.vegaprotocol.io/vega/types/num"
	"github.com/golang/mock/gomock"
)

type testHarness struct {
	AuctionState          *mocks.MockAuctionState
	TargetStakeCalculator *mocks.MockTargetStakeCalculator
}

func newTestHarness(t *testing.T) *testHarness {
	t.Helper()
	ctrl := gomock.NewController(t)
	return &testHarness{
		AuctionState:          mocks.NewMockAuctionState(ctrl),
		TargetStakeCalculator: mocks.NewMockTargetStakeCalculator(ctrl),
	}
}

func (h *testHarness) WhenInLiquidityAuction(v bool) *testHarness {
	h.AuctionState.EXPECT().IsLiquidityAuction().AnyTimes().Return(v)
	return h
}

func TestEngineWhenInLiquidityAuction(t *testing.T) {
	now := time.Now()

	tests := []struct {
		desc string
		// when
		current             *num.Uint
		target              *num.Uint
		bestStaticBidVolume uint64
		bestStaticAskVolume uint64
		// expect
		auctionShouldEnd bool
	}{
		{"Current >  Target", num.NewUint(20), num.NewUint(15), 1, 1, true},
		{"Current == Target", num.NewUint(15), num.NewUint(15), 1, 1, true},
		{"Current <  Target", num.NewUint(14), num.NewUint(15), 1, 1, false},
		{"Current >  Target, no best bid", num.NewUint(20), num.NewUint(15), 0, 1, false},
		{"Current == Target, no best ask", num.NewUint(15), num.NewUint(15), 1, 0, false},
		{"Current == Target, no best bid and ask", num.NewUint(15), num.NewUint(15), 0, 0, false},
	}

	h := newTestHarness(t).WhenInLiquidityAuction(true)
	exp := now.Add(-1 * time.Second)
	keep := now.Add(time.Second)
	mon := liquidity.NewMonitor(h.TargetStakeCalculator, &types.LiquidityMonitoringParameters{
		TriggeringRatio: num.DecimalFromFloat(.7),
	})
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			var trades []*types.Trade = nil
			rf := types.RiskFactor{}
			markPrice := num.NewUint(100)
			if test.auctionShouldEnd {
				h.AuctionState.EXPECT().SetReadyToLeave().Times(1)
				h.AuctionState.EXPECT().ExpiresAt().Times(1).Return(&exp)
				h.TargetStakeCalculator.EXPECT().GetTheoreticalTargetStake(rf, now, markPrice.Clone(), trades).Times(1).Return(test.target)
			} else {
				h.AuctionState.EXPECT().ExpiresAt().Times(1).Return(&keep)
			}

			mon.CheckLiquidity(h.AuctionState, now, test.current, trades, rf, markPrice.Clone(), test.bestStaticBidVolume, test.bestStaticAskVolume)
		})
	}
}

func TestEngineWhenNotInLiquidityAuction(t *testing.T) {
	now := time.Now()

	tests := []struct {
		desc string
		// when
		current             *num.Uint
		target              *num.Uint
		bestStaticBidVolume uint64
		bestStaticAskVolume uint64
		// expect
		auctionShouldStart bool
	}{
		{"Current <  (Target * c1)", num.NewUint(10), num.NewUint(30), 1, 1, true},
		{"Current >  (Target * c1)", num.NewUint(15), num.NewUint(15), 1, 1, false},
		{"Current == (Target * c1)", num.NewUint(10), num.NewUint(20), 1, 1, false},
		{"Current >  (Target * c1), no best bid", num.NewUint(15), num.NewUint(15), 0, 1, true},
		{"Current == (Target * c1), no best ask", num.NewUint(10), num.NewUint(20), 1, 0, true},
		{"Current == (Target * c1), no best bid and ask", num.NewUint(10), num.NewUint(20), 0, 0, true},
	}

	h := newTestHarness(t).WhenInLiquidityAuction(false)
	mon := liquidity.NewMonitor(h.TargetStakeCalculator, &types.LiquidityMonitoringParameters{
		TriggeringRatio: num.DecimalFromFloat(.5),
	})
	h.AuctionState.EXPECT().ExpiresAt().Times(len(tests)).Return(nil)
	for _, test := range tests {
		t.Run(test.desc, func(t *testing.T) {
			if test.auctionShouldStart {
				h.AuctionState.EXPECT().StartLiquidityAuction(now, gomock.Any()).Times(1)
			}
			var trades []*types.Trade = nil
			rf := types.RiskFactor{}
			markPrice := num.NewUint(100)
			h.TargetStakeCalculator.EXPECT().GetTheoreticalTargetStake(rf, now, markPrice.Clone(), trades).Return(test.target)
			mon.CheckLiquidity(h.AuctionState, now, test.current, trades, rf, markPrice.Clone(), test.bestStaticBidVolume, test.bestStaticAskVolume)
		})
	}
}

func TestEngineAfterParametersUpdate(t *testing.T) {
	h := newTestHarness(t).WhenInLiquidityAuction(false)

	now := time.Now()
	target := num.NewUint(100)
	bestStaticBidVolume := uint64(1)
	bestStaticAskVolume := uint64(1)
	var trades []*types.Trade
	rf := types.RiskFactor{}
	markPrice := num.NewUint(100)
	params := &types.LiquidityMonitoringParameters{
		TriggeringRatio:  num.DecimalFromFloat(.5),
		AuctionExtension: 40,
	}

	mon := liquidity.NewMonitor(h.TargetStakeCalculator, params)

	expiresAt := now.Add(-24 * time.Hour)
	h.AuctionState.EXPECT().ExpiresAt().Times(1).Return(&expiresAt)
	h.AuctionState.EXPECT().ExtendAuctionLiquidity(types.AuctionDuration{
		Duration: params.AuctionExtension,
	}).Times(1)
	h.TargetStakeCalculator.EXPECT().GetTheoreticalTargetStake(rf, now, markPrice.Clone(), trades).Return(target)

	mon.CheckLiquidity(h.AuctionState, now, num.NewUint(40), trades, rf, markPrice.Clone(), bestStaticBidVolume, bestStaticAskVolume)

	updatedParams := &types.LiquidityMonitoringParameters{
		TriggeringRatio:  num.DecimalFromFloat(.8),
		AuctionExtension: 80,
	}

	mon.UpdateParameters(updatedParams)

	h.AuctionState.EXPECT().ExpiresAt().Times(1).Return(&expiresAt)
	// Verify the auction extension is called with update parameters.
	h.AuctionState.EXPECT().ExtendAuctionLiquidity(types.AuctionDuration{
		Duration: updatedParams.AuctionExtension,
	}).Times(1)

	h.TargetStakeCalculator.EXPECT().GetTheoreticalTargetStake(rf, now, markPrice.Clone(), trades).Return(target)
	// Higher current stake to test the updated Triggering Ratio is reached.
	mon.CheckLiquidity(h.AuctionState, now, num.NewUint(70), nil, rf, markPrice.Clone(), bestStaticBidVolume, bestStaticAskVolume)
}

func TestEngineAfterParametersUpdateWithAuctionExtension0(t *testing.T) {
	h := newTestHarness(t).WhenInLiquidityAuction(false)

	now := time.Now()
	target := num.NewUint(100)
	bestStaticBidVolume := uint64(1)
	bestStaticAskVolume := uint64(1)
	var trades []*types.Trade
	rf := types.RiskFactor{}
	markPrice := num.NewUint(100)
	params := &types.LiquidityMonitoringParameters{
		TriggeringRatio:  num.DecimalFromFloat(.5),
		AuctionExtension: 0,
	}

	mon := liquidity.NewMonitor(h.TargetStakeCalculator, params)

	expiresAt := now.Add(-24 * time.Hour)
	h.AuctionState.EXPECT().ExpiresAt().Times(1).Return(&expiresAt)
	h.AuctionState.EXPECT().ExtendAuctionLiquidity(types.AuctionDuration{
		Duration: 1, // to test the patch.
	}).Times(1)
	h.TargetStakeCalculator.EXPECT().GetTheoreticalTargetStake(rf, now, markPrice.Clone(), trades).Return(target)

	mon.CheckLiquidity(h.AuctionState, now, num.NewUint(40), trades, rf, markPrice.Clone(), bestStaticBidVolume, bestStaticAskVolume)

	updatedParams := &types.LiquidityMonitoringParameters{
		TriggeringRatio:  num.DecimalFromFloat(.8),
		AuctionExtension: 0,
	}

	mon.UpdateParameters(updatedParams)

	h.AuctionState.EXPECT().ExpiresAt().Times(1).Return(&expiresAt)
	// Verify the auction extension is called with update parameters.
	h.AuctionState.EXPECT().ExtendAuctionLiquidity(types.AuctionDuration{
		Duration: 1, // to test the patch.
	}).Times(1)

	h.TargetStakeCalculator.EXPECT().GetTheoreticalTargetStake(rf, now, markPrice.Clone(), trades).Return(target)
	// Higher current stake to test the updated Triggering Ratio is reached.
	mon.CheckLiquidity(h.AuctionState, now, num.NewUint(70), nil, rf, markPrice.Clone(), bestStaticBidVolume, bestStaticAskVolume)
}
