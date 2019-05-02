package gql_test

import (
	"context"
	"testing"

	"code.vegaprotocol.io/vega/proto"
	types "code.vegaprotocol.io/vega/proto"

	"code.vegaprotocol.io/vega/internal/api"
	gql "code.vegaprotocol.io/vega/internal/gateway/graphql"
	"code.vegaprotocol.io/vega/internal/logging"
	"code.vegaprotocol.io/vega/internal/monitoring"

	"github.com/golang/mock/gomock"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

func TestNewResolverRoot_ConstructAndResolve(t *testing.T) {
	root := buildTestResolverRoot(t)
	defer root.Finish()
	assert.NotNil(t, root)

	partyResolver := root.Party()
	assert.NotNil(t, partyResolver)

	marketResolver := root.Market()
	assert.NotNil(t, marketResolver)

	depthResolver := root.MarketDepth()
	assert.NotNil(t, depthResolver)

	candleResolver := root.Candle()
	assert.NotNil(t, candleResolver)

	orderResolver := root.Order()
	assert.NotNil(t, orderResolver)

	tradeResolver := root.Trade()
	assert.NotNil(t, tradeResolver)

	priceLevelResolver := root.PriceLevel()
	assert.NotNil(t, priceLevelResolver)

	mutationResolver := root.Mutation()
	assert.NotNil(t, mutationResolver)

	positionResolver := root.Position()
	assert.NotNil(t, positionResolver)

	queryResolver := root.Query()
	assert.NotNil(t, queryResolver)

	subsResolver := root.Subscription()
	assert.NotNil(t, subsResolver)
}

func TestNewResolverRoot_QueryResolver(t *testing.T) {
	root := buildTestResolverRoot(t)
	defer root.Finish()
	assert.NotNil(t, root)

	queryResolver := root.Query()
	assert.NotNil(t, queryResolver)

}

func getTestMarket() *types.Market {
	return &types.Market{
		Id: "BTC/DEC19",
		TradableInstrument: &proto.TradableInstrument{
			Instrument: &proto.Instrument{
				Id:   "Crypto/BTCUSD/Futures/Dec19",
				Code: "FX:BTCUSD/DEC19",
				Name: "December 2019 BTC vs USD future",
				Metadata: &proto.InstrumentMetadata{
					Tags: []string{
						"asset_class:fx/crypto",
						"product:futures",
					},
				},
				Product: &proto.Instrument_Future{
					Future: &proto.Future{
						Maturity: "2019-12-31",
						Oracle: &proto.Future_EthereumEvent{
							EthereumEvent: &proto.EthereumEvent{
								ContractID: "0x0B484706fdAF3A4F24b2266446B1cb6d648E3cC1",
								Event:      "price_changed",
							},
						},
						Asset: "Ethereum/Ether",
					},
				},
			},
			RiskModel: &proto.TradableInstrument_BuiltinFutures{
				BuiltinFutures: &proto.BuiltinFutures{
					HistoricVolatility: 0.15,
				},
			},
		},
		TradingMode: &proto.Market_Continuous{
			Continuous: &proto.ContinuousTrading{},
		},
	}
}

func getTestParty() *types.Party {
	return &types.Party{
		Name: "barney",
	}
}

func TestNewResolverRoot_Resolver(t *testing.T) {
	root := buildTestResolverRoot(t)
	defer root.Finish()
	ctx := context.Background()

	marketNotExistsErr := errors.New("market does not exist")
	markets := map[string]*types.Market{
		"BTC/DEC19": getTestMarket(),
		"ETH/USD18": nil,
	}

	partyNotExistsErr := errors.New("party does not exist")
	parties := map[string]*types.Party{
		"barney": getTestParty(),
	}

	root.party.EXPECT().GetByID(gomock.Any(), gomock.Any()).Times(len(parties)).DoAndReturn(func(_ context.Context, k string) (*types.Party, error) {
		m, ok := parties[k]
		assert.True(t, ok)
		if m == nil {
			return nil, partyNotExistsErr
		}
		return m, nil
	})

	root.market.EXPECT().GetByID(gomock.Any(), gomock.Any()).Times(len(markets)).DoAndReturn(func(_ context.Context, k string) (*types.Market, error) {
		m, ok := markets[k]
		assert.True(t, ok)
		if m == nil {
			return nil, marketNotExistsErr
		}
		return m, nil
	})
	incompleteMarket := &types.Market{
		Id: "foobar",
	}
	root.market.EXPECT().GetAll(gomock.Any()).Times(1).Return([]*types.Market{incompleteMarket}, nil)

	name := "BTC/DEC19"
	vMarkets, err := root.Query().Markets(ctx, &name)
	assert.Nil(t, err)
	assert.NotNil(t, vMarkets)
	assert.Len(t, vMarkets, 1)

	name = "ETH/USD18"
	vMarkets, err = root.Query().Markets(ctx, &name)
	assert.Error(t, err)
	assert.Nil(t, vMarkets)

	vMarkets, err = root.Query().Markets(ctx, nil)
	assert.Error(t, err)
	assert.Nil(t, vMarkets)

	name = "barney"
	vParties, err := root.Query().Parties(ctx, &name)
	assert.Nil(t, err)
	assert.NotNil(t, vParties)
	assert.Len(t, vParties, 1)

	vParties, err = root.Query().Parties(ctx, nil)
	assert.Error(t, err)
	assert.Nil(t, vParties)
}

func TestNewResolverRoot_MarketResolver(t *testing.T) {
	root := buildTestResolverRoot(t)
	defer root.Finish()
	ctx := context.Background()

	notExistsErr := errors.New("market does not exist")
	markets := map[string]*types.Market{
		"BTC/DEC19": &types.Market{
			Id: "BTC/DEC19",
		},
	}
	marketId := "BTC/DEC19"
	market := &gql.Market{
		ID: marketId,
	}

	root.market.EXPECT().GetByID(gomock.Any(), gomock.Any()).Times(len(markets)).DoAndReturn(func(_ context.Context, k string) (*types.Market, error) {
		m, ok := markets[k]
		assert.True(t, ok)
		if m == nil {
			return nil, notExistsErr
		}
		return m, nil
	})
	var ui0 uint64
	root.order.EXPECT().GetByMarket(gomock.Any(), marketId, ui0, ui0, false, nil).Times(1).Return([]*types.Order{
		{
			Id:        "order-id-1",
			Price:     1000,
			CreatedAt: 1,
		},
		{
			Id:        "order-id-2",
			Price:     2000,
			CreatedAt: 2,
		},
	}, nil)

	marketResolver := root.Market()
	assert.NotNil(t, marketResolver)

	orders, err := marketResolver.Orders(ctx, market, nil, nil, nil, nil)
	assert.NotNil(t, orders)
	assert.Nil(t, err)
	assert.Len(t, orders, 2)
}

type resolverRoot interface {
	Query() gql.QueryResolver
	Mutation() gql.MutationResolver
	Candle() gql.CandleResolver
	MarketDepth() gql.MarketDepthResolver
	PriceLevel() gql.PriceLevelResolver
	Market() gql.MarketResolver
	Order() gql.OrderResolver
	Trade() gql.TradeResolver
	Position() gql.PositionResolver
	Party() gql.PartyResolver
	Subscription() gql.SubscriptionResolver
}

type testResolver struct {
	resolverRoot
	log  *logging.Logger
	ctrl *gomock.Controller
}

func buildTestResolverRoot(t *testing.T) *testResolver {
	ctrl := gomock.NewController(t)
	log := logging.NewTestLogger()
	conf := api.NewDefaultConfig()
	statusChecker := &monitoring.Status{}
	resolver := gql.NewResolverRoot(
		log,
		conf,
	)
	return &testResolver{
		resolverRoot: resolver,
		log:          log,
		ctrl:         ctrl,
	}
}

func (t *testResolver) Finish() {
	t.log.Sync()
	t.ctrl.Finish()
}