package server

import (
	"context"

	"github.com/F24-CSE535/2pcbyz/client/pkg/rpc/database"

	"google.golang.org/protobuf/types/known/emptypb"
)

// serve4 is a partial gRPC service in client side.
type server struct {
	database.UnimplementedDatabaseServer
}

// Reply accepts all reply messages.
func (s *server) Reply(_ context.Context, msg *database.ReplyMsg) (*emptypb.Empty, error) {
	return &emptypb.Empty{}, nil
}
