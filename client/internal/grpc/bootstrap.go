package server

import (
	"fmt"
	"log"
	"net"

	"github.com/F24-CSE535/2pcbyz/client/pkg/rpc/database"

	"google.golang.org/grpc"
)

// StartNewServer accepts a port and channel to get and forward packets to the processor.
func StartNewServer(port int) error {
	// on the local network, listen to a port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to start the listener server: %v", err)
	}

	// create a new gRPC instance
	server := grpc.NewServer()

	// register all gRPC services
	database.RegisterDatabaseServer(server, &Server{})

	// start gRPC server
	log.Printf("grpc server started on %d ...\n", port)
	if err := server.Serve(listener); err != nil {
		return fmt.Errorf("failed to start the server: %v", err)
	}

	return nil
}
