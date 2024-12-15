package server

import (
	"context"

	"github.com/F24-CSE535/2pcbyz/client/pkg/rpc/database"

	"google.golang.org/protobuf/types/known/emptypb"
)

// server is a partial gRPC service in client side.
type server struct {
	database.UnimplementedDatabaseServer

	limit  int
	memory map[int]int

	output chan string
}

// Reply accepts all reply messages from system nodes.
func (s *server) Reply(_ context.Context, msg *database.ReplyMsg) (*emptypb.Empty, error) {
	// get the message sessionId
	sid := int(msg.GetSessionId())

	// add the response message to the memory
	if _, ok := s.memory[sid]; !ok {
		s.memory[sid] = 1
	} else {
		s.memory[sid]++
	}

	// check the limit, if there are enough responses, return the response to user
	if s.memory[sid] >= s.limit {
		s.output <- msg.GetText()
		delete(s.memory, sid)
	}

	return &emptypb.Empty{}, nil
}
