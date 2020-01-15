package api

import (
	"context"
	"fmt"
	"sync"

	"code.vegaprotocol.io/vega/auth"
	"code.vegaprotocol.io/vega/logging"
	"code.vegaprotocol.io/vega/monitoring"
	types "code.vegaprotocol.io/vega/proto"
	protoapi "code.vegaprotocol.io/vega/proto/api"
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

var (
	// ErrAuthDisabled signals to the caller that the authentication is disabled
	ErrAuthDisabled = errors.New("auth disabled")
	// ErrInvalidCredentials signals that the credentials specified by the client are invalid
	ErrInvalidCredentials = errors.New("invalid credentials")
	// ErrInvalidMarketID signals that the market ID does not exists
	ErrInvalidMarketID = errors.New("invalid market ID")
	// ErrAuthRequired signals that authentication is required to reach this endpoint
	ErrAuthRequired = errors.New("auth required")
	// ErrMissingOrder signals that the actual payload is expected to contains an order
	ErrMissingOrder = errors.New("missing order in request payload")
	// ErrMissingTraderID signals that the payload is expected to contain a trader id
	ErrMissingTraderID = errors.New("missing trader id")
	// ErrMissingPartyID signals that the payload is expected to contain a party id
	ErrMissingPartyID = errors.New("missing party id")
	// ErrMissingToken signals that a token was required but is missing with this request
	ErrMissingToken = errors.New("missing token")
	// ErrMalformedRequest signals that the request was malformed
	ErrMalformedRequest = errors.New("malformed request")
	// ErrInvalidWithdrawAmount signals that the amount of money to withdraw is invalid
	// usually the party specified an amount of 0
	ErrInvalidWithdrawAmount = errors.New("invalid withdraw amount (must be > 0)")
	// ErrMissingAsset signals that an asset was required but not specified
	ErrMissingAsset = errors.New("missing asset")
	// ErrInvalidPartyID signals that the partyID does not exist
	ErrInvalidPartyID = errors.New("invalid party id")
)

// TradeOrderService ...
//go:generate go run github.com/golang/mock/mockgen -destination mocks/trade_order_service_mock.go -package mocks code.vegaprotocol.io/vega/api TradeOrderService
type TradeOrderService interface {
	CreateOrder(ctx context.Context, submission *types.OrderSubmission) (*types.PendingOrder, error)
	CancelOrder(ctx context.Context, cancellation *types.OrderCancellation) (*types.PendingOrder, error)
	AmendOrder(ctx context.Context, amendment *types.OrderAmendment) (*types.PendingOrder, error)
}

// AccountService ...
//go:generate go run github.com/golang/mock/mockgen -destination mocks/account_service_mock.go -package mocks code.vegaprotocol.io/vega/api  AccountService
type AccountService interface {
	NotifyTraderAccount(ctx context.Context, notify *types.NotifyTraderAccount) (bool, error)
	Withdraw(context.Context, *types.Withdraw) (bool, error)
}

type tradingService struct {
	log               *logging.Logger
	tradeOrderService TradeOrderService
	accountService    AccountService
	marketService     MarketService
	statusChecker     *monitoring.Status

	authEnabled bool
	parties     []auth.PartyInfo
	mu          sync.Mutex
}

func (s *tradingService) UpdateParties(parties []auth.PartyInfo) {
	s.mu.Lock()
	s.parties = parties
	s.mu.Unlock()
}

func (s *tradingService) validateToken(partyID string, tkn string) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	for _, v := range s.parties {
		if v.ID == partyID && v.Token == tkn {
			return nil
		}
	}
	return ErrInvalidCredentials
}

func (s *tradingService) CheckToken(
	ctx context.Context, req *protoapi.CheckTokenRequest,
) (*protoapi.CheckTokenResponse, error) {
	if req == nil {
		return nil, ErrMalformedRequest
	}
	if !s.authEnabled {
		return nil, ErrAuthDisabled
	}
	if len(req.PartyID) <= 0 {
		return nil, ErrMissingPartyID
	}
	if len(req.Token) <= 0 {
		return nil, ErrMissingToken
	}

	err := s.validateToken(req.PartyID, req.Token)
	if err != nil {
		if err == ErrInvalidCredentials {
			return &protoapi.CheckTokenResponse{Ok: false}, nil
		}
		return nil, err
	}

	return &protoapi.CheckTokenResponse{Ok: true}, nil
}

func (s *tradingService) SignIn(
	ctx context.Context, req *protoapi.SignInRequest,
) (*protoapi.SignInResponse, error) {
	if req == nil {
		return nil, ErrMalformedRequest
	}
	if len(req.Id) <= 0 {
		return nil, ErrInvalidCredentials
	}
	if len(req.Password) <= 0 {
		return nil, ErrInvalidCredentials
	}

	var tkn string
	saltpass := fmt.Sprintf("vega%v", req.Password)

	s.mu.Lock()
	defer s.mu.Unlock()

	if !s.authEnabled {
		return nil, ErrAuthDisabled
	}
	for _, v := range s.parties {
		if v.ID == req.Id {
			if err := bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(saltpass)); err != nil {
				s.log.Debug("invalid password",
					logging.String("user-id", v.ID),
					logging.Error(err),
				)
				return nil, ErrInvalidCredentials
			}
			tkn = v.Token
		}
	}

	if len(tkn) <= 0 {
		return nil, ErrInvalidCredentials
	}

	return &protoapi.SignInResponse{
		Token: tkn,
	}, nil
}

// CreateOrder is used to request sending an order into the VEGA platform, via consensus.
func (s *tradingService) SubmitOrder(
	ctx context.Context, req *protoapi.SubmitOrderRequest,
) (*types.PendingOrder, error) {
	if req == nil {
		return nil, ErrMalformedRequest
	}
	if s.statusChecker.ChainStatus() != types.ChainStatus_CONNECTED {
		return nil, ErrChainNotConnected
	}
	if req.Submission == nil {
		return nil, ErrMissingOrder
	}

	// check auth if required
	if s.authEnabled {
		if len(req.Token) <= 0 {
			s.log.Debug("missing token")
			return nil, errors.New("missing auth token")
		}
		if err := s.validateToken(req.Submission.PartyID, req.Token); err != nil {
			s.log.Debug("token error", logging.Error(err))
			return nil, err
		}
	}

	// Validate market early
	_, err := s.marketService.GetByID(ctx, req.Submission.MarketID)
	if err != nil {
		s.log.Error("Invalid Market ID during SubmitOrder",
			logging.String("marketID", req.Submission.MarketID),
		)
		return nil, ErrInvalidMarketID
	}

	return s.tradeOrderService.CreateOrder(ctx, req.Submission)
}

// CancelOrder is used to request cancelling an order into the VEGA platform, via consensus.
func (s *tradingService) CancelOrder(
	ctx context.Context, req *protoapi.CancelOrderRequest,
) (*types.PendingOrder, error) {
	if req == nil {
		return nil, ErrMalformedRequest
	}
	if s.statusChecker.ChainStatus() != types.ChainStatus_CONNECTED {
		return nil, ErrChainNotConnected
	}
	if req.Cancellation == nil {
		return nil, ErrMissingOrder
	}

	// check auth if required
	if s.authEnabled {
		if len(req.Token) <= 0 {
			s.log.Debug("missing token")
			return nil, errors.New("missing auth token")
		}
		if err := s.validateToken(req.Cancellation.PartyID, req.Token); err != nil {
			s.log.Debug("token error", logging.Error(err))
			return nil, err
		}
	}

	return s.tradeOrderService.CancelOrder(ctx, req.Cancellation)
}

// AmendOrder is used to request editing an order onto the VEGA platform, via consensus.
func (s *tradingService) AmendOrder(
	ctx context.Context, req *protoapi.AmendOrderRequest,
) (*types.PendingOrder, error) {
	if req == nil {
		return nil, ErrMalformedRequest
	}
	if req.Amendment == nil {
		return nil, ErrMissingOrder
	}

	// check auth if required
	if s.authEnabled {
		if len(req.Token) <= 0 {
			s.log.Debug("missing token")
			return nil, errors.New("missing auth token")
		}
		if err := s.validateToken(req.Amendment.PartyID, req.Token); err != nil {
			s.log.Debug("token error", logging.Error(err))
			return nil, err
		}
	}

	return s.tradeOrderService.AmendOrder(ctx, req.Amendment)
}

func (s *tradingService) NotifyTraderAccount(
	ctx context.Context, req *protoapi.NotifyTraderAccountRequest,
) (*protoapi.NotifyTraderAccountResponse, error) {
	if req == nil || req.Notif == nil {
		return nil, ErrMalformedRequest
	}
	if len(req.Notif.TraderID) <= 0 {
		return nil, ErrMissingTraderID
	}

	submitted, err := s.accountService.NotifyTraderAccount(ctx, req.Notif)
	if err != nil {
		return nil, err
	}

	return &protoapi.NotifyTraderAccountResponse{
		Submitted: submitted,
	}, nil
}

func (s *tradingService) Withdraw(
	ctx context.Context, req *protoapi.WithdrawRequest,
) (*protoapi.WithdrawResponse, error) {
	if len(req.Withdraw.PartyID) <= 0 {
		return nil, ErrMissingTraderID
	}
	if len(req.Withdraw.Asset) <= 0 {
		return nil, ErrMissingAsset
	}

	if req.Withdraw.Amount == 0 {
		return nil, ErrInvalidWithdrawAmount
	}

	ok, err := s.accountService.Withdraw(ctx, req.Withdraw)
	if err != nil {
		return nil, err
	}

	return &protoapi.WithdrawResponse{
		Success: ok,
	}, nil
}
