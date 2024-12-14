package server

import (
	"context"

	"github.com/F24-CSE535/2pcbyz/client/pkg/rpc/database"

	"google.golang.org/protobuf/types/known/emptypb"
)

// Server is a partial gRPC server in client side.
type Server struct {
	database.UnimplementedDatabaseServer
}

// Reply accepts all reply messages.
func (s *Server) Reply(_ context.Context, msg *database.ReplyMsg) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
