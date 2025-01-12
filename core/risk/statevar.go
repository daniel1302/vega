// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE.VEGA file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package risk

import (
	"context"

	"code.vegaprotocol.io/vega/core/events"
	"code.vegaprotocol.io/vega/core/types"
	"code.vegaprotocol.io/vega/core/types/statevar"
	"code.vegaprotocol.io/vega/libs/num"
	"code.vegaprotocol.io/vega/logging"
)

type FactorConverter struct{}

var riskFactorTolerance = num.MustDecimalFromString("1e-6")

func (FactorConverter) BundleToInterface(kvb *statevar.KeyValueBundle) statevar.StateVariableResult {
	return &types.RiskFactor{
		Short: kvb.KVT[0].Val.(*statevar.DecimalScalar).Val,
		Long:  kvb.KVT[1].Val.(*statevar.DecimalScalar).Val,
	}
}

func (FactorConverter) InterfaceToBundle(res statevar.StateVariableResult) *statevar.KeyValueBundle {
	value := res.(*types.RiskFactor)
	return &statevar.KeyValueBundle{
		KVT: []statevar.KeyValueTol{
			{Key: "short", Val: &statevar.DecimalScalar{Val: value.Short}, Tolerance: riskFactorTolerance},
			{Key: "long", Val: &statevar.DecimalScalar{Val: value.Long}, Tolerance: riskFactorTolerance},
		},
	}
}

// startRiskFactorsCalculation kicks off the risk factors calculation, done asynchronously for illustration.
func (e *Engine) startRiskFactorsCalculation(eventID string, endOfCalcCallback statevar.FinaliseCalculation) {
	rf := e.model.CalculateRiskFactors()
	e.log.Info("risk factors calculated", logging.String("event-id", eventID), logging.Decimal("short", rf.Short), logging.Decimal("long", rf.Long))
	endOfCalcCallback.CalculationFinished(eventID, rf, nil)
}

// CalculateRiskFactorsForTest is a hack for testing for setting directly the risk factors for a market.
func (e *Engine) CalculateRiskFactorsForTest() {
	e.factors = e.model.CalculateRiskFactors()
	e.factors.Market = e.mktID
}

// updateRiskFactor sets the risk factor value to that of the decimal consensus value.
func (e *Engine) updateRiskFactor(ctx context.Context, res statevar.StateVariableResult) error {
	e.factors = res.(*types.RiskFactor)
	e.factors.Market = e.mktID
	e.riskFactorsInitialised = true
	e.log.Info("consensus reached for risk factors", logging.String("market", e.mktID), logging.Decimal("short", e.factors.Short), logging.Decimal("long", e.factors.Long))
	// then we can send in the broker
	e.broker.Send(events.NewRiskFactorEvent(ctx, *e.factors))
	return nil
}

func (e *Engine) IsRiskFactorInitialised() bool {
	return e.riskFactorsInitialised
}
