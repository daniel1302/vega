package rest

import (
	"fmt"
	"net/http"
	"vega/api"
	"vega/core"
)

type restServer struct{
	orderService api.OrderService
	tradeService api.TradeService
}

func NewRestServer(vega *core.Vega) *restServer {
	return &restServer{
		orderService: vega.OrdersService,
		tradeService: vega.TradesService,
	}
}

func (s *restServer) Start() {
	var port= 3003
	var addr= fmt.Sprintf(":%d", port)
	fmt.Printf("Starting REST based HTTP server on port %d...\n", port)
	router := NewRouter(s.orderService, s.tradeService)
	http.ListenAndServe(addr, router)
}

