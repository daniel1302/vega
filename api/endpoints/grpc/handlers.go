package grpc

import (
	"context"
	"vega/proto"
	"vega/api"
	"time"
	"github.com/pkg/errors"
)

type Handlers struct {
	OrderService api.OrderService
	TradeService api.TradeService
}

// If no limit is provided at the gRPC API level, the system will use this limit instead.
// (Prevent returning all results every time a careless query is made)
const defaultLimit = uint64(1000)

// CreateOrder is used to send an order into the VEGA platform, via consensus. TODO pre-validation.
func (h *Handlers) CreateOrder(ctx context.Context, order *msg.Order) (*api.OrderResponse, error) {
	success, err := h.OrderService.CreateOrder(ctx, order)
	return &api.OrderResponse{Success: success}, err
}

// OrdersByMarket provides a list of orders for a given market. Optional limits can be provided. Most recent first.
func (h *Handlers) OrdersByMarket(ctx context.Context, request *api.OrdersByMarketRequest) (*api.OrdersByMarketResponse, error) {
	if request.Market == "" {
		return nil, errors.New("Market empty or missing")
	}
	limit := defaultLimit
	if request.Params != nil && request.Params.Limit > 0 {
		limit = request.Params.Limit
	}
	orders, err := h.OrderService.GetByMarket(ctx, request.Market, limit)
	if err != nil {
		return nil, err
	}
	var response = &api.OrdersByMarketResponse{}
	if len(orders) > 0 {
	   response.Orders = orders
	}
	return response, nil
}

// OrdersByParty provides a list of orders for a given party. Optional limits can be provided. Most recent first.
func (h *Handlers) OrdersByParty(ctx context.Context, request *api.OrdersByPartyRequest) (*api.OrdersByPartyResponse, error) {
	if request.Party == "" {
		return nil, errors.New("Party empty or missing")
	}
	limit := defaultLimit
	if request.Params != nil && request.Params.Limit > 0 {
		limit = request.Params.Limit
	}
	orders, err := h.OrderService.GetByParty(ctx, request.Party, limit)
	if err != nil {
		return nil, err
	}
	var response = &api.OrdersByPartyResponse{}
	if len(orders) > 0 {
		response.Orders = orders
	}
	return response, nil
}

// Markets provides a list of all current markets that exist on the VEGA platform.
func (h *Handlers) Markets(ctx context.Context, request *api.MarketsRequest) (*api.MarketsResponse, error) {
	markets, err := h.OrderService.GetMarkets(ctx)
	if err != nil {
		return nil, err
	}
	var response = &api.MarketsResponse{}
	if len(markets) > 0 {
		response.Markets = markets
	}
	return response, nil
}

// OrdersByMarketAndId searches for the given order by Id and Market. If found it will return
// an Order msg otherwise it will return an error.
func (h *Handlers) OrderByMarketAndId(ctx context.Context, request *api.OrderByMarketAndIdRequest) (*api.OrderByMarketAndIdResponse, error) {
	if request.Market == "" {
		return nil, errors.New("Market empty or missing")
	}
	if request.Id == "" {
		return nil, errors.New("Id empty or missing")
	}
	order, err := h.OrderService.GetByMarketAndId(ctx, request.Market, request.Id)
	if err != nil {
		return nil, err
	}
	var response = &api.OrderByMarketAndIdResponse{}
	response.Order = order
	return response, nil
}

// TradeCandles returns trade open/close/volume data for the given time period and interval.
// It will fill in any tradeless intervals with zero based candles. Since time period must be in RFC3339 string format.
func (h *Handlers) TradeCandles(ctx context.Context, request *api.TradeCandlesRequest) (*api.TradeCandlesResponse, error) {
	market := request.Market
	if market == "" {
		return nil, errors.New("Market empty or missing")
	}
	if request.Since == "" {
		request.Since = "2018-07-09T12:00:00Z"
	}
	since, err := time.Parse(time.RFC3339, request.Since)
	if err != nil {
		return nil, err
	}
	interval := request.Interval
	if interval < 1 {
		interval = 2
	}
	res, err := h.TradeService.GetCandles(ctx, market, since, interval)
	if err != nil {
		return nil, err
	}
	var response = &api.TradeCandlesResponse{}
	if len(res.Candles) > 0 {
		response.Candles = res.Candles
	}
	return response, nil
}


func (h *Handlers) OrderBookDepth(ctx context.Context, request *api.OrderBookDepthRequest) (*api.OrderBookDepthResponse, error) {
	if request.Market == "" {
		return nil, errors.New("Market empty or missing")
	}
	depth, err := h.OrderService.GetOrderBookDepth(ctx, request.Market)
	if err != nil {
		return nil, err
	}
	var response = &api.OrderBookDepthResponse{}
	response.Buy = depth.Buy
	response.Name = depth.Name
	response.Sell = depth.Sell
	return response, nil
}