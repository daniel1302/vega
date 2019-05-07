// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/api/trading.proto

package api

import (
	fmt "fmt"
	math "math"

	_ "code.vegaprotocol.io/vega/proto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/golang/protobuf/ptypes/empty"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *OrderResponse) Validate() error {
	return nil
}
func (this *LastTradeRequest) Validate() error {
	return nil
}
func (this *LastTradeResponse) Validate() error {
	if this.Trade != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Trade); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Trade", err)
		}
	}
	return nil
}
func (this *MarketByIDRequest) Validate() error {
	return nil
}
func (this *MarketByIDResponse) Validate() error {
	if this.Market != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Market); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Market", err)
		}
	}
	return nil
}
func (this *PartyByIDRequest) Validate() error {
	return nil
}
func (this *PartyByIDResponse) Validate() error {
	if this.Party != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Party); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Party", err)
		}
	}
	return nil
}
func (this *PartiesResponse) Validate() error {
	for _, item := range this.Parties {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Parties", err)
			}
		}
	}
	return nil
}
func (this *TradesByPartyRequest) Validate() error {
	if this.Pagination != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pagination); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pagination", err)
		}
	}
	return nil
}
func (this *TradesByPartyResponse) Validate() error {
	for _, item := range this.Trades {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Trades", err)
			}
		}
	}
	return nil
}
func (this *TradesByOrderRequest) Validate() error {
	return nil
}
func (this *TradesByOrderResponse) Validate() error {
	for _, item := range this.Trades {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Trades", err)
			}
		}
	}
	return nil
}
func (this *OrdersSubscribeRequest) Validate() error {
	return nil
}
func (this *TradesSubscribeRequest) Validate() error {
	return nil
}
func (this *CandlesSubscribeRequest) Validate() error {
	return nil
}
func (this *MarketDepthSubscribeRequest) Validate() error {
	return nil
}
func (this *PositionsSubscribeRequest) Validate() error {
	return nil
}
func (this *OrdersByMarketRequest) Validate() error {
	if this.Pagination != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pagination); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pagination", err)
		}
	}
	return nil
}
func (this *OrdersByMarketResponse) Validate() error {
	for _, item := range this.Orders {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Orders", err)
			}
		}
	}
	return nil
}
func (this *OrdersByPartyRequest) Validate() error {
	if this.Pagination != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pagination); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pagination", err)
		}
	}
	return nil
}
func (this *OrdersByPartyResponse) Validate() error {
	for _, item := range this.Orders {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Orders", err)
			}
		}
	}
	return nil
}
func (this *OrderByMarketAndIdRequest) Validate() error {
	return nil
}
func (this *OrderByMarketAndIdResponse) Validate() error {
	if this.Order != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Order); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Order", err)
		}
	}
	return nil
}
func (this *MarketsResponse) Validate() error {
	for _, item := range this.Markets {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Markets", err)
			}
		}
	}
	return nil
}
func (this *CandlesRequest) Validate() error {
	if !(this.SinceTimestamp > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("SinceTimestamp", fmt.Errorf(`value '%v' must be greater than '0'`, this.SinceTimestamp))
	}
	return nil
}
func (this *CandlesResponse) Validate() error {
	for _, item := range this.Candles {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Candles", err)
			}
		}
	}
	return nil
}
func (this *MarketDepthRequest) Validate() error {
	return nil
}
func (this *MarketDepthResponse) Validate() error {
	for _, item := range this.Buy {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Buy", err)
			}
		}
	}
	for _, item := range this.Sell {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Sell", err)
			}
		}
	}
	if this.LastTrade != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.LastTrade); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("LastTrade", err)
		}
	}
	return nil
}
func (this *TradesByMarketRequest) Validate() error {
	if this.Pagination != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pagination); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pagination", err)
		}
	}
	return nil
}
func (this *TradesByMarketResponse) Validate() error {
	for _, item := range this.Trades {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Trades", err)
			}
		}
	}
	return nil
}
func (this *PositionsByPartyRequest) Validate() error {
	return nil
}
func (this *PositionsByPartyResponse) Validate() error {
	for _, item := range this.Positions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Positions", err)
			}
		}
	}
	return nil
}
func (this *VegaTimeResponse) Validate() error {
	return nil
}
func (this *Pagination) Validate() error {
	return nil
}
func (this *OrdersStream) Validate() error {
	for _, item := range this.Orders {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Orders", err)
			}
		}
	}
	return nil
}
func (this *TradesStream) Validate() error {
	for _, item := range this.Trades {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Trades", err)
			}
		}
	}
	return nil
}
