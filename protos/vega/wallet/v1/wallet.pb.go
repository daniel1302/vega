// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.27.1
// 	protoc        (unknown)
// source: vega/wallet/v1/wallet.proto

package v1

import (
	v1 "code.vegaprotocol.io/vega/protos/vega/commands/v1"
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

type SubmitTransactionRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	PubKey    string `protobuf:"bytes,1,opt,name=pub_key,json=pubKey,proto3" json:"pub_key,omitempty"`
	Propagate bool   `protobuf:"varint,2,opt,name=propagate,proto3" json:"propagate,omitempty"`
	// Types that are assignable to Command:
	//	*SubmitTransactionRequest_OrderSubmission
	//	*SubmitTransactionRequest_OrderCancellation
	//	*SubmitTransactionRequest_OrderAmendment
	//	*SubmitTransactionRequest_WithdrawSubmission
	//	*SubmitTransactionRequest_ProposalSubmission
	//	*SubmitTransactionRequest_VoteSubmission
	//	*SubmitTransactionRequest_LiquidityProvisionSubmission
	//	*SubmitTransactionRequest_DelegateSubmission
	//	*SubmitTransactionRequest_UndelegateSubmission
	//	*SubmitTransactionRequest_LiquidityProvisionCancellation
	//	*SubmitTransactionRequest_LiquidityProvisionAmendment
	//	*SubmitTransactionRequest_Transfer
	//	*SubmitTransactionRequest_CancelTransfer
	//	*SubmitTransactionRequest_AnnounceNode
	//	*SubmitTransactionRequest_NodeVote
	//	*SubmitTransactionRequest_NodeSignature
	//	*SubmitTransactionRequest_ChainEvent
	//	*SubmitTransactionRequest_KeyRotateSubmission
	//	*SubmitTransactionRequest_StateVariableProposal
	//	*SubmitTransactionRequest_ValidatorHeartbeat
	//	*SubmitTransactionRequest_EthereumKeyRotateSubmission
	//	*SubmitTransactionRequest_ProtocolUpgradeProposal
	//	*SubmitTransactionRequest_IssueSignatures
	//	*SubmitTransactionRequest_OracleDataSubmission
	Command isSubmitTransactionRequest_Command `protobuf_oneof:"command"`
}

func (x *SubmitTransactionRequest) Reset() {
	*x = SubmitTransactionRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_vega_wallet_v1_wallet_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *SubmitTransactionRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*SubmitTransactionRequest) ProtoMessage() {}

func (x *SubmitTransactionRequest) ProtoReflect() protoreflect.Message {
	mi := &file_vega_wallet_v1_wallet_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use SubmitTransactionRequest.ProtoReflect.Descriptor instead.
func (*SubmitTransactionRequest) Descriptor() ([]byte, []int) {
	return file_vega_wallet_v1_wallet_proto_rawDescGZIP(), []int{0}
}

func (x *SubmitTransactionRequest) GetPubKey() string {
	if x != nil {
		return x.PubKey
	}
	return ""
}

func (x *SubmitTransactionRequest) GetPropagate() bool {
	if x != nil {
		return x.Propagate
	}
	return false
}

func (m *SubmitTransactionRequest) GetCommand() isSubmitTransactionRequest_Command {
	if m != nil {
		return m.Command
	}
	return nil
}

func (x *SubmitTransactionRequest) GetOrderSubmission() *v1.OrderSubmission {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_OrderSubmission); ok {
		return x.OrderSubmission
	}
	return nil
}

func (x *SubmitTransactionRequest) GetOrderCancellation() *v1.OrderCancellation {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_OrderCancellation); ok {
		return x.OrderCancellation
	}
	return nil
}

func (x *SubmitTransactionRequest) GetOrderAmendment() *v1.OrderAmendment {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_OrderAmendment); ok {
		return x.OrderAmendment
	}
	return nil
}

func (x *SubmitTransactionRequest) GetWithdrawSubmission() *v1.WithdrawSubmission {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_WithdrawSubmission); ok {
		return x.WithdrawSubmission
	}
	return nil
}

func (x *SubmitTransactionRequest) GetProposalSubmission() *v1.ProposalSubmission {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_ProposalSubmission); ok {
		return x.ProposalSubmission
	}
	return nil
}

func (x *SubmitTransactionRequest) GetVoteSubmission() *v1.VoteSubmission {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_VoteSubmission); ok {
		return x.VoteSubmission
	}
	return nil
}

func (x *SubmitTransactionRequest) GetLiquidityProvisionSubmission() *v1.LiquidityProvisionSubmission {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_LiquidityProvisionSubmission); ok {
		return x.LiquidityProvisionSubmission
	}
	return nil
}

func (x *SubmitTransactionRequest) GetDelegateSubmission() *v1.DelegateSubmission {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_DelegateSubmission); ok {
		return x.DelegateSubmission
	}
	return nil
}

func (x *SubmitTransactionRequest) GetUndelegateSubmission() *v1.UndelegateSubmission {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_UndelegateSubmission); ok {
		return x.UndelegateSubmission
	}
	return nil
}

func (x *SubmitTransactionRequest) GetLiquidityProvisionCancellation() *v1.LiquidityProvisionCancellation {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_LiquidityProvisionCancellation); ok {
		return x.LiquidityProvisionCancellation
	}
	return nil
}

func (x *SubmitTransactionRequest) GetLiquidityProvisionAmendment() *v1.LiquidityProvisionAmendment {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_LiquidityProvisionAmendment); ok {
		return x.LiquidityProvisionAmendment
	}
	return nil
}

func (x *SubmitTransactionRequest) GetTransfer() *v1.Transfer {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_Transfer); ok {
		return x.Transfer
	}
	return nil
}

func (x *SubmitTransactionRequest) GetCancelTransfer() *v1.CancelTransfer {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_CancelTransfer); ok {
		return x.CancelTransfer
	}
	return nil
}

func (x *SubmitTransactionRequest) GetAnnounceNode() *v1.AnnounceNode {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_AnnounceNode); ok {
		return x.AnnounceNode
	}
	return nil
}

func (x *SubmitTransactionRequest) GetNodeVote() *v1.NodeVote {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_NodeVote); ok {
		return x.NodeVote
	}
	return nil
}

func (x *SubmitTransactionRequest) GetNodeSignature() *v1.NodeSignature {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_NodeSignature); ok {
		return x.NodeSignature
	}
	return nil
}

func (x *SubmitTransactionRequest) GetChainEvent() *v1.ChainEvent {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_ChainEvent); ok {
		return x.ChainEvent
	}
	return nil
}

func (x *SubmitTransactionRequest) GetKeyRotateSubmission() *v1.KeyRotateSubmission {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_KeyRotateSubmission); ok {
		return x.KeyRotateSubmission
	}
	return nil
}

func (x *SubmitTransactionRequest) GetStateVariableProposal() *v1.StateVariableProposal {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_StateVariableProposal); ok {
		return x.StateVariableProposal
	}
	return nil
}

func (x *SubmitTransactionRequest) GetValidatorHeartbeat() *v1.ValidatorHeartbeat {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_ValidatorHeartbeat); ok {
		return x.ValidatorHeartbeat
	}
	return nil
}

func (x *SubmitTransactionRequest) GetEthereumKeyRotateSubmission() *v1.EthereumKeyRotateSubmission {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_EthereumKeyRotateSubmission); ok {
		return x.EthereumKeyRotateSubmission
	}
	return nil
}

func (x *SubmitTransactionRequest) GetProtocolUpgradeProposal() *v1.ProtocolUpgradeProposal {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_ProtocolUpgradeProposal); ok {
		return x.ProtocolUpgradeProposal
	}
	return nil
}

func (x *SubmitTransactionRequest) GetIssueSignatures() *v1.IssueSignatures {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_IssueSignatures); ok {
		return x.IssueSignatures
	}
	return nil
}

func (x *SubmitTransactionRequest) GetOracleDataSubmission() *v1.OracleDataSubmission {
	if x, ok := x.GetCommand().(*SubmitTransactionRequest_OracleDataSubmission); ok {
		return x.OracleDataSubmission
	}
	return nil
}

type isSubmitTransactionRequest_Command interface {
	isSubmitTransactionRequest_Command()
}

type SubmitTransactionRequest_OrderSubmission struct {
	// User commands
	OrderSubmission *v1.OrderSubmission `protobuf:"bytes,1001,opt,name=order_submission,json=orderSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_OrderCancellation struct {
	OrderCancellation *v1.OrderCancellation `protobuf:"bytes,1002,opt,name=order_cancellation,json=orderCancellation,proto3,oneof"`
}

type SubmitTransactionRequest_OrderAmendment struct {
	OrderAmendment *v1.OrderAmendment `protobuf:"bytes,1003,opt,name=order_amendment,json=orderAmendment,proto3,oneof"`
}

type SubmitTransactionRequest_WithdrawSubmission struct {
	WithdrawSubmission *v1.WithdrawSubmission `protobuf:"bytes,1004,opt,name=withdraw_submission,json=withdrawSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_ProposalSubmission struct {
	ProposalSubmission *v1.ProposalSubmission `protobuf:"bytes,1005,opt,name=proposal_submission,json=proposalSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_VoteSubmission struct {
	VoteSubmission *v1.VoteSubmission `protobuf:"bytes,1006,opt,name=vote_submission,json=voteSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_LiquidityProvisionSubmission struct {
	LiquidityProvisionSubmission *v1.LiquidityProvisionSubmission `protobuf:"bytes,1007,opt,name=liquidity_provision_submission,json=liquidityProvisionSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_DelegateSubmission struct {
	DelegateSubmission *v1.DelegateSubmission `protobuf:"bytes,1008,opt,name=delegate_submission,json=delegateSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_UndelegateSubmission struct {
	UndelegateSubmission *v1.UndelegateSubmission `protobuf:"bytes,1009,opt,name=undelegate_submission,json=undelegateSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_LiquidityProvisionCancellation struct {
	LiquidityProvisionCancellation *v1.LiquidityProvisionCancellation `protobuf:"bytes,1010,opt,name=liquidity_provision_cancellation,json=liquidityProvisionCancellation,proto3,oneof"`
}

type SubmitTransactionRequest_LiquidityProvisionAmendment struct {
	LiquidityProvisionAmendment *v1.LiquidityProvisionAmendment `protobuf:"bytes,1011,opt,name=liquidity_provision_amendment,json=liquidityProvisionAmendment,proto3,oneof"`
}

type SubmitTransactionRequest_Transfer struct {
	Transfer *v1.Transfer `protobuf:"bytes,1012,opt,name=transfer,proto3,oneof"`
}

type SubmitTransactionRequest_CancelTransfer struct {
	CancelTransfer *v1.CancelTransfer `protobuf:"bytes,1013,opt,name=cancel_transfer,json=cancelTransfer,proto3,oneof"`
}

type SubmitTransactionRequest_AnnounceNode struct {
	AnnounceNode *v1.AnnounceNode `protobuf:"bytes,1014,opt,name=announce_node,json=announceNode,proto3,oneof"`
}

type SubmitTransactionRequest_NodeVote struct {
	// Validator commands
	NodeVote *v1.NodeVote `protobuf:"bytes,2002,opt,name=node_vote,json=nodeVote,proto3,oneof"`
}

type SubmitTransactionRequest_NodeSignature struct {
	NodeSignature *v1.NodeSignature `protobuf:"bytes,2003,opt,name=node_signature,json=nodeSignature,proto3,oneof"`
}

type SubmitTransactionRequest_ChainEvent struct {
	ChainEvent *v1.ChainEvent `protobuf:"bytes,2004,opt,name=chain_event,json=chainEvent,proto3,oneof"`
}

type SubmitTransactionRequest_KeyRotateSubmission struct {
	KeyRotateSubmission *v1.KeyRotateSubmission `protobuf:"bytes,2005,opt,name=key_rotate_submission,json=keyRotateSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_StateVariableProposal struct {
	StateVariableProposal *v1.StateVariableProposal `protobuf:"bytes,2006,opt,name=state_variable_proposal,json=stateVariableProposal,proto3,oneof"`
}

type SubmitTransactionRequest_ValidatorHeartbeat struct {
	ValidatorHeartbeat *v1.ValidatorHeartbeat `protobuf:"bytes,2007,opt,name=validator_heartbeat,json=validatorHeartbeat,proto3,oneof"`
}

type SubmitTransactionRequest_EthereumKeyRotateSubmission struct {
	EthereumKeyRotateSubmission *v1.EthereumKeyRotateSubmission `protobuf:"bytes,2008,opt,name=ethereum_key_rotate_submission,json=ethereumKeyRotateSubmission,proto3,oneof"`
}

type SubmitTransactionRequest_ProtocolUpgradeProposal struct {
	ProtocolUpgradeProposal *v1.ProtocolUpgradeProposal `protobuf:"bytes,2009,opt,name=protocol_upgrade_proposal,json=protocolUpgradeProposal,proto3,oneof"`
}

type SubmitTransactionRequest_IssueSignatures struct {
	IssueSignatures *v1.IssueSignatures `protobuf:"bytes,2010,opt,name=issue_signatures,json=issueSignatures,proto3,oneof"`
}

type SubmitTransactionRequest_OracleDataSubmission struct {
	// Oracle commands
	OracleDataSubmission *v1.OracleDataSubmission `protobuf:"bytes,3001,opt,name=oracle_data_submission,json=oracleDataSubmission,proto3,oneof"`
}

func (*SubmitTransactionRequest_OrderSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_OrderCancellation) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_OrderAmendment) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_WithdrawSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_ProposalSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_VoteSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_LiquidityProvisionSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_DelegateSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_UndelegateSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_LiquidityProvisionCancellation) isSubmitTransactionRequest_Command() {
}

func (*SubmitTransactionRequest_LiquidityProvisionAmendment) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_Transfer) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_CancelTransfer) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_AnnounceNode) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_NodeVote) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_NodeSignature) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_ChainEvent) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_KeyRotateSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_StateVariableProposal) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_ValidatorHeartbeat) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_EthereumKeyRotateSubmission) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_ProtocolUpgradeProposal) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_IssueSignatures) isSubmitTransactionRequest_Command() {}

func (*SubmitTransactionRequest_OracleDataSubmission) isSubmitTransactionRequest_Command() {}

var File_vega_wallet_v1_wallet_proto protoreflect.FileDescriptor

var file_vega_wallet_v1_wallet_proto_rawDesc = []byte{
	0x0a, 0x1b, 0x76, 0x65, 0x67, 0x61, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x76, 0x31,
	0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x0e, 0x76,
	0x65, 0x67, 0x61, 0x2e, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2e, 0x76, 0x31, 0x1a, 0x1f, 0x76,
	0x65, 0x67, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x29,
	0x76, 0x65, 0x67, 0x61, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31,
	0x2f, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x1e, 0x76, 0x65, 0x67, 0x61, 0x2f,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2f, 0x76, 0x31, 0x2f, 0x6f, 0x72, 0x61, 0x63,
	0x6c, 0x65, 0x73, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xcd, 0x11, 0x0a, 0x18, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x74, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x17, 0x0a, 0x07, 0x70, 0x75, 0x62, 0x5f, 0x6b, 0x65,
	0x79, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x70, 0x75, 0x62, 0x4b, 0x65, 0x79, 0x12,
	0x1c, 0x0a, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x08, 0x52, 0x09, 0x70, 0x72, 0x6f, 0x70, 0x61, 0x67, 0x61, 0x74, 0x65, 0x12, 0x4f, 0x0a,
	0x10, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0xe9, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x21, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65,
	0x72, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0f, 0x6f,
	0x72, 0x64, 0x65, 0x72, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x55,
	0x0a, 0x12, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x18, 0xea, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x23, 0x2e, 0x76, 0x65,
	0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f,
	0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x11, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x0f, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x5f, 0x61,
	0x6d, 0x65, 0x6e, 0x64, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0xeb, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x20, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x6d, 0x65, 0x6e, 0x64, 0x6d, 0x65, 0x6e,
	0x74, 0x48, 0x00, 0x52, 0x0e, 0x6f, 0x72, 0x64, 0x65, 0x72, 0x41, 0x6d, 0x65, 0x6e, 0x64, 0x6d,
	0x65, 0x6e, 0x74, 0x12, 0x58, 0x0a, 0x13, 0x77, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x5f,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0xec, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x24, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x57, 0x69, 0x74, 0x68, 0x64, 0x72, 0x61, 0x77, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x12, 0x77, 0x69, 0x74, 0x68, 0x64,
	0x72, 0x61, 0x77, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x58, 0x0a,
	0x13, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73,
	0x73, 0x69, 0x6f, 0x6e, 0x18, 0xed, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x65,
	0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x50,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x48, 0x00, 0x52, 0x12, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x53, 0x75, 0x62,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x4c, 0x0a, 0x0f, 0x76, 0x6f, 0x74, 0x65, 0x5f,
	0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0xee, 0x07, 0x20, 0x01, 0x28,
	0x0b, 0x32, 0x20, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64,
	0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x6f, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73,
	0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x0e, 0x76, 0x6f, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x77, 0x0a, 0x1e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x73, 0x75, 0x62,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0xef, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2e,
	0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x1c, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x58,
	0x0a, 0x13, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0xf0, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76,
	0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x44, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69,
	0x6f, 0x6e, 0x48, 0x00, 0x52, 0x12, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x5e, 0x0a, 0x15, 0x75, 0x6e, 0x64, 0x65,
	0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f,
	0x6e, 0x18, 0xf1, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x55, 0x6e, 0x64, 0x65,
	0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x48, 0x00, 0x52, 0x14, 0x75, 0x6e, 0x64, 0x65, 0x6c, 0x65, 0x67, 0x61, 0x74, 0x65, 0x53, 0x75,
	0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x7d, 0x0a, 0x20, 0x6c, 0x69, 0x71, 0x75,
	0x69, 0x64, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f,
	0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0xf2, 0x07, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x30, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79,
	0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x6c,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x1e, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69,
	0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x43, 0x61, 0x6e, 0x63, 0x65,
	0x6c, 0x6c, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x74, 0x0a, 0x1d, 0x6c, 0x69, 0x71, 0x75, 0x69,
	0x64, 0x69, 0x74, 0x79, 0x5f, 0x70, 0x72, 0x6f, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x61,
	0x6d, 0x65, 0x6e, 0x64, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0xf3, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x2d, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x4c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x65, 0x6e, 0x64, 0x6d, 0x65, 0x6e, 0x74, 0x48, 0x00,
	0x52, 0x1b, 0x6c, 0x69, 0x71, 0x75, 0x69, 0x64, 0x69, 0x74, 0x79, 0x50, 0x72, 0x6f, 0x76, 0x69,
	0x73, 0x69, 0x6f, 0x6e, 0x41, 0x6d, 0x65, 0x6e, 0x64, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x39, 0x0a,
	0x08, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0xf4, 0x07, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x1a, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73,
	0x2e, 0x76, 0x31, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x48, 0x00, 0x52, 0x08,
	0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x4c, 0x0a, 0x0f, 0x63, 0x61, 0x6e, 0x63,
	0x65, 0x6c, 0x5f, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x18, 0xf5, 0x07, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x20, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72, 0x61, 0x6e,
	0x73, 0x66, 0x65, 0x72, 0x48, 0x00, 0x52, 0x0e, 0x63, 0x61, 0x6e, 0x63, 0x65, 0x6c, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x12, 0x46, 0x0a, 0x0d, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e,
	0x63, 0x65, 0x5f, 0x6e, 0x6f, 0x64, 0x65, 0x18, 0xf6, 0x07, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1e,
	0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x41, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x48, 0x00,
	0x52, 0x0c, 0x61, 0x6e, 0x6e, 0x6f, 0x75, 0x6e, 0x63, 0x65, 0x4e, 0x6f, 0x64, 0x65, 0x12, 0x3a,
	0x0a, 0x09, 0x6e, 0x6f, 0x64, 0x65, 0x5f, 0x76, 0x6f, 0x74, 0x65, 0x18, 0xd2, 0x0f, 0x20, 0x01,
	0x28, 0x0b, 0x32, 0x1a, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x48, 0x00,
	0x52, 0x08, 0x6e, 0x6f, 0x64, 0x65, 0x56, 0x6f, 0x74, 0x65, 0x12, 0x49, 0x0a, 0x0e, 0x6e, 0x6f,
	0x64, 0x65, 0x5f, 0x73, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0xd3, 0x0f, 0x20,
	0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61,
	0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4e, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x48, 0x00, 0x52, 0x0d, 0x6e, 0x6f, 0x64, 0x65, 0x53, 0x69, 0x67, 0x6e,
	0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x40, 0x0a, 0x0b, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x5f, 0x65,
	0x76, 0x65, 0x6e, 0x74, 0x18, 0xd4, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1c, 0x2e, 0x76, 0x65,
	0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x43,
	0x68, 0x61, 0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x48, 0x00, 0x52, 0x0a, 0x63, 0x68, 0x61,
	0x69, 0x6e, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x5c, 0x0a, 0x15, 0x6b, 0x65, 0x79, 0x5f, 0x72,
	0x6f, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e,
	0x18, 0xd5, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x25, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4b, 0x65, 0x79, 0x52, 0x6f,
	0x74, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x13, 0x6b, 0x65, 0x79, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x53, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x62, 0x0a, 0x17, 0x73, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x76,
	0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x5f, 0x70, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x18, 0xd6, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c, 0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c,
	0x48, 0x00, 0x52, 0x15, 0x73, 0x74, 0x61, 0x74, 0x65, 0x56, 0x61, 0x72, 0x69, 0x61, 0x62, 0x6c,
	0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x12, 0x58, 0x0a, 0x13, 0x76, 0x61, 0x6c,
	0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x5f, 0x68, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74,
	0x18, 0xd7, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x56, 0x61, 0x6c, 0x69, 0x64,
	0x61, 0x74, 0x6f, 0x72, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62, 0x65, 0x61, 0x74, 0x48, 0x00, 0x52,
	0x12, 0x76, 0x61, 0x6c, 0x69, 0x64, 0x61, 0x74, 0x6f, 0x72, 0x48, 0x65, 0x61, 0x72, 0x74, 0x62,
	0x65, 0x61, 0x74, 0x12, 0x75, 0x0a, 0x1e, 0x65, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x5f,
	0x6b, 0x65, 0x79, 0x5f, 0x72, 0x6f, 0x74, 0x61, 0x74, 0x65, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69,
	0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0xd8, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x2d, 0x2e, 0x76,
	0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e,
	0x45, 0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x4b, 0x65, 0x79, 0x52, 0x6f, 0x74, 0x61, 0x74,
	0x65, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00, 0x52, 0x1b, 0x65,
	0x74, 0x68, 0x65, 0x72, 0x65, 0x75, 0x6d, 0x4b, 0x65, 0x79, 0x52, 0x6f, 0x74, 0x61, 0x74, 0x65,
	0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x12, 0x68, 0x0a, 0x19, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x5f, 0x75, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x5f, 0x70,
	0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x18, 0xd9, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x29,
	0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76,
	0x31, 0x2e, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64,
	0x65, 0x50, 0x72, 0x6f, 0x70, 0x6f, 0x73, 0x61, 0x6c, 0x48, 0x00, 0x52, 0x17, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x55, 0x70, 0x67, 0x72, 0x61, 0x64, 0x65, 0x50, 0x72, 0x6f, 0x70,
	0x6f, 0x73, 0x61, 0x6c, 0x12, 0x4f, 0x0a, 0x10, 0x69, 0x73, 0x73, 0x75, 0x65, 0x5f, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x73, 0x18, 0xda, 0x0f, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x21, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e,
	0x76, 0x31, 0x2e, 0x49, 0x73, 0x73, 0x75, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72,
	0x65, 0x73, 0x48, 0x00, 0x52, 0x0f, 0x69, 0x73, 0x73, 0x75, 0x65, 0x53, 0x69, 0x67, 0x6e, 0x61,
	0x74, 0x75, 0x72, 0x65, 0x73, 0x12, 0x5f, 0x0a, 0x16, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x5f,
	0x64, 0x61, 0x74, 0x61, 0x5f, 0x73, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x18,
	0xb9, 0x17, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x26, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x61, 0x6e, 0x64, 0x73, 0x2e, 0x76, 0x31, 0x2e, 0x4f, 0x72, 0x61, 0x63, 0x6c, 0x65,
	0x44, 0x61, 0x74, 0x61, 0x53, 0x75, 0x62, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x48, 0x00,
	0x52, 0x14, 0x6f, 0x72, 0x61, 0x63, 0x6c, 0x65, 0x44, 0x61, 0x74, 0x61, 0x53, 0x75, 0x62, 0x6d,
	0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x42, 0x09, 0x0a, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x61, 0x6e,
	0x64, 0x4a, 0x06, 0x08, 0xd1, 0x0f, 0x10, 0xd2, 0x0f, 0x42, 0x31, 0x5a, 0x2f, 0x63, 0x6f, 0x64,
	0x65, 0x2e, 0x76, 0x65, 0x67, 0x61, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x63, 0x6f, 0x6c, 0x2e, 0x69,
	0x6f, 0x2f, 0x76, 0x65, 0x67, 0x61, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2f, 0x76, 0x65,
	0x67, 0x61, 0x2f, 0x77, 0x61, 0x6c, 0x6c, 0x65, 0x74, 0x2f, 0x76, 0x31, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_vega_wallet_v1_wallet_proto_rawDescOnce sync.Once
	file_vega_wallet_v1_wallet_proto_rawDescData = file_vega_wallet_v1_wallet_proto_rawDesc
)

func file_vega_wallet_v1_wallet_proto_rawDescGZIP() []byte {
	file_vega_wallet_v1_wallet_proto_rawDescOnce.Do(func() {
		file_vega_wallet_v1_wallet_proto_rawDescData = protoimpl.X.CompressGZIP(file_vega_wallet_v1_wallet_proto_rawDescData)
	})
	return file_vega_wallet_v1_wallet_proto_rawDescData
}

var file_vega_wallet_v1_wallet_proto_msgTypes = make([]protoimpl.MessageInfo, 1)
var file_vega_wallet_v1_wallet_proto_goTypes = []interface{}{
	(*SubmitTransactionRequest)(nil),          // 0: vega.wallet.v1.SubmitTransactionRequest
	(*v1.OrderSubmission)(nil),                // 1: vega.commands.v1.OrderSubmission
	(*v1.OrderCancellation)(nil),              // 2: vega.commands.v1.OrderCancellation
	(*v1.OrderAmendment)(nil),                 // 3: vega.commands.v1.OrderAmendment
	(*v1.WithdrawSubmission)(nil),             // 4: vega.commands.v1.WithdrawSubmission
	(*v1.ProposalSubmission)(nil),             // 5: vega.commands.v1.ProposalSubmission
	(*v1.VoteSubmission)(nil),                 // 6: vega.commands.v1.VoteSubmission
	(*v1.LiquidityProvisionSubmission)(nil),   // 7: vega.commands.v1.LiquidityProvisionSubmission
	(*v1.DelegateSubmission)(nil),             // 8: vega.commands.v1.DelegateSubmission
	(*v1.UndelegateSubmission)(nil),           // 9: vega.commands.v1.UndelegateSubmission
	(*v1.LiquidityProvisionCancellation)(nil), // 10: vega.commands.v1.LiquidityProvisionCancellation
	(*v1.LiquidityProvisionAmendment)(nil),    // 11: vega.commands.v1.LiquidityProvisionAmendment
	(*v1.Transfer)(nil),                       // 12: vega.commands.v1.Transfer
	(*v1.CancelTransfer)(nil),                 // 13: vega.commands.v1.CancelTransfer
	(*v1.AnnounceNode)(nil),                   // 14: vega.commands.v1.AnnounceNode
	(*v1.NodeVote)(nil),                       // 15: vega.commands.v1.NodeVote
	(*v1.NodeSignature)(nil),                  // 16: vega.commands.v1.NodeSignature
	(*v1.ChainEvent)(nil),                     // 17: vega.commands.v1.ChainEvent
	(*v1.KeyRotateSubmission)(nil),            // 18: vega.commands.v1.KeyRotateSubmission
	(*v1.StateVariableProposal)(nil),          // 19: vega.commands.v1.StateVariableProposal
	(*v1.ValidatorHeartbeat)(nil),             // 20: vega.commands.v1.ValidatorHeartbeat
	(*v1.EthereumKeyRotateSubmission)(nil),    // 21: vega.commands.v1.EthereumKeyRotateSubmission
	(*v1.ProtocolUpgradeProposal)(nil),        // 22: vega.commands.v1.ProtocolUpgradeProposal
	(*v1.IssueSignatures)(nil),                // 23: vega.commands.v1.IssueSignatures
	(*v1.OracleDataSubmission)(nil),           // 24: vega.commands.v1.OracleDataSubmission
}
var file_vega_wallet_v1_wallet_proto_depIdxs = []int32{
	1,  // 0: vega.wallet.v1.SubmitTransactionRequest.order_submission:type_name -> vega.commands.v1.OrderSubmission
	2,  // 1: vega.wallet.v1.SubmitTransactionRequest.order_cancellation:type_name -> vega.commands.v1.OrderCancellation
	3,  // 2: vega.wallet.v1.SubmitTransactionRequest.order_amendment:type_name -> vega.commands.v1.OrderAmendment
	4,  // 3: vega.wallet.v1.SubmitTransactionRequest.withdraw_submission:type_name -> vega.commands.v1.WithdrawSubmission
	5,  // 4: vega.wallet.v1.SubmitTransactionRequest.proposal_submission:type_name -> vega.commands.v1.ProposalSubmission
	6,  // 5: vega.wallet.v1.SubmitTransactionRequest.vote_submission:type_name -> vega.commands.v1.VoteSubmission
	7,  // 6: vega.wallet.v1.SubmitTransactionRequest.liquidity_provision_submission:type_name -> vega.commands.v1.LiquidityProvisionSubmission
	8,  // 7: vega.wallet.v1.SubmitTransactionRequest.delegate_submission:type_name -> vega.commands.v1.DelegateSubmission
	9,  // 8: vega.wallet.v1.SubmitTransactionRequest.undelegate_submission:type_name -> vega.commands.v1.UndelegateSubmission
	10, // 9: vega.wallet.v1.SubmitTransactionRequest.liquidity_provision_cancellation:type_name -> vega.commands.v1.LiquidityProvisionCancellation
	11, // 10: vega.wallet.v1.SubmitTransactionRequest.liquidity_provision_amendment:type_name -> vega.commands.v1.LiquidityProvisionAmendment
	12, // 11: vega.wallet.v1.SubmitTransactionRequest.transfer:type_name -> vega.commands.v1.Transfer
	13, // 12: vega.wallet.v1.SubmitTransactionRequest.cancel_transfer:type_name -> vega.commands.v1.CancelTransfer
	14, // 13: vega.wallet.v1.SubmitTransactionRequest.announce_node:type_name -> vega.commands.v1.AnnounceNode
	15, // 14: vega.wallet.v1.SubmitTransactionRequest.node_vote:type_name -> vega.commands.v1.NodeVote
	16, // 15: vega.wallet.v1.SubmitTransactionRequest.node_signature:type_name -> vega.commands.v1.NodeSignature
	17, // 16: vega.wallet.v1.SubmitTransactionRequest.chain_event:type_name -> vega.commands.v1.ChainEvent
	18, // 17: vega.wallet.v1.SubmitTransactionRequest.key_rotate_submission:type_name -> vega.commands.v1.KeyRotateSubmission
	19, // 18: vega.wallet.v1.SubmitTransactionRequest.state_variable_proposal:type_name -> vega.commands.v1.StateVariableProposal
	20, // 19: vega.wallet.v1.SubmitTransactionRequest.validator_heartbeat:type_name -> vega.commands.v1.ValidatorHeartbeat
	21, // 20: vega.wallet.v1.SubmitTransactionRequest.ethereum_key_rotate_submission:type_name -> vega.commands.v1.EthereumKeyRotateSubmission
	22, // 21: vega.wallet.v1.SubmitTransactionRequest.protocol_upgrade_proposal:type_name -> vega.commands.v1.ProtocolUpgradeProposal
	23, // 22: vega.wallet.v1.SubmitTransactionRequest.issue_signatures:type_name -> vega.commands.v1.IssueSignatures
	24, // 23: vega.wallet.v1.SubmitTransactionRequest.oracle_data_submission:type_name -> vega.commands.v1.OracleDataSubmission
	24, // [24:24] is the sub-list for method output_type
	24, // [24:24] is the sub-list for method input_type
	24, // [24:24] is the sub-list for extension type_name
	24, // [24:24] is the sub-list for extension extendee
	0,  // [0:24] is the sub-list for field type_name
}

func init() { file_vega_wallet_v1_wallet_proto_init() }
func file_vega_wallet_v1_wallet_proto_init() {
	if File_vega_wallet_v1_wallet_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_vega_wallet_v1_wallet_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*SubmitTransactionRequest); i {
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
	file_vega_wallet_v1_wallet_proto_msgTypes[0].OneofWrappers = []interface{}{
		(*SubmitTransactionRequest_OrderSubmission)(nil),
		(*SubmitTransactionRequest_OrderCancellation)(nil),
		(*SubmitTransactionRequest_OrderAmendment)(nil),
		(*SubmitTransactionRequest_WithdrawSubmission)(nil),
		(*SubmitTransactionRequest_ProposalSubmission)(nil),
		(*SubmitTransactionRequest_VoteSubmission)(nil),
		(*SubmitTransactionRequest_LiquidityProvisionSubmission)(nil),
		(*SubmitTransactionRequest_DelegateSubmission)(nil),
		(*SubmitTransactionRequest_UndelegateSubmission)(nil),
		(*SubmitTransactionRequest_LiquidityProvisionCancellation)(nil),
		(*SubmitTransactionRequest_LiquidityProvisionAmendment)(nil),
		(*SubmitTransactionRequest_Transfer)(nil),
		(*SubmitTransactionRequest_CancelTransfer)(nil),
		(*SubmitTransactionRequest_AnnounceNode)(nil),
		(*SubmitTransactionRequest_NodeVote)(nil),
		(*SubmitTransactionRequest_NodeSignature)(nil),
		(*SubmitTransactionRequest_ChainEvent)(nil),
		(*SubmitTransactionRequest_KeyRotateSubmission)(nil),
		(*SubmitTransactionRequest_StateVariableProposal)(nil),
		(*SubmitTransactionRequest_ValidatorHeartbeat)(nil),
		(*SubmitTransactionRequest_EthereumKeyRotateSubmission)(nil),
		(*SubmitTransactionRequest_ProtocolUpgradeProposal)(nil),
		(*SubmitTransactionRequest_IssueSignatures)(nil),
		(*SubmitTransactionRequest_OracleDataSubmission)(nil),
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_vega_wallet_v1_wallet_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   1,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_vega_wallet_v1_wallet_proto_goTypes,
		DependencyIndexes: file_vega_wallet_v1_wallet_proto_depIdxs,
		MessageInfos:      file_vega_wallet_v1_wallet_proto_msgTypes,
	}.Build()
	File_vega_wallet_v1_wallet_proto = out.File
	file_vega_wallet_v1_wallet_proto_rawDesc = nil
	file_vega_wallet_v1_wallet_proto_goTypes = nil
	file_vega_wallet_v1_wallet_proto_depIdxs = nil
}
