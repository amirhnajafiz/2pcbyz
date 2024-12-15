package network

import (
	"context"
	"fmt"

	"github.com/F24-CSE535/2pcbyz/client/pkg/rpc/api"

	"google.golang.org/protobuf/types/known/emptypb"
)

// Block calls a Block RPC on the given address.
func Block(address string) error {
	// base connection
	conn, err := connect(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	// call Block RPC
	if _, err = api.NewAPIClient(conn).Block(context.Background(), &emptypb.Empty{}); err != nil {
		return fmt.Errorf("failed to call Block rpc: %v", err)
	}

	return nil
}

// Unblock calls a Unblock RPC on the given address.
func Unblock(address string) error {
	// base connection
	conn, err := connect(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	// call Unblock RPC
	if _, err = api.NewAPIClient(conn).Unblock(context.Background(), &emptypb.Empty{}); err != nil {
		return fmt.Errorf("failed to call Unblock rpc: %v", err)
	}

	return nil
}

// Byzantine calls a Byzantine RPC on the given address.
func Byzantine(address string) error {
	// base connection
	conn, err := connect(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	// call Byzantine RPC
	if _, err = api.NewAPIClient(conn).Byzantine(context.Background(), &emptypb.Empty{}); err != nil {
		return fmt.Errorf("failed to call Byzantine rpc: %v", err)
	}

	return nil
}

// NonByzantine calls a NonByzantine RPC on the given address.
func NonByzantine(address string) error {
	// base connection
	conn, err := connect(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	// call NonByzantine RPC
	if _, err = api.NewAPIClient(conn).NonByzantine(context.Background(), &emptypb.Empty{}); err != nil {
		return fmt.Errorf("failed to call NonByzantine rpc: %v", err)
	}

	return nil
}
