// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: api.proto

package api

import (
	context "context"
	grpc "google.golang.org/grpc"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the grpc package it is being compiled against.
// Requires gRPC-Go v1.64.0 or later.
const _ = grpc.SupportPackageIsVersion9

const (
	API_PrintBalance_FullMethodName   = "/api.API/PrintBalance"
	API_PrintDatastore_FullMethodName = "/api.API/PrintDatastore"
	API_Block_FullMethodName          = "/api.API/Block"
	API_Unblock_FullMethodName        = "/api.API/Unblock"
	API_Byzantine_FullMethodName      = "/api.API/Byzantine"
	API_NonByzantine_FullMethodName   = "/api.API/NonByzantine"
)

// APIClient is the client API for API service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// creating rpc services for api system.
type APIClient interface {
	PrintBalance(ctx context.Context, in *PrintBalanceMsg, opts ...grpc.CallOption) (*PrintBalanceRsp, error)
	PrintDatastore(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DatastoreRsp], error)
	Block(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Unblock(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Byzantine(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
	NonByzantine(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type aPIClient struct {
	cc grpc.ClientConnInterface
}

func NewAPIClient(cc grpc.ClientConnInterface) APIClient {
	return &aPIClient{cc}
}

func (c *aPIClient) PrintBalance(ctx context.Context, in *PrintBalanceMsg, opts ...grpc.CallOption) (*PrintBalanceRsp, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(PrintBalanceRsp)
	err := c.cc.Invoke(ctx, API_PrintBalance_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) PrintDatastore(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (grpc.ServerStreamingClient[DatastoreRsp], error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	stream, err := c.cc.NewStream(ctx, &API_ServiceDesc.Streams[0], API_PrintDatastore_FullMethodName, cOpts...)
	if err != nil {
		return nil, err
	}
	x := &grpc.GenericClientStream[emptypb.Empty, DatastoreRsp]{ClientStream: stream}
	if err := x.ClientStream.SendMsg(in); err != nil {
		return nil, err
	}
	if err := x.ClientStream.CloseSend(); err != nil {
		return nil, err
	}
	return x, nil
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type API_PrintDatastoreClient = grpc.ServerStreamingClient[DatastoreRsp]

func (c *aPIClient) Block(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, API_Block_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Unblock(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, API_Unblock_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) Byzantine(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, API_Byzantine_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *aPIClient) NonByzantine(ctx context.Context, in *emptypb.Empty, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, API_NonByzantine_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// APIServer is the server API for API service.
// All implementations must embed UnimplementedAPIServer
// for forward compatibility.
//
// creating rpc services for api system.
type APIServer interface {
	PrintBalance(context.Context, *PrintBalanceMsg) (*PrintBalanceRsp, error)
	PrintDatastore(*emptypb.Empty, grpc.ServerStreamingServer[DatastoreRsp]) error
	Block(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Unblock(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	Byzantine(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	NonByzantine(context.Context, *emptypb.Empty) (*emptypb.Empty, error)
	mustEmbedUnimplementedAPIServer()
}

// UnimplementedAPIServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedAPIServer struct{}

func (UnimplementedAPIServer) PrintBalance(context.Context, *PrintBalanceMsg) (*PrintBalanceRsp, error) {
	return nil, status.Errorf(codes.Unimplemented, "method PrintBalance not implemented")
}
func (UnimplementedAPIServer) PrintDatastore(*emptypb.Empty, grpc.ServerStreamingServer[DatastoreRsp]) error {
	return status.Errorf(codes.Unimplemented, "method PrintDatastore not implemented")
}
func (UnimplementedAPIServer) Block(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Block not implemented")
}
func (UnimplementedAPIServer) Unblock(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Unblock not implemented")
}
func (UnimplementedAPIServer) Byzantine(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Byzantine not implemented")
}
func (UnimplementedAPIServer) NonByzantine(context.Context, *emptypb.Empty) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method NonByzantine not implemented")
}
func (UnimplementedAPIServer) mustEmbedUnimplementedAPIServer() {}
func (UnimplementedAPIServer) testEmbeddedByValue()             {}

// UnsafeAPIServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to APIServer will
// result in compilation errors.
type UnsafeAPIServer interface {
	mustEmbedUnimplementedAPIServer()
}

func RegisterAPIServer(s grpc.ServiceRegistrar, srv APIServer) {
	// If the following call pancis, it indicates UnimplementedAPIServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&API_ServiceDesc, srv)
}

func _API_PrintBalance_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(PrintBalanceMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).PrintBalance(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: API_PrintBalance_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).PrintBalance(ctx, req.(*PrintBalanceMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_PrintDatastore_Handler(srv interface{}, stream grpc.ServerStream) error {
	m := new(emptypb.Empty)
	if err := stream.RecvMsg(m); err != nil {
		return err
	}
	return srv.(APIServer).PrintDatastore(m, &grpc.GenericServerStream[emptypb.Empty, DatastoreRsp]{ServerStream: stream})
}

// This type alias is provided for backwards compatibility with existing code that references the prior non-generic stream type by name.
type API_PrintDatastoreServer = grpc.ServerStreamingServer[DatastoreRsp]

func _API_Block_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Block(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: API_Block_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Block(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Unblock_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Unblock(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: API_Unblock_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Unblock(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_Byzantine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).Byzantine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: API_Byzantine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).Byzantine(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

func _API_NonByzantine_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(emptypb.Empty)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(APIServer).NonByzantine(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: API_NonByzantine_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(APIServer).NonByzantine(ctx, req.(*emptypb.Empty))
	}
	return interceptor(ctx, in, info, handler)
}

// API_ServiceDesc is the grpc.ServiceDesc for API service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var API_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "api.API",
	HandlerType: (*APIServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "PrintBalance",
			Handler:    _API_PrintBalance_Handler,
		},
		{
			MethodName: "Block",
			Handler:    _API_Block_Handler,
		},
		{
			MethodName: "Unblock",
			Handler:    _API_Unblock_Handler,
		},
		{
			MethodName: "Byzantine",
			Handler:    _API_Byzantine_Handler,
		},
		{
			MethodName: "NonByzantine",
			Handler:    _API_NonByzantine_Handler,
		},
	},
	Streams: []grpc.StreamDesc{
		{
			StreamName:    "PrintDatastore",
			Handler:       _API_PrintDatastore_Handler,
			ServerStreams: true,
		},
	},
	Metadata: "api.proto",
}
