package network

import (
	"context"
	"fmt"

	"github.com/F24-CSE535/2pcbyz/client/pkg/rpc/database"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// connect should be called in the beginning of each method to establish a connection.
func connect(address string) (*grpc.ClientConn, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to open connection to %s: %v", address, err)
	}

	return conn, nil
}

// Request calls a Request RPC on the given address.
func Request(address string, in *database.RequestMsg) error {
	// base connection
	conn, err := connect(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	// call Request RPC
	if _, err = database.NewDatabaseClient(conn).Request(context.Background(), in); err != nil {
		return fmt.Errorf("failed to call Request rpc: %v", err)
	}

	return nil
}
