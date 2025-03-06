package network

import (
	"context"
	"fmt"

	"github.com/F24-CSE535/2pcbyz/cluster/pkg/rpc/pbft"
)

// PrePrepare calls the preprepare RPC on a target.
func PrePrepare(address string, in *pbft.PrePrepareMsg) error {
	// base connection
	conn, err := connect(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	// call PrePrepare RPC
	if _, err := pbft.NewPBFTClient(conn).PrePrepare(context.Background(), in); err != nil {
		return fmt.Errorf("failed to call reply RPC: %v", err)
	}

	return nil
}
