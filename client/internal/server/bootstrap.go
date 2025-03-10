package server

import (
	"fmt"
	"log"
	"net"
	"sync"

	"github.com/F24-CSE535/2pcbyz/client/pkg/rpc/database"

	"google.golang.org/grpc"
)

// ListenAndServe accepts a port and starts a gRPC server.
func ListenAndServe(port, limit int, output chan string) error {
	// on the local network, listen to a port
	listener, err := net.Listen("tcp", fmt.Sprintf(":%d", port))
	if err != nil {
		return fmt.Errorf("failed to start the listener server: %v", err)
	}

	// create a new gRPC instance
	srv := grpc.NewServer()

	// register all gRPC services
	database.RegisterDatabaseServer(srv, &server{
		lock:   sync.Mutex{},
		limit:  limit,
		output: output,
		memory: make(map[int]int),
	})

	// start gRPC server
	log.Printf("grpc server started on %d ...\n", port)
	if err := srv.Serve(listener); err != nil {
		return fmt.Errorf("failed to start the server: %v", err)
	}

	return nil
}
