package network

import (
	"context"
	"fmt"

	"github.com/F24-CSE535/2pcbyz/cluster/pkg/rpc/database"

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

// Reply calls the reply RPC on the client.
func Reply(address, raddress, paddress, text string, sessionId int) error {
	// base connection
	conn, err := connect(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	// call Reply RPC
	if _, err := database.NewDatabaseClient(conn).Reply(context.Background(), &database.ReplyMsg{
		SessionId:          int64(sessionId),
		Text:               text,
		ReturnAddress:      raddress,
		ParticipantAddress: paddress,
	}); err != nil {
		return fmt.Errorf("failed to call reply RPC: %v", err)
	}

	return nil
}
