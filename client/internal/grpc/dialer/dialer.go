package grpc

import (
	"context"
	"fmt"
	"io"

	"github.com/F24-CSE535/2pc/client/pkg/rpc/database"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

// Dialer is a module for making RPC calls from client to clusters.
type Dialer struct {
	Nodes    map[string]string
	contacts map[string]string
}

// connect should be called in the beginning of each method to establish a connection.
func (d *Dialer) connect(target string) (*grpc.ClientConn, error) {
	address := d.Nodes[target]

	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("failed to open connection to %s: %v", address, err)
	}

	return conn, nil
}

// SetContacts updates the dialer contacts.
func (d *Dialer) SetContacts(c map[string]string) {
	d.contacts = c
}

// Request accepts a transaction parameters for an inter-shard transaction.
func (d *Dialer) Request(target, sender, receiver string, amount, sessionId int) error {
	// base connection
	conn, err := d.connect(d.contacts[target])
	if err != nil {
		return err
	}
	defer conn.Close()

	// call Request RPC
	if _, err = database.NewDatabaseClient(conn).Request(context.Background(), &database.RequestMsg{
		Transaction: &database.TransactionMsg{ // initialize a new transaction
			Sender:    sender,
			Receiver:  receiver,
			Amount:    int64(amount),
			SessionId: int64(sessionId),
		},
		ReturnAddress: d.Nodes["client"], // set the return address
	}); err != nil {
		return err
	}

	return nil
}

// PrintBalance accepts a target and client to return the client balance.
func (d *Dialer) PrintBalance(target string, client string) (int, error) {
	// base connection
	conn, err := d.connect(target)
	if err != nil {
		return 0, err
	}
	defer conn.Close()

	// call PrintBalance RPC
	resp, err := database.NewDatabaseClient(conn).PrintBalance(context.Background(), &database.PrintBalanceMsg{
		Client: client,
	})
	if err != nil {
		return 0, err
	}

	return int(resp.GetBalance()), nil
}

// PrintDatastore accepts a target and calls PrintDatastore RPC on the target.
func (d *Dialer) PrintDatastore(target string) ([]*database.DatastoreRsp, error) {
	// base connection
	conn, err := d.connect(target)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	// open a stream on PrintDatastore to get blocks
	stream, err := database.NewDatabaseClient(conn).PrintDatastore(context.Background(), &emptypb.Empty{})
	if err != nil {
		return nil, fmt.Errorf("failed to process printdatastore: %v", err)
	}

	// create a list to store datastore
	list := make([]*database.DatastoreRsp, 0)

	for {
		// read logs one by one
		in, err := stream.Recv()
		if err != nil {
			if err == io.EOF { // send a response once the stream is closed
				return list, nil
			}

			return nil, fmt.Errorf("failed to receive datastore: %v", err)
		}

		// append to the list of blocks
		list = append(list, in)
	}
}

// Block accepts a target and blocks it.
func (d *Dialer) Block(target string) error {
	// base connection
	conn, err := d.connect(target)
	if err != nil {
		return err
	}
	defer conn.Close()

	// call Block RPC
	if _, err = database.NewDatabaseClient(conn).Block(context.Background(), &emptypb.Empty{}); err != nil {
		return err
	}

	return nil
}

// Unblock accepts a target and blocks it.
func (d *Dialer) Unblock(target string) error {
	// base connection
	conn, err := d.connect(target)
	if err != nil {
		return err
	}
	defer conn.Close()

	// call Unblock RPC
	if _, err = database.NewDatabaseClient(conn).Unblock(context.Background(), &emptypb.Empty{}); err != nil {
		return err
	}

	return nil
}
