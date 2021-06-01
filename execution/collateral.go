package execution

import (
	"context"
	"errors"
	"math/big"

	"code.vegaprotocol.io/vega/events"
	"code.vegaprotocol.io/vega/types"

	"github.com/shopspring/decimal"
)

var (
	// ErrBondSlashing - just indicates that we had to penalize the trader due to insufficient funds, and as such, we have to cancel their LP
	ErrBondSlashing = errors.New("bond slashing")
)

// this will transfer funds calculated for a party amending a liquidity
// provision during auction.
func (m *Market) transferMarginsLiquidityProvisionAmendAuction(
	ctx context.Context, risk events.Risk) error {
	market := m.GetID()
	// This is ultimately the same behaviour than update on order
	// all or nothing of margin needsto be transferred
	tsfr, _, err := m.collateral.MarginUpdateOnOrder(ctx, market, risk)
	if err != nil {
		return err
	}

	m.broker.Send(events.NewTransferResponse(ctx, []*types.TransferResponse{tsfr}))
	return nil
}

func (m *Market) transferMargins(ctx context.Context, risk []events.Risk, closed []events.MarketPosition) error {
	if m.as.InAuction() {
		return m.transferMarginsAuction(ctx, risk, closed)
	}
	return m.transferMarginsContinuous(ctx, risk)
}

func (m *Market) transferMarginsAuction(ctx context.Context, risk []events.Risk, distressed []events.MarketPosition) error {
	evts := make([]events.Event, 0, len(risk))
	mID := m.GetID()
	// first, update the margin accounts for all traders who have enough balance
	for _, re := range risk {
		tr, _, err := m.collateral.MarginUpdateOnOrder(ctx, mID, re)
		if err != nil {
			// @TODO handle this
			return err
		}
		evts = append(evts, events.NewTransferResponse(ctx, []*types.TransferResponse{tr}))
	}
	m.broker.SendBatch(evts)
	rmorders, err := m.matching.RemoveDistressedOrders(distressed)
	if err != nil {
		return err
	}
	evts = make([]events.Event, 0, len(rmorders))
	for _, o := range rmorders {
		// cancel order
		o.Status = types.Order_STATUS_CANCELLED
		o.Reason = types.OrderError_ORDER_ERROR_INSUFFICIENT_ASSET_BALANCE
		// create event
		evts = append(evts, events.NewOrderEvent(ctx, o))
		// remove order from positions
		_ = m.position.UnregisterOrder(o)
	}
	m.broker.SendBatch(evts)
	return nil
}

func (m *Market) transferMarginsContinuous(ctx context.Context, risk []events.Risk) error {
	if len(risk) > 1 {
		return errors.New("transferMarginsContinuous should not be possible when len(risk) > 1")
	}
	if len(risk) == 0 {
		return nil
	}
	mID := m.GetID()
	tr, closed, err := m.collateral.MarginUpdateOnOrder(ctx, mID, risk[0])
	if err != nil {
		return err
	}
	// if LP shortfall is not empty, this trader will have to pay the LP penalty
	responses := make([]*types.TransferResponse, 0, len(risk))
	if tr != nil {
		responses = append(responses, tr)
	}
	// margin shortfall && liquidity provider -> bond slashing
	if closed != nil && closed.MarginShortFall() > 0 {
		// we pay the bond penalty if the order was not pending
		if !m.liquidity.IsPending(closed.Party()) {
			// get bond penalty
			resp, err := m.bondSlashing(ctx, closed)
			if err != nil {
				return err
			}
			responses = append(responses, resp...)
		}
	}
	m.broker.Send(events.NewTransferResponse(ctx, responses))
	return nil
}

func (m *Market) bondSlashing(ctx context.Context, closed ...events.Margin) ([]*types.TransferResponse, error) {
	mID := m.GetID()
	asset, _ := m.mkt.GetAsset()
	factor := decimal.NewFromFloatWithExponent(m.bondPenaltyFactor, 0)
	ret := make([]*types.TransferResponse, 0, len(closed))
	for _, c := range closed {
		shortfall := decimal.NewFromBigInt(new(big.Int).SetUint64(c.MarginShortFall()), 0)
		penalty := shortfall.Mul(factor).Floor().BigInt().Uint64()
		resp, err := m.collateral.BondUpdate(ctx, mID, c.Party(), &types.Transfer{
			Owner: c.Party(),
			Amount: &types.FinancialAmount{
				Amount: penalty,
				Asset:  asset,
			},
			Type: types.TransferType_TRANSFER_TYPE_BOND_SLASHING,
		})
		if err != nil {
			return nil, err
		}
		ret = append(ret, resp)
	}
	return ret, nil
}
