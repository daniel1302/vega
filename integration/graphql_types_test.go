package integration_test

import (
	"strings"
	"time"
)

// These types map to the results of GraphQL queries on the data node schema; it could probably be
// autogenerated from schema.graphql, but I haven't found a tool that does it nicely.

type Market struct {
	Id                            HexString
	Name                          string
	Fees                          Fees
	Trades                        []Trade
	TradableInstrument            TradableInstrument
	DecimalPlaces                 int
	OpeningAuction                AuctionDuration
	PriceMonitoringSettings       PriceMonitoringSettings
	LiquidityMonitoringParameters LiquidityMonitoringParameters
	TradingMode                   string
	State                         string
	Proposal                      Proposal
	Orders                        []Order
	Accounts                      []Account
}

type Fees struct {
	MakerFee          string
	InfrastructureFee string
	LiquidityFee      string
}

type TradableInstrument struct {
	Instrument Instrument
	// RiskModel        RiskModel
	MarginCalculator MarginCalculator
}

type Instrument struct {
	Id       string
	Code     string
	Name     string
	Metadata Metadata
}

type MarginCalculator struct {
	ScalingFactors ScalingFactors
}

type ScalingFactors struct {
	SearchLevel       float64
	InitialMargin     float64
	CollateralRelease float64
}

type Metadata struct {
	Tags []string
}

type AuctionDuration struct {
	DurationSecs int
	Volume       int
}

type Trade struct {
	Id                 HexString
	Price              string
	Size               string
	CreatedAt          TimeString
	Market             Market
	BuyOrder           HexString
	SellOrder          HexString
	Buyer              Party
	Seller             Party
	Aggressor          string
	Type               string
	BuyerFee           TradeFee
	SellerFee          TradeFee
	BuyerAuctionBatch  int
	SellerAuctionBatch int
}

type TradeFee struct {
	MakerFee          string
	InfrastructureFee string
	LiquidityFee      string
}

type PriceMonitoringSettings struct {
	Parameters          PriceMonitoringParameters
	UpdateFrequencySecs int
}

type PriceMonitoringParameters struct {
	Triggers []PriceMonitoringTrigger
}

type PriceMonitoringTrigger struct {
	HorizonSecs          int
	Probability          float64
	AuctionExtensionSecs float64
}

type LiquidityMonitoringParameters struct {
	TargetStakeParameters TargetStakeParameters
	TriggeringRatio       float64
}

type TargetStakeParameters struct {
	TimeWindow    int
	ScalingFactor float64
}

type Proposal struct {
	Id              HexString
	Reference       string
	Party           Party
	State           string
	Datetime        TimeString
	Terms           ProposalTerms
	Votes           ProposalVotes
	RejectionReason string
}

type Party struct {
	Id                 HexString
	Orders             []Order
	Trades             []Trade
	Accounts           []Account
	Proposals          []Proposal
	Votes              []Vote
	LiquidityProvision []LiquidityProvision
	Positions          []Position
	// TODO:
	// Margins []MarginLevels
	Withdrawals []Withdrawal
	Deposits    []Deposit
	// Delegations []Delegation
	// Stake PartyStake
	// Rewards []Reward
	// RewardSummaries []RewardSummary
}

type ProposalTerms struct {
	ClosingDateTime   TimeString
	EnactmentDatetime TimeString
	// TODO: Change (can't to ...on Foo yet)
}

type ProposalVotes struct {
	Yes ProposalVoteSide
	No  ProposalVoteSide
}

type ProposalVoteSide struct {
	Votes       []Vote
	TotalNumber string
	TotalWeight string
	TotalTokens string
}

type Vote struct {
	Value                  string
	Party                  Party
	Datetime               TimeString
	ProposalId             HexString
	GovernanceTokenBalance string
	GovernanceTokenWeight  string
}

type Order struct {
	Id                 HexString
	Price              string
	Side               string
	Timeinforce        string
	Market             Market
	Size               string
	Remaining          string
	Party              Party
	CreatedAt          TimeString
	ExpiresAt          TimeString
	Status             string
	Reference          string
	Trades             []Trade
	Type               string
	RejectionReason    string
	Version            string
	UpdatedAt          string
	PeggedOrder        PeggedOrder
	LiquidityProvision LiquidityProvision
}

type PeggedOrder struct {
	Reference string
	Offset    string
}

type LiquidityProvision struct {
	Id               HexString
	Party            Party
	CreatedAt        TimeString
	UpdatedAt        TimeString
	Market           Market
	CommitmentAmount string
	Fee              string
	Sells            []LiquidityOrderReference
	Buys             []LiquidityOrderReference
	Version          string
	Status           string
	Reference        string
}

type LiquidityOrderReference struct {
	Order          Order
	LiquidityOrder LiquidityOrder
}

type LiquidityOrder struct {
	Reference  PeggedReference
	Proportion int
	Offset     string
}

type PeggedReference = HexString

type Account struct {
	Balance string
	Asset   Asset
	Type    string
	Market  Market
}

type Asset struct {
	Id          HexString
	Name        string
	Symbol      string
	TotalSupply string
	Decimals    int
	Quantum     string
	// TODO: source
	InfrastructureFeeAccount *Account
	GlobalRewardPoolAccount  *Account
}

// ----------------------------------------------------------------------------
// Some wrappers around standard types to provide non-standard comparisons,
// where the output from the API might differ slightly but we don't care.

type HexString string

func (s HexString) Equal(other HexString) bool {
	return strings.ToLower(string(s)) == strings.ToLower(string(other))
}

type TimeString string

func (s TimeString) Equal(other TimeString) bool {
	if s == "" && other == "" {
		return true
	}
	// Postgres doesn't store nanoseconds, so only compare the millisecond portion
	t1, err1 := time.Parse(time.RFC3339Nano, string(s))
	t2, err2 := time.Parse(time.RFC3339Nano, string(other))

	if err1 != nil || err2 != nil {
		return false
	}

	t1t := t1.Truncate(time.Microsecond)
	t2t := t2.Truncate(time.Microsecond)
	if t1t != t2t {
		_ = "foo"
	}
	return t1t == t2t
}

type Deposit struct {
	ID                HexString
	Party             Party
	Amount            string
	Asset             Asset
	Status            string
	CreatedTimestamp  TimeString
	CreditedTimestamp TimeString
	TxHash            string
}

type NetworkParameter struct {
	Key   string
	Value string
}

type Epoch struct {
	ID          HexString
	Timestamps  EpochTimestamps
	Delegations []Delegation
}

type EpochTimestamps struct {
	Start  string
	End    string
	Expiry string
}

type Delegation struct {
	Amount string
	Party  Party
	Node   Node
	Epoch  int
}

type NodeData struct {
	StakedTotal     string
	TotalNodes      uint32
	InactiveNodes   uint32
	ValidatingNodes uint32
	Uptime          float64
}

type Node struct {
	Id                HexString
	Pubkey            string
	TmPubkey          string
	EthereumAdddress  string
	InfoUrl           string
	Location          string
	StakedByOperator  string
	StakedByDelegates string
	StakedTotal       string
	PendingStake      string
	EpochData         EpochData
	Status            string
	Delegations       []Delegation
	Name              string
	AvatarUrl         string
	RewardScore       RewardScore
	RankingScore      RankingScore
}

type EpochData struct {
	Total   int
	Offline int
	Online  int
}

type RewardScore struct {
	ValidatorNodeStatus string

	PerformanceScore  string
	MultisigScore     string
	RawValidatorScore string
	ValidatorScore    string
	NormalisedScore   string
}

type RankingScore struct {
	Status         string
	PreviousStatus string

	VotingPower      string
	StakeScore       string
	PerformanceScore string
	RankingScore     string
}

type Withdrawal struct {
	ID                 HexString
	Party              Party
	Amount             string
	Asset              Asset
	Status             string
	Ref                string
	Expiry             TimeString
	TxHash             string
	CreatedTimestamp   TimeString
	WithdrawnTimeStamp TimeString
}

type Transfer struct {
	Id              string
	From            string
	FromAccountType string
	To              string
	ToAccountType   string
	Amount          string
	Reference       string
	Status          string
	Timestamp       time.Time
	Asset           Asset
}

type Property struct {
	Name  string
	Value string
}

type OracleData struct {
	PubKeys []string
	Data    []Property
}

type PropertyKeyType = string

type PropertyKey struct {
	Name string
	Type PropertyKeyType
}

type ConditionOperator = string

type Condition struct {
	Operator ConditionOperator
	Value    string
}

type Filter struct {
	Key        PropertyKey
	Conditions []Condition
}

type OracleSpecStatus = string

type OracleSpec struct {
	ID        HexString
	CreatedAt TimeString
	UpdatedAt TimeString
	PubKeys   []string
	Filters   []Filter
	Status    OracleSpecStatus
	Data      []OracleData
}

type Position struct {
	Market            Market
	Party             Party
	OpenVolume        string
	RealisedPNL       string
	UnrealisedPNL     string
	AverageEntryPrice string
	UpdatedAt         TimeString
}

type NodeSignature struct {
	ID        HexString
	Signature HexString
	Kind      string
}

type ERC20WithdrawalApproval struct {
	AssetSource   string
	Amount        string
	Nonce         string
	Signatures    string
	TargetAddress string
}
