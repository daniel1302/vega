package gql

import (
	"context"
	"encoding/base64"
	"fmt"
	"io"
	"strconv"
	"time"

	"github.com/golang/protobuf/ptypes/empty"
	"google.golang.org/grpc"

	"code.vegaprotocol.io/vega/gateway"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/proto"
	types "code.vegaprotocol.io/vega/proto"
	protoapi "code.vegaprotocol.io/vega/proto/api"
	"code.vegaprotocol.io/vega/vegatime"

	"github.com/pkg/errors"
)

var (
	// ErrNilPendingOrder a pending order was nil when expected not to be
	ErrNilPendingOrder = errors.New("nil pending order")
	// ErrUnknownAccountType a account type specified does not exist
	ErrUnknownAccountType = errors.New("unknown account type")
	// ErrMissingIDOrReference is returned when neither id nor reference has been supplied in the query
	ErrMissingIDOrReference = errors.New("missing id or reference")
	// ErrInvalidVotesSubscription is returned if neither proposal ID nor party ID is specified
	ErrInvalidVotesSubscription = errors.New("invlid subscription, either proposal or party ID required")
	// ErrInvalidProposal is returned when invalid governance data is received by proposal resolver
	ErrInvalidProposal = errors.New("invalid proposal")
)

// TradingClient ...
//go:generate go run github.com/golang/mock/mockgen -destination mocks/trading_client_mock.go -package mocks code.vegaprotocol.io/vega/gateway/graphql TradingClient
type TradingClient interface {
	// prepare calls (unary-like calls)
	PrepareSubmitOrder(ctx context.Context, in *protoapi.SubmitOrderRequest, opts ...grpc.CallOption) (*protoapi.PrepareSubmitOrderResponse, error)
	PrepareAmendOrder(ctx context.Context, in *protoapi.AmendOrderRequest, opts ...grpc.CallOption) (*protoapi.PrepareAmendOrderResponse, error)
	PrepareCancelOrder(ctx context.Context, in *protoapi.CancelOrderRequest, opts ...grpc.CallOption) (*protoapi.PrepareCancelOrderResponse, error)
	PrepareProposal(ctx context.Context, in *protoapi.PrepareProposalRequest, opts ...grpc.CallOption) (*protoapi.PrepareProposalResponse, error)

	PrepareVote(ctx context.Context, in *protoapi.PrepareVoteRequest, opts ...grpc.CallOption) (*protoapi.PrepareVoteResponse, error)
	PrepareWithdraw(ctx context.Context, in *protoapi.PrepareWithdrawRequest, opts ...grpc.CallOption) (*protoapi.PrepareWithdrawResponse, error)
	// unary calls - writes
	SubmitTransaction(ctx context.Context, in *protoapi.SubmitTransactionRequest, opts ...grpc.CallOption) (*protoapi.SubmitTransactionResponse, error)
}

// TradingDataClient ...
//go:generate go run github.com/golang/mock/mockgen -destination mocks/trading_data_client_mock.go -package mocks code.vegaprotocol.io/vega/gateway/graphql TradingDataClient
type TradingDataClient interface {
	// orders
	OrdersByMarket(ctx context.Context, in *protoapi.OrdersByMarketRequest, opts ...grpc.CallOption) (*protoapi.OrdersByMarketResponse, error)
	OrderByReference(ctx context.Context, in *protoapi.OrderByReferenceRequest, opts ...grpc.CallOption) (*protoapi.OrderByReferenceResponse, error)
	OrdersByParty(ctx context.Context, in *protoapi.OrdersByPartyRequest, opts ...grpc.CallOption) (*protoapi.OrdersByPartyResponse, error)
	OrderByMarketAndID(ctx context.Context, in *protoapi.OrderByMarketAndIdRequest, opts ...grpc.CallOption) (*protoapi.OrderByMarketAndIdResponse, error)
	OrderByID(ctx context.Context, in *protoapi.OrderByIDRequest, opts ...grpc.CallOption) (*types.Order, error)
	OrderVersionsByID(ctx context.Context, in *protoapi.OrderVersionsByIDRequest, opts ...grpc.CallOption) (*protoapi.OrderVersionsResponse, error)
	// markets
	MarketByID(ctx context.Context, in *protoapi.MarketByIDRequest, opts ...grpc.CallOption) (*protoapi.MarketByIDResponse, error)
	Markets(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*protoapi.MarketsResponse, error)
	MarketDepth(ctx context.Context, in *protoapi.MarketDepthRequest, opts ...grpc.CallOption) (*protoapi.MarketDepthResponse, error)
	LastTrade(ctx context.Context, in *protoapi.LastTradeRequest, opts ...grpc.CallOption) (*protoapi.LastTradeResponse, error)
	MarketDataByID(ctx context.Context, in *protoapi.MarketDataByIDRequest, opts ...grpc.CallOption) (*protoapi.MarketDataByIDResponse, error)
	// parties
	PartyByID(ctx context.Context, in *protoapi.PartyByIDRequest, opts ...grpc.CallOption) (*protoapi.PartyByIDResponse, error)
	Parties(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*protoapi.PartiesResponse, error)
	// trades
	TradesByMarket(ctx context.Context, in *protoapi.TradesByMarketRequest, opts ...grpc.CallOption) (*protoapi.TradesByMarketResponse, error)
	TradesByParty(ctx context.Context, in *protoapi.TradesByPartyRequest, opts ...grpc.CallOption) (*protoapi.TradesByPartyResponse, error)
	TradesByOrder(ctx context.Context, in *protoapi.TradesByOrderRequest, opts ...grpc.CallOption) (*protoapi.TradesByOrderResponse, error)
	// positions
	PositionsByParty(ctx context.Context, in *protoapi.PositionsByPartyRequest, opts ...grpc.CallOption) (*protoapi.PositionsByPartyResponse, error)
	// candles
	Candles(ctx context.Context, in *protoapi.CandlesRequest, opts ...grpc.CallOption) (*protoapi.CandlesResponse, error)
	// metrics
	Statistics(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*types.Statistics, error)
	GetVegaTime(ctx context.Context, in *empty.Empty, opts ...grpc.CallOption) (*protoapi.VegaTimeResponse, error)
	// streams
	AccountsSubscribe(ctx context.Context, in *protoapi.AccountsSubscribeRequest, opts ...grpc.CallOption) (protoapi.TradingData_AccountsSubscribeClient, error)
	OrdersSubscribe(ctx context.Context, in *protoapi.OrdersSubscribeRequest, opts ...grpc.CallOption) (protoapi.TradingData_OrdersSubscribeClient, error)
	TradesSubscribe(ctx context.Context, in *protoapi.TradesSubscribeRequest, opts ...grpc.CallOption) (protoapi.TradingData_TradesSubscribeClient, error)
	CandlesSubscribe(ctx context.Context, in *protoapi.CandlesSubscribeRequest, opts ...grpc.CallOption) (protoapi.TradingData_CandlesSubscribeClient, error)
	MarketDepthSubscribe(ctx context.Context, in *protoapi.MarketDepthSubscribeRequest, opts ...grpc.CallOption) (protoapi.TradingData_MarketDepthSubscribeClient, error)
	PositionsSubscribe(ctx context.Context, in *protoapi.PositionsSubscribeRequest, opts ...grpc.CallOption) (protoapi.TradingData_PositionsSubscribeClient, error)
	MarketsDataSubscribe(ctx context.Context, in *protoapi.MarketsDataSubscribeRequest, opts ...grpc.CallOption) (protoapi.TradingData_MarketsDataSubscribeClient, error)
	MarginLevelsSubscribe(ctx context.Context, in *protoapi.MarginLevelsSubscribeRequest, opts ...grpc.CallOption) (protoapi.TradingData_MarginLevelsSubscribeClient, error)
	// accounts
	PartyAccounts(ctx context.Context, req *protoapi.PartyAccountsRequest, opts ...grpc.CallOption) (*protoapi.PartyAccountsResponse, error)
	MarketAccounts(ctx context.Context, req *protoapi.MarketAccountsRequest, opts ...grpc.CallOption) (*protoapi.MarketAccountsResponse, error)
	// margins
	MarginLevels(ctx context.Context, in *protoapi.MarginLevelsRequest, opts ...grpc.CallOption) (*protoapi.MarginLevelsResponse, error)
	// governance
	GetProposals(ctx context.Context, in *protoapi.GetProposalsRequest, opts ...grpc.CallOption) (*protoapi.GetProposalsResponse, error)
	GetProposalsByParty(ctx context.Context, in *protoapi.GetProposalsByPartyRequest, opts ...grpc.CallOption) (*protoapi.GetProposalsByPartyResponse, error)
	GetVotesByParty(ctx context.Context, in *protoapi.GetVotesByPartyRequest, opts ...grpc.CallOption) (*protoapi.GetVotesByPartyResponse, error)
	GetNewMarketProposals(ctx context.Context, in *protoapi.GetNewMarketProposalsRequest, opts ...grpc.CallOption) (*protoapi.GetNewMarketProposalsResponse, error)
	GetUpdateMarketProposals(ctx context.Context, in *protoapi.GetUpdateMarketProposalsRequest, opts ...grpc.CallOption) (*protoapi.GetUpdateMarketProposalsResponse, error)
	GetNetworkParametersProposals(ctx context.Context, in *protoapi.GetNetworkParametersProposalsRequest, opts ...grpc.CallOption) (*protoapi.GetNetworkParametersProposalsResponse, error)
	GetNewAssetProposals(ctx context.Context, in *protoapi.GetNewAssetProposalsRequest, opts ...grpc.CallOption) (*protoapi.GetNewAssetProposalsResponse, error)
	GetProposalByID(ctx context.Context, in *protoapi.GetProposalByIDRequest, opts ...grpc.CallOption) (*protoapi.GetProposalByIDResponse, error)
	GetProposalByReference(ctx context.Context, in *protoapi.GetProposalByReferenceRequest, opts ...grpc.CallOption) (*protoapi.GetProposalByReferenceResponse, error)

	ObserveGovernance(ctx context.Context, _ *empty.Empty, opts ...grpc.CallOption) (protoapi.TradingData_ObserveGovernanceClient, error)
	ObservePartyProposals(ctx context.Context, in *protoapi.ObservePartyProposalsRequest, opts ...grpc.CallOption) (protoapi.TradingData_ObservePartyProposalsClient, error)
	ObservePartyVotes(ctx context.Context, in *protoapi.ObservePartyVotesRequest, opts ...grpc.CallOption) (protoapi.TradingData_ObservePartyVotesClient, error)
	ObserveProposalVotes(ctx context.Context, in *protoapi.ObserveProposalVotesRequest, opts ...grpc.CallOption) (protoapi.TradingData_ObserveProposalVotesClient, error)
	GetNodeSignaturesAggregate(ctx context.Context, in *protoapi.GetNodeSignaturesAggregateRequest, opts ...grpc.CallOption) (*protoapi.GetNodeSignaturesAggregateResponse, error)
	AssetByID(ctx context.Context, in *protoapi.AssetByIDRequest, opts ...grpc.CallOption) (*protoapi.AssetByIDResponse, error)
	Assets(ctx context.Context, in *protoapi.AssetsRequest, opts ...grpc.CallOption) (*protoapi.AssetsResponse, error)
	FeeInfrastructureAccounts(ctx context.Context, in *protoapi.FeeInfrastructureAccountsRequest, opts ...grpc.CallOption) (*protoapi.FeeInfrastructureAccountsResponse, error)
	EstimateFee(ctx context.Context, in *protoapi.EstimateFeeRequest, opts ...grpc.CallOption) (*protoapi.EstimateFeeResponse, error)
	Withdrawal(ctx context.Context, in *protoapi.WithdrawalRequest, opts ...grpc.CallOption) (*protoapi.WithdrawalResponse, error)
	Withdrawals(ctx context.Context, in *protoapi.WithdrawalsRequest, opts ...grpc.CallOption) (*protoapi.WithdrawalsResponse, error)
	ERC20WithdrawalApproval(ctx context.Context, in *protoapi.ERC20WithdrawalApprovalRequest, opts ...grpc.CallOption) (*protoapi.ERC20WithdrawalApprovalResponse, error)
	Deposit(ctx context.Context, in *protoapi.DepositRequest, opts ...grpc.CallOption) (*protoapi.DepositResponse, error)
	Deposits(ctx context.Context, in *protoapi.DepositsRequest, opts ...grpc.CallOption) (*protoapi.DepositsResponse, error)

	ObserveEventBus(ctx context.Context, in *protoapi.ObserveEventsRequest, opts ...grpc.CallOption) (protoapi.TradingData_ObserveEventBusClient, error)
}

// VegaResolverRoot is the root resolver for all graphql types
type VegaResolverRoot struct {
	gateway.Config

	log               *logging.Logger
	tradingClient     TradingClient
	tradingDataClient TradingDataClient
}

// NewResolverRoot instantiate a graphql root resolver
func NewResolverRoot(
	log *logging.Logger,
	config gateway.Config,
	tradingClient TradingClient,
	tradingDataClient TradingDataClient,
) *VegaResolverRoot {

	return &VegaResolverRoot{
		log:               log,
		Config:            config,
		tradingClient:     tradingClient,
		tradingDataClient: tradingDataClient,
	}
}

// Query returns the query resolver
func (r *VegaResolverRoot) Query() QueryResolver {
	return (*myQueryResolver)(r)
}

// Mutation returns the mutations resolver
func (r *VegaResolverRoot) Mutation() MutationResolver {
	return (*myMutationResolver)(r)
}

// Candle returns the candles resolver
func (r *VegaResolverRoot) Candle() CandleResolver {
	return (*myCandleResolver)(r)
}

// MarketDepth returns the market depth resolver
func (r *VegaResolverRoot) MarketDepth() MarketDepthResolver {
	return (*myMarketDepthResolver)(r)
}

// MarketData returns the market data resolver
func (r *VegaResolverRoot) MarketData() MarketDataResolver {
	return (*myMarketDataResolver)(r)
}

// MarginLevels returns the market levels resolver
func (r *VegaResolverRoot) MarginLevels() MarginLevelsResolver {
	return (*myMarginLevelsResolver)(r)
}

// PriceLevel returns the price levels resolver
func (r *VegaResolverRoot) PriceLevel() PriceLevelResolver {
	return (*myPriceLevelResolver)(r)
}

// Market returns the markets resolver
func (r *VegaResolverRoot) Market() MarketResolver {
	return (*myMarketResolver)(r)
}

// Order returns the order resolver
func (r *VegaResolverRoot) Order() OrderResolver {
	return (*myOrderResolver)(r)
}

// Trade returns the trades resolver
func (r *VegaResolverRoot) Trade() TradeResolver {
	return (*myTradeResolver)(r)
}

// Position returns the positions resolver
func (r *VegaResolverRoot) Position() PositionResolver {
	return (*myPositionResolver)(r)
}

// Party returns the parties resolver
func (r *VegaResolverRoot) Party() PartyResolver {
	return (*myPartyResolver)(r)
}

// Subscription returns the subscriptions resolver
func (r *VegaResolverRoot) Subscription() SubscriptionResolver {
	return (*mySubscriptionResolver)(r)
}

// Account returns the accounts resolver
func (r *VegaResolverRoot) Account() AccountResolver {
	return (*myAccountResolver)(r)
}

// Statistics returns the statistics resolver
func (r *VegaResolverRoot) Statistics() StatisticsResolver {
	return (*myStatisticsResolver)(r)
}

// Proposal returns the proposal resolver
func (r *VegaResolverRoot) Proposal() ProposalResolver {
	return (*myProposalResolver)(r)
}

// NodeSignature ...
func (r *VegaResolverRoot) NodeSignature() NodeSignatureResolver {
	return (*myNodeSignatureResolver)(r)
}

// Asset ...
func (r *VegaResolverRoot) Asset() AssetResolver {
	return (*myAssetResolver)(r)
}

// Deposit ...
func (r *VegaResolverRoot) Deposit() DepositResolver {
	return (*myDepositResolver)(r)
}

// deposit resolver

type myDepositResolver VegaResolverRoot

func (r *myDepositResolver) Asset(ctx context.Context, obj *types.Deposit) (*Asset, error) {
	if len(obj.Asset) <= 0 {
		return nil, errors.New("missign asset ID")
	}
	req := &protoapi.AssetByIDRequest{
		ID: obj.Asset,
	}
	res, err := r.tradingDataClient.AssetByID(ctx, req)
	if err != nil {
		return nil, err
	}
	return AssetFromProto(res.Asset)
}

func (r *myDepositResolver) Party(ctx context.Context, obj *types.Deposit) (*types.Party, error) {
	if len(obj.PartyID) <= 0 {
		return nil, errors.New("missing party ID")
	}
	return &types.Party{Id: obj.PartyID}, nil
}

func (r *myDepositResolver) CreatedTimestamp(ctx context.Context, obj *types.Deposit) (string, error) {
	if obj.CreatedTimestamp == 0 {
		return "", errors.New("invalid timestamp")
	}
	return vegatime.Format(vegatime.UnixNano(obj.CreatedTimestamp)), nil
}

func (r *myDepositResolver) CreditedTimestamp(ctx context.Context, obj *types.Deposit) (*string, error) {
	if obj.CreatedTimestamp == 0 {
		return nil, nil
	}
	t := vegatime.Format(vegatime.UnixNano(obj.CreatedTimestamp))
	return &t, nil
}

func (r *myDepositResolver) Status(ctx context.Context, obj *types.Deposit) (DepositStatus, error) {
	return convertDepositStatusFromProto(obj.Status)
}

// asset resolver

type myAssetResolver VegaResolverRoot

func (r *myAssetResolver) InfrastructureFeeAccount(ctx context.Context, obj *Asset) (*proto.Account, error) {
	if len(obj.ID) <= 0 {
		return nil, ErrMissingIDOrReference
	}
	req := &protoapi.FeeInfrastructureAccountsRequest{
		Asset: obj.ID,
	}
	res, err := r.tradingDataClient.FeeInfrastructureAccounts(ctx, req)
	if err != nil {
		return nil, err
	}

	var acc *types.Account
	if len(res.Accounts) > 0 {
		acc = res.Accounts[0]
	}

	return acc, nil
}

// BEGIN: Query Resolver

type myQueryResolver VegaResolverRoot

func (r *myQueryResolver) Erc20WithdrawalApproval(ctx context.Context, wid string) (*Erc20WithdrawalApproval, error) {
	res, err := r.tradingDataClient.ERC20WithdrawalApproval(
		ctx, &protoapi.ERC20WithdrawalApprovalRequest{WithdrawalID: wid},
	)
	if err != nil {
		return nil, err
	}

	return &Erc20WithdrawalApproval{
		AssetSource: res.AssetSource,
		Amount:      res.Amount,
		Expiry:      vegatime.Format(vegatime.UnixNano(res.Expiry)),
		Nonce:       res.Nonce,
		Signatures:  res.Signatures,
	}, nil
}

func (r *myQueryResolver) Withdrawal(ctx context.Context, wid string) (*Withdrawal, error) {
	res, err := r.tradingDataClient.Withdrawal(
		ctx, &protoapi.WithdrawalRequest{ID: wid},
	)
	if err != nil {
		return nil, err
	}

	w, err := NewWithdrawalFromProto(res.Withdrawal)
	if err != nil {
		return nil, err
	}

	asset, err := r.Asset(ctx, res.Withdrawal.Asset)
	if err != nil {
		return nil, err
	}

	w.Asset = asset
	return w, nil
}

func (r *myQueryResolver) Deposit(ctx context.Context, did string) (*types.Deposit, error) {
	res, err := r.tradingDataClient.Deposit(
		ctx, &protoapi.DepositRequest{ID: did},
	)
	if err != nil {
		return nil, err
	}

	return res.Deposit, nil
}

func (r *myQueryResolver) EstimateFeeForOrder(ctx context.Context, market, party string, price *string, size string, side Side,
	timeInForce OrderTimeInForce, expiration *string, ty OrderType) (*OrderFeeEstimate, error) {
	order := &types.Order{}

	var (
		p   uint64
		err error
	)

	// We need to convert strings to uint64 (JS doesn't yet support uint64)
	if price != nil {
		p, err = safeStringUint64(*price)
		if err != nil {
			return nil, err
		}
	}
	order.Price = p
	s, err := safeStringUint64(size)
	if err != nil {
		return nil, err
	}
	order.Size = s
	if len(market) <= 0 {
		return nil, errors.New("market missing or empty")
	}
	order.MarketID = market
	if len(party) <= 0 {
		return nil, errors.New("party missing or empty")
	}

	order.PartyID = party
	if order.TimeInForce, err = convertOrderTimeInForceToProto(timeInForce); err != nil {
		return nil, err
	}
	if order.Side, err = convertSideToProto(side); err != nil {
		return nil, err
	}
	if order.Type, err = convertOrderTypeToProto(ty); err != nil {
		return nil, err
	}

	// GTT must have an expiration value
	if order.TimeInForce == types.Order_TIF_GTT && expiration != nil {
		var expiresAt time.Time
		expiresAt, err = vegatime.Parse(*expiration)
		if err != nil {
			return nil, fmt.Errorf("cannot parse expiration time: %s - invalid format sent to create order (example: 2018-01-02T15:04:05Z)", *expiration)
		}

		// move to pure timestamps or convert an RFC format shortly
		order.ExpiresAt = expiresAt.UnixNano()
	}

	req := protoapi.EstimateFeeRequest{
		Order: order,
	}

	// Pass the order over for consensus (service layer will use RPC client internally and handle errors etc)
	resp, err := r.tradingDataClient.EstimateFee(ctx, &req)
	if err != nil {
		r.log.Error("Failed to create order using rpc client in graphQL resolver", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}

	// calclate the fee total amount
	ttf := resp.Fee.MakerFee + resp.Fee.InfrastructureFee + resp.Fee.LiquidityFee

	fee := TradeFee{
		MakerFee:          fmt.Sprintf("%d", resp.Fee.MakerFee),
		InfrastructureFee: fmt.Sprintf("%d", resp.Fee.InfrastructureFee),
		LiquidityFee:      fmt.Sprintf("%d", resp.Fee.LiquidityFee),
	}

	return &OrderFeeEstimate{
		Fee:            &fee,
		TotalFeeAmount: fmt.Sprintf("%d", ttf),
	}, nil

}

func (r *myQueryResolver) Asset(ctx context.Context, id string) (*Asset, error) {
	if len(id) <= 0 {
		return nil, ErrMissingIDOrReference
	}
	req := &protoapi.AssetByIDRequest{
		ID: id,
	}
	res, err := r.tradingDataClient.AssetByID(ctx, req)
	if err != nil {
		return nil, err
	}
	return AssetFromProto(res.Asset)
}

func (r *myQueryResolver) Assets(ctx context.Context) ([]*Asset, error) {
	req := &protoapi.AssetsRequest{}
	res, err := r.tradingDataClient.Assets(ctx, req)
	if err != nil {
		return nil, err
	}
	out := make([]*Asset, 0, len(res.Assets))
	for _, v := range res.Assets {
		a, err := AssetFromProto(v)
		if err != nil {
			return nil, err
		}
		out = append(out, a)
	}

	return out, nil

}

func (r *myQueryResolver) NodeSignatures(ctx context.Context, resourceID string) ([]*types.NodeSignature, error) {
	if len(resourceID) <= 0 {
		return nil, ErrMissingIDOrReference
	}

	req := &protoapi.GetNodeSignaturesAggregateRequest{
		ID: resourceID,
	}
	res, err := r.tradingDataClient.GetNodeSignaturesAggregate(ctx, req)
	if err != nil {
		return nil, err
	}
	return res.Signatures, nil
}

func (r *myQueryResolver) Markets(ctx context.Context, id *string) ([]*Market, error) {
	if id != nil {
		mkt, err := r.Market(ctx, *id)
		if err != nil {
			return nil, err
		}
		return []*Market{mkt}, nil
	}
	res, err := r.tradingDataClient.Markets(ctx, &empty.Empty{})
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}

	m := make([]*Market, 0, len(res.Markets))
	for _, pmarket := range res.Markets {
		market, err := MarketFromProto(pmarket)
		if err != nil {
			r.log.Error("unable to convert market from proto", logging.Error(err))
			return nil, err
		}
		// add the market name explicitly here as it does not
		// come anymore from the market framework
		market.Name = market.TradableInstrument.Instrument.Name
		// set Asset here too as well
		switch p := market.TradableInstrument.Instrument.Product.(type) {
		case *Future:
			req := protoapi.AssetByIDRequest{ID: p.Asset.ID}
			res, err := r.tradingDataClient.AssetByID(context.Background(), &req)
			if err != nil {
				r.log.Error("tradingData client", logging.Error(err))
				return nil, customErrorFromStatus(err)
			}
			p.Asset, err = AssetFromProto(res.Asset)
		}

		m = append(m, market)
	}

	return m, nil
}

func (r *myQueryResolver) Market(ctx context.Context, id string) (*Market, error) {
	req := protoapi.MarketByIDRequest{MarketID: id}
	res, err := r.tradingDataClient.MarketByID(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	market, err := MarketFromProto(res.Market)
	if err != nil {
		r.log.Error("unable to convert market from proto", logging.Error(err))
		return nil, err
	}
	// add the market name explicitly here as it does not
	// come anymore from the market framework
	market.Name = market.TradableInstrument.Instrument.Name
	// set Asset here too
	switch p := market.TradableInstrument.Instrument.Product.(type) {
	case *Future:
		req := protoapi.AssetByIDRequest{ID: p.Asset.ID}
		res, err := r.tradingDataClient.AssetByID(context.Background(), &req)
		if err != nil {
			r.log.Error("tradingData client", logging.Error(err))
			return nil, customErrorFromStatus(err)
		}
		p.Asset, err = AssetFromProto(res.Asset)
	}

	return market, nil
}

func (r *myQueryResolver) Parties(ctx context.Context, name *string) ([]*types.Party, error) {
	if name == nil {
		var empty empty.Empty
		resp, err := r.tradingDataClient.Parties(ctx, &empty)
		if err != nil {
			return nil, err
		}
		if resp.Parties == nil {
			return []*types.Party{}, nil
		}
		return resp.Parties, nil
	}
	party, err := r.Party(ctx, *name)
	if err != nil {
		return nil, err
	}

	// if we asked for a single party it may be null
	// so then we return an empty slice
	if party == nil {
		return []*types.Party{}, nil
	}

	return []*types.Party{party}, nil
}

func (r *myQueryResolver) Party(ctx context.Context, name string) (*types.Party, error) {
	return getParty(ctx, r.log, r.tradingDataClient, name)
}

func (r *myQueryResolver) Statistics(ctx context.Context) (*types.Statistics, error) {
	res, err := r.tradingDataClient.Statistics(ctx, &empty.Empty{})
	if err != nil {
		r.log.Error("tradingCore client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	return res, nil
}

func (r *myQueryResolver) OrderByID(ctx context.Context, orderID string, version *int) (*types.Order, error) {
	v, err := convertVersion(version)
	if err != nil {
		r.log.Error("tradingCore client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	orderReq := &protoapi.OrderByIDRequest{
		OrderID: orderID,
		Version: v,
	}
	order, err := r.tradingDataClient.OrderByID(ctx, orderReq)
	return order, err
}

func (r *myQueryResolver) OrderVersions(
	ctx context.Context, orderID string, skip, first, last *int) ([]*types.Order, error) {

	p := makePagination(skip, first, last)
	reqest := &protoapi.OrderVersionsByIDRequest{
		OrderID:    orderID,
		Pagination: p,
	}
	res, err := r.tradingDataClient.OrderVersionsByID(ctx, reqest)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	return res.Orders, nil
}

func (r *myQueryResolver) OrderByReference(ctx context.Context, reference string) (*types.Order, error) {
	req := &protoapi.OrderByReferenceRequest{
		Reference: reference,
	}
	res, err := r.tradingDataClient.OrderByReference(ctx, req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	return res.Order, err
}

func (r *myQueryResolver) Proposals(ctx context.Context, inState *ProposalState) ([]*types.GovernanceData, error) {
	filter, err := inState.ToOptionalProposalState()
	if err != nil {
		return nil, err
	}
	resp, err := r.tradingDataClient.GetProposals(ctx, &protoapi.GetProposalsRequest{
		SelectInState: filter,
	})
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *myQueryResolver) Proposal(ctx context.Context, id *string, reference *string) (*types.GovernanceData, error) {
	if id != nil {
		resp, err := r.tradingDataClient.GetProposalByID(ctx, &protoapi.GetProposalByIDRequest{
			ProposalID: *id,
		})
		if err != nil {
			return nil, err
		}
		return resp.Data, nil
	} else if reference != nil {
		resp, err := r.tradingDataClient.GetProposalByReference(ctx, &protoapi.GetProposalByReferenceRequest{
			Reference: *reference,
		})
		if err != nil {
			return nil, err
		}
		return resp.Data, nil
	}

	return nil, ErrMissingIDOrReference
}

func (r *myQueryResolver) NewMarketProposals(ctx context.Context, inState *ProposalState) ([]*types.GovernanceData, error) {
	filter, err := inState.ToOptionalProposalState()
	if err != nil {
		return nil, err
	}
	resp, err := r.tradingDataClient.GetNewMarketProposals(ctx, &protoapi.GetNewMarketProposalsRequest{
		SelectInState: filter,
	})
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *myQueryResolver) UpdateMarketProposals(ctx context.Context, marketID *string, inState *ProposalState) ([]*types.GovernanceData, error) {
	filter, err := inState.ToOptionalProposalState()
	if err != nil {
		return nil, err
	}
	var market string
	if marketID != nil {
		market = *marketID
	}
	resp, err := r.tradingDataClient.GetUpdateMarketProposals(ctx, &protoapi.GetUpdateMarketProposalsRequest{
		MarketID:      market,
		SelectInState: filter,
	})
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *myQueryResolver) NetworkParametersProposals(ctx context.Context, inState *ProposalState) ([]*types.GovernanceData, error) {
	filter, err := inState.ToOptionalProposalState()
	if err != nil {
		return nil, err
	}
	resp, err := r.tradingDataClient.GetNetworkParametersProposals(ctx, &protoapi.GetNetworkParametersProposalsRequest{
		SelectInState: filter,
	})
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *myQueryResolver) NewAssetProposals(ctx context.Context, inState *ProposalState) ([]*types.GovernanceData, error) {
	filter, err := inState.ToOptionalProposalState()
	if err != nil {
		return nil, err
	}
	resp, err := r.tradingDataClient.GetNewAssetProposals(ctx, &protoapi.GetNewAssetProposalsRequest{
		SelectInState: filter,
	})
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

// END: Root Resolver

// BEGIN: Market Resolver

type myMarketResolver VegaResolverRoot

func (r *myMarketResolver) Data(ctx context.Context, market *Market) (*types.MarketData, error) {
	req := protoapi.MarketDataByIDRequest{
		MarketID: market.ID,
	}
	res, err := r.tradingDataClient.MarketDataByID(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	return res.MarketData, nil
}

func (r *myMarketResolver) Orders(ctx context.Context, market *Market,
	skip, first, last *int) ([]*types.Order, error) {
	p := makePagination(skip, first, last)
	req := protoapi.OrdersByMarketRequest{
		MarketID:   market.ID,
		Pagination: p,
	}
	res, err := r.tradingDataClient.OrdersByMarket(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	return res.Orders, nil
}

func (r *myMarketResolver) Trades(ctx context.Context, market *Market,
	skip, first, last *int) ([]*types.Trade, error) {
	p := makePagination(skip, first, last)
	req := protoapi.TradesByMarketRequest{
		MarketID:   market.ID,
		Pagination: p,
	}
	res, err := r.tradingDataClient.TradesByMarket(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}

	return res.Trades, nil
}

func (r *myMarketResolver) Depth(ctx context.Context, market *Market, maxDepth *int) (*types.MarketDepth, error) {

	if market == nil {
		return nil, errors.New("market missing or empty")
	}

	req := protoapi.MarketDepthRequest{MarketID: market.ID}
	if maxDepth != nil {
		if *maxDepth <= 0 {
			return nil, errors.New("invalid maxDepth, must be a positive number")
		}
		req.MaxDepth = uint64(*maxDepth)
	}

	// Look for market depth for the given market (will validate market internally)
	// Note: Market depth is also known as OrderBook depth within the matching-engine
	res, err := r.tradingDataClient.MarketDepth(ctx, &req)
	if err != nil {
		r.log.Error("trading data client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}

	return &types.MarketDepth{
		MarketID: res.MarketID,
		Buy:      res.Buy,
		Sell:     res.Sell,
	}, nil
}

func (r *myMarketResolver) Candles(ctx context.Context, market *Market,
	sinceRaw string, interval Interval) ([]*types.Candle, error) {
	pinterval, err := convertIntervalToProto(interval)
	if err != nil {
		r.log.Debug("interval convert error", logging.Error(err))
	}

	since, err := vegatime.Parse(sinceRaw)
	if err != nil {
		return nil, err
	}

	var mkt string
	if market != nil {
		mkt = market.ID
	}

	req := protoapi.CandlesRequest{
		MarketID:       mkt,
		SinceTimestamp: since.UnixNano(),
		Interval:       pinterval,
	}
	res, err := r.tradingDataClient.Candles(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	return res.Candles, nil
}

// Accounts ...
// if partyID specified get margin account for the given market
// if nil return the insurance pool for the market
func (r *myMarketResolver) Accounts(ctx context.Context, market *Market, partyID *string) ([]*types.Account, error) {
	// get margin account for a party
	if partyID != nil {
		req := protoapi.PartyAccountsRequest{
			PartyID:  *partyID,
			MarketID: market.ID,
			Type:     types.AccountType_ACCOUNT_TYPE_MARGIN,
			Asset:    "",
		}
		res, err := r.tradingDataClient.PartyAccounts(ctx, &req)
		if err != nil {
			r.log.Error("unable to get PartyAccounts",
				logging.Error(err),
				logging.String("market-id", market.ID),
				logging.String("party-id", *partyID))
			return []*types.Account{}, customErrorFromStatus(err)
		}
		return res.Accounts, nil
	}
	// get accounts for the market
	req := protoapi.MarketAccountsRequest{
		MarketID: market.ID,
		Asset:    "", // all assets
	}
	res, err := r.tradingDataClient.MarketAccounts(ctx, &req)
	if err != nil {
		r.log.Error("unable to get MarketAccounts",
			logging.Error(err),
			logging.String("market-id", market.ID))
		return []*types.Account{}, customErrorFromStatus(err)
	}
	return res.Accounts, nil
}

// END: Market Resolver

type myNodeSignatureResolver VegaResolverRoot

func (r *myNodeSignatureResolver) Signature(ctx context.Context, obj *types.NodeSignature) (*string, error) {
	sig := base64.StdEncoding.EncodeToString(obj.Sig)
	return &sig, nil
}

func (r *myNodeSignatureResolver) Kind(ctx context.Context, obj *types.NodeSignature) (*NodeSignatureKind, error) {
	kind, err := convertNodeSignatureKindFromProto(obj.Kind)
	if err != nil {
		return nil, err
	}
	return &kind, nil
}

// BEGIN: Party Resolver

type myPartyResolver VegaResolverRoot

func makePagination(skip, first, last *int) *protoapi.Pagination {
	var (
		offset, limit uint64
		descending    bool
	)
	if skip != nil {
		offset = uint64(*skip)
	}
	if last != nil {
		limit = uint64(*last)
		descending = true
	} else if first != nil {
		limit = uint64(*first)
	}
	return &protoapi.Pagination{
		Skip:       offset,
		Limit:      limit,
		Descending: descending,
	}
}

func (r *myPartyResolver) Margins(ctx context.Context,
	party *types.Party, marketID *string) ([]*types.MarginLevels, error) {

	var marketId string
	if marketID != nil {
		marketId = *marketID
	}
	req := protoapi.MarginLevelsRequest{
		PartyID:  party.Id,
		MarketID: marketId,
	}
	res, err := r.tradingDataClient.MarginLevels(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	out := make([]*types.MarginLevels, 0, len(res.MarginLevels))
	out = append(out, res.MarginLevels...)
	return out, nil
}

func (r *myPartyResolver) Orders(ctx context.Context, party *types.Party,
	skip, first, last *int) ([]*types.Order, error) {

	p := makePagination(skip, first, last)
	req := protoapi.OrdersByPartyRequest{
		PartyID:    party.Id,
		Pagination: p,
	}
	res, err := r.tradingDataClient.OrdersByParty(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}

	if len(res.Orders) > 0 {
		return res.Orders, nil
	}
	// mandatory return field in schema
	return []*types.Order{}, nil
}

func (r *myPartyResolver) Trades(ctx context.Context, party *types.Party,
	market *string, skip, first, last *int) ([]*types.Trade, error) {

	var mkt string
	if market != nil {
		mkt = *market
	}

	p := makePagination(skip, first, last)
	req := protoapi.TradesByPartyRequest{
		PartyID:    party.Id,
		MarketID:   mkt,
		Pagination: p,
	}

	res, err := r.tradingDataClient.TradesByParty(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}

	if len(res.Trades) > 0 {
		return res.Trades, nil
	}
	// mandatory return field in schema
	return []*types.Trade{}, nil
}

func (r *myPartyResolver) Positions(ctx context.Context, party *types.Party) ([]*types.Position, error) {
	if party == nil {
		return nil, errors.New("nil party")
	}
	req := protoapi.PositionsByPartyRequest{PartyID: party.Id}
	res, err := r.tradingDataClient.PositionsByParty(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	if len(res.Positions) > 0 {
		return res.Positions, nil
	}
	// mandatory return field in schema
	return []*types.Position{}, nil
}

func (r *myPartyResolver) Accounts(ctx context.Context, party *types.Party,
	marketID *string, asset *string, accType *AccountType) ([]*types.Account, error) {
	if party == nil {
		return nil, errors.New("a party must be specified when querying accounts")
	}
	var (
		mktid = ""
		asst  = ""
		accTy = types.AccountType_ACCOUNT_TYPE_UNSPECIFIED
		err   error
	)

	if marketID != nil {
		mktid = *marketID
	}
	if asset != nil {
		asst = *asset
	}
	if accType != nil {
		accTy, err = convertAccountTypeToProto(*accType)
		if err != nil || (accTy != types.AccountType_ACCOUNT_TYPE_GENERAL && accTy != types.AccountType_ACCOUNT_TYPE_MARGIN) {
			return nil, fmt.Errorf("invalid account type for party %v", accType)
		}
	}
	req := protoapi.PartyAccountsRequest{
		PartyID:  party.Id,
		MarketID: mktid,
		Asset:    asst,
		Type:     accTy,
	}
	res, err := r.tradingDataClient.PartyAccounts(ctx, &req)
	if err != nil {
		r.log.Error("unable to get Party account",
			logging.Error(err),
			logging.String("party-id", party.Id),
			logging.String("market-id", mktid),
			logging.String("asset", asst),
			logging.String("type", accTy.String()))
		return nil, customErrorFromStatus(err)
	}

	if len(res.Accounts) > 0 {
		return res.Accounts, nil
	}
	// mandatory return field in schema
	return []*types.Account{}, nil
}

func (r *myPartyResolver) Proposals(ctx context.Context, party *types.Party, inState *ProposalState) ([]*types.GovernanceData, error) {
	filter, err := inState.ToOptionalProposalState()
	if err != nil {
		return nil, err
	}
	resp, err := r.tradingDataClient.GetProposalsByParty(ctx, &protoapi.GetProposalsByPartyRequest{
		PartyID:       party.Id,
		SelectInState: filter,
	})
	if err != nil {
		return nil, err
	}
	return resp.Data, nil
}

func (r *myPartyResolver) Withdrawals(ctx context.Context, party *types.Party) ([]*Withdrawal, error) {
	res, err := r.tradingDataClient.Withdrawals(
		ctx, &protoapi.WithdrawalsRequest{PartyID: party.Id},
	)
	if err != nil {
		return nil, err
	}

	out := make([]*Withdrawal, 0, len(res.Withdrawals))
	for _, v := range res.Withdrawals {
		w, err := NewWithdrawalFromProto(v)
		if err != nil {
			return nil, err
		}
		req := protoapi.AssetByIDRequest{ID: v.Asset}
		res, err := r.tradingDataClient.AssetByID(ctx, &req)
		if err != nil {
			r.log.Error("tradingData client", logging.Error(err))
			return nil, err
		}
		w.Asset, err = AssetFromProto(res.Asset)
		if err != nil {
			return nil, err
		}
		out = append(out, w)
	}

	return out, nil
}

func (r *myPartyResolver) Deposits(ctx context.Context, party *types.Party) ([]*types.Deposit, error) {
	res, err := r.tradingDataClient.Deposits(
		ctx, &protoapi.DepositsRequest{PartyID: party.Id},
	)
	if err != nil {
		return nil, err
	}

	return res.Deposits, nil
}

func (r *myPartyResolver) Votes(ctx context.Context, party *types.Party) ([]*ProposalVote, error) {
	resp, err := r.tradingDataClient.GetVotesByParty(ctx, &protoapi.GetVotesByPartyRequest{
		PartyID: party.Id,
	})
	if err != nil {
		return nil, err
	}
	result := make([]*ProposalVote, len(resp.Votes))
	for i, vote := range resp.Votes {
		//TODO: voter might be shortcut to party once all testing is done, vote.PartyID should be party.ID
		voter, err := getParty(ctx, r.log, r.tradingDataClient, vote.PartyID)
		if err != nil {
			return nil, err
		}
		result[i] = ProposalVoteFromProto(vote, voter)
	}
	return result, nil
}

// END: Party Resolver

// BEGIN: Proposal Resolver

type myProposalResolver VegaResolverRoot

func (r *myProposalResolver) RejectionReason(_ context.Context, data *types.GovernanceData) (*ProposalRejectionReason, error) {
	if data == nil || data.Proposal == nil {
		return nil, ErrInvalidProposal
	}
	p := data.Proposal
	if p.Reason == types.ProposalError_PROPOSAL_ERROR_UNSPECIFIED {
		return nil, nil
	}

	reason, err := convertProposalRejectionReasonFromProto(p.Reason)
	if err != nil {
		return nil, err
	}
	return &reason, nil
}

func (r *myProposalResolver) ID(ctx context.Context, data *types.GovernanceData) (*string, error) {
	if data == nil || data.Proposal == nil {
		return nil, ErrInvalidProposal
	}
	return &data.Proposal.ID, nil
}

func (r *myProposalResolver) Reference(ctx context.Context, data *types.GovernanceData) (string, error) {
	if data == nil || data.Proposal == nil {
		return "", ErrInvalidProposal
	}
	return data.Proposal.Reference, nil
}

func (r *myProposalResolver) Party(ctx context.Context, data *types.GovernanceData) (*types.Party, error) {
	if data == nil || data.Proposal == nil {
		return nil, ErrInvalidProposal
	}
	p, err := getParty(ctx, r.log, r.tradingDataClient, data.Proposal.PartyID)
	if p == nil && err == nil {
		// the api could return an nil party in some cases
		// e.g: when a party does not exists in the stores
		// this is not an error, but here we are not checking
		// if a party exists or not, but what party did propose
		p = &types.Party{Id: data.Proposal.PartyID}
	}
	return p, err
}

func (r *myProposalResolver) State(ctx context.Context, data *types.GovernanceData) (ProposalState, error) {
	if data == nil || data.Proposal == nil {
		return "", ErrInvalidProposal
	}
	return convertProposalStateFromProto(data.Proposal.State)
}

func (r *myProposalResolver) Datetime(ctx context.Context, data *types.GovernanceData) (string, error) {
	if data == nil || data.Proposal == nil {
		return "", ErrInvalidProposal
	}
	if data.Proposal.Timestamp == 0 {
		// no timestamp for prepared proposals
		return "", nil
	}
	return nanoTSToDatetime(data.Proposal.Timestamp), nil
}

func (r *myProposalResolver) Terms(ctx context.Context, data *types.GovernanceData) (*ProposalTerms, error) {
	if data == nil || data.Proposal == nil {
		return nil, ErrInvalidProposal
	}
	return ProposalTermsFromProto(data.Proposal.Terms)
}

func (r *myProposalResolver) convertVotes(ctx context.Context, data []*types.Vote) ([]*Vote, error) {
	result := make([]*Vote, len(data))
	for i, v := range data {
		voter, err := getParty(ctx, r.log, r.tradingDataClient, v.PartyID)
		if err != nil {
			return nil, err
		}
		value, err := convertVoteValueFromProto(v.Value)
		if err != nil {
			return nil, err
		}
		result[i] = &Vote{
			Value:    value,
			Party:    voter,
			Datetime: nanoTSToDatetime(v.Timestamp),
		}
	}
	return result, nil
}

func (r *myProposalResolver) YesVotes(ctx context.Context, data *types.GovernanceData) ([]*Vote, error) {
	if data == nil || data.Proposal == nil {
		return nil, ErrInvalidProposal
	}
	return r.convertVotes(ctx, data.Yes)
}

func (r *myProposalResolver) NoVotes(ctx context.Context, data *types.GovernanceData) ([]*Vote, error) {
	if data == nil || data.Proposal == nil {
		return nil, ErrInvalidProposal
	}
	return r.convertVotes(ctx, data.No)
}

// END: Proposal Resolver

// BEGIN: MarginLevels Resolver

type myMarginLevelsResolver VegaResolverRoot

func (r *myMarginLevelsResolver) Market(ctx context.Context, m *types.MarginLevels) (*Market, error) {
	req := protoapi.MarketByIDRequest{MarketID: m.MarketID}
	res, err := r.tradingDataClient.MarketByID(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}

	market, err := MarketFromProto(res.Market)
	if err != nil {
		r.log.Error("unable to convert market from proto", logging.Error(err))
		return nil, err
	}
	// set Asset here too as well
	switch p := market.TradableInstrument.Instrument.Product.(type) {
	case *Future:
		req := protoapi.AssetByIDRequest{ID: p.Asset.ID}
		res, err := r.tradingDataClient.AssetByID(context.Background(), &req)
		if err != nil {
			r.log.Error("tradingData client", logging.Error(err))
			return nil, customErrorFromStatus(err)
		}
		p.Asset, err = AssetFromProto(res.Asset)
	}

	return market, nil
}

func (r *myMarginLevelsResolver) Party(ctx context.Context, m *types.MarginLevels) (*types.Party, error) {
	if m == nil {
		return nil, errors.New("nil order")
	}
	if len(m.PartyID) == 0 {
		return nil, errors.New("invalid party")
	}
	req := protoapi.PartyByIDRequest{PartyID: m.PartyID}
	res, err := r.tradingDataClient.PartyByID(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	return res.Party, nil
}

func (r *myMarginLevelsResolver) Asset(ctx context.Context, m *types.MarginLevels) (*Asset, error) {
	req := protoapi.AssetByIDRequest{ID: m.Asset}
	res, err := r.tradingDataClient.AssetByID(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	return AssetFromProto(res.Asset)
}

func (r *myMarginLevelsResolver) CollateralReleaseLevel(_ context.Context, m *types.MarginLevels) (string, error) {
	return strconv.FormatUint(m.CollateralReleaseLevel, 10), nil
}

func (r *myMarginLevelsResolver) InitialLevel(_ context.Context, m *types.MarginLevels) (string, error) {
	return strconv.FormatUint(m.InitialMargin, 10), nil
}

func (r *myMarginLevelsResolver) SearchLevel(_ context.Context, m *types.MarginLevels) (string, error) {
	return strconv.FormatUint(m.SearchLevel, 10), nil
}

func (r *myMarginLevelsResolver) MaintenanceLevel(_ context.Context, m *types.MarginLevels) (string, error) {
	return strconv.FormatUint(m.MaintenanceMargin, 10), nil
}

func (r *myMarginLevelsResolver) Timestamp(_ context.Context, m *types.MarginLevels) (string, error) {
	return vegatime.Format(vegatime.UnixNano(m.Timestamp)), nil
}

// END: MarginLevels Resolver

// BEGIN: MarketData resolver

type myMarketDataResolver VegaResolverRoot

func (r *myMarketDataResolver) MarketState(_ context.Context, m *types.MarketData) (MarketState, error) {
	return convertMarketStateFromProto(m.MarketState)
}

func (r *myMarketDataResolver) IndicativePrice(_ context.Context, m *types.MarketData) (string, error) {
	return strconv.FormatUint(m.IndicativePrice, 10), nil
}

func (r *myMarketDataResolver) IndicativeVolume(_ context.Context, m *types.MarketData) (string, error) {
	return strconv.FormatUint(m.IndicativeVolume, 10), nil
}

func (r *myMarketDataResolver) BestBidPrice(_ context.Context, m *types.MarketData) (string, error) {
	return strconv.FormatUint(m.BestBidPrice, 10), nil
}

func (r *myMarketDataResolver) OpenInterest(_ context.Context, m *types.MarketData) (string, error) {
	return strconv.FormatUint(m.OpenInterest, 10), nil
}

func (r *myMarketDataResolver) BestBidVolume(_ context.Context, m *types.MarketData) (string, error) {
	return strconv.FormatUint(m.BestBidVolume, 10), nil
}

func (r *myMarketDataResolver) BestOfferPrice(_ context.Context, m *types.MarketData) (string, error) {
	return strconv.FormatUint(m.BestOfferPrice, 10), nil
}

func (r *myMarketDataResolver) BestOfferVolume(_ context.Context, m *types.MarketData) (string, error) {
	return strconv.FormatUint(m.BestOfferVolume, 10), nil
}

func (r *myMarketDataResolver) MidPrice(_ context.Context, m *types.MarketData) (string, error) {
	return strconv.FormatUint(m.MidPrice, 10), nil
}

func (r *myMarketDataResolver) MarkPrice(_ context.Context, m *types.MarketData) (string, error) {
	return strconv.FormatUint(m.MarkPrice, 10), nil
}

func (r *myMarketDataResolver) Timestamp(_ context.Context, m *types.MarketData) (string, error) {
	return vegatime.Format(vegatime.UnixNano(m.Timestamp)), nil
}

func (r *myMarketDataResolver) Market(ctx context.Context, m *types.MarketData) (*Market, error) {
	req := protoapi.MarketByIDRequest{MarketID: m.Market}
	res, err := r.tradingDataClient.MarketByID(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}

	market, err := MarketFromProto(res.Market)
	if err != nil {
		r.log.Error("unable to convert market from proto", logging.Error(err))
		return nil, err
	}
	// set Asset here too as well
	switch p := market.TradableInstrument.Instrument.Product.(type) {
	case *Future:
		req := protoapi.AssetByIDRequest{ID: p.Asset.ID}
		res, err := r.tradingDataClient.AssetByID(context.Background(), &req)
		if err != nil {
			r.log.Error("tradingData client", logging.Error(err))
			return nil, customErrorFromStatus(err)
		}
		p.Asset, err = AssetFromProto(res.Asset)
	}

	return market, nil
}

// END: MarketData resolver

// BEGIN: Market Depth Resolver

type myMarketDepthResolver VegaResolverRoot

func (r *myMarketDepthResolver) Buy(ctx context.Context, obj *types.MarketDepth) ([]types.PriceLevel, error) {
	valBuyLevels := make([]types.PriceLevel, 0)
	for _, v := range obj.Buy {
		valBuyLevels = append(valBuyLevels, *v)
	}
	return valBuyLevels, nil
}
func (r *myMarketDepthResolver) Sell(ctx context.Context, obj *types.MarketDepth) ([]types.PriceLevel, error) {
	valBuyLevels := make([]types.PriceLevel, 0)
	for _, v := range obj.Sell {
		valBuyLevels = append(valBuyLevels, *v)
	}
	return valBuyLevels, nil
}

func (r *myMarketDepthResolver) LastTrade(ctx context.Context, md *types.MarketDepth) (*types.Trade, error) {
	if md == nil {
		return nil, errors.New("invalid market depth")
	}

	req := protoapi.LastTradeRequest{MarketID: md.MarketID}
	res, err := r.tradingDataClient.LastTrade(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}

	return res.Trade, nil
}

func (r *myMarketDepthResolver) Market(ctx context.Context, md *types.MarketDepth) (*Market, error) {
	if md == nil {
		return nil, errors.New("invalid market depth")
	}

	req := protoapi.MarketByIDRequest{MarketID: md.MarketID}
	res, err := r.tradingDataClient.MarketByID(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	market, err := MarketFromProto(res.Market)
	// set Asset here too as well
	switch p := market.TradableInstrument.Instrument.Product.(type) {
	case *Future:
		req := protoapi.AssetByIDRequest{ID: p.Asset.ID}
		res, err := r.tradingDataClient.AssetByID(context.Background(), &req)
		if err != nil {
			r.log.Error("tradingData client", logging.Error(err))
			return nil, customErrorFromStatus(err)
		}
		p.Asset, err = AssetFromProto(res.Asset)
	}
	return market, err
}

// END: Market Depth Resolver

// BEGIN: Order Resolver

type myOrderResolver VegaResolverRoot

func (r *myOrderResolver) RejectionReason(_ context.Context, o *types.Order) (*OrderRejectionReason, error) {
	if o.Reason == types.OrderError_ORDER_ERROR_NONE {
		return nil, nil
	}
	reason, err := convertOrderRejectionReasonFromProto(o.Reason)
	if err != nil {
		return nil, err
	}
	return &reason, nil
}

func (r *myOrderResolver) Price(ctx context.Context, obj *types.Order) (string, error) {
	return strconv.FormatUint(obj.Price, 10), nil
}
func (r *myOrderResolver) TimeInForce(ctx context.Context, obj *types.Order) (OrderTimeInForce, error) {
	return convertOrderTimeInForceFromProto(obj.TimeInForce)
}

func (r *myOrderResolver) Type(ctx context.Context, obj *types.Order) (*OrderType, error) {
	t, err := convertOrderTypeFromProto(obj.Type)
	if err != nil {
		return nil, err
	}
	return &t, nil
}

func (r *myOrderResolver) Side(ctx context.Context, obj *types.Order) (Side, error) {
	return convertSideFromProto(obj.Side)
}

func (r *myOrderResolver) Market(ctx context.Context, obj *types.Order) (*Market, error) {
	if obj == nil {
		return nil, errors.New("invalid order")
	}
	req := protoapi.MarketByIDRequest{MarketID: obj.MarketID}
	res, err := r.tradingDataClient.MarketByID(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	market, err := MarketFromProto(res.Market)
	// set Asset here too as well
	switch p := market.TradableInstrument.Instrument.Product.(type) {
	case *Future:
		req := protoapi.AssetByIDRequest{ID: p.Asset.ID}
		res, err := r.tradingDataClient.AssetByID(context.Background(), &req)
		if err != nil {
			r.log.Error("tradingData client", logging.Error(err))
			return nil, customErrorFromStatus(err)
		}
		p.Asset, err = AssetFromProto(res.Asset)
	}
	return market, err
}
func (r *myOrderResolver) Size(ctx context.Context, obj *types.Order) (string, error) {
	return strconv.FormatUint(obj.Size, 10), nil
}
func (r *myOrderResolver) Remaining(ctx context.Context, obj *types.Order) (string, error) {
	return strconv.FormatUint(obj.Remaining, 10), nil
}

func (r *myOrderResolver) Status(ctx context.Context, obj *types.Order) (OrderStatus, error) {
	return convertOrderStatusFromProto(obj.Status)
}

func (r *myOrderResolver) CreatedAt(ctx context.Context, obj *types.Order) (string, error) {
	return vegatime.Format(vegatime.UnixNano(obj.CreatedAt)), nil
}

func (r *myOrderResolver) UpdatedAt(ctx context.Context, obj *types.Order) (string, error) {
	if obj.UpdatedAt <= 0 {
		return "", nil
	}
	return vegatime.Format(vegatime.UnixNano(obj.UpdatedAt)), nil
}

func (r *myOrderResolver) Version(ctx context.Context, obj *types.Order) (string, error) {
	return strconv.FormatUint(obj.Version, 10), nil
}

func (r *myOrderResolver) ExpiresAt(ctx context.Context, obj *types.Order) (*string, error) {
	if obj.ExpiresAt <= 0 {
		return nil, nil
	}
	expiresAt := vegatime.Format(vegatime.UnixNano(obj.ExpiresAt))
	return &expiresAt, nil
}

func (r *myOrderResolver) Trades(ctx context.Context, ord *types.Order) ([]*types.Trade, error) {
	if ord == nil {
		return nil, errors.New("nil order")
	}
	req := protoapi.TradesByOrderRequest{OrderID: ord.Id}
	res, err := r.tradingDataClient.TradesByOrder(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	return res.Trades, nil
}
func (r *myOrderResolver) Party(ctx context.Context, order *types.Order) (*types.Party, error) {
	if order == nil {
		return nil, errors.New("nil order")
	}
	if len(order.PartyID) == 0 {
		return nil, errors.New("invalid party")
	}
	return &types.Party{Id: order.PartyID}, nil
}

// END: Order Resolver

// BEGIN: Trade Resolver

type myTradeResolver VegaResolverRoot

func (r *myTradeResolver) Market(ctx context.Context, obj *types.Trade) (*Market, error) {
	if obj == nil {
		return nil, errors.New("invalid trade")
	}
	req := protoapi.MarketByIDRequest{MarketID: obj.MarketID}
	res, err := r.tradingDataClient.MarketByID(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	market, err := MarketFromProto(res.Market)
	// set Asset here too as well
	switch p := market.TradableInstrument.Instrument.Product.(type) {
	case *Future:
		req := protoapi.AssetByIDRequest{ID: p.Asset.ID}
		res, err := r.tradingDataClient.AssetByID(context.Background(), &req)
		if err != nil {
			r.log.Error("tradingData client", logging.Error(err))
			return nil, customErrorFromStatus(err)
		}
		p.Asset, err = AssetFromProto(res.Asset)
	}
	return market, err

}
func (r *myTradeResolver) Aggressor(ctx context.Context, obj *types.Trade) (Side, error) {
	return Side(obj.Aggressor.String()), nil
}
func (r *myTradeResolver) Price(ctx context.Context, obj *types.Trade) (string, error) {
	return strconv.FormatUint(obj.Price, 10), nil
}
func (r *myTradeResolver) Size(ctx context.Context, obj *types.Trade) (string, error) {
	return strconv.FormatUint(obj.Size, 10), nil
}
func (r *myTradeResolver) CreatedAt(ctx context.Context, obj *types.Trade) (string, error) {
	return vegatime.Format(vegatime.UnixNano(obj.Timestamp)), nil
}
func (r *myTradeResolver) Buyer(ctx context.Context, obj *types.Trade) (*types.Party, error) {
	if obj == nil {
		return nil, errors.New("invalid trade")
	}
	if len(obj.Buyer) == 0 {
		return nil, errors.New("invalid buyer")
	}
	req := protoapi.PartyByIDRequest{PartyID: obj.Buyer}
	res, err := r.tradingDataClient.PartyByID(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	return res.Party, nil
}
func (r *myTradeResolver) Seller(ctx context.Context, obj *types.Trade) (*types.Party, error) {
	if obj == nil {
		return nil, errors.New("invalid trade")
	}
	if len(obj.Seller) == 0 {
		return nil, errors.New("invalid seller")
	}
	req := protoapi.PartyByIDRequest{PartyID: obj.Seller}
	res, err := r.tradingDataClient.PartyByID(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	return res.Party, nil
}

func (r *myTradeResolver) Type(ctx context.Context, obj *proto.Trade) (TradeType, error) {
	return convertTradeTypeFromProto(obj.Type)
}

func (r *myTradeResolver) BuyerAuctionBatch(ctx context.Context, obj *types.Trade) (*int, error) {
	i := int(obj.BuyerAuctionBatch)
	return &i, nil
}

func (r *myTradeResolver) BuyerFee(ctx context.Context, obj *types.Trade) (*TradeFee, error) {
	fee := TradeFee{
		MakerFee:          "0",
		InfrastructureFee: "0",
		LiquidityFee:      "0",
	}
	if obj.BuyerFee != nil {
		fee.MakerFee = strconv.FormatUint(obj.BuyerFee.MakerFee, 10)
		fee.InfrastructureFee = strconv.FormatUint(obj.BuyerFee.InfrastructureFee, 10)
		fee.LiquidityFee = strconv.FormatUint(obj.BuyerFee.LiquidityFee, 10)
	}
	return &fee, nil
}

func (r *myTradeResolver) SellerAuctionBatch(ctx context.Context, obj *types.Trade) (*int, error) {
	i := int(obj.SellerAuctionBatch)
	return &i, nil
}

func (r *myTradeResolver) SellerFee(ctx context.Context, obj *types.Trade) (*TradeFee, error) {
	fee := TradeFee{
		MakerFee:          "0",
		InfrastructureFee: "0",
		LiquidityFee:      "0",
	}
	if obj.SellerFee != nil {
		fee.MakerFee = strconv.FormatUint(obj.SellerFee.MakerFee, 10)
		fee.InfrastructureFee = strconv.FormatUint(obj.SellerFee.InfrastructureFee, 10)
		fee.LiquidityFee = strconv.FormatUint(obj.SellerFee.LiquidityFee, 10)
	}
	return &fee, nil
}

// END: Trade Resolver

// BEGIN: Candle Resolver

type myCandleResolver VegaResolverRoot

func (r *myCandleResolver) High(ctx context.Context, obj *types.Candle) (string, error) {
	return strconv.FormatUint(obj.High, 10), nil
}
func (r *myCandleResolver) Low(ctx context.Context, obj *types.Candle) (string, error) {
	return strconv.FormatUint(obj.Low, 10), nil
}
func (r *myCandleResolver) Open(ctx context.Context, obj *types.Candle) (string, error) {
	return strconv.FormatUint(obj.Open, 10), nil
}
func (r *myCandleResolver) Close(ctx context.Context, obj *types.Candle) (string, error) {
	return strconv.FormatUint(obj.Close, 10), nil
}
func (r *myCandleResolver) Volume(ctx context.Context, obj *types.Candle) (string, error) {
	return strconv.FormatUint(obj.Volume, 10), nil
}
func (r *myCandleResolver) Datetime(ctx context.Context, obj *types.Candle) (string, error) {
	return vegatime.Format(vegatime.UnixNano(obj.Timestamp)), nil
}
func (r *myCandleResolver) Timestamp(ctx context.Context, obj *types.Candle) (string, error) {
	return strconv.FormatInt(obj.Timestamp, 10), nil
}
func (r *myCandleResolver) Interval(ctx context.Context, obj *types.Candle) (Interval, error) {
	return convertIntervalFromProto(obj.Interval)
}

// END: Candle Resolver

// BEGIN: Price Level Resolver

type myPriceLevelResolver VegaResolverRoot

func (r *myPriceLevelResolver) Price(ctx context.Context, obj *types.PriceLevel) (string, error) {
	return strconv.FormatUint(obj.Price, 10), nil
}

func (r *myPriceLevelResolver) Volume(ctx context.Context, obj *types.PriceLevel) (string, error) {
	return strconv.FormatUint(obj.Volume, 10), nil
}

func (r *myPriceLevelResolver) NumberOfOrders(ctx context.Context, obj *types.PriceLevel) (string, error) {
	return strconv.FormatUint(obj.Price, 10), nil
}

// END: Price Level Resolver

// BEGIN: Position Resolver

type myPositionResolver VegaResolverRoot

func (r *myPositionResolver) Market(ctx context.Context, obj *types.Position) (*Market, error) {
	if obj == nil {
		return nil, errors.New("invalid position")
	}
	if len(obj.MarketID) <= 0 {
		return nil, nil
	}
	req := protoapi.MarketByIDRequest{MarketID: obj.MarketID}
	res, err := r.tradingDataClient.MarketByID(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	market, err := MarketFromProto(res.Market)
	if err != nil {
		r.log.Error("unable to convert market from proto", logging.Error(err))
		return nil, err
	}
	// add the market name explicitly here as it does not
	// come anymore from the market framework
	market.Name = market.TradableInstrument.Instrument.Name
	// set Asset here too
	switch p := market.TradableInstrument.Instrument.Product.(type) {
	case *Future:
		req := protoapi.AssetByIDRequest{ID: p.Asset.ID}
		res, err := r.tradingDataClient.AssetByID(context.Background(), &req)
		if err != nil {
			r.log.Error("tradingData client", logging.Error(err))
			return nil, customErrorFromStatus(err)
		}
		p.Asset, err = AssetFromProto(res.Asset)
	}

	return market, nil
}

func (r *myPositionResolver) UpdatedAt(ctx context.Context, obj *types.Position) (string, error) {
	return vegatime.Format(vegatime.UnixNano(obj.UpdatedAt)), nil
}

func (r *myPositionResolver) OpenVolume(ctx context.Context, obj *types.Position) (string, error) {
	return strconv.FormatInt(obj.OpenVolume, 10), nil
}

func (r *myPositionResolver) RealisedPnl(ctx context.Context, obj *types.Position) (string, error) {
	return strconv.FormatInt(obj.RealisedPNL, 10), nil
}

func (r *myPositionResolver) UnrealisedPnl(ctx context.Context, obj *types.Position) (string, error) {
	return strconv.FormatInt(obj.UnrealisedPNL, 10), nil
}

func (r *myPositionResolver) AverageEntryPrice(ctx context.Context, obj *types.Position) (string, error) {
	return strconv.FormatUint(obj.AverageEntryPrice, 10), nil
}

func (r *myPositionResolver) Margins(ctx context.Context, obj *types.Position) ([]*types.MarginLevels, error) {
	if obj == nil {
		return nil, errors.New("invalid position")
	}
	if len(obj.PartyID) <= 0 {
		return nil, errors.New("missing party id")
	}
	req := protoapi.MarginLevelsRequest{
		PartyID:  obj.PartyID,
		MarketID: obj.MarketID,
	}
	res, err := r.tradingDataClient.MarginLevels(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	out := make([]*types.MarginLevels, 0, len(res.MarginLevels))
	out = append(out, res.MarginLevels...)
	return out, nil
}

// END: Position Resolver

// BEGIN: Mutation Resolver

type myMutationResolver VegaResolverRoot

func (r *myMutationResolver) PrepareWithdrawal(
	ctx context.Context,
	partyID, amount, asset string,
	erc20Details *Erc20WithdrawalDetailsInput,
) (*PreparedWithdrawal, error) {
	var ext *types.WithdrawExt
	if erc20Details != nil {
		ext = erc20Details.IntoProtoExt()
	}

	amountU, err := safeStringUint64(amount)
	if err != nil {
		return nil, err
	}

	req := protoapi.PrepareWithdrawRequest{
		Withdraw: &types.WithdrawSubmission{
			PartyID: partyID,
			Asset:   asset,
			Amount:  amountU,
			Ext:     ext,
		},
	}

	res, err := r.tradingClient.PrepareWithdraw(ctx, &req)
	if err != nil {
		return nil, err
	}

	return &PreparedWithdrawal{
		Blob: base64.StdEncoding.EncodeToString(res.Blob),
	}, nil
}

func (r *myMutationResolver) SubmitTransaction(ctx context.Context, data string, sig SignatureInput) (*TransactionSubmitted, error) {
	decodedData, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return nil, err
	}
	decodedSig, err := base64.StdEncoding.DecodeString(sig.Sig)
	if err != nil {
		return nil, err
	}
	req := &protoapi.SubmitTransactionRequest{
		Tx: &types.SignedBundle{
			Tx: decodedData,
			Sig: &types.Signature{
				Sig:     decodedSig,
				Version: uint64(sig.Version),
				Algo:    sig.Algo,
			},
		},
	}
	res, err := r.tradingClient.SubmitTransaction(ctx, req)
	if err != nil {
		r.log.Error("Failed to submit transaction", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}

	return &TransactionSubmitted{
		Success: res.Success,
	}, nil
}

func (r *myMutationResolver) PrepareOrderSubmit(ctx context.Context, market, party string, price *string, size string, side Side,
	timeInForce OrderTimeInForce, expiration *string, ty OrderType, reference *string) (*PreparedSubmitOrder, error) {

	order := &types.OrderSubmission{}

	var (
		p   uint64
		err error
	)

	// We need to convert strings to uint64 (JS doesn't yet support uint64)
	if price != nil {
		p, err = safeStringUint64(*price)
		if err != nil {
			return nil, err
		}
	}
	order.Price = p
	s, err := safeStringUint64(size)
	if err != nil {
		return nil, err
	}
	order.Size = s
	if len(market) <= 0 {
		return nil, errors.New("market missing or empty")
	}
	order.MarketID = market
	if len(party) <= 0 {
		return nil, errors.New("party missing or empty")
	}

	order.PartyID = party
	if order.TimeInForce, err = convertOrderTimeInForceToProto(timeInForce); err != nil {
		return nil, err
	}
	if order.Side, err = convertSideToProto(side); err != nil {
		return nil, err
	}
	if order.Type, err = convertOrderTypeToProto(ty); err != nil {
		return nil, err
	}

	// GTT must have an expiration value
	if order.TimeInForce == types.Order_TIF_GTT && expiration != nil {
		var expiresAt time.Time
		expiresAt, err = vegatime.Parse(*expiration)
		if err != nil {
			return nil, fmt.Errorf("cannot parse expiration time: %s - invalid format sent to create order (example: 2018-01-02T15:04:05Z)", *expiration)
		}

		// move to pure timestamps or convert an RFC format shortly
		order.ExpiresAt = expiresAt.UnixNano()
	}
	if reference != nil {
		order.Reference = *reference
	}

	req := protoapi.SubmitOrderRequest{
		Submission: order,
	}

	// Pass the order over for consensus (service layer will use RPC client internally and handle errors etc)
	resp, err := r.tradingClient.PrepareSubmitOrder(ctx, &req)
	if err != nil {
		r.log.Error("Failed to create order using rpc client in graphQL resolver", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	return &PreparedSubmitOrder{
		Blob: base64.StdEncoding.EncodeToString(resp.Blob),
	}, nil
}

func (r *myMutationResolver) PrepareOrderCancel(ctx context.Context, id *string, party string, market *string) (*PreparedCancelOrder, error) {
	order := &types.OrderCancellation{}

	if market != nil {
		order.MarketID = *market
	}
	if id != nil {
		order.OrderID = *id
	}
	if len(party) == 0 {
		return nil, errors.New("party missing or empty")
	}
	order.PartyID = party

	// Pass the cancellation over for consensus (service layer will use RPC client internally and handle errors etc)

	req := protoapi.CancelOrderRequest{
		Cancellation: order,
	}
	pendingOrder, err := r.tradingClient.PrepareCancelOrder(ctx, &req)
	if err != nil {
		return nil, customErrorFromStatus(err)
	}
	return &PreparedCancelOrder{
		Blob: base64.StdEncoding.EncodeToString(pendingOrder.Blob),
	}, nil

}

func (r *myMutationResolver) PrepareProposal(
	ctx context.Context, partyID string, reference *string, proposalTerms ProposalTermsInput) (*PreparedProposal, error) {
	var ref string
	if reference != nil {
		ref = *reference
	}

	terms, err := proposalTerms.IntoProto()
	if err != nil {
		return nil, err
	}

	pendingProposal, err := r.tradingClient.PrepareProposal(ctx, &protoapi.PrepareProposalRequest{
		PartyID:   partyID,
		Reference: ref,
		Proposal:  terms,
	})
	if err != nil {
		return nil, customErrorFromStatus(err)
	}
	return &PreparedProposal{
		Blob: base64.StdEncoding.EncodeToString(pendingProposal.Blob),
		PendingProposal: &types.GovernanceData{
			Proposal: pendingProposal.PendingProposal,
		},
	}, nil
}

func (r *myMutationResolver) PrepareVote(ctx context.Context, value VoteValue, partyID, proposalID string) (*PreparedVote, error) {
	party, err := getParty(ctx, r.log, r.tradingDataClient, partyID)
	if err != nil {
		return nil, err
	}
	protoValue, err := convertVoteValueToProto(value)
	if err != nil {
		return nil, err
	}
	req := &protoapi.PrepareVoteRequest{
		Vote: &types.Vote{
			Value:      protoValue,
			PartyID:    partyID,
			ProposalID: proposalID,
		},
	}
	resp, err := r.tradingClient.PrepareVote(ctx, req)
	if err != nil {
		return nil, err
	}
	gqlValue, err := convertVoteValueFromProto(resp.Vote.Value)
	if err != nil {
		return nil, err
	}
	return &PreparedVote{
		Blob: base64.StdEncoding.EncodeToString(resp.Blob),
		Vote: &ProposalVote{
			Vote: &Vote{
				Party: party,
				Value: gqlValue,
			},
			ProposalID: resp.Vote.ProposalID,
		},
	}, nil
}

func (r *myMutationResolver) PrepareOrderAmend(ctx context.Context, id string, party string, price, size string, expiration *string, tif OrderTimeInForce) (*PreparedAmendOrder, error) {
	order := &types.OrderAmendment{}

	// Cancellation currently only requires ID and Market to be set, all other fields will be added
	if len(id) == 0 {
		return nil, errors.New("id missing or empty")
	}
	order.OrderID = id
	if len(party) == 0 {
		return nil, errors.New("party missing or empty")
	}
	order.PartyID = party

	var err error
	pricevalue, err := strconv.ParseUint(price, 10, 64)
	if err != nil {
		r.log.Error("unable to convert price from string in order amend",
			logging.Error(err))
		return nil, errors.New("invalid price, could not convert to unsigned int")
	}
	order.Price = &proto.Price{Value: pricevalue}

	order.SizeDelta, err = strconv.ParseInt(size, 10, 64)
	if err != nil {
		r.log.Error("unable to convert size from string in order amend",
			logging.Error(err))
		return nil, errors.New("invalid size, could not convert to unsigned int")
	}

	order.TimeInForce, err = convertOrderTimeInForceToProto(tif)
	if err != nil {
		r.log.Error("unable to parse time in force in order amend",
			logging.Error(err))
		return nil, errors.New("invalid time in force, could not convert to vega time in force")
	}

	if expiration != nil {
		expiresAt, err := vegatime.Parse(*expiration)
		if err != nil {
			return nil, fmt.Errorf("cannot parse expiration time: %s - invalid format sent to create order (example: 2018-01-02T15:04:05Z)", *expiration)
		}
		// move to pure timestamps or convert an RFC format shortly
		order.ExpiresAt = &proto.Timestamp{Value: expiresAt.UnixNano()}
	}

	req := protoapi.AmendOrderRequest{
		Amendment: order,
	}
	pendingOrder, err := r.tradingClient.PrepareAmendOrder(ctx, &req)
	if err != nil {
		return nil, customErrorFromStatus(err)
	}
	return &PreparedAmendOrder{
		Blob: base64.StdEncoding.EncodeToString(pendingOrder.Blob),
	}, nil
}

// END: Mutation Resolver

// BEGIN: Subscription Resolver

type mySubscriptionResolver VegaResolverRoot

func (r *mySubscriptionResolver) Margins(ctx context.Context, partyID string, marketID *string) (<-chan *types.MarginLevels, error) {
	var mktid string
	if marketID != nil {
		mktid = *marketID
	}
	req := &protoapi.MarginLevelsSubscribeRequest{
		MarketID: mktid,
		PartyID:  partyID,
	}
	stream, err := r.tradingDataClient.MarginLevelsSubscribe(ctx, req)
	if err != nil {
		return nil, customErrorFromStatus(err)
	}

	ch := make(chan *types.MarginLevels)
	go func() {
		defer func() {
			stream.CloseSend()
			close(ch)
		}()
		for {
			m, err := stream.Recv()
			if err == io.EOF {
				r.log.Error("margin levels: stream closed by server", logging.Error(err))
				break
			}
			if err != nil {
				r.log.Error("margin levls: stream closed", logging.Error(err))
				break
			}
			ch <- m
		}
	}()

	return ch, nil
}

func (r *mySubscriptionResolver) MarketData(ctx context.Context, marketID *string) (<-chan *types.MarketData, error) {
	var mktid string
	if marketID != nil {
		mktid = *marketID
	}
	req := &protoapi.MarketsDataSubscribeRequest{
		MarketID: mktid,
	}
	stream, err := r.tradingDataClient.MarketsDataSubscribe(ctx, req)
	if err != nil {
		return nil, customErrorFromStatus(err)
	}

	ch := make(chan *types.MarketData)
	go func() {
		defer func() {
			stream.CloseSend()
			close(ch)
		}()
		for {
			m, err := stream.Recv()
			if err == io.EOF {
				r.log.Error("marketdata: stream closed by server", logging.Error(err))
				break
			}
			if err != nil {
				r.log.Error("marketdata: stream closed", logging.Error(err))
				break
			}
			ch <- m
		}
	}()

	return ch, nil
}

func (r *mySubscriptionResolver) Accounts(ctx context.Context, marketID *string, partyID *string, asset *string, typeArg *AccountType) (<-chan *types.Account, error) {
	var (
		mkt, pty string
		ty       types.AccountType
	)

	if marketID == nil && partyID == nil && asset == nil && typeArg == nil {
		// Updates on every balance update, on every account, for everyone and shouldn't be allowed for GraphQL.
		return nil, errors.New("at least one query filter must be applied for this subscription")
	}
	if marketID != nil {
		mkt = *marketID
	}
	if partyID != nil {
		pty = *partyID
	}
	if typeArg != nil {
		ty = typeArg.IntoProto()
	}

	req := &protoapi.AccountsSubscribeRequest{
		MarketID: mkt,
		PartyID:  pty,
		Type:     ty,
	}
	stream, err := r.tradingDataClient.AccountsSubscribe(ctx, req)
	if err != nil {
		return nil, customErrorFromStatus(err)
	}

	c := make(chan *types.Account)
	go func() {
		defer func() {
			stream.CloseSend()
			close(c)
		}()
		for {
			a, err := stream.Recv()
			if err == io.EOF {
				r.log.Error("accounts: stream closed by server", logging.Error(err))
				break
			}
			if err != nil {
				r.log.Error("accounts: stream closed", logging.Error(err))
				break
			}
			c <- a
		}
	}()

	return c, nil
}

func (r *mySubscriptionResolver) Orders(ctx context.Context, market *string, party *string) (<-chan []*types.Order, error) {
	var (
		mkt, pty string
	)
	if market != nil {
		mkt = *market
	}
	if party != nil {
		pty = *party
	}

	req := &protoapi.OrdersSubscribeRequest{
		MarketID: mkt,
		PartyID:  pty,
	}
	stream, err := r.tradingDataClient.OrdersSubscribe(ctx, req)
	if err != nil {
		return nil, customErrorFromStatus(err)
	}

	c := make(chan []*types.Order)
	go func() {
		defer func() {
			stream.CloseSend()
			close(c)
		}()
		for {
			o, err := stream.Recv()
			if err == io.EOF {
				r.log.Error("orders: stream closed by server", logging.Error(err))
				break
			}
			if err != nil {
				r.log.Error("orders: stream closed", logging.Error(err))
				break
			}
			c <- o.Orders
		}
	}()

	return c, nil
}

func (r *mySubscriptionResolver) Trades(ctx context.Context, market *string, party *string) (<-chan []*types.Trade, error) {
	var (
		mkt, pty string
	)
	if market != nil {
		mkt = *market
	}
	if party != nil {
		pty = *party
	}

	req := &protoapi.TradesSubscribeRequest{
		MarketID: mkt,
		PartyID:  pty,
	}
	stream, err := r.tradingDataClient.TradesSubscribe(ctx, req)
	if err != nil {
		return nil, customErrorFromStatus(err)
	}

	c := make(chan []*types.Trade)
	go func() {
		defer func() {
			stream.CloseSend()
			close(c)
		}()
		for {
			t, err := stream.Recv()
			if err == io.EOF {
				r.log.Error("trades: stream closed by server", logging.Error(err))
				break
			}
			if err != nil {
				r.log.Error("trades: stream closed", logging.Error(err))
				break
			}
			c <- t.Trades
		}
	}()

	return c, nil
}

func (r *mySubscriptionResolver) Positions(ctx context.Context, party string) (<-chan *types.Position, error) {
	req := &protoapi.PositionsSubscribeRequest{
		PartyID: party,
	}
	stream, err := r.tradingDataClient.PositionsSubscribe(ctx, req)
	if err != nil {
		return nil, customErrorFromStatus(err)
	}

	c := make(chan *types.Position)
	go func() {
		defer func() {
			stream.CloseSend()
			close(c)
		}()
		for {
			t, err := stream.Recv()
			if err == io.EOF {
				r.log.Error("positions: stream closed by server", logging.Error(err))
				break
			}
			if err != nil {
				r.log.Error("positions: stream closed", logging.Error(err))
				break
			}
			c <- t
		}
	}()

	return c, nil
}

func (r *mySubscriptionResolver) MarketDepth(ctx context.Context, market string) (<-chan *types.MarketDepth, error) {
	req := &protoapi.MarketDepthSubscribeRequest{
		MarketID: market,
	}
	stream, err := r.tradingDataClient.MarketDepthSubscribe(ctx, req)
	if err != nil {
		return nil, customErrorFromStatus(err)
	}

	c := make(chan *types.MarketDepth)
	go func() {
		defer func() {
			stream.CloseSend()
			close(c)
		}()
		for {
			md, err := stream.Recv()
			if err == io.EOF {
				r.log.Error("marketDepth: stream closed by server", logging.Error(err))
				break
			}
			if err != nil {
				r.log.Error("marketDepth: stream closed", logging.Error(err))
				break
			}
			c <- md
		}
	}()

	return c, nil
}

func (r *mySubscriptionResolver) Candles(ctx context.Context, market string, interval Interval) (<-chan *types.Candle, error) {

	pinterval, err := convertIntervalToProto(interval)
	if err != nil {
		r.log.Debug("invalid interval for candles subscriptions", logging.Error(err))
	}

	req := &protoapi.CandlesSubscribeRequest{
		MarketID: market,
		Interval: pinterval,
	}
	stream, err := r.tradingDataClient.CandlesSubscribe(ctx, req)
	if err != nil {
		return nil, customErrorFromStatus(err)
	}

	c := make(chan *types.Candle)
	go func() {
		defer func() {
			stream.CloseSend()
			close(c)
		}()
		for {
			cdl, err := stream.Recv()
			if err == io.EOF {
				r.log.Error("candles: stream closed by server", logging.Error(err))
				break
			}
			if err != nil {
				r.log.Error("candles: stream closed", logging.Error(err))
				break
			}
			c <- cdl
		}
	}()
	return c, nil
}

func isStreamClosed(err error, log *logging.Logger) bool {
	if err == io.EOF {
		log.Error("stream closed by server", logging.Error(err))
		return true
	}
	if err != nil {
		log.Error("stream closed", logging.Error(err))
		return true
	}
	return false
}

func (r *mySubscriptionResolver) subscribeAllProposals(ctx context.Context) (<-chan *types.GovernanceData, error) {
	stream, err := r.tradingDataClient.ObserveGovernance(ctx, &empty.Empty{})
	if err != nil {
		return nil, customErrorFromStatus(err)
	}
	output := make(chan *types.GovernanceData)
	go func() {
		defer func() {
			stream.CloseSend()
			close(output)
		}()
		for data, err := stream.Recv(); !isStreamClosed(err, r.log); data, err = stream.Recv() {
			output <- data
		}
	}()
	return output, nil
}

func (r *mySubscriptionResolver) subscribePartyProposals(ctx context.Context, partyID string) (<-chan *types.GovernanceData, error) {
	stream, err := r.tradingDataClient.ObservePartyProposals(ctx, &protoapi.ObservePartyProposalsRequest{
		PartyID: partyID,
	})
	if err != nil {
		return nil, customErrorFromStatus(err)
	}
	output := make(chan *types.GovernanceData)
	go func() {
		defer func() {
			stream.CloseSend()
			close(output)
		}()
		for data, err := stream.Recv(); !isStreamClosed(err, r.log); data, err = stream.Recv() {
			output <- data
		}
	}()
	return output, nil
}

func (r *mySubscriptionResolver) Proposals(ctx context.Context, partyID *string) (<-chan *types.GovernanceData, error) {
	if partyID != nil && len(*partyID) > 0 {
		return r.subscribePartyProposals(ctx, *partyID)
	}
	return r.subscribeAllProposals(ctx)
}

func (r *mySubscriptionResolver) subscribeProposalVotes(ctx context.Context, proposalID string) (<-chan *ProposalVote, error) {
	output := make(chan *ProposalVote)
	stream, err := r.tradingDataClient.ObserveProposalVotes(ctx, &protoapi.ObserveProposalVotesRequest{
		ProposalID: proposalID,
	})
	if err != nil {
		return nil, customErrorFromStatus(err)
	}
	go func() {
		defer func() {
			stream.CloseSend()
			close(output)
		}()
		for {
			data, err := stream.Recv()
			if isStreamClosed(err, r.log) {
				break
			}
			party, err := getParty(ctx, r.log, r.tradingDataClient, data.PartyID)
			if err != nil {
				r.log.Error("Votes subscriber. getParty error", logging.Error(err))
				break
			}
			output <- ProposalVoteFromProto(data, party)
		}
	}()
	return output, nil
}

func (r *mySubscriptionResolver) subscribePartyVotes(ctx context.Context, partyID string) (<-chan *ProposalVote, error) {
	output := make(chan *ProposalVote)
	stream, err := r.tradingDataClient.ObservePartyVotes(ctx, &protoapi.ObservePartyVotesRequest{
		PartyID: partyID,
	})
	if err != nil {
		return nil, customErrorFromStatus(err)
	}
	go func() {
		defer func() {
			stream.CloseSend()
			close(output)
		}()
		for {
			data, err := stream.Recv()
			if isStreamClosed(err, r.log) {
				break
			}
			party, err := getParty(ctx, r.log, r.tradingDataClient, data.PartyID)
			if err != nil {
				r.log.Error("Votes subscriber. getParty error", logging.Error(err))
				break
			}
			output <- ProposalVoteFromProto(data, party)
		}
	}()
	return output, nil
}

func (r *mySubscriptionResolver) Votes(ctx context.Context, proposalID *string, partyID *string) (<-chan *ProposalVote, error) {

	if proposalID != nil && len(*proposalID) == 0 {
		return r.subscribeProposalVotes(ctx, *proposalID)
	} else if partyID != nil && len(*partyID) == 0 {
		return r.subscribePartyVotes(ctx, *partyID)
	}
	return nil, ErrInvalidVotesSubscription
}

func (r *mySubscriptionResolver) BusEvents(ctx context.Context, types []BusEventType, marketID, partyID *string) (<-chan []*BusEvent, error) {
	t := eventTypeToProto(types...)
	req := protoapi.ObserveEventsRequest{
		Type: t,
	}
	if marketID != nil {
		req.MarketID = *marketID
	}
	if partyID != nil {
		req.PartyID = *partyID
	}
	stream, err := r.tradingDataClient.ObserveEventBus(ctx, &req)
	if err != nil {
		return nil, customErrorFromStatus(err)
	}
	out := make(chan []*BusEvent)
	go func() {
		defer func() {
			stream.CloseSend()
			close(out)
		}()
		for {
			data, err := stream.Recv()
			if isStreamClosed(err, r.log) {
				return
			}
			be := busEventFromProto(data.Events...)
			out <- be
		}
	}()
	return out, nil
}

// START: Account Resolver

type myAccountResolver VegaResolverRoot

func (r *myAccountResolver) Balance(ctx context.Context, acc *types.Account) (string, error) {
	bal := fmt.Sprintf("%d", acc.Balance)
	return bal, nil
}

func (r *myAccountResolver) Market(ctx context.Context, acc *types.Account) (*Market, error) {
	if acc == nil {
		return nil, errors.New("invalid account")
	}

	// Only margin accounts have a market relation
	if acc.Type == types.AccountType_ACCOUNT_TYPE_MARGIN {
		req := protoapi.MarketByIDRequest{MarketID: acc.MarketID}
		res, err := r.tradingDataClient.MarketByID(ctx, &req)
		if err != nil {
			r.log.Error("tradingData client", logging.Error(err))
			return nil, customErrorFromStatus(err)
		}
		market, err := MarketFromProto(res.Market)
		if err != nil {
			r.log.Error("unable to convert market from proto", logging.Error(err))
			return nil, err
		}
		// add the market name explicitly here as it does not
		// come anymore from the market framework
		market.Name = market.TradableInstrument.Instrument.Name
		// set Asset here too
		switch p := market.TradableInstrument.Instrument.Product.(type) {
		case *Future:
			req := protoapi.AssetByIDRequest{ID: p.Asset.ID}
			res, err := r.tradingDataClient.AssetByID(context.Background(), &req)
			if err != nil {
				r.log.Error("tradingData client", logging.Error(err))
				return nil, customErrorFromStatus(err)
			}
			p.Asset, err = AssetFromProto(res.Asset)
		}
		return market, nil
	}

	return nil, nil
}

func (r *myAccountResolver) Type(ctx context.Context, obj *types.Account) (AccountType, error) {
	return convertAccountTypeFromProto(obj.Type)
}

func (r *myAccountResolver) Asset(ctx context.Context, obj *types.Account) (*Asset, error) {
	req := protoapi.AssetByIDRequest{ID: obj.Asset}
	res, err := r.tradingDataClient.AssetByID(ctx, &req)
	if err != nil {
		r.log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	return AssetFromProto(res.Asset)
}

// END: Account Resolver

type myStatisticsResolver VegaResolverRoot

func (r *myStatisticsResolver) BlockHeight(ctx context.Context, obj *types.Statistics) (int, error) {
	return int(obj.BlockHeight), nil
}

func (r *myStatisticsResolver) BacklogLength(ctx context.Context, obj *types.Statistics) (int, error) {
	return int(obj.BacklogLength), nil
}

func (r *myStatisticsResolver) TotalPeers(ctx context.Context, obj *types.Statistics) (int, error) {
	return int(obj.TotalPeers), nil
}

func (r *myStatisticsResolver) Status(ctx context.Context, obj *types.Statistics) (string, error) {
	return obj.Status.String(), nil
}

func (r *myStatisticsResolver) TxPerBlock(ctx context.Context, obj *types.Statistics) (int, error) {
	return int(obj.TxPerBlock), nil
}

func (r *myStatisticsResolver) AverageTxBytes(ctx context.Context, obj *types.Statistics) (int, error) {
	return int(obj.AverageTxBytes), nil
}

func (r *myStatisticsResolver) AverageOrdersPerBlock(ctx context.Context, obj *types.Statistics) (int, error) {
	return int(obj.AverageOrdersPerBlock), nil
}

func (r *myStatisticsResolver) TradesPerSecond(ctx context.Context, obj *types.Statistics) (int, error) {
	return int(obj.TradesPerSecond), nil
}

func (r *myStatisticsResolver) OrdersPerSecond(ctx context.Context, obj *types.Statistics) (int, error) {
	return int(obj.OrdersPerSecond), nil
}

func (r *myStatisticsResolver) TotalMarkets(ctx context.Context, obj *types.Statistics) (int, error) {
	return int(obj.TotalMarkets), nil
}

func (r *myStatisticsResolver) TotalAmendOrder(ctx context.Context, obj *types.Statistics) (int, error) {
	return int(obj.TotalAmendOrder), nil
}

func (r *myStatisticsResolver) TotalCancelOrder(ctx context.Context, obj *types.Statistics) (int, error) {
	return int(obj.TotalCancelOrder), nil
}

func (r *myStatisticsResolver) TotalCreateOrder(ctx context.Context, obj *types.Statistics) (int, error) {
	return int(obj.TotalCreateOrder), nil
}

func (r *myStatisticsResolver) TotalOrders(ctx context.Context, obj *types.Statistics) (int, error) {
	return int(obj.TotalOrders), nil
}

func (r *myStatisticsResolver) TotalTrades(ctx context.Context, obj *types.Statistics) (int, error) {
	return int(obj.TotalTrades), nil
}

func (r *myStatisticsResolver) BlockDuration(ctx context.Context, obj *types.Statistics) (int, error) {
	return int(obj.BlockDuration), nil
}

func (r *myStatisticsResolver) CandleSubscriptions(ctx context.Context, obj *types.Statistics) (int, error) {
	return int(obj.CandleSubscriptions), nil
}

func (r *myStatisticsResolver) MarketDepthSubscriptions(ctx context.Context, obj *types.Statistics) (int, error) {
	return int(obj.MarketDepthSubscriptions), nil
}

func (r *myStatisticsResolver) OrderSubscriptions(ctx context.Context, obj *types.Statistics) (int, error) {
	return int(obj.OrderSubscriptions), nil
}

func (r *myStatisticsResolver) PositionsSubscriptions(ctx context.Context, obj *types.Statistics) (int, error) {
	return int(obj.PositionsSubscriptions), nil
}

func (r *myStatisticsResolver) TradeSubscriptions(ctx context.Context, obj *types.Statistics) (int, error) {
	return int(obj.TradeSubscriptions), nil
}

func getParty(ctx context.Context, log *logging.Logger, client TradingDataClient, id string) (*types.Party, error) {
	if len(id) == 0 {
		return nil, errors.New("invalid party id")
	}
	res, err := client.PartyByID(ctx, &protoapi.PartyByIDRequest{PartyID: id})
	if err != nil {
		log.Error("tradingData client", logging.Error(err))
		return nil, customErrorFromStatus(err)
	}
	return res.Party, nil
}
