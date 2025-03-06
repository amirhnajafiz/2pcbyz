package network

import (
	"context"
	"fmt"
	"io"

	"github.com/F24-CSE535/2pcbyz/client/pkg/rpc/api"
	"github.com/F24-CSE535/2pcbyz/client/pkg/rpc/database"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
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

// PrintBalance accepts an address and client to get the client balance on the target address.
func PrintBalance(address, client string) (int, error) {
	// base connection
	conn, err := connect(address)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	// call PrintBalance RPC
	if resp, err := api.NewAPIClient(conn).PrintBalance(context.Background(), &api.PrintBalanceMsg{
		Client: client,
	}); err != nil {
		return 0, fmt.Errorf("failed to call PrintBalance rpc: %v", err)
	} else {
		return int(resp.GetBalance()), nil
	}
}

// PrintDatastore accepts an address and gets the datastore of the target.
func PrintDatastore(address string) ([]string, error) {
	// base connection
	conn, err := connect(address)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// open a stream on PrintDatastore to get blocks
	stream, err := api.NewAPIClient(conn).PrintDatastore(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to process PrintDatastore: %v", err)
	}

	// create a list to store datastore items
	list := make([]string, 0)

	for {
		// read items one by one
		in, err := stream.Recv()
		if err != nil {
			if err == io.EOF { // send a response once the stream is closed
				return list, nil
			}

			return nil, fmt.Errorf("failed to receive item: %v", err)
		}

		// append to the list of items
		list = append(list, fmt.Sprintf("%d (%d, %s, %s, %d)", in.GetSessionId(), in.GetSequence(), in.GetSender(), in.GetReceiver(), in.GetAmount()))
	}
}
