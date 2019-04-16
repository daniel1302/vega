// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: proto/vega.proto

package proto

import (
	fmt "fmt"
	math "math"

	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *Party) Validate() error {
	for _, item := range this.Positions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Positions", err)
			}
		}
	}
	return nil
}
func (this *RiskFactor) Validate() error {
	return nil
}
func (this *RiskResult) Validate() error {
	// Validation of proto3 map<> fields is unsupported.
	// Validation of proto3 map<> fields is unsupported.
	return nil
}
func (this *Order) Validate() error {
	return nil
}
func (this *OrderCancellationConfirmation) Validate() error {
	if this.Order != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Order); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Order", err)
		}
	}
	return nil
}
func (this *OrderConfirmation) Validate() error {
	if this.Order != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Order); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Order", err)
		}
	}
	for _, item := range this.Trades {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Trades", err)
			}
		}
	}
	for _, item := range this.PassiveOrdersAffected {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("PassiveOrdersAffected", err)
			}
		}
	}
	return nil
}
func (this *Trade) Validate() error {
	return nil
}
func (this *TradeSet) Validate() error {
	for _, item := range this.Trades {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Trades", err)
			}
		}
	}
	return nil
}
func (this *MarketData) Validate() error {
	return nil
}
func (this *Candle) Validate() error {
	return nil
}
func (this *PriceLevel) Validate() error {
	return nil
}
func (this *MarketDepth) Validate() error {
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
	return nil
}
func (this *MarketPosition) Validate() error {
	return nil
}
func (this *Statistics) Validate() error {
	return nil
}
func (this *OrderAmendment) Validate() error {
	if !(this.Size > 0) {
		return github_com_mwitkow_go_proto_validators.FieldError("Size_", fmt.Errorf(`value '%v' must be greater than '0'`, this.Size))
	}
	return nil
}
func (this *OrderSubmission) Validate() error {
	return nil
}
func (this *OrderCancellation) Validate() error {
	return nil
}
