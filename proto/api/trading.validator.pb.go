// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: api/trading.proto

package api

import (
	fmt "fmt"
	math "math"

	_ "code.vegaprotocol.io/vega/proto"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	github_com_mwitkow_go_proto_validators "github.com/mwitkow/go-proto-validators"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

func (this *PropagateChainEventRequest) Validate() error {
	if this.Evt != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Evt); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Evt", err)
		}
	}
	return nil
}
func (this *PropagateChainEventResponse) Validate() error {
	return nil
}
func (this *SubmitTransactionRequest) Validate() error {
	if this.Tx != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Tx); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Tx", err)
		}
	}
	return nil
}
func (this *SubmitTransactionResponse) Validate() error {
	return nil
}
func (this *PrepareWithdrawRequest) Validate() error {
	if this.Withdraw != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Withdraw); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Withdraw", err)
		}
	}
	return nil
}
func (this *PrepareWithdrawResponse) Validate() error {
	return nil
}
func (this *PrepareSubmitOrderResponse) Validate() error {
	return nil
}
func (this *PrepareCancelOrderResponse) Validate() error {
	return nil
}
func (this *PrepareAmendOrderResponse) Validate() error {
	return nil
}
func (this *PrepareSubmitOrderRequest) Validate() error {
	if this.Submission != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Submission); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Submission", err)
		}
	}
	return nil
}
func (this *PrepareCancelOrderRequest) Validate() error {
	if this.Cancellation != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Cancellation); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Cancellation", err)
		}
	}
	return nil
}
func (this *PrepareAmendOrderRequest) Validate() error {
	if this.Amendment != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Amendment); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Amendment", err)
		}
	}
	return nil
}
func (this *AssetsRequest) Validate() error {
	return nil
}
func (this *AssetsResponse) Validate() error {
	for _, item := range this.Assets {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Assets", err)
			}
		}
	}
	return nil
}
func (this *AssetByIDRequest) Validate() error {
	if this.ID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ID", fmt.Errorf(`value '%v' must not be an empty string`, this.ID))
	}
	return nil
}
func (this *AssetByIDResponse) Validate() error {
	if this.Asset != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Asset); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Asset", err)
		}
	}
	return nil
}
func (this *GetNodeSignaturesAggregateRequest) Validate() error {
	if this.ID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ID", fmt.Errorf(`value '%v' must not be an empty string`, this.ID))
	}
	return nil
}
func (this *GetNodeSignaturesAggregateResponse) Validate() error {
	for _, item := range this.Signatures {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Signatures", err)
			}
		}
	}
	return nil
}
func (this *OptionalProposalState) Validate() error {
	return nil
}
func (this *GetProposalsRequest) Validate() error {
	if this.SelectInState != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SelectInState); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SelectInState", err)
		}
	}
	return nil
}
func (this *GetProposalsResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetProposalsByPartyRequest) Validate() error {
	if this.PartyID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyID", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyID))
	}
	if this.SelectInState != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SelectInState); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SelectInState", err)
		}
	}
	return nil
}
func (this *GetProposalsByPartyResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetVotesByPartyRequest) Validate() error {
	if this.PartyID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyID", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyID))
	}
	return nil
}
func (this *GetVotesByPartyResponse) Validate() error {
	for _, item := range this.Votes {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Votes", err)
			}
		}
	}
	return nil
}
func (this *GetNewMarketProposalsRequest) Validate() error {
	if this.SelectInState != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SelectInState); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SelectInState", err)
		}
	}
	return nil
}
func (this *GetNewMarketProposalsResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetUpdateMarketProposalsRequest) Validate() error {
	if this.MarketID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketID", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketID))
	}
	if this.SelectInState != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SelectInState); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SelectInState", err)
		}
	}
	return nil
}
func (this *GetUpdateMarketProposalsResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetNetworkParametersProposalsRequest) Validate() error {
	if this.SelectInState != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SelectInState); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SelectInState", err)
		}
	}
	return nil
}
func (this *GetNetworkParametersProposalsResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetNewAssetProposalsRequest) Validate() error {
	if this.SelectInState != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.SelectInState); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("SelectInState", err)
		}
	}
	return nil
}
func (this *GetNewAssetProposalsResponse) Validate() error {
	for _, item := range this.Data {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
			}
		}
	}
	return nil
}
func (this *GetProposalByIDRequest) Validate() error {
	if this.ProposalID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProposalID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProposalID))
	}
	return nil
}
func (this *GetProposalByIDResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *GetProposalByReferenceRequest) Validate() error {
	if this.Reference == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Reference", fmt.Errorf(`value '%v' must not be an empty string`, this.Reference))
	}
	return nil
}
func (this *GetProposalByReferenceResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ObserveGovernanceRequest) Validate() error {
	return nil
}
func (this *ObserveGovernanceResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ObservePartyProposalsRequest) Validate() error {
	if this.PartyID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyID", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyID))
	}
	return nil
}
func (this *ObservePartyProposalsResponse) Validate() error {
	if this.Data != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Data); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Data", err)
		}
	}
	return nil
}
func (this *ObserveProposalVotesRequest) Validate() error {
	if this.ProposalID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ProposalID", fmt.Errorf(`value '%v' must not be an empty string`, this.ProposalID))
	}
	return nil
}
func (this *ObserveProposalVotesResponse) Validate() error {
	if this.Vote != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Vote); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Vote", err)
		}
	}
	return nil
}
func (this *ObservePartyVotesRequest) Validate() error {
	if this.PartyID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyID", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyID))
	}
	return nil
}
func (this *ObservePartyVotesResponse) Validate() error {
	if this.Vote != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Vote); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Vote", err)
		}
	}
	return nil
}
func (this *MarginLevelsSubscribeRequest) Validate() error {
	if this.PartyID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyID", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyID))
	}
	return nil
}
func (this *MarginLevelsSubscribeResponse) Validate() error {
	if this.MarginLevels != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MarginLevels); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MarginLevels", err)
		}
	}
	return nil
}
func (this *MarginLevelsRequest) Validate() error {
	if this.PartyID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyID", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyID))
	}
	return nil
}
func (this *MarginLevelsResponse) Validate() error {
	for _, item := range this.MarginLevels {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MarginLevels", err)
			}
		}
	}
	return nil
}
func (this *MarketsDataSubscribeRequest) Validate() error {
	return nil
}
func (this *MarketsDataSubscribeResponse) Validate() error {
	if this.MarketData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MarketData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MarketData", err)
		}
	}
	return nil
}
func (this *MarketDataByIDRequest) Validate() error {
	if this.MarketID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketID", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketID))
	}
	return nil
}
func (this *MarketDataByIDResponse) Validate() error {
	if this.MarketData != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MarketData); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MarketData", err)
		}
	}
	return nil
}
func (this *MarketsDataRequest) Validate() error {
	return nil
}
func (this *MarketsDataResponse) Validate() error {
	for _, item := range this.MarketsData {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("MarketsData", err)
			}
		}
	}
	return nil
}
func (this *LastTradeRequest) Validate() error {
	if this.MarketID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketID", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketID))
	}
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
	if this.MarketID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketID", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketID))
	}
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
	if this.PartyID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyID", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyID))
	}
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
func (this *PartiesRequest) Validate() error {
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
func (this *AccountsSubscribeRequest) Validate() error {
	return nil
}
func (this *AccountsSubscribeResponse) Validate() error {
	if this.Account != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Account); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Account", err)
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
	if this.MarketID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketID", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketID))
	}
	return nil
}
func (this *CandlesSubscribeResponse) Validate() error {
	if this.Candle != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Candle); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Candle", err)
		}
	}
	return nil
}
func (this *MarketDepthSubscribeRequest) Validate() error {
	if this.MarketID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketID", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketID))
	}
	return nil
}
func (this *MarketDepthSubscribeResponse) Validate() error {
	if this.MarketDepth != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MarketDepth); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MarketDepth", err)
		}
	}
	return nil
}
func (this *MarketDepthUpdatesSubscribeRequest) Validate() error {
	if this.MarketID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketID", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketID))
	}
	return nil
}
func (this *MarketDepthUpdatesSubscribeResponse) Validate() error {
	if this.Update != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Update); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Update", err)
		}
	}
	return nil
}
func (this *PositionsSubscribeRequest) Validate() error {
	return nil
}
func (this *PositionsSubscribeResponse) Validate() error {
	if this.Position != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Position); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Position", err)
		}
	}
	return nil
}
func (this *OrdersByMarketRequest) Validate() error {
	if this.MarketID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketID", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketID))
	}
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
	if this.PartyID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyID", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyID))
	}
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
func (this *OrderByMarketAndIDRequest) Validate() error {
	if this.MarketID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketID", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketID))
	}
	if this.OrderID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("OrderID", fmt.Errorf(`value '%v' must not be an empty string`, this.OrderID))
	}
	return nil
}
func (this *OrderByMarketAndIDResponse) Validate() error {
	if this.Order != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Order); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Order", err)
		}
	}
	return nil
}
func (this *OrderByReferenceRequest) Validate() error {
	if this.Reference == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("Reference", fmt.Errorf(`value '%v' must not be an empty string`, this.Reference))
	}
	return nil
}
func (this *OrderByReferenceResponse) Validate() error {
	if this.Order != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Order); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Order", err)
		}
	}
	return nil
}
func (this *MarketsRequest) Validate() error {
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
	if this.MarketID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketID", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketID))
	}
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
	if this.MarketID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketID", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketID))
	}
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
	if this.MarketID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("MarketID", fmt.Errorf(`value '%v' must not be an empty string`, this.MarketID))
	}
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
	if this.PartyID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyID", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyID))
	}
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
func (this *GetVegaTimeRequest) Validate() error {
	return nil
}
func (this *GetVegaTimeResponse) Validate() error {
	return nil
}
func (this *Pagination) Validate() error {
	return nil
}
func (this *OrdersSubscribeResponse) Validate() error {
	for _, item := range this.Orders {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Orders", err)
			}
		}
	}
	return nil
}
func (this *TradesSubscribeResponse) Validate() error {
	for _, item := range this.Trades {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Trades", err)
			}
		}
	}
	return nil
}
func (this *TransferResponsesSubscribeRequest) Validate() error {
	return nil
}
func (this *TransferResponsesSubscribeResponse) Validate() error {
	if this.Response != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Response); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Response", err)
		}
	}
	return nil
}
func (this *PartyAccountsRequest) Validate() error {
	return nil
}
func (this *PartyAccountsResponse) Validate() error {
	for _, item := range this.Accounts {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Accounts", err)
			}
		}
	}
	return nil
}
func (this *MarketAccountsRequest) Validate() error {
	return nil
}
func (this *MarketAccountsResponse) Validate() error {
	for _, item := range this.Accounts {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Accounts", err)
			}
		}
	}
	return nil
}
func (this *FeeInfrastructureAccountsRequest) Validate() error {
	return nil
}
func (this *FeeInfrastructureAccountsResponse) Validate() error {
	for _, item := range this.Accounts {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Accounts", err)
			}
		}
	}
	return nil
}
func (this *PrepareProposalRequest) Validate() error {
	if this.PartyID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyID", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyID))
	}
	if nil == this.Proposal {
		return github_com_mwitkow_go_proto_validators.FieldError("Proposal", fmt.Errorf("message must exist"))
	}
	if this.Proposal != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Proposal); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Proposal", err)
		}
	}
	return nil
}
func (this *PrepareProposalResponse) Validate() error {
	if this.PendingProposal != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.PendingProposal); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("PendingProposal", err)
		}
	}
	return nil
}
func (this *PrepareVoteRequest) Validate() error {
	if nil == this.Vote {
		return github_com_mwitkow_go_proto_validators.FieldError("Vote", fmt.Errorf("message must exist"))
	}
	if this.Vote != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Vote); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Vote", err)
		}
	}
	return nil
}
func (this *PrepareVoteResponse) Validate() error {
	if this.Vote != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Vote); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Vote", err)
		}
	}
	return nil
}
func (this *PrepareLiquidityProvisionRequest) Validate() error {
	if nil == this.Submission {
		return github_com_mwitkow_go_proto_validators.FieldError("Submission", fmt.Errorf("message must exist"))
	}
	if this.Submission != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Submission); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Submission", err)
		}
	}
	return nil
}
func (this *PrepareLiquidityProvisionResponse) Validate() error {
	return nil
}
func (this *OrderByIDRequest) Validate() error {
	return nil
}
func (this *OrderByIDResponse) Validate() error {
	if this.Order != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Order); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Order", err)
		}
	}
	return nil
}
func (this *OrderVersionsByIDRequest) Validate() error {
	if this.Pagination != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Pagination); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Pagination", err)
		}
	}
	return nil
}
func (this *OrderVersionsByIDResponse) Validate() error {
	for _, item := range this.Orders {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Orders", err)
			}
		}
	}
	return nil
}
func (this *EstimateFeeRequest) Validate() error {
	if this.Order != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Order); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Order", err)
		}
	}
	return nil
}
func (this *EstimateFeeResponse) Validate() error {
	if this.Fee != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Fee); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Fee", err)
		}
	}
	return nil
}
func (this *EstimateMarginRequest) Validate() error {
	if this.Order != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Order); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Order", err)
		}
	}
	return nil
}
func (this *EstimateMarginResponse) Validate() error {
	if this.MarginLevels != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.MarginLevels); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("MarginLevels", err)
		}
	}
	return nil
}
func (this *ObserveEventBusRequest) Validate() error {
	return nil
}
func (this *ObserveEventBusResponse) Validate() error {
	for _, item := range this.Events {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Events", err)
			}
		}
	}
	return nil
}
func (this *StatisticsRequest) Validate() error {
	return nil
}
func (this *StatisticsResponse) Validate() error {
	if this.Statistics != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Statistics); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Statistics", err)
		}
	}
	return nil
}
func (this *WithdrawalsRequest) Validate() error {
	if this.PartyID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyID", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyID))
	}
	return nil
}
func (this *WithdrawalsResponse) Validate() error {
	for _, item := range this.Withdrawals {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Withdrawals", err)
			}
		}
	}
	return nil
}
func (this *WithdrawalRequest) Validate() error {
	if this.ID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ID", fmt.Errorf(`value '%v' must not be an empty string`, this.ID))
	}
	return nil
}
func (this *WithdrawalResponse) Validate() error {
	if this.Withdrawal != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Withdrawal); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Withdrawal", err)
		}
	}
	return nil
}
func (this *ERC20WithdrawalApprovalRequest) Validate() error {
	if this.WithdrawalID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("WithdrawalID", fmt.Errorf(`value '%v' must not be an empty string`, this.WithdrawalID))
	}
	return nil
}
func (this *ERC20WithdrawalApprovalResponse) Validate() error {
	return nil
}
func (this *DepositsRequest) Validate() error {
	if this.PartyID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("PartyID", fmt.Errorf(`value '%v' must not be an empty string`, this.PartyID))
	}
	return nil
}
func (this *DepositsResponse) Validate() error {
	for _, item := range this.Deposits {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("Deposits", err)
			}
		}
	}
	return nil
}
func (this *DepositRequest) Validate() error {
	if this.ID == "" {
		return github_com_mwitkow_go_proto_validators.FieldError("ID", fmt.Errorf(`value '%v' must not be an empty string`, this.ID))
	}
	return nil
}
func (this *DepositResponse) Validate() error {
	if this.Deposit != nil {
		if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(this.Deposit); err != nil {
			return github_com_mwitkow_go_proto_validators.FieldError("Deposit", err)
		}
	}
	return nil
}
func (this *NetworkParametersRequest) Validate() error {
	return nil
}
func (this *NetworkParametersResponse) Validate() error {
	for _, item := range this.NetworkParameters {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("NetworkParameters", err)
			}
		}
	}
	return nil
}
func (this *LiquidityProvisionsRequest) Validate() error {
	return nil
}
func (this *LiquidityProvisionsResponse) Validate() error {
	for _, item := range this.LiquidityProvisions {
		if item != nil {
			if err := github_com_mwitkow_go_proto_validators.CallValidatorIfExists(item); err != nil {
				return github_com_mwitkow_go_proto_validators.FieldError("LiquidityProvisions", err)
			}
		}
	}
	return nil
}
