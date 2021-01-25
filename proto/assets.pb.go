// Code generated by protoc-gen-go. DO NOT EDIT.
// source: assets.proto

package proto

import (
	fmt "fmt"
	proto "github.com/golang/protobuf/proto"
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

// The Vega representation of an external asset
type Asset struct {
	// Internal identifier of the asset
	ID string `protobuf:"bytes,1,opt,name=ID,proto3" json:"ID,omitempty"`
	// Name of the asset (e.g: Great British Pound)
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	// Symbol of the asset (e.g: GBP)
	Symbol string `protobuf:"bytes,3,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// Total circulating supply for the asset
	TotalSupply string `protobuf:"bytes,4,opt,name=totalSupply,proto3" json:"totalSupply,omitempty"`
	// Number of decimals / precision handled by this asset
	Decimals uint64 `protobuf:"varint,5,opt,name=decimals,proto3" json:"decimals,omitempty"`
	// The definition of the external source for this asset
	Source               *AssetSource `protobuf:"bytes,7,opt,name=source,proto3" json:"source,omitempty"`
	XXX_NoUnkeyedLiteral struct{}     `json:"-"`
	XXX_unrecognized     []byte       `json:"-"`
	XXX_sizecache        int32        `json:"-"`
}

func (m *Asset) Reset()         { *m = Asset{} }
func (m *Asset) String() string { return proto.CompactTextString(m) }
func (*Asset) ProtoMessage()    {}
func (*Asset) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{0}
}

func (m *Asset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_Asset.Unmarshal(m, b)
}
func (m *Asset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_Asset.Marshal(b, m, deterministic)
}
func (m *Asset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_Asset.Merge(m, src)
}
func (m *Asset) XXX_Size() int {
	return xxx_messageInfo_Asset.Size(m)
}
func (m *Asset) XXX_DiscardUnknown() {
	xxx_messageInfo_Asset.DiscardUnknown(m)
}

var xxx_messageInfo_Asset proto.InternalMessageInfo

func (m *Asset) GetID() string {
	if m != nil {
		return m.ID
	}
	return ""
}

func (m *Asset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *Asset) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *Asset) GetTotalSupply() string {
	if m != nil {
		return m.TotalSupply
	}
	return ""
}

func (m *Asset) GetDecimals() uint64 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *Asset) GetSource() *AssetSource {
	if m != nil {
		return m.Source
	}
	return nil
}

// Asset source definition
type AssetSource struct {
	// The source
	//
	// Types that are valid to be assigned to Source:
	//	*AssetSource_BuiltinAsset
	//	*AssetSource_Erc20
	Source               isAssetSource_Source `protobuf_oneof:"source"`
	XXX_NoUnkeyedLiteral struct{}             `json:"-"`
	XXX_unrecognized     []byte               `json:"-"`
	XXX_sizecache        int32                `json:"-"`
}

func (m *AssetSource) Reset()         { *m = AssetSource{} }
func (m *AssetSource) String() string { return proto.CompactTextString(m) }
func (*AssetSource) ProtoMessage()    {}
func (*AssetSource) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{1}
}

func (m *AssetSource) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_AssetSource.Unmarshal(m, b)
}
func (m *AssetSource) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_AssetSource.Marshal(b, m, deterministic)
}
func (m *AssetSource) XXX_Merge(src proto.Message) {
	xxx_messageInfo_AssetSource.Merge(m, src)
}
func (m *AssetSource) XXX_Size() int {
	return xxx_messageInfo_AssetSource.Size(m)
}
func (m *AssetSource) XXX_DiscardUnknown() {
	xxx_messageInfo_AssetSource.DiscardUnknown(m)
}

var xxx_messageInfo_AssetSource proto.InternalMessageInfo

type isAssetSource_Source interface {
	isAssetSource_Source()
}

type AssetSource_BuiltinAsset struct {
	BuiltinAsset *BuiltinAsset `protobuf:"bytes,1,opt,name=builtinAsset,proto3,oneof"`
}

type AssetSource_Erc20 struct {
	Erc20 *ERC20 `protobuf:"bytes,2,opt,name=erc20,proto3,oneof"`
}

func (*AssetSource_BuiltinAsset) isAssetSource_Source() {}

func (*AssetSource_Erc20) isAssetSource_Source() {}

func (m *AssetSource) GetSource() isAssetSource_Source {
	if m != nil {
		return m.Source
	}
	return nil
}

func (m *AssetSource) GetBuiltinAsset() *BuiltinAsset {
	if x, ok := m.GetSource().(*AssetSource_BuiltinAsset); ok {
		return x.BuiltinAsset
	}
	return nil
}

func (m *AssetSource) GetErc20() *ERC20 {
	if x, ok := m.GetSource().(*AssetSource_Erc20); ok {
		return x.Erc20
	}
	return nil
}

// XXX_OneofWrappers is for the internal use of the proto package.
func (*AssetSource) XXX_OneofWrappers() []interface{} {
	return []interface{}{
		(*AssetSource_BuiltinAsset)(nil),
		(*AssetSource_Erc20)(nil),
	}
}

// A Vega internal asset
type BuiltinAsset struct {
	// Name of the asset (e.g: Great British Pound)
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Symbol of the asset (e.g: GBP)
	Symbol string `protobuf:"bytes,2,opt,name=symbol,proto3" json:"symbol,omitempty"`
	// Total circulating supply for the asset
	TotalSupply string `protobuf:"bytes,3,opt,name=totalSupply,proto3" json:"totalSupply,omitempty"`
	// Number of decimal / precision handled by this asset
	Decimals uint64 `protobuf:"varint,4,opt,name=decimals,proto3" json:"decimals,omitempty"`
	// Maximum amount that can be requested by a party through the built-in asset faucet at a time
	MaxFaucetAmountMint  string   `protobuf:"bytes,5,opt,name=maxFaucetAmountMint,proto3" json:"maxFaucetAmountMint,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *BuiltinAsset) Reset()         { *m = BuiltinAsset{} }
func (m *BuiltinAsset) String() string { return proto.CompactTextString(m) }
func (*BuiltinAsset) ProtoMessage()    {}
func (*BuiltinAsset) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{2}
}

func (m *BuiltinAsset) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_BuiltinAsset.Unmarshal(m, b)
}
func (m *BuiltinAsset) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_BuiltinAsset.Marshal(b, m, deterministic)
}
func (m *BuiltinAsset) XXX_Merge(src proto.Message) {
	xxx_messageInfo_BuiltinAsset.Merge(m, src)
}
func (m *BuiltinAsset) XXX_Size() int {
	return xxx_messageInfo_BuiltinAsset.Size(m)
}
func (m *BuiltinAsset) XXX_DiscardUnknown() {
	xxx_messageInfo_BuiltinAsset.DiscardUnknown(m)
}

var xxx_messageInfo_BuiltinAsset proto.InternalMessageInfo

func (m *BuiltinAsset) GetName() string {
	if m != nil {
		return m.Name
	}
	return ""
}

func (m *BuiltinAsset) GetSymbol() string {
	if m != nil {
		return m.Symbol
	}
	return ""
}

func (m *BuiltinAsset) GetTotalSupply() string {
	if m != nil {
		return m.TotalSupply
	}
	return ""
}

func (m *BuiltinAsset) GetDecimals() uint64 {
	if m != nil {
		return m.Decimals
	}
	return 0
}

func (m *BuiltinAsset) GetMaxFaucetAmountMint() string {
	if m != nil {
		return m.MaxFaucetAmountMint
	}
	return ""
}

// An ERC20 token based asset, living on the ethereum network
type ERC20 struct {
	// The address of the contract for the token, on the ethereum network
	ContractAddress      string   `protobuf:"bytes,1,opt,name=contractAddress,proto3" json:"contractAddress,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ERC20) Reset()         { *m = ERC20{} }
func (m *ERC20) String() string { return proto.CompactTextString(m) }
func (*ERC20) ProtoMessage()    {}
func (*ERC20) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{3}
}

func (m *ERC20) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ERC20.Unmarshal(m, b)
}
func (m *ERC20) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ERC20.Marshal(b, m, deterministic)
}
func (m *ERC20) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ERC20.Merge(m, src)
}
func (m *ERC20) XXX_Size() int {
	return xxx_messageInfo_ERC20.Size(m)
}
func (m *ERC20) XXX_DiscardUnknown() {
	xxx_messageInfo_ERC20.DiscardUnknown(m)
}

var xxx_messageInfo_ERC20 proto.InternalMessageInfo

func (m *ERC20) GetContractAddress() string {
	if m != nil {
		return m.ContractAddress
	}
	return ""
}

// Dev assets are for use in development networks only
type DevAssets struct {
	// Asset sources for development networks
	Sources              []*AssetSource `protobuf:"bytes,1,rep,name=sources,proto3" json:"sources,omitempty"`
	XXX_NoUnkeyedLiteral struct{}       `json:"-"`
	XXX_unrecognized     []byte         `json:"-"`
	XXX_sizecache        int32          `json:"-"`
}

func (m *DevAssets) Reset()         { *m = DevAssets{} }
func (m *DevAssets) String() string { return proto.CompactTextString(m) }
func (*DevAssets) ProtoMessage()    {}
func (*DevAssets) Descriptor() ([]byte, []int) {
	return fileDescriptor_610ca40ce07a87fe, []int{4}
}

func (m *DevAssets) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DevAssets.Unmarshal(m, b)
}
func (m *DevAssets) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DevAssets.Marshal(b, m, deterministic)
}
func (m *DevAssets) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DevAssets.Merge(m, src)
}
func (m *DevAssets) XXX_Size() int {
	return xxx_messageInfo_DevAssets.Size(m)
}
func (m *DevAssets) XXX_DiscardUnknown() {
	xxx_messageInfo_DevAssets.DiscardUnknown(m)
}

var xxx_messageInfo_DevAssets proto.InternalMessageInfo

func (m *DevAssets) GetSources() []*AssetSource {
	if m != nil {
		return m.Sources
	}
	return nil
}

func init() {
	proto.RegisterType((*Asset)(nil), "vega.Asset")
	proto.RegisterType((*AssetSource)(nil), "vega.AssetSource")
	proto.RegisterType((*BuiltinAsset)(nil), "vega.BuiltinAsset")
	proto.RegisterType((*ERC20)(nil), "vega.ERC20")
	proto.RegisterType((*DevAssets)(nil), "vega.DevAssets")
}

func init() { proto.RegisterFile("assets.proto", fileDescriptor_610ca40ce07a87fe) }

var fileDescriptor_610ca40ce07a87fe = []byte{
	// 347 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x7c, 0x92, 0x4f, 0x4f, 0x83, 0x40,
	0x10, 0xc5, 0x0b, 0x85, 0xfe, 0x19, 0x1a, 0x8d, 0x6b, 0x62, 0x88, 0x17, 0x11, 0x2f, 0x18, 0x13,
	0x5a, 0xf1, 0xd2, 0x6b, 0x6b, 0x35, 0xed, 0xc1, 0xcb, 0xf6, 0xe6, 0x6d, 0xbb, 0x6c, 0x0c, 0x09,
	0xb0, 0x0d, 0xbb, 0x34, 0xf6, 0x2b, 0x19, 0x3f, 0xa4, 0x61, 0x68, 0x1b, 0x6c, 0x9a, 0x9e, 0x60,
	0x7e, 0xef, 0x2d, 0x79, 0xf3, 0x58, 0x18, 0x30, 0xa5, 0x84, 0x56, 0xe1, 0xba, 0x90, 0x5a, 0x12,
	0x6b, 0x23, 0xbe, 0x98, 0xff, 0x6b, 0x80, 0x3d, 0xa9, 0x30, 0xb9, 0x00, 0x73, 0x31, 0x73, 0x0d,
	0xcf, 0x08, 0xfa, 0xd4, 0x5c, 0xcc, 0x08, 0x01, 0x2b, 0x67, 0x99, 0x70, 0x4d, 0x24, 0xf8, 0x4e,
	0x6e, 0xa0, 0xa3, 0xb6, 0xd9, 0x4a, 0xa6, 0x6e, 0x1b, 0xe9, 0x6e, 0x22, 0x1e, 0x38, 0x5a, 0x6a,
	0x96, 0x2e, 0xcb, 0xf5, 0x3a, 0xdd, 0xba, 0x16, 0x8a, 0x4d, 0x44, 0x6e, 0xa1, 0x17, 0x0b, 0x9e,
	0x64, 0x2c, 0x55, 0xae, 0xed, 0x19, 0x81, 0x45, 0x0f, 0x33, 0x79, 0x84, 0x8e, 0x92, 0x65, 0xc1,
	0x85, 0xdb, 0xf5, 0x8c, 0xc0, 0x89, 0xae, 0xc2, 0x2a, 0x5a, 0x88, 0xb1, 0x96, 0x28, 0xd0, 0x9d,
	0xc1, 0xdf, 0x80, 0xd3, 0xc0, 0x64, 0x0c, 0x83, 0x55, 0x99, 0xa4, 0x3a, 0xc9, 0x91, 0x62, 0x7a,
	0x27, 0x22, 0xf5, 0xf9, 0x69, 0x43, 0x99, 0xb7, 0xe8, 0x3f, 0x27, 0x79, 0x00, 0x5b, 0x14, 0x3c,
	0x1a, 0xe1, 0x7a, 0x4e, 0xe4, 0xd4, 0x47, 0xde, 0xe8, 0x6b, 0x34, 0x9a, 0xb7, 0x68, 0xad, 0x4d,
	0x7b, 0xfb, 0x60, 0xfe, 0x8f, 0x01, 0x83, 0xe6, 0xf7, 0x0e, 0xed, 0x18, 0x27, 0xdb, 0x31, 0xcf,
	0xb5, 0xd3, 0x3e, 0xdf, 0x8e, 0x75, 0xd4, 0xce, 0x08, 0xae, 0x33, 0xf6, 0xfd, 0xce, 0x4a, 0x2e,
	0xf4, 0x24, 0x93, 0x65, 0xae, 0x3f, 0x92, 0x5c, 0x63, 0x89, 0x7d, 0x7a, 0x4a, 0xf2, 0x9f, 0xc1,
	0xc6, 0x45, 0x48, 0x00, 0x97, 0x5c, 0xe6, 0xba, 0x60, 0x5c, 0x4f, 0xe2, 0xb8, 0x10, 0x4a, 0xed,
	0xf2, 0x1e, 0x63, 0x7f, 0x0c, 0xfd, 0x99, 0xd8, 0xe0, 0x6a, 0x8a, 0x3c, 0x41, 0xb7, 0x5e, 0xbb,
	0xb2, 0xb7, 0x4f, 0xff, 0x90, 0xbd, 0x63, 0x7a, 0xff, 0x79, 0xc7, 0x65, 0x2c, 0xd0, 0x81, 0x17,
	0x8b, 0xcb, 0x34, 0x4c, 0xe4, 0xb0, 0x9a, 0x87, 0x08, 0x56, 0x1d, 0x7c, 0xbc, 0xfc, 0x05, 0x00,
	0x00, 0xff, 0xff, 0x83, 0x8b, 0xd5, 0x0d, 0x80, 0x02, 0x00, 0x00,
}
