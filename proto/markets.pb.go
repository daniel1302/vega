// Code generated by protoc-gen-go. DO NOT EDIT.
// source: proto/markets.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
	_ "github.com/mwitkow/go-proto-validators"
	math "math"
)

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion3 // please upgrade the proto package

type ContinuousTrading struct {
	// Duration in nanoseconds, maximum 1 month
	DurationNs           int64    `protobuf:"varint,1,opt,name=durationNs,proto3" json:"durationNs,omitempty"`
	TickSize             uint64   `protobuf:"varint,2,opt,name=tickSize,proto3" json:"tickSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ContinuousTrading) Reset()         { *m = ContinuousTrading{} }
func (m *ContinuousTrading) String() string { return proto.CompactTextString(m) }
func (*ContinuousTrading) ProtoMessage()    {}
func (*ContinuousTrading) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{0}
}

func (m *ContinuousTrading) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ContinuousTrading.Unmarshal(m, b)
}
func (m *ContinuousTrading) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ContinuousTrading.Marshal(b, m, deterministic)
}
func (m *ContinuousTrading) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ContinuousTrading.Merge(m, src)
}
func (m *ContinuousTrading) XXX_Size() int {
	return xxx_messageInfo_ContinuousTrading.Size(m)
}
func (m *ContinuousTrading) XXX_DiscardUnknown() {
	xxx_messageInfo_ContinuousTrading.DiscardUnknown(m)
}

var xxx_messageInfo_ContinuousTrading proto.InternalMessageInfo

func (m *ContinuousTrading) GetDurationNs() int64 {
	if m != nil {
		return m.DurationNs
	}
	return 0
}

func (m *ContinuousTrading) GetTickSize() uint64 {
	if m != nil {
		return m.TickSize
	}
	return 0
}

type DiscreteTrading struct {
	// Duration in nanoseconds, maximum 1 month
	DurationNs           int64    `protobuf:"varint,1,opt,name=durationNs,proto3" json:"durationNs,omitempty"`
	TickSize             uint64   `protobuf:"varint,2,opt,name=tickSize,proto3" json:"tickSize,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *DiscreteTrading) Reset()         { *m = DiscreteTrading{} }
func (m *DiscreteTrading) String() string { return proto.CompactTextString(m) }
func (*DiscreteTrading) ProtoMessage()    {}
func (*DiscreteTrading) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{1}
}

func (m *DiscreteTrading) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DiscreteTrading.Unmarshal(m, b)
}
func (m *DiscreteTrading) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DiscreteTrading.Marshal(b, m, deterministic)
}
func (m *DiscreteTrading) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DiscreteTrading.Merge(m, src)
}
func (m *DiscreteTrading) XXX_Size() int {
	return xxx_messageInfo_DiscreteTrading.Size(m)
}
func (m *DiscreteTrading) XXX_DiscardUnknown() {
	xxx_messageInfo_DiscreteTrading.DiscardUnknown(m)
}

var xxx_messageInfo_DiscreteTrading proto.InternalMessageInfo

func (m *DiscreteTrading) GetDurationNs() int64 {
	if m != nil {
		return m.DurationNs
	}
	return 0
}

func (m *DiscreteTrading) GetTickSize() uint64 {
	if m != nil {
		return m.TickSize
	}
	return 0
}

type Future struct {
	Maturity string `protobuf:"bytes,1,opt,name=maturity,proto3" json:"maturity,omitempty"`
	Asset    string `protobuf:"bytes,2,opt,name=asset,proto3" json:"asset,omitempty"`
	// Types that are valid to be assigned to Oracle:
	//	*Future_EthereumEvent
	Oracle               isFuture_Oracle `protobuf_oneof:"oracle"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *Future) Reset()         { *m = Future{} }
func (m *Future) String() string { return proto.CompactTextString(m) }
func (*Future) ProtoMessage()    {}
func (*Future) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{2}
}

func (m *Future) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Future.Unmarshal(m, b)
}
func (m *Future) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Future.Marshal(b, m, deterministic)
}
func (m *Future) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Future.Merge(m, src)
}
func (m *Future) XXX_Size() int {
	return xxx_messageInfo_Future.Size(m)
}
func (m *Future) XXX_DiscardUnknown() {
	xxx_messageInfo_Future.DiscardUnknown(m)
}

var xxx_messageInfo_Future proto.InternalMessageInfo

func (m *Future) GetMaturity() string {
	if m != nil {
		return m.Maturity
	}
	return ""
}

func (m *Future) GetAsset() string {
	if m != nil {
		return m.Asset
	}
	return ""
}

type isFuture_Oracle interface {
	isFuture_Oracle()
}

type Future_EthereumEvent struct {
	EthereumEvent *EthereumEvent `protobuf:"bytes,100,opt,name=ethereumEvent,proto3,oneof"`
}

func (*Future_EthereumEvent) isFuture_Oracle() {}

func (m *Future) GetOracle() isFuture_Oracle {
	if m != nil {
		return m.Oracle
	}
	return nil
}

func (m *Future) GetEthereumEvent() *EthereumEvent {
	if x, ok := m.GetOracle().(*Future_EthereumEvent); ok {
		return x.EthereumEvent
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Future) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Future_EthereumEvent)(nil),
	}
}

type EthereumEvent struct {
	ContractID           string   `protobuf:"bytes,1,opt,name=contractID,proto3" json:"contractID,omitempty"`
	Event                string   `protobuf:"bytes,2,opt,name=event,proto3" json:"event,omitempty"`
	Value                uint64   `protobuf:"varint,3,opt,name=value,proto3" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EthereumEvent) Reset()         { *m = EthereumEvent{} }
func (m *EthereumEvent) String() string { return proto.CompactTextString(m) }
func (*EthereumEvent) ProtoMessage()    {}
func (*EthereumEvent) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{3}
}

func (m *EthereumEvent) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EthereumEvent.Unmarshal(m, b)
}
func (m *EthereumEvent) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EthereumEvent.Marshal(b, m, deterministic)
}
func (m *EthereumEvent) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EthereumEvent.Merge(m, src)
}
func (m *EthereumEvent) XXX_Size() int {
	return xxx_messageInfo_EthereumEvent.Size(m)
}
func (m *EthereumEvent) XXX_DiscardUnknown() {
	xxx_messageInfo_EthereumEvent.DiscardUnknown(m)
}

var xxx_messageInfo_EthereumEvent proto.InternalMessageInfo

func (m *EthereumEvent) GetContractID() string {
	if m != nil {
		return m.ContractID
	}
	return ""
}

func (m *EthereumEvent) GetEvent() string {
	if m != nil {
		return m.Event
	}
	return ""
}

func (m *EthereumEvent) GetValue() uint64 {
	if m != nil {
		return m.Value
	}
	return 0
}

type InstrumentMetadata struct {
	Tags                 []string `protobuf:"bytes,1,rep,name=tags,proto3" json:"tags,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *InstrumentMetadata) Reset()         { *m = InstrumentMetadata{} }
func (m *InstrumentMetadata) String() string { return proto.CompactTextString(m) }
func (*InstrumentMetadata) ProtoMessage()    {}
func (*InstrumentMetadata) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{4}
}

func (m *InstrumentMetadata) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_InstrumentMetadata.Unmarshal(m, b)
}
func (m *InstrumentMetadata) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_InstrumentMetadata.Marshal(b, m, deterministic)
}
func (m *InstrumentMetadata) XXX_Merge(src proto.Message) {
	xxx_messageInfo_InstrumentMetadata.Merge(m, src)
}
func (m *InstrumentMetadata) XXX_Size() int {
	return xxx_messageInfo_InstrumentMetadata.Size(m)
}
func (m *InstrumentMetadata) XXX_DiscardUnknown() {
	xxx_messageInfo_InstrumentMetadata.DiscardUnknown(m)
}

var xxx_messageInfo_InstrumentMetadata proto.InternalMessageInfo

func (m *InstrumentMetadata) GetTags() []string {
	if m != nil {
		return m.Tags
	}
	return nil
}

type Instrument struct {
	Id               string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Code             string              `protobuf:"bytes,2,opt,name=code,proto3" json:"code,omitempty"`
	Name             string              `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	BaseName         string              `protobuf:"bytes,4,opt,name=baseName,proto3" json:"baseName,omitempty"`
	QuoteName        string              `protobuf:"bytes,5,opt,name=quoteName,proto3" json:"quoteName,omitempty"`
	Metadata         *InstrumentMetadata `protobuf:"bytes,6,opt,name=metadata,proto3" json:"metadata,omitempty"`
	InitialMarkPrice uint64              `protobuf:"varint,7,opt,name=initialMarkPrice,proto3" json:"initialMarkPrice,omitempty"`
	// Types that are valid to be assigned to Product:
	//	*Instrument_Future
	Product              isInstrument_Product `protobuf_oneof:"product"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Instrument) Reset()         { *m = Instrument{} }
func (m *Instrument) String() string { return proto.CompactTextString(m) }
func (*Instrument) ProtoMessage()    {}
func (*Instrument) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{5}
}

func (m *Instrument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Instrument.Unmarshal(m, b)
}
func (m *Instrument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Instrument.Marshal(b, m, deterministic)
}
func (m *Instrument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Instrument.Merge(m, src)
}
func (m *Instrument) XXX_Size() int {
	return xxx_messageInfo_Instrument.Size(m)
}
func (m *Instrument) XXX_DiscardUnknown() {
	xxx_messageInfo_Instrument.DiscardUnknown(m)
}

var xxx_messageInfo_Instrument proto.InternalMessageInfo

func (m *Instrument) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Instrument) GetCode() string {
	if m != nil {
		return m.Code
	}
	return ""
}

func (m *Instrument) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Instrument) GetBaseName() string {
	if m != nil {
		return m.BaseName
	}
	return ""
}

func (m *Instrument) GetQuoteName() string {
	if m != nil {
		return m.QuoteName
	}
	return ""
}

func (m *Instrument) GetMetadata() *InstrumentMetadata {
	if m != nil {
		return m.Metadata
	}
	return nil
}

func (m *Instrument) GetInitialMarkPrice() uint64 {
	if m != nil {
		return m.InitialMarkPrice
	}
	return 0
}

type isInstrument_Product interface {
	isInstrument_Product()
}

type Instrument_Future struct {
	Future *Future `protobuf:"bytes,100,opt,name=future,proto3,oneof"`
}

func (*Instrument_Future) isInstrument_Product() {}

func (m *Instrument) GetProduct() isInstrument_Product {
	if m != nil {
		return m.Product
	}
	return nil
}

func (m *Instrument) GetFuture() *Future {
	if x, ok := m.GetProduct().(*Instrument_Future); ok {
		return x.Future
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Instrument) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Instrument_Future)(nil),
	}
}

type LogNormalRiskModel struct {
	RiskAversionParameter float64               `protobuf:"fixed64,1,opt,name=riskAversionParameter,proto3" json:"riskAversionParameter,omitempty"`
	Tau                   float64               `protobuf:"fixed64,2,opt,name=tau,proto3" json:"tau,omitempty"`
	Params                *LogNormalModelParams `protobuf:"bytes,3,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral  struct{}              `json:"-"`
	XXX_unrecognized      []byte                `json:"-"`
	XXX_sizecache         int32                 `json:"-"`
}

func (m *LogNormalRiskModel) Reset()         { *m = LogNormalRiskModel{} }
func (m *LogNormalRiskModel) String() string { return proto.CompactTextString(m) }
func (*LogNormalRiskModel) ProtoMessage()    {}
func (*LogNormalRiskModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{6}
}

func (m *LogNormalRiskModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogNormalRiskModel.Unmarshal(m, b)
}
func (m *LogNormalRiskModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogNormalRiskModel.Marshal(b, m, deterministic)
}
func (m *LogNormalRiskModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogNormalRiskModel.Merge(m, src)
}
func (m *LogNormalRiskModel) XXX_Size() int {
	return xxx_messageInfo_LogNormalRiskModel.Size(m)
}
func (m *LogNormalRiskModel) XXX_DiscardUnknown() {
	xxx_messageInfo_LogNormalRiskModel.DiscardUnknown(m)
}

var xxx_messageInfo_LogNormalRiskModel proto.InternalMessageInfo

func (m *LogNormalRiskModel) GetRiskAversionParameter() float64 {
	if m != nil {
		return m.RiskAversionParameter
	}
	return 0
}

func (m *LogNormalRiskModel) GetTau() float64 {
	if m != nil {
		return m.Tau
	}
	return 0
}

func (m *LogNormalRiskModel) GetParams() *LogNormalModelParams {
	if m != nil {
		return m.Params
	}
	return nil
}

type LogNormalModelParams struct {
	Mu                   float64  `protobuf:"fixed64,1,opt,name=mu,proto3" json:"mu,omitempty"`
	R                    float64  `protobuf:"fixed64,2,opt,name=r,proto3" json:"r,omitempty"`
	Sigma                float64  `protobuf:"fixed64,3,opt,name=sigma,proto3" json:"sigma,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *LogNormalModelParams) Reset()         { *m = LogNormalModelParams{} }
func (m *LogNormalModelParams) String() string { return proto.CompactTextString(m) }
func (*LogNormalModelParams) ProtoMessage()    {}
func (*LogNormalModelParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{7}
}

func (m *LogNormalModelParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_LogNormalModelParams.Unmarshal(m, b)
}
func (m *LogNormalModelParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_LogNormalModelParams.Marshal(b, m, deterministic)
}
func (m *LogNormalModelParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_LogNormalModelParams.Merge(m, src)
}
func (m *LogNormalModelParams) XXX_Size() int {
	return xxx_messageInfo_LogNormalModelParams.Size(m)
}
func (m *LogNormalModelParams) XXX_DiscardUnknown() {
	xxx_messageInfo_LogNormalModelParams.DiscardUnknown(m)
}

var xxx_messageInfo_LogNormalModelParams proto.InternalMessageInfo

func (m *LogNormalModelParams) GetMu() float64 {
	if m != nil {
		return m.Mu
	}
	return 0
}

func (m *LogNormalModelParams) GetR() float64 {
	if m != nil {
		return m.R
	}
	return 0
}

func (m *LogNormalModelParams) GetSigma() float64 {
	if m != nil {
		return m.Sigma
	}
	return 0
}

type SimpleRiskModel struct {
	Params               *SimpleModelParams `protobuf:"bytes,1,opt,name=params,proto3" json:"params,omitempty"`
	XXX_NoUnkeyedLiteral struct{}           `json:"-"`
	XXX_unrecognized     []byte             `json:"-"`
	XXX_sizecache        int32              `json:"-"`
}

func (m *SimpleRiskModel) Reset()         { *m = SimpleRiskModel{} }
func (m *SimpleRiskModel) String() string { return proto.CompactTextString(m) }
func (*SimpleRiskModel) ProtoMessage()    {}
func (*SimpleRiskModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{8}
}

func (m *SimpleRiskModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleRiskModel.Unmarshal(m, b)
}
func (m *SimpleRiskModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleRiskModel.Marshal(b, m, deterministic)
}
func (m *SimpleRiskModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleRiskModel.Merge(m, src)
}
func (m *SimpleRiskModel) XXX_Size() int {
	return xxx_messageInfo_SimpleRiskModel.Size(m)
}
func (m *SimpleRiskModel) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleRiskModel.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleRiskModel proto.InternalMessageInfo

func (m *SimpleRiskModel) GetParams() *SimpleModelParams {
	if m != nil {
		return m.Params
	}
	return nil
}

type SimpleModelParams struct {
	FactorLong           float64  `protobuf:"fixed64,1,opt,name=factorLong,proto3" json:"factorLong,omitempty"`
	FactorShort          float64  `protobuf:"fixed64,2,opt,name=factorShort,proto3" json:"factorShort,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *SimpleModelParams) Reset()         { *m = SimpleModelParams{} }
func (m *SimpleModelParams) String() string { return proto.CompactTextString(m) }
func (*SimpleModelParams) ProtoMessage()    {}
func (*SimpleModelParams) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{9}
}

func (m *SimpleModelParams) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SimpleModelParams.Unmarshal(m, b)
}
func (m *SimpleModelParams) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SimpleModelParams.Marshal(b, m, deterministic)
}
func (m *SimpleModelParams) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SimpleModelParams.Merge(m, src)
}
func (m *SimpleModelParams) XXX_Size() int {
	return xxx_messageInfo_SimpleModelParams.Size(m)
}
func (m *SimpleModelParams) XXX_DiscardUnknown() {
	xxx_messageInfo_SimpleModelParams.DiscardUnknown(m)
}

var xxx_messageInfo_SimpleModelParams proto.InternalMessageInfo

func (m *SimpleModelParams) GetFactorLong() float64 {
	if m != nil {
		return m.FactorLong
	}
	return 0
}

func (m *SimpleModelParams) GetFactorShort() float64 {
	if m != nil {
		return m.FactorShort
	}
	return 0
}

type ExternalRiskModel struct {
	Name                 string            `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Socket               string            `protobuf:"bytes,2,opt,name=socket,proto3" json:"socket,omitempty"`
	Config               map[string]string `protobuf:"bytes,3,rep,name=config,proto3" json:"config,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	XXX_NoUnkeyedLiteral struct{}          `json:"-"`
	XXX_unrecognized     []byte            `json:"-"`
	XXX_sizecache        int32             `json:"-"`
}

func (m *ExternalRiskModel) Reset()         { *m = ExternalRiskModel{} }
func (m *ExternalRiskModel) String() string { return proto.CompactTextString(m) }
func (*ExternalRiskModel) ProtoMessage()    {}
func (*ExternalRiskModel) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{10}
}

func (m *ExternalRiskModel) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExternalRiskModel.Unmarshal(m, b)
}
func (m *ExternalRiskModel) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExternalRiskModel.Marshal(b, m, deterministic)
}
func (m *ExternalRiskModel) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExternalRiskModel.Merge(m, src)
}
func (m *ExternalRiskModel) XXX_Size() int {
	return xxx_messageInfo_ExternalRiskModel.Size(m)
}
func (m *ExternalRiskModel) XXX_DiscardUnknown() {
	xxx_messageInfo_ExternalRiskModel.DiscardUnknown(m)
}

var xxx_messageInfo_ExternalRiskModel proto.InternalMessageInfo

func (m *ExternalRiskModel) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *ExternalRiskModel) GetSocket() string {
	if m != nil {
		return m.Socket
	}
	return ""
}

func (m *ExternalRiskModel) GetConfig() map[string]string {
	if m != nil {
		return m.Config
	}
	return nil
}

type ScalingFactors struct {
	SearchLevel          float64  `protobuf:"fixed64,1,opt,name=searchLevel,proto3" json:"searchLevel,omitempty"`
	InitialMargin        float64  `protobuf:"fixed64,2,opt,name=initialMargin,proto3" json:"initialMargin,omitempty"`
	CollateralRelease    float64  `protobuf:"fixed64,3,opt,name=collateralRelease,proto3" json:"collateralRelease,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ScalingFactors) Reset()         { *m = ScalingFactors{} }
func (m *ScalingFactors) String() string { return proto.CompactTextString(m) }
func (*ScalingFactors) ProtoMessage()    {}
func (*ScalingFactors) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{11}
}

func (m *ScalingFactors) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ScalingFactors.Unmarshal(m, b)
}
func (m *ScalingFactors) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ScalingFactors.Marshal(b, m, deterministic)
}
func (m *ScalingFactors) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ScalingFactors.Merge(m, src)
}
func (m *ScalingFactors) XXX_Size() int {
	return xxx_messageInfo_ScalingFactors.Size(m)
}
func (m *ScalingFactors) XXX_DiscardUnknown() {
	xxx_messageInfo_ScalingFactors.DiscardUnknown(m)
}

var xxx_messageInfo_ScalingFactors proto.InternalMessageInfo

func (m *ScalingFactors) GetSearchLevel() float64 {
	if m != nil {
		return m.SearchLevel
	}
	return 0
}

func (m *ScalingFactors) GetInitialMargin() float64 {
	if m != nil {
		return m.InitialMargin
	}
	return 0
}

func (m *ScalingFactors) GetCollateralRelease() float64 {
	if m != nil {
		return m.CollateralRelease
	}
	return 0
}

type MarginCalculator struct {
	ScalingFactors       *ScalingFactors `protobuf:"bytes,1,opt,name=scalingFactors,proto3" json:"scalingFactors,omitempty"`
	XXX_NoUnkeyedLiteral struct{}        `json:"-"`
	XXX_unrecognized     []byte          `json:"-"`
	XXX_sizecache        int32           `json:"-"`
}

func (m *MarginCalculator) Reset()         { *m = MarginCalculator{} }
func (m *MarginCalculator) String() string { return proto.CompactTextString(m) }
func (*MarginCalculator) ProtoMessage()    {}
func (*MarginCalculator) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{12}
}

func (m *MarginCalculator) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_MarginCalculator.Unmarshal(m, b)
}
func (m *MarginCalculator) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_MarginCalculator.Marshal(b, m, deterministic)
}
func (m *MarginCalculator) XXX_Merge(src proto.Message) {
	xxx_messageInfo_MarginCalculator.Merge(m, src)
}
func (m *MarginCalculator) XXX_Size() int {
	return xxx_messageInfo_MarginCalculator.Size(m)
}
func (m *MarginCalculator) XXX_DiscardUnknown() {
	xxx_messageInfo_MarginCalculator.DiscardUnknown(m)
}

var xxx_messageInfo_MarginCalculator proto.InternalMessageInfo

func (m *MarginCalculator) GetScalingFactors() *ScalingFactors {
	if m != nil {
		return m.ScalingFactors
	}
	return nil
}

type TradableInstrument struct {
	Instrument       *Instrument       `protobuf:"bytes,1,opt,name=instrument,proto3" json:"instrument,omitempty"`
	MarginCalculator *MarginCalculator `protobuf:"bytes,2,opt,name=marginCalculator,proto3" json:"marginCalculator,omitempty"`
	// Types that are valid to be assigned to RiskModel:
	//	*TradableInstrument_LogNormalRiskModel
	//	*TradableInstrument_ExternalRiskModel
	//	*TradableInstrument_SimpleRiskModel
	RiskModel            isTradableInstrument_RiskModel `protobuf_oneof:"riskModel"`
	XXX_NoUnkeyedLiteral struct{}                       `json:"-"`
	XXX_unrecognized     []byte                         `json:"-"`
	XXX_sizecache        int32                          `json:"-"`
}

func (m *TradableInstrument) Reset()         { *m = TradableInstrument{} }
func (m *TradableInstrument) String() string { return proto.CompactTextString(m) }
func (*TradableInstrument) ProtoMessage()    {}
func (*TradableInstrument) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{13}
}

func (m *TradableInstrument) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_TradableInstrument.Unmarshal(m, b)
}
func (m *TradableInstrument) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_TradableInstrument.Marshal(b, m, deterministic)
}
func (m *TradableInstrument) XXX_Merge(src proto.Message) {
	xxx_messageInfo_TradableInstrument.Merge(m, src)
}
func (m *TradableInstrument) XXX_Size() int {
	return xxx_messageInfo_TradableInstrument.Size(m)
}
func (m *TradableInstrument) XXX_DiscardUnknown() {
	xxx_messageInfo_TradableInstrument.DiscardUnknown(m)
}

var xxx_messageInfo_TradableInstrument proto.InternalMessageInfo

func (m *TradableInstrument) GetInstrument() *Instrument {
	if m != nil {
		return m.Instrument
	}
	return nil
}

func (m *TradableInstrument) GetMarginCalculator() *MarginCalculator {
	if m != nil {
		return m.MarginCalculator
	}
	return nil
}

type isTradableInstrument_RiskModel interface {
	isTradableInstrument_RiskModel()
}

type TradableInstrument_LogNormalRiskModel struct {
	LogNormalRiskModel *LogNormalRiskModel `protobuf:"bytes,100,opt,name=logNormalRiskModel,proto3,oneof"`
}

type TradableInstrument_ExternalRiskModel struct {
	ExternalRiskModel *ExternalRiskModel `protobuf:"bytes,101,opt,name=externalRiskModel,proto3,oneof"`
}

type TradableInstrument_SimpleRiskModel struct {
	SimpleRiskModel *SimpleRiskModel `protobuf:"bytes,102,opt,name=simpleRiskModel,proto3,oneof"`
}

func (*TradableInstrument_LogNormalRiskModel) isTradableInstrument_RiskModel() {}

func (*TradableInstrument_ExternalRiskModel) isTradableInstrument_RiskModel() {}

func (*TradableInstrument_SimpleRiskModel) isTradableInstrument_RiskModel() {}

func (m *TradableInstrument) GetRiskModel() isTradableInstrument_RiskModel {
	if m != nil {
		return m.RiskModel
	}
	return nil
}

func (m *TradableInstrument) GetLogNormalRiskModel() *LogNormalRiskModel {
	if x, ok := m.GetRiskModel().(*TradableInstrument_LogNormalRiskModel); ok {
		return x.LogNormalRiskModel
	}
	return nil
}

func (m *TradableInstrument) GetExternalRiskModel() *ExternalRiskModel {
	if x, ok := m.GetRiskModel().(*TradableInstrument_ExternalRiskModel); ok {
		return x.ExternalRiskModel
	}
	return nil
}

func (m *TradableInstrument) GetSimpleRiskModel() *SimpleRiskModel {
	if x, ok := m.GetRiskModel().(*TradableInstrument_SimpleRiskModel); ok {
		return x.SimpleRiskModel
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*TradableInstrument) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*TradableInstrument_LogNormalRiskModel)(nil),
		(*TradableInstrument_ExternalRiskModel)(nil),
		(*TradableInstrument_SimpleRiskModel)(nil),
	}
}

type Market struct {
	Id                 string              `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	TradableInstrument *TradableInstrument `protobuf:"bytes,2,opt,name=tradableInstrument,proto3" json:"tradableInstrument,omitempty"`
	// the number of decimal places that a price must be shifted by in order to get a correct price denominated in the currency of the Market. ie `realPrice = price / 10^decimalPlaces`
	DecimalPlaces uint64 `protobuf:"varint,3,opt,name=decimalPlaces,proto3" json:"decimalPlaces,omitempty"`
	// Types that are valid to be assigned to TradingMode:
	//	*Market_Continuous
	//	*Market_Discrete
	TradingMode          isMarket_TradingMode `protobuf_oneof:"tradingMode"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *Market) Reset()         { *m = Market{} }
func (m *Market) String() string { return proto.CompactTextString(m) }
func (*Market) ProtoMessage()    {}
func (*Market) Descriptor() ([]byte, []int) {
	return fileDescriptor_ef38c4b9a7594dbd, []int{14}
}

func (m *Market) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Market.Unmarshal(m, b)
}
func (m *Market) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Market.Marshal(b, m, deterministic)
}
func (m *Market) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Market.Merge(m, src)
}
func (m *Market) XXX_Size() int {
	return xxx_messageInfo_Market.Size(m)
}
func (m *Market) XXX_DiscardUnknown() {
	xxx_messageInfo_Market.DiscardUnknown(m)
}

var xxx_messageInfo_Market proto.InternalMessageInfo

func (m *Market) GetId() string {
	if m != nil {
		return m.Id
	}
	return ""
}

func (m *Market) GetTradableInstrument() *TradableInstrument {
	if m != nil {
		return m.TradableInstrument
	}
	return nil
}

func (m *Market) GetDecimalPlaces() uint64 {
	if m != nil {
		return m.DecimalPlaces
	}
	return 0
}

type isMarket_TradingMode interface {
	isMarket_TradingMode()
}

type Market_Continuous struct {
	Continuous *ContinuousTrading `protobuf:"bytes,100,opt,name=continuous,proto3,oneof"`
}

type Market_Discrete struct {
	Discrete *DiscreteTrading `protobuf:"bytes,101,opt,name=discrete,proto3,oneof"`
}

func (*Market_Continuous) isMarket_TradingMode() {}

func (*Market_Discrete) isMarket_TradingMode() {}

func (m *Market) GetTradingMode() isMarket_TradingMode {
	if m != nil {
		return m.TradingMode
	}
	return nil
}

func (m *Market) GetContinuous() *ContinuousTrading {
	if x, ok := m.GetTradingMode().(*Market_Continuous); ok {
		return x.Continuous
	}
	return nil
}

func (m *Market) GetDiscrete() *DiscreteTrading {
	if x, ok := m.GetTradingMode().(*Market_Discrete); ok {
		return x.Discrete
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*Market) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*Market_Continuous)(nil),
		(*Market_Discrete)(nil),
	}
}

func init() {
	proto.RegisterType((*ContinuousTrading)(nil), "vega.ContinuousTrading")
	proto.RegisterType((*DiscreteTrading)(nil), "vega.DiscreteTrading")
	proto.RegisterType((*Future)(nil), "vega.Future")
	proto.RegisterType((*EthereumEvent)(nil), "vega.EthereumEvent")
	proto.RegisterType((*InstrumentMetadata)(nil), "vega.InstrumentMetadata")
	proto.RegisterType((*Instrument)(nil), "vega.Instrument")
	proto.RegisterType((*LogNormalRiskModel)(nil), "vega.LogNormalRiskModel")
	proto.RegisterType((*LogNormalModelParams)(nil), "vega.LogNormalModelParams")
	proto.RegisterType((*SimpleRiskModel)(nil), "vega.SimpleRiskModel")
	proto.RegisterType((*SimpleModelParams)(nil), "vega.SimpleModelParams")
	proto.RegisterType((*ExternalRiskModel)(nil), "vega.ExternalRiskModel")
	proto.RegisterMapType((map[string]string)(nil), "vega.ExternalRiskModel.ConfigEntry")
	proto.RegisterType((*ScalingFactors)(nil), "vega.ScalingFactors")
	proto.RegisterType((*MarginCalculator)(nil), "vega.MarginCalculator")
	proto.RegisterType((*TradableInstrument)(nil), "vega.TradableInstrument")
	proto.RegisterType((*Market)(nil), "vega.Market")
}

func init() { proto.RegisterFile("proto/markets.proto", fileDescriptor_ef38c4b9a7594dbd) }

var fileDescriptor_ef38c4b9a7594dbd = []byte{
	// 1007 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0xb4, 0x56, 0xdd, 0x6e, 0x1b, 0x45,
	0x14, 0xf6, 0xda, 0xa9, 0x1b, 0x1f, 0x37, 0x89, 0x33, 0x4d, 0xcb, 0x2a, 0x42, 0xd4, 0x2c, 0x08,
	0x59, 0x88, 0xda, 0xc8, 0xad, 0x10, 0xa5, 0xdc, 0xd4, 0x69, 0x8a, 0x5b, 0x25, 0x51, 0x34, 0x81,
	0x1b, 0x90, 0x10, 0x93, 0xdd, 0xc9, 0x66, 0xe4, 0xd9, 0x9d, 0x30, 0x33, 0xeb, 0x52, 0xae, 0xca,
	0x0d, 0x70, 0x59, 0x89, 0xe7, 0xe0, 0x1a, 0x89, 0x77, 0xe0, 0x19, 0x90, 0x78, 0x01, 0x5e, 0x01,
	0xcd, 0x4f, 0xd6, 0x6b, 0xaf, 0x6f, 0xb9, 0xf2, 0x9c, 0x73, 0xbe, 0xf3, 0x33, 0xdf, 0x39, 0x67,
	0xc7, 0x70, 0xfb, 0x4a, 0x0a, 0x2d, 0x46, 0x19, 0x91, 0x33, 0xaa, 0xd5, 0xd0, 0x4a, 0x68, 0x63,
	0x4e, 0x53, 0xb2, 0xff, 0x49, 0xca, 0xf4, 0x65, 0x71, 0x3e, 0x8c, 0x45, 0x36, 0xca, 0x5e, 0x32,
	0x3d, 0x13, 0x2f, 0x47, 0xa9, 0xb8, 0x6f, 0x21, 0xf7, 0xe7, 0x84, 0xb3, 0x84, 0x68, 0x21, 0xd5,
	0xa8, 0x3c, 0x3a, 0xef, 0xe8, 0x3b, 0xd8, 0x3d, 0x10, 0xb9, 0x66, 0x79, 0x21, 0x0a, 0xf5, 0xa5,
	0x24, 0x09, 0xcb, 0x53, 0x34, 0x02, 0x48, 0x0a, 0x49, 0x34, 0x13, 0xf9, 0x89, 0x0a, 0x83, 0x7e,
	0x30, 0x68, 0x4d, 0x76, 0xfe, 0xf9, 0xfb, 0x5e, 0xb7, 0xd7, 0x08, 0x5f, 0xbf, 0xfe, 0xf5, 0xcf,
	0x37, 0xbf, 0xff, 0xb5, 0x81, 0x2b, 0x10, 0xb4, 0x0f, 0x9b, 0x9a, 0xc5, 0xb3, 0x33, 0xf6, 0x23,
	0x0d, 0x9b, 0xfd, 0x60, 0xb0, 0x81, 0x4b, 0x39, 0xfa, 0x16, 0x76, 0x9e, 0x32, 0x15, 0x4b, 0xaa,
	0xe9, 0xff, 0x12, 0xff, 0xa7, 0x00, 0xda, 0xcf, 0x0a, 0x5d, 0x48, 0x6a, 0x60, 0x19, 0xd1, 0x85,
	0x64, 0xfa, 0x95, 0x8d, 0xda, 0xc1, 0xa5, 0x8c, 0xf6, 0xe0, 0x06, 0x51, 0x8a, 0x6a, 0xeb, 0xdf,
	0xc1, 0x4e, 0x40, 0x8f, 0x61, 0x8b, 0xea, 0x4b, 0x2a, 0x69, 0x91, 0x1d, 0xce, 0x69, 0xae, 0xc3,
	0xa4, 0x1f, 0x0c, 0xba, 0xe3, 0xdb, 0x43, 0x43, 0xea, 0xf0, 0xb0, 0x6a, 0x9a, 0x36, 0xf0, 0x32,
	0x76, 0xb2, 0x09, 0x6d, 0x21, 0x49, 0xcc, 0x69, 0xf4, 0x0d, 0x6c, 0x2d, 0x61, 0xd1, 0x3b, 0x00,
	0xb1, 0xc8, 0xb5, 0x24, 0xb1, 0x7e, 0xfe, 0xd4, 0xd7, 0x52, 0xd1, 0x98, 0x6a, 0xa8, 0xcd, 0xe7,
	0xab, 0xb1, 0x82, 0xd1, 0xce, 0x09, 0x2f, 0x68, 0xd8, 0xb2, 0x77, 0x74, 0x42, 0x34, 0x00, 0xf4,
	0x3c, 0x57, 0x5a, 0x16, 0x19, 0xcd, 0xf5, 0x31, 0xd5, 0x24, 0x21, 0x9a, 0x20, 0x04, 0x1b, 0x9a,
	0xa4, 0x86, 0xbd, 0xd6, 0xa0, 0x83, 0xed, 0x39, 0xfa, 0xad, 0x09, 0xb0, 0x80, 0xa2, 0x6d, 0x68,
	0xb2, 0xc4, 0x27, 0x6f, 0xb2, 0xc4, 0xb8, 0xc4, 0x22, 0xa1, 0x3e, 0xa7, 0x3d, 0x1b, 0x5d, 0x4e,
	0x32, 0x97, 0xb1, 0x83, 0xed, 0xd9, 0xd0, 0x78, 0x4e, 0x14, 0x3d, 0x31, 0xfa, 0x0d, 0x47, 0xe3,
	0xb5, 0x8c, 0xde, 0x86, 0xce, 0xf7, 0x85, 0xd0, 0xce, 0x78, 0xc3, 0x1a, 0x17, 0x0a, 0xf4, 0x10,
	0x36, 0x33, 0x5f, 0x60, 0xd8, 0xb6, 0x4c, 0x86, 0x8e, 0xc9, 0xfa, 0x05, 0x70, 0x89, 0x44, 0x1f,
	0x42, 0x8f, 0xe5, 0x4c, 0x33, 0xc2, 0x8f, 0x89, 0x9c, 0x9d, 0x4a, 0x16, 0xd3, 0xf0, 0xa6, 0x65,
	0xa0, 0xa6, 0x47, 0x1f, 0x40, 0xfb, 0xc2, 0x36, 0xdb, 0x77, 0xea, 0x96, 0x8b, 0xef, 0x06, 0x60,
	0xda, 0xc0, 0xde, 0x3a, 0xe9, 0xc0, 0xcd, 0x2b, 0x29, 0x92, 0x22, 0xd6, 0xd1, 0x9b, 0x00, 0xd0,
	0x91, 0x48, 0x4f, 0x84, 0xcc, 0x08, 0xc7, 0x4c, 0xcd, 0x8e, 0x45, 0x42, 0x39, 0x7a, 0x08, 0x77,
	0x24, 0x53, 0xb3, 0x27, 0x73, 0x2a, 0x15, 0x13, 0xf9, 0x29, 0x91, 0x24, 0xa3, 0x9a, 0x4a, 0x4b,
	0x58, 0x80, 0xd7, 0x1b, 0x51, 0x0f, 0x5a, 0x9a, 0x14, 0x96, 0xc2, 0x00, 0x9b, 0x23, 0x1a, 0x43,
	0xfb, 0xca, 0x98, 0x95, 0xe5, 0xb0, 0x3b, 0xde, 0x77, 0x15, 0x95, 0x19, 0x6d, 0x36, 0x1b, 0x40,
	0x61, 0x8f, 0x8c, 0x5e, 0xc0, 0xde, 0x3a, 0xbb, 0xe9, 0x58, 0x56, 0xf8, 0x02, 0x9a, 0x59, 0x81,
	0x6e, 0x41, 0x20, 0x7d, 0xae, 0x40, 0x9a, 0xf1, 0x50, 0x2c, 0xcd, 0x88, 0x4d, 0x14, 0x60, 0x27,
	0x44, 0x13, 0xd8, 0x39, 0x63, 0xd9, 0x15, 0xa7, 0x8b, 0xab, 0x8d, 0xca, 0x92, 0x02, 0x5b, 0xd2,
	0x5b, 0xae, 0x24, 0x07, 0x5b, 0x57, 0xcf, 0x57, 0xb0, 0x5b, 0x33, 0x9a, 0x19, 0xbe, 0x20, 0xb1,
	0x16, 0xf2, 0x48, 0xe4, 0xa9, 0x2f, 0xaa, 0xa2, 0x41, 0x7d, 0xe8, 0x3a, 0xe9, 0xec, 0x52, 0x48,
	0xed, 0xcb, 0xac, 0xaa, 0xa2, 0x3f, 0x02, 0xd8, 0x3d, 0xfc, 0x41, 0x53, 0x99, 0x57, 0x89, 0xbf,
	0x1e, 0xb9, 0xa0, 0x32, 0x72, 0x77, 0xa1, 0xad, 0x44, 0x3c, 0x2b, 0xd7, 0xd3, 0x4b, 0xe8, 0x31,
	0xb4, 0x63, 0x91, 0x5f, 0xb0, 0x34, 0x6c, 0xf5, 0x5b, 0x83, 0xee, 0xf8, 0x3d, 0xbf, 0x98, 0xab,
	0x41, 0x87, 0x07, 0x16, 0x75, 0x98, 0x6b, 0xf9, 0x0a, 0x7b, 0x97, 0xfd, 0x47, 0xd0, 0xad, 0xa8,
	0x4d, 0xeb, 0x66, 0xf4, 0xfa, 0xc3, 0x60, 0x8e, 0x8b, 0x7d, 0xf3, 0x5b, 0x68, 0x85, 0xcf, 0x9a,
	0x9f, 0x06, 0xd1, 0xcf, 0x01, 0x6c, 0x9f, 0xc5, 0x84, 0xb3, 0x3c, 0x7d, 0x66, 0x2f, 0xa4, 0xcc,
	0x75, 0x15, 0x25, 0x32, 0xbe, 0x3c, 0xa2, 0x73, 0xca, 0x3d, 0x1f, 0x55, 0x15, 0x7a, 0x1f, 0xb6,
	0x16, 0xf3, 0x9a, 0xb2, 0xdc, 0x53, 0xb2, 0xac, 0x44, 0x1f, 0xc1, 0x6e, 0x2c, 0x38, 0x27, 0x9a,
	0x4a, 0xc2, 0x31, 0xe5, 0x94, 0x28, 0xea, 0x3b, 0x5a, 0x37, 0x44, 0xa7, 0xd0, 0x73, 0x7e, 0x07,
	0x84, 0xc7, 0x05, 0x37, 0x5f, 0x6e, 0xf4, 0x39, 0x6c, 0xab, 0xa5, 0xda, 0x7c, 0x9b, 0xf7, 0x7c,
	0x9b, 0x97, 0x6c, 0x78, 0x05, 0x1b, 0xfd, 0xdb, 0x04, 0x64, 0x3e, 0xc4, 0xe4, 0x9c, 0xd3, 0xca,
	0xc7, 0xe2, 0x63, 0x00, 0x56, 0x4a, 0x3e, 0x60, 0x6f, 0x75, 0x79, 0x71, 0x05, 0x83, 0x26, 0xd0,
	0xcb, 0x56, 0x4a, 0xb3, 0x37, 0xee, 0x8e, 0xef, 0x3a, 0xbf, 0xd5, 0xc2, 0x71, 0x0d, 0x8f, 0x5e,
	0x00, 0xe2, 0xb5, 0xd5, 0xf4, 0xab, 0x1d, 0xae, 0x2c, 0x52, 0x69, 0x9f, 0x36, 0xf0, 0x1a, 0x2f,
	0xf4, 0x05, 0xec, 0xd2, 0xd5, 0xb9, 0x08, 0x69, 0x75, 0x01, 0x6a, 0x63, 0x33, 0x6d, 0xe0, 0xba,
	0x0f, 0x7a, 0x02, 0x3b, 0x6a, 0x79, 0xa3, 0xc2, 0x0b, 0x1b, 0xe6, 0x4e, 0x75, 0x8f, 0xaa, 0x41,
	0x56, 0xf1, 0x93, 0x2e, 0x74, 0xe4, 0xb5, 0x10, 0xfd, 0xd2, 0x84, 0xf6, 0xb1, 0x7d, 0xb3, 0x6b,
	0x9f, 0xe4, 0x29, 0x20, 0x5d, 0xeb, 0x85, 0x67, 0xd1, 0xdf, 0xbf, 0xde, 0x2b, 0xbc, 0xc6, 0xc7,
	0x0c, 0x5f, 0x42, 0x63, 0x96, 0x11, 0x7e, 0xca, 0x49, 0x4c, 0x95, 0x7f, 0x43, 0x96, 0x95, 0xe8,
	0x91, 0x7b, 0x97, 0xdc, 0x73, 0xef, 0x79, 0xf6, 0xe4, 0xd4, 0xfe, 0x06, 0x4c, 0x1b, 0xb8, 0x02,
	0x46, 0x0f, 0x60, 0x33, 0xf1, 0xef, 0xb8, 0x67, 0xd5, 0xd3, 0xb1, 0xf2, 0xba, 0x4f, 0x1b, 0xb8,
	0x04, 0x4e, 0xb6, 0xa0, 0xab, 0x9d, 0xda, 0x50, 0x31, 0x79, 0xf7, 0xeb, 0x7b, 0xe6, 0xd5, 0xb1,
	0x7e, 0xf6, 0xff, 0x47, 0x2c, 0xf8, 0x90, 0x89, 0x91, 0x91, 0x47, 0x56, 0x71, 0xde, 0xb6, 0x3f,
	0x0f, 0xfe, 0x0b, 0x00, 0x00, 0xff, 0xff, 0x4b, 0xe5, 0xa0, 0xe5, 0xec, 0x08, 0x00, 0x00,
}
