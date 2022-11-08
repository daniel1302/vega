// Copyright (c) 2022 Gobalsky Labs Limited
//
// Use of this software is governed by the Business Source License included
// in the LICENSE.DATANODE file and at https://www.mariadb.com/bsl11.
//
// Change Date: 18 months from the later of the date of the first publicly
// available Distribution of this version of the repository, and 25 June 2022.
//
// On the date above, in accordance with the Business Source License, use
// of this software will be governed by version 3 or later of the GNU General
// Public License.

package gql

import (
	"context"

	types "code.vegaprotocol.io/vega/protos/vega"
	vega "code.vegaprotocol.io/vega/protos/vega"
)

type myFutureProductResolver VegaResolverRoot

func (r *myFutureProductResolver) SettlementAsset(ctx context.Context, obj *types.FutureProduct) (*types.Asset, error) {
	return r.r.getAssetByID(ctx, obj.SettlementAsset)
}

func (r *myFutureProductResolver) DataSourceSpecForSettlementData(ctx context.Context, obj *vega.FutureProduct) (*DataSourceDefinition, error) {
	if obj.DataSourceSpecForSettlementData == nil {
		return nil, nil
	}

	dataSourceSpec := obj.DataSourceSpecForSettlementData
	return &DataSourceDefinition{
		SourceType: dataSourceSpec.SourceType.(DataSourceKind),
	}, nil
}

func (r *myFutureProductResolver) DataSourceSpecForTradingTermination(ctx context.Context, obj *vega.FutureProduct) (*DataSourceDefinition, error) {
	if obj.DataSourceSpecForTradingTermination == nil {
		return nil, nil
	}

	dataSourceSpec := obj.DataSourceSpecForTradingTermination
	return &DataSourceDefinition{
		SourceType: dataSourceSpec.SourceType.(DataSourceKind),
	}, nil
}
