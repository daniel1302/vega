// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.2.0
// - protoc             (unknown)
// source: blockexplorer/blockexplorer.proto

package v1

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.32.0 or later.
const _ = grpc.SupportPackageIsVersion7

// BlockExplorerServiceClient is the client API for BlockExplorerService service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
type BlockExplorerServiceClient interface {
	// Get transaction
	//
	// Get a transaction from the Vega blockchain
	GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error)
	// List transactions
	//
	// List transactions from the Vega blockchain
	ListTransactions(ctx context.Context, in *ListTransactionsRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error)
	// Info
	//
	// Retrieves information about the block explorer.
	// Response contains a semver formatted version of the data node and the commit hash, from which the block explorer was built,
	Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error)
}

type blockExplorerServiceClient struct {
	cc grpc.ClientConnInterface
}

func NewBlockExplorerServiceClient(cc grpc.ClientConnInterface) BlockExplorerServiceClient {
	return &blockExplorerServiceClient{cc}
}

func (c *blockExplorerServiceClient) GetTransaction(ctx context.Context, in *GetTransactionRequest, opts ...grpc.CallOption) (*GetTransactionResponse, error) {
	out := new(GetTransactionResponse)
	err := c.cc.Invoke(ctx, "/blockexplorer.api.v1.BlockExplorerService/GetTransaction", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockExplorerServiceClient) ListTransactions(ctx context.Context, in *ListTransactionsRequest, opts ...grpc.CallOption) (*ListTransactionsResponse, error) {
	out := new(ListTransactionsResponse)
	err := c.cc.Invoke(ctx, "/blockexplorer.api.v1.BlockExplorerService/ListTransactions", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *blockExplorerServiceClient) Info(ctx context.Context, in *InfoRequest, opts ...grpc.CallOption) (*InfoResponse, error) {
	out := new(InfoResponse)
	err := c.cc.Invoke(ctx, "/blockexplorer.api.v1.BlockExplorerService/Info", in, out, opts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// BlockExplorerServiceServer is the server API for BlockExplorerService service.
// All implementations must embed UnimplementedBlockExplorerServiceServer
// for forward compatibility
type BlockExplorerServiceServer interface {
	// Get transaction
	//
	// Get a transaction from the Vega blockchain
	GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error)
	// List transactions
	//
	// List transactions from the Vega blockchain
	ListTransactions(context.Context, *ListTransactionsRequest) (*ListTransactionsResponse, error)
	// Info
	//
	// Retrieves information about the block explorer.
	// Response contains a semver formatted version of the data node and the commit hash, from which the block explorer was built,
	Info(context.Context, *InfoRequest) (*InfoResponse, error)
	mustEmbedUnimplementedBlockExplorerServiceServer()
}

// UnimplementedBlockExplorerServiceServer must be embedded to have forward compatible implementations.
type UnimplementedBlockExplorerServiceServer struct {
}

func (UnimplementedBlockExplorerServiceServer) GetTransaction(context.Context, *GetTransactionRequest) (*GetTransactionResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method GetTransaction not implemented")
}
func (UnimplementedBlockExplorerServiceServer) ListTransactions(context.Context, *ListTransactionsRequest) (*ListTransactionsResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method ListTransactions not implemented")
}
func (UnimplementedBlockExplorerServiceServer) Info(context.Context, *InfoRequest) (*InfoResponse, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Info not implemented")
}
func (UnimplementedBlockExplorerServiceServer) mustEmbedUnimplementedBlockExplorerServiceServer() {}

// UnsafeBlockExplorerServiceServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to BlockExplorerServiceServer will
// result in compilation errors.
type UnsafeBlockExplorerServiceServer interface {
	mustEmbedUnimplementedBlockExplorerServiceServer()
}

func RegisterBlockExplorerServiceServer(s grpc.ServiceRegistrar, srv BlockExplorerServiceServer) {
	s.RegisterService(&BlockExplorerService_ServiceDesc, srv)
}

func _BlockExplorerService_GetTransaction_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(GetTransactionRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockExplorerServiceServer).GetTransaction(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockexplorer.api.v1.BlockExplorerService/GetTransaction",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockExplorerServiceServer).GetTransaction(ctx, req.(*GetTransactionRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockExplorerService_ListTransactions_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ListTransactionsRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockExplorerServiceServer).ListTransactions(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockexplorer.api.v1.BlockExplorerService/ListTransactions",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockExplorerServiceServer).ListTransactions(ctx, req.(*ListTransactionsRequest))
	}
	return interceptor(ctx, in, info, handler)
}

func _BlockExplorerService_Info_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(InfoRequest)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(BlockExplorerServiceServer).Info(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: "/blockexplorer.api.v1.BlockExplorerService/Info",
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(BlockExplorerServiceServer).Info(ctx, req.(*InfoRequest))
	}
	return interceptor(ctx, in, info, handler)
}

// BlockExplorerService_ServiceDesc is the grpc.ServiceDesc for BlockExplorerService service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var BlockExplorerService_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "blockexplorer.api.v1.BlockExplorerService",
	HandlerType: (*BlockExplorerServiceServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "GetTransaction",
			Handler:    _BlockExplorerService_GetTransaction_Handler,
		},
		{
			MethodName: "ListTransactions",
			Handler:    _BlockExplorerService_ListTransactions_Handler,
		},
		{
			MethodName: "Info",
			Handler:    _BlockExplorerService_Info_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "blockexplorer/blockexplorer.proto",
}
