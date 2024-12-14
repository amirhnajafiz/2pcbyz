package dialer

import (
	"context"
	"fmt"

	"github.com/F24-CSE535/2pcbyz/client/pkg/rpc/database"
)

// Request dialer calls the Request RPC on a given target and passes the input request.
func Request(address string, in *database.RequestMsg) error {
	// base connection
	conn, err := connect(address)
	if err != nil {
		return err
	}
	defer conn.Close()

	// call Request RPC
	if _, err := database.NewDatabaseClient(conn).Request(context.Background(), in); err != nil {
		return fmt.Errorf("failed to call request rpc: %v", err)
	}

	return nil
}
