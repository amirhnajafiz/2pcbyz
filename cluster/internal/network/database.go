package network

import (
	"context"

	"github.com/F24-CSE535/2pcbyz/cluster/pkg/rpc/database"
)

// Prepare accepts a transaction parameters for a cross-shard transaction.
func Prepare(address string, in *database.PrepareMsg) error {
	// base connection
	conn, err := connect(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	// call Prepare RPC
	if _, err = database.NewDatabaseClient(conn).Prepare(context.Background(), in); err != nil {
		return err
	}

	return nil
}

// Commit accepts a target and sessionId to send a commit message.
func Commit(address, returnAddress string, sessionId int) error {
	// base connection
	conn, err := connect(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	// call Commit RPC
	if _, err = database.NewDatabaseClient(conn).Commit(context.Background(), &database.CommitMsg{
		SessionId:     int64(sessionId), // set the session id
		ReturnAddress: returnAddress,    // set the return address
	}); err != nil {
		return err
	}

	return nil
}

// Abort accepts a target and sessionId to send an abort message.
func Abort(address, returnAddress string, sessionId int) error {
	// base connection
	conn, err := connect(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	// call Abort RPC
	if _, err = database.NewDatabaseClient(conn).Abort(context.Background(), &database.AbortMsg{
		SessionId:     int64(sessionId), // set the session id
		ReturnAddress: returnAddress,    // set the return address
	}); err != nil {
		return err
	}

	return nil
}
