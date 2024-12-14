package grpc

import (
	"context"
	"strings"

	"go.uber.org/zap"
	"google.golang.org/grpc"
)

// stream interceptor is used to print a log on each stream RPC.
func (b *Bootstrap) allStreamInterceptor(
	srv interface{},
	ss grpc.ServerStream,
	info *grpc.StreamServerInfo,
	handler grpc.StreamHandler,
) error {
	b.Logger.Info("stream rpc called", zap.String("method", info.FullMethod))
	return handler(srv, ss)
}

// allUnaryInterceptor interceptor checks the status of a service before running the gRPC function.
func (b *Bootstrap) allUnaryInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	b.Logger.Info("rpc called", zap.String("method", info.FullMethod))
	return b.checkEmptyReturnCallsInterceptor(ctx, req, info, handler)
}

// checkEmptyReturnCallsInterceptor accepts requests from the all unary interceptor
// and passes the database and pbft requests inside the handler queue.
func (b *Bootstrap) checkEmptyReturnCallsInterceptor(
	ctx context.Context,
	req interface{},
	info *grpc.UnaryServerInfo,
	handler grpc.UnaryHandler,
) (interface{}, error) {
	// parse to get service and method
	if svc, method, err := parseFullMethod(info.FullMethod); err == nil {
		// if the service is database or pbft, it is an empty return call
		if strings.HasPrefix(svc, "database") || strings.HasPrefix(svc, "pbft") {
			b.Queue <- context.WithValue(
				context.WithValue(
					context.Background(),
					contextKey("method"),
					method,
				),
				contextKey("request"),
				req,
			)

			return nil, nil
		}
	}

	return handler(ctx, req)
}
