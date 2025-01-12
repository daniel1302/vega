// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: vega/data_source.proto

package vega

import (
	v1 "code.vegaprotocol.io/vega/protos/vega/data/v1"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// Status describe the status of the data source spec
type DataSourceSpec_Status int32

const (
	// Default value.
	DataSourceSpec_STATUS_UNSPECIFIED DataSourceSpec_Status = 0
	// STATUS_ACTIVE describes an active data source spec.
	DataSourceSpec_STATUS_ACTIVE DataSourceSpec_Status = 1
	// STATUS_DEACTIVATED describes an data source spec that is not listening to data
	// anymore.
	DataSourceSpec_STATUS_DEACTIVATED DataSourceSpec_Status = 2
)

// Enum value maps for DataSourceSpec_Status.
var (
	DataSourceSpec_Status_name = map[int32]string{
		0: "STATUS_UNSPECIFIED",
		1: "STATUS_ACTIVE",
		2: "STATUS_DEACTIVATED",
	}
	DataSourceSpec_Status_value = map[string]int32{
		"STATUS_UNSPECIFIED": 0,
		"STATUS_ACTIVE":      1,
		"STATUS_DEACTIVATED": 2,
	}
)

func (x DataSourceSpec_Status) Enum() *DataSourceSpec_Status {
	p := new(DataSourceSpec_Status)
	*p = x
	return p
}

func (x DataSourceSpec_Status) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (DataSourceSpec_Status) Descriptor() protoreflect.EnumDescriptor {
	return file_vega_data_source_proto_enumTypes[0].Descriptor()
}

func (DataSourceSpec_Status) Type() protoreflect.EnumType {
	return &file_vega_data_source_proto_enumTypes[0]
}

func (x DataSourceSpec_Status) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use DataSourceSpec_Status.Descriptor instead.
func (DataSourceSpec_Status) EnumDescriptor() ([]byte, []int) {
	return file_vega_data_source_proto_rawDescGZIP(), []int{5, 0}
}

// DataSourceDefinition represents the top level object that deals with data sources.
// DataSourceDefinition can be external or internal, with whatever number of data sources are defined
// for each type in the child objects below.
type DataSourceDefinition struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types that are assignable to SourceType:
	//
	//	*DataSourceDefinition_Internal
	//	*DataSourceDefinition_External
	SourceType isDataSourceDefinition_SourceType `protobuf_oneof:"source_type"`
}

func (x *DataSourceDefinition) Reset() {
	*x = DataSourceDefinition{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_data_source_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSourceDefinition) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceDefinition) ProtoMessage() {}

func (x *DataSourceDefinition) ProtoReflect() protoreflect.Message {
	mi := &file_vega_data_source_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceDefinition.ProtoReflect.Descriptor instead.
func (*DataSourceDefinition) Descriptor() ([]byte, []int) {
	return file_vega_data_source_proto_rawDescGZIP(), []int{0}
}

func (m *DataSourceDefinition) GetSourceType() isDataSourceDefinition_SourceType {
	if m != nil {
		return m.SourceType
	}
	return nil
}

func (x *DataSourceDefinition) GetInternal() *DataSourceDefinitionInternal {
	if x, ok := x.GetSourceType().(*DataSourceDefinition_Internal); ok {
		return x.Internal
	}
	return nil
}

func (x *DataSourceDefinition) GetExternal() *DataSourceDefinitionExternal {
	if x, ok := x.GetSourceType().(*DataSourceDefinition_External); ok {
		return x.External
	}
	return nil
}

type isDataSourceDefinition_SourceType interface {
	isDataSourceDefinition_SourceType()
}

type DataSourceDefinition_Internal struct {
	Internal *DataSourceDefinitionInternal `protobuf:"bytes,1,opt,name=internal,proto3,oneof"`
}

type DataSourceDefinition_External struct {
	External *DataSourceDefinitionExternal `protobuf:"bytes,2,opt,name=external,proto3,oneof"`
}

func (*DataSourceDefinition_Internal) isDataSourceDefinition_SourceType() {}

func (*DataSourceDefinition_External) isDataSourceDefinition_SourceType() {}

// DataSourceSpecConfigurationTime is the internal data source used for emitting timestamps.
type DataSourceSpecConfigurationTime struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Conditions that the timestamps should meet in order to be considered.
	Conditions []*v1.Condition `protobuf:"bytes,1,rep,name=conditions,proto3" json:"conditions,omitempty"`
}

func (x *DataSourceSpecConfigurationTime) Reset() {
	*x = DataSourceSpecConfigurationTime{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_data_source_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSourceSpecConfigurationTime) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceSpecConfigurationTime) ProtoMessage() {}

func (x *DataSourceSpecConfigurationTime) ProtoReflect() protoreflect.Message {
	mi := &file_vega_data_source_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceSpecConfigurationTime.ProtoReflect.Descriptor instead.
func (*DataSourceSpecConfigurationTime) Descriptor() ([]byte, []int) {
	return file_vega_data_source_proto_rawDescGZIP(), []int{1}
}

func (x *DataSourceSpecConfigurationTime) GetConditions() []*v1.Condition {
	if x != nil {
		return x.Conditions
	}
	return nil
}

// DataSourceDefinitionInternal is the top level object used for all internal data sources.
// It contains one of any of the defined `SourceType` variants.
type DataSourceDefinitionInternal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types of internal data sources
	//
	// Types that are assignable to SourceType:
	//
	//	*DataSourceDefinitionInternal_Time
	SourceType isDataSourceDefinitionInternal_SourceType `protobuf_oneof:"source_type"`
}

func (x *DataSourceDefinitionInternal) Reset() {
	*x = DataSourceDefinitionInternal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_data_source_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSourceDefinitionInternal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceDefinitionInternal) ProtoMessage() {}

func (x *DataSourceDefinitionInternal) ProtoReflect() protoreflect.Message {
	mi := &file_vega_data_source_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceDefinitionInternal.ProtoReflect.Descriptor instead.
func (*DataSourceDefinitionInternal) Descriptor() ([]byte, []int) {
	return file_vega_data_source_proto_rawDescGZIP(), []int{2}
}

func (m *DataSourceDefinitionInternal) GetSourceType() isDataSourceDefinitionInternal_SourceType {
	if m != nil {
		return m.SourceType
	}
	return nil
}

func (x *DataSourceDefinitionInternal) GetTime() *DataSourceSpecConfigurationTime {
	if x, ok := x.GetSourceType().(*DataSourceDefinitionInternal_Time); ok {
		return x.Time
	}
	return nil
}

type isDataSourceDefinitionInternal_SourceType interface {
	isDataSourceDefinitionInternal_SourceType()
}

type DataSourceDefinitionInternal_Time struct {
	Time *DataSourceSpecConfigurationTime `protobuf:"bytes,1,opt,name=time,proto3,oneof"`
}

func (*DataSourceDefinitionInternal_Time) isDataSourceDefinitionInternal_SourceType() {}

// DataSourceDefinitionExternal is the top level object used for all external data sources.
// It contains one of any of the defined `SourceType` variants.
type DataSourceDefinitionExternal struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Types of External data sources
	//
	// Types that are assignable to SourceType:
	//
	//	*DataSourceDefinitionExternal_Oracle
	SourceType isDataSourceDefinitionExternal_SourceType `protobuf_oneof:"source_type"`
}

func (x *DataSourceDefinitionExternal) Reset() {
	*x = DataSourceDefinitionExternal{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_data_source_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSourceDefinitionExternal) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceDefinitionExternal) ProtoMessage() {}

func (x *DataSourceDefinitionExternal) ProtoReflect() protoreflect.Message {
	mi := &file_vega_data_source_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceDefinitionExternal.ProtoReflect.Descriptor instead.
func (*DataSourceDefinitionExternal) Descriptor() ([]byte, []int) {
	return file_vega_data_source_proto_rawDescGZIP(), []int{3}
}

func (m *DataSourceDefinitionExternal) GetSourceType() isDataSourceDefinitionExternal_SourceType {
	if m != nil {
		return m.SourceType
	}
	return nil
}

func (x *DataSourceDefinitionExternal) GetOracle() *DataSourceSpecConfiguration {
	if x, ok := x.GetSourceType().(*DataSourceDefinitionExternal_Oracle); ok {
		return x.Oracle
	}
	return nil
}

type isDataSourceDefinitionExternal_SourceType interface {
	isDataSourceDefinitionExternal_SourceType()
}

type DataSourceDefinitionExternal_Oracle struct {
	Oracle *DataSourceSpecConfiguration `protobuf:"bytes,1,opt,name=oracle,proto3,oneof"`
}

func (*DataSourceDefinitionExternal_Oracle) isDataSourceDefinitionExternal_SourceType() {}

// All types of external data sources use the same configuration set for meeting requirements
// in order for the data to be useful for Vega - valid signatures and matching filters.
type DataSourceSpecConfiguration struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Signers is the list of authorized signatures that signed the data for this
	// source. All the signatures in the data source data should be contained in this
	// external source. All the signatures in the data should be contained in this list.
	Signers []*v1.Signer `protobuf:"bytes,1,rep,name=signers,proto3" json:"signers,omitempty"`
	// Filters describes which source data are considered of interest or not for
	// the product (or the risk model).
	Filters []*v1.Filter `protobuf:"bytes,2,rep,name=filters,proto3" json:"filters,omitempty"`
}

func (x *DataSourceSpecConfiguration) Reset() {
	*x = DataSourceSpecConfiguration{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_data_source_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSourceSpecConfiguration) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceSpecConfiguration) ProtoMessage() {}

func (x *DataSourceSpecConfiguration) ProtoReflect() protoreflect.Message {
	mi := &file_vega_data_source_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceSpecConfiguration.ProtoReflect.Descriptor instead.
func (*DataSourceSpecConfiguration) Descriptor() ([]byte, []int) {
	return file_vega_data_source_proto_rawDescGZIP(), []int{4}
}

func (x *DataSourceSpecConfiguration) GetSigners() []*v1.Signer {
	if x != nil {
		return x.Signers
	}
	return nil
}

func (x *DataSourceSpecConfiguration) GetFilters() []*v1.Filter {
	if x != nil {
		return x.Filters
	}
	return nil
}

// Data source spec describes the data source base that a product or a risk model
// wants to get from the data source engine.
// This message contains additional information used by the API.
type DataSourceSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Hash generated from the DataSpec data.
	Id string `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	// Creation date and time
	CreatedAt int64 `protobuf:"varint,2,opt,name=created_at,json=createdAt,proto3" json:"created_at,omitempty"`
	// Last Updated timestamp
	UpdatedAt int64                 `protobuf:"varint,3,opt,name=updated_at,json=updatedAt,proto3" json:"updated_at,omitempty"`
	Data      *DataSourceDefinition `protobuf:"bytes,4,opt,name=data,proto3" json:"data,omitempty"`
	// Status describes the status of the data source spec
	Status DataSourceSpec_Status `protobuf:"varint,5,opt,name=status,proto3,enum=vega.DataSourceSpec_Status" json:"status,omitempty"`
}

func (x *DataSourceSpec) Reset() {
	*x = DataSourceSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_data_source_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DataSourceSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DataSourceSpec) ProtoMessage() {}

func (x *DataSourceSpec) ProtoReflect() protoreflect.Message {
	mi := &file_vega_data_source_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DataSourceSpec.ProtoReflect.Descriptor instead.
func (*DataSourceSpec) Descriptor() ([]byte, []int) {
	return file_vega_data_source_proto_rawDescGZIP(), []int{5}
}

func (x *DataSourceSpec) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *DataSourceSpec) GetCreatedAt() int64 {
	if x != nil {
		return x.CreatedAt
	}
	return 0
}

func (x *DataSourceSpec) GetUpdatedAt() int64 {
	if x != nil {
		return x.UpdatedAt
	}
	return 0
}

func (x *DataSourceSpec) GetData() *DataSourceDefinition {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *DataSourceSpec) GetStatus() DataSourceSpec_Status {
	if x != nil {
		return x.Status
	}
	return DataSourceSpec_STATUS_UNSPECIFIED
}

type ExternalDataSourceSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Spec *DataSourceSpec `protobuf:"bytes,1,opt,name=spec,proto3" json:"spec,omitempty"`
}

func (x *ExternalDataSourceSpec) Reset() {
	*x = ExternalDataSourceSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_data_source_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ExternalDataSourceSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ExternalDataSourceSpec) ProtoMessage() {}

func (x *ExternalDataSourceSpec) ProtoReflect() protoreflect.Message {
	mi := &file_vega_data_source_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ExternalDataSourceSpec.ProtoReflect.Descriptor instead.
func (*ExternalDataSourceSpec) Descriptor() ([]byte, []int) {
	return file_vega_data_source_proto_rawDescGZIP(), []int{6}
}

func (x *ExternalDataSourceSpec) GetSpec() *DataSourceSpec {
	if x != nil {
		return x.Spec
	}
	return nil
}

var File_vega_data_source_proto protoreflect.FileDescriptor

var file_vega_data_source_proto_rawDesc = []byte{
	0x0a, 0x16, 0x76, 0x65, 0x67, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x76, 0x65, 0x67, 0x61, 0x1a, 0x17,
	0x76, 0x65, 0x67, 0x61, 0x2f, 0x64, 0x61, 0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x64, 0x61, 0x74,
	0x61, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x17, 0x76, 0x65, 0x67, 0x61, 0x2f, 0x64, 0x61,
	0x74, 0x61, 0x2f, 0x76, 0x31, 0x2f, 0x73, 0x70, 0x65, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f,
	0x22, 0xa9, 0x01, 0x0a, 0x14, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x40, 0x0a, 0x08, 0x69, 0x6e, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e, 0x76, 0x65,
	0x67, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x66,
	0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x48,
	0x00, 0x52, 0x08, 0x69, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x40, 0x0a, 0x08, 0x65,
	0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x22, 0x2e,
	0x76, 0x65, 0x67, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44,
	0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61,
	0x6c, 0x48, 0x00, 0x52, 0x08, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x42, 0x0d, 0x0a,
	0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65, 0x22, 0x5a, 0x0a, 0x1f,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x43, 0x6f,
	0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x12,
	0x37, 0x0a, 0x0a, 0x63, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20,
	0x03, 0x28, 0x0b, 0x32, 0x17, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e,
	0x76, 0x31, 0x2e, 0x43, 0x6f, 0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x0a, 0x63, 0x6f,
	0x6e, 0x64, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x22, 0x6a, 0x0a, 0x1c, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e,
	0x49, 0x6e, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x12, 0x3b, 0x0a, 0x04, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x43, 0x6f, 0x6e, 0x66,
	0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x48, 0x00, 0x52,
	0x04, 0x74, 0x69, 0x6d, 0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x74, 0x79, 0x70, 0x65, 0x22, 0x6a, 0x0a, 0x1c, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x78, 0x74, 0x65,
	0x72, 0x6e, 0x61, 0x6c, 0x12, 0x3b, 0x0a, 0x06, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61,
	0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67,
	0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x06, 0x6f, 0x72, 0x61, 0x63, 0x6c,
	0x65, 0x42, 0x0d, 0x0a, 0x0b, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x74, 0x79, 0x70, 0x65,
	0x22, 0x7d, 0x0a, 0x1b, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x70,
	0x65, 0x63, 0x43, 0x6f, 0x6e, 0x66, 0x69, 0x67, 0x75, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12,
	0x2e, 0x0a, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x53, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x52, 0x07, 0x73, 0x69, 0x67, 0x6e, 0x65, 0x72, 0x73, 0x12,
	0x2e, 0x0a, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b,
	0x32, 0x14, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x64, 0x61, 0x74, 0x61, 0x2e, 0x76, 0x31, 0x2e,
	0x46, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x52, 0x07, 0x66, 0x69, 0x6c, 0x74, 0x65, 0x72, 0x73, 0x22,
	0x90, 0x02, 0x0a, 0x0e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x70,
	0x65, 0x63, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x74, 0x12, 0x1d, 0x0a, 0x0a, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x5f, 0x61, 0x74, 0x18,
	0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x75, 0x70, 0x64, 0x61, 0x74, 0x65, 0x64, 0x41, 0x74,
	0x12, 0x2e, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1a,
	0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x44, 0x65, 0x66, 0x69, 0x6e, 0x69, 0x74, 0x69, 0x6f, 0x6e, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61,
	0x12, 0x33, 0x0a, 0x06, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0e,
	0x32, 0x1b, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72,
	0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x52, 0x06, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x22, 0x4b, 0x0a, 0x06, 0x53, 0x74, 0x61, 0x74, 0x75, 0x73, 0x12,
	0x16, 0x0a, 0x12, 0x53, 0x54, 0x41, 0x54, 0x55, 0x53, 0x5f, 0x55, 0x4e, 0x53, 0x50, 0x45, 0x43,
	0x49, 0x46, 0x49, 0x45, 0x44, 0x10, 0x00, 0x12, 0x11, 0x0a, 0x0d, 0x53, 0x54, 0x41, 0x54, 0x55,
	0x53, 0x5f, 0x41, 0x43, 0x54, 0x49, 0x56, 0x45, 0x10, 0x01, 0x12, 0x16, 0x0a, 0x12, 0x53, 0x54,
	0x41, 0x54, 0x55, 0x53, 0x5f, 0x44, 0x45, 0x41, 0x43, 0x54, 0x49, 0x56, 0x41, 0x54, 0x45, 0x44,
	0x10, 0x02, 0x22, 0x42, 0x0a, 0x16, 0x45, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x28, 0x0a, 0x04,
	0x73, 0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x76, 0x65, 0x67,
	0x61, 0x2e, 0x44, 0x61, 0x74, 0x61, 0x53, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x04, 0x73, 0x70, 0x65, 0x63, 0x42, 0x27, 0x5a, 0x25, 0x63, 0x6f, 0x64, 0x65, 0x2e, 0x76,
	0x65, 0x67, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x69, 0x6f, 0x2f, 0x76,
	0x65, 0x67, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x65, 0x67, 0x61, 0x62,
	0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vega_data_source_proto_rawDescOnce sync.Once
	file_vega_data_source_proto_rawDescData = file_vega_data_source_proto_rawDesc
)

func file_vega_data_source_proto_rawDescGZIP() []byte {
	file_vega_data_source_proto_rawDescOnce.Do(func() {
		file_vega_data_source_proto_rawDescData = protoimpl.X.CompressGZIP(file_vega_data_source_proto_rawDescData)
	})
	return file_vega_data_source_proto_rawDescData
}

var file_vega_data_source_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_vega_data_source_proto_msgTypes = make([]protoimpl.MessageInfo, 7)
var file_vega_data_source_proto_goTypes = []interface{}{
	(DataSourceSpec_Status)(0),              // 0: vega.DataSourceSpec.Status
	(*DataSourceDefinition)(nil),            // 1: vega.DataSourceDefinition
	(*DataSourceSpecConfigurationTime)(nil), // 2: vega.DataSourceSpecConfigurationTime
	(*DataSourceDefinitionInternal)(nil),    // 3: vega.DataSourceDefinitionInternal
	(*DataSourceDefinitionExternal)(nil),    // 4: vega.DataSourceDefinitionExternal
	(*DataSourceSpecConfiguration)(nil),     // 5: vega.DataSourceSpecConfiguration
	(*DataSourceSpec)(nil),                  // 6: vega.DataSourceSpec
	(*ExternalDataSourceSpec)(nil),          // 7: vega.ExternalDataSourceSpec
	(*v1.Condition)(nil),                    // 8: vega.data.v1.Condition
	(*v1.Signer)(nil),                       // 9: vega.data.v1.Signer
	(*v1.Filter)(nil),                       // 10: vega.data.v1.Filter
}
var file_vega_data_source_proto_depIdxs = []int32{
	3,  // 0: vega.DataSourceDefinition.internal:type_name -> vega.DataSourceDefinitionInternal
	4,  // 1: vega.DataSourceDefinition.external:type_name -> vega.DataSourceDefinitionExternal
	8,  // 2: vega.DataSourceSpecConfigurationTime.conditions:type_name -> vega.data.v1.Condition
	2,  // 3: vega.DataSourceDefinitionInternal.time:type_name -> vega.DataSourceSpecConfigurationTime
	5,  // 4: vega.DataSourceDefinitionExternal.oracle:type_name -> vega.DataSourceSpecConfiguration
	9,  // 5: vega.DataSourceSpecConfiguration.signers:type_name -> vega.data.v1.Signer
	10, // 6: vega.DataSourceSpecConfiguration.filters:type_name -> vega.data.v1.Filter
	1,  // 7: vega.DataSourceSpec.data:type_name -> vega.DataSourceDefinition
	0,  // 8: vega.DataSourceSpec.status:type_name -> vega.DataSourceSpec.Status
	6,  // 9: vega.ExternalDataSourceSpec.spec:type_name -> vega.DataSourceSpec
	10, // [10:10] is the sub-list for method output_type
	10, // [10:10] is the sub-list for method input_type
	10, // [10:10] is the sub-list for extension type_name
	10, // [10:10] is the sub-list for extension extendee
	0,  // [0:10] is the sub-list for field type_name
}

func init() { file_vega_data_source_proto_init() }
func file_vega_data_source_proto_init() {
	if File_vega_data_source_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vega_data_source_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSourceDefinition); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_data_source_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSourceSpecConfigurationTime); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_data_source_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSourceDefinitionInternal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_data_source_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSourceDefinitionExternal); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_data_source_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSourceSpecConfiguration); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_data_source_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DataSourceSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_vega_data_source_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ExternalDataSourceSpec); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	file_vega_data_source_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*DataSourceDefinition_Internal)(nil),
		(*DataSourceDefinition_External)(nil),
	}
	file_vega_data_source_proto_msgTypes[2].OneofWrappers = []interface{}{
		(*DataSourceDefinitionInternal_Time)(nil),
	}
	file_vega_data_source_proto_msgTypes[3].OneofWrappers = []interface{}{
		(*DataSourceDefinitionExternal_Oracle)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vega_data_source_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   7,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vega_data_source_proto_goTypes,
		DependencyIndexes: file_vega_data_source_proto_depIdxs,
		EnumInfos:         file_vega_data_source_proto_enumTypes,
		MessageInfos:      file_vega_data_source_proto_msgTypes,
	}.Build()
	File_vega_data_source_proto = out.File
	file_vega_data_source_proto_rawDesc = nil
	file_vega_data_source_proto_goTypes = nil
	file_vega_data_source_proto_depIdxs = nil
}
