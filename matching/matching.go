package matching

import (
	"vega/proto"
)

type MatchingEngine interface {
	CreateMarket(id string)
	SubmitOrder(order *msg.Order) (*msg.OrderConfirmation, msg.OrderError)
	DeleteOrder(order *msg.Order)
}

type matchingEngine struct {
	//matchingEngine MatchingEngine
	markets map[string]*OrderBook
	config  Config
}

func NewMatchingEngine() MatchingEngine {
	return &matchingEngine{markets: make(map[string]*OrderBook)}
}

func (me *matchingEngine) CreateMarket(marketName string) {
	if _, exists := me.markets[marketName]; !exists {
		book := NewBook(marketName, me.config)
		me.markets[marketName] = book
	}
}

func (me *matchingEngine) SubmitOrder(order *msg.Order) (*msg.OrderConfirmation, msg.OrderError) {
	market, exists := me.markets[order.Market]
	if !exists {
		return nil, msg.OrderError_INVALID_MARKET_ID
	}

	confirmationMessage, err := market.AddOrder(order)
	if err != msg.OrderError_NONE {
		return nil, err
	}
	return confirmationMessage, msg.OrderError_NONE
}

func (me *matchingEngine) DeleteOrder(order *msg.Order) {
	if market, exists := me.markets[order.Market]; exists {
		market.RemoveOrder(order)
	}
}