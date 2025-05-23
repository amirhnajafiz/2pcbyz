// Code generated by protoc-gen-go-grpc. DO NOT EDIT.
// versions:
// - protoc-gen-go-grpc v1.5.1
// - protoc             v5.28.2
// source: database.proto

package database

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
	Database_Request_FullMethodName = "/database.Database/Request"
	Database_Reply_FullMethodName   = "/database.Database/Reply"
	Database_Prepare_FullMethodName = "/database.Database/Prepare"
	Database_Ack_FullMethodName     = "/database.Database/Ack"
	Database_Commit_FullMethodName  = "/database.Database/Commit"
	Database_Abort_FullMethodName   = "/database.Database/Abort"
)

// DatabaseClient is the client API for Database service.
//
// For semantics around ctx use and closing/ending streaming RPCs, please refer to https://pkg.go.dev/google.golang.org/grpc/?tab=doc#ClientConn.NewStream.
//
// creating rpc services for database management system.
type DatabaseClient interface {
	Request(ctx context.Context, in *RequestMsg, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Reply(ctx context.Context, in *ReplyMsg, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Prepare(ctx context.Context, in *RequestMsg, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Ack(ctx context.Context, in *AckMsg, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Commit(ctx context.Context, in *CommitMsg, opts ...grpc.CallOption) (*emptypb.Empty, error)
	Abort(ctx context.Context, in *AbortMsg, opts ...grpc.CallOption) (*emptypb.Empty, error)
}

type databaseClient struct {
	cc grpc.ClientConnInterface
}

func NewDatabaseClient(cc grpc.ClientConnInterface) DatabaseClient {
	return &databaseClient{cc}
}

func (c *databaseClient) Request(ctx context.Context, in *RequestMsg, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Database_Request_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Reply(ctx context.Context, in *ReplyMsg, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Database_Reply_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Prepare(ctx context.Context, in *RequestMsg, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Database_Prepare_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Ack(ctx context.Context, in *AckMsg, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Database_Ack_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Commit(ctx context.Context, in *CommitMsg, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Database_Commit_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

func (c *databaseClient) Abort(ctx context.Context, in *AbortMsg, opts ...grpc.CallOption) (*emptypb.Empty, error) {
	cOpts := append([]grpc.CallOption{grpc.StaticMethod()}, opts...)
	out := new(emptypb.Empty)
	err := c.cc.Invoke(ctx, Database_Abort_FullMethodName, in, out, cOpts...)
	if err != nil {
		return nil, err
	}
	return out, nil
}

// DatabaseServer is the server API for Database service.
// All implementations must embed UnimplementedDatabaseServer
// for forward compatibility.
//
// creating rpc services for database management system.
type DatabaseServer interface {
	Request(context.Context, *RequestMsg) (*emptypb.Empty, error)
	Reply(context.Context, *ReplyMsg) (*emptypb.Empty, error)
	Prepare(context.Context, *RequestMsg) (*emptypb.Empty, error)
	Ack(context.Context, *AckMsg) (*emptypb.Empty, error)
	Commit(context.Context, *CommitMsg) (*emptypb.Empty, error)
	Abort(context.Context, *AbortMsg) (*emptypb.Empty, error)
	mustEmbedUnimplementedDatabaseServer()
}

// UnimplementedDatabaseServer must be embedded to have
// forward compatible implementations.
//
// NOTE: this should be embedded by value instead of pointer to avoid a nil
// pointer dereference when methods are called.
type UnimplementedDatabaseServer struct{}

func (UnimplementedDatabaseServer) Request(context.Context, *RequestMsg) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Request not implemented")
}
func (UnimplementedDatabaseServer) Reply(context.Context, *ReplyMsg) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Reply not implemented")
}
func (UnimplementedDatabaseServer) Prepare(context.Context, *RequestMsg) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Prepare not implemented")
}
func (UnimplementedDatabaseServer) Ack(context.Context, *AckMsg) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Ack not implemented")
}
func (UnimplementedDatabaseServer) Commit(context.Context, *CommitMsg) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Commit not implemented")
}
func (UnimplementedDatabaseServer) Abort(context.Context, *AbortMsg) (*emptypb.Empty, error) {
	return nil, status.Errorf(codes.Unimplemented, "method Abort not implemented")
}
func (UnimplementedDatabaseServer) mustEmbedUnimplementedDatabaseServer() {}
func (UnimplementedDatabaseServer) testEmbeddedByValue()                  {}

// UnsafeDatabaseServer may be embedded to opt out of forward compatibility for this service.
// Use of this interface is not recommended, as added methods to DatabaseServer will
// result in compilation errors.
type UnsafeDatabaseServer interface {
	mustEmbedUnimplementedDatabaseServer()
}

func RegisterDatabaseServer(s grpc.ServiceRegistrar, srv DatabaseServer) {
	// If the following call pancis, it indicates UnimplementedDatabaseServer was
	// embedded by pointer and is nil.  This will cause panics if an
	// unimplemented method is ever invoked, so we test this at initialization
	// time to prevent it from happening at runtime later due to I/O.
	if t, ok := srv.(interface{ testEmbeddedByValue() }); ok {
		t.testEmbeddedByValue()
	}
	s.RegisterService(&Database_ServiceDesc, srv)
}

func _Database_Request_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Request(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_Request_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Request(ctx, req.(*RequestMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Reply_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(ReplyMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Reply(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_Reply_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Reply(ctx, req.(*ReplyMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Prepare_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(RequestMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Prepare(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_Prepare_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Prepare(ctx, req.(*RequestMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Ack_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AckMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Ack(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_Ack_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Ack(ctx, req.(*AckMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Commit_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(CommitMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Commit(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_Commit_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Commit(ctx, req.(*CommitMsg))
	}
	return interceptor(ctx, in, info, handler)
}

func _Database_Abort_Handler(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
	in := new(AbortMsg)
	if err := dec(in); err != nil {
		return nil, err
	}
	if interceptor == nil {
		return srv.(DatabaseServer).Abort(ctx, in)
	}
	info := &grpc.UnaryServerInfo{
		Server:     srv,
		FullMethod: Database_Abort_FullMethodName,
	}
	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		return srv.(DatabaseServer).Abort(ctx, req.(*AbortMsg))
	}
	return interceptor(ctx, in, info, handler)
}

// Database_ServiceDesc is the grpc.ServiceDesc for Database service.
// It's only intended for direct use with grpc.RegisterService,
// and not to be introspected or modified (even as a copy)
var Database_ServiceDesc = grpc.ServiceDesc{
	ServiceName: "database.Database",
	HandlerType: (*DatabaseServer)(nil),
	Methods: []grpc.MethodDesc{
		{
			MethodName: "Request",
			Handler:    _Database_Request_Handler,
		},
		{
			MethodName: "Reply",
			Handler:    _Database_Reply_Handler,
		},
		{
			MethodName: "Prepare",
			Handler:    _Database_Prepare_Handler,
		},
		{
			MethodName: "Ack",
			Handler:    _Database_Ack_Handler,
		},
		{
			MethodName: "Commit",
			Handler:    _Database_Commit_Handler,
		},
		{
			MethodName: "Abort",
			Handler:    _Database_Abort_Handler,
		},
	},
	Streams:  []grpc.StreamDesc{},
	Metadata: "database.proto",
}
