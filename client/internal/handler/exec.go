package handler

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/F24-CSE535/2pcbyz/client/internal/network"
	"github.com/F24-CSE535/2pcbyz/client/pkg/rpc/database"
)

// request accepts all parameters for a transaction and calls a Request RPC.
func (h *Handler) request(argc int, argv []string) error {
	if argc < 3 {
		return errors.New("not enough arguments (sender receiver amount)")
	}

	// parse input arguments
	sender := argv[0]
	receiver := argv[1]
	amount, _ := strconv.Atoi(argv[2])

	// set a sessionId for the transaction
	session := h.session
	h.session++

	// find the clinet shard
	shard := findClientShard(sender, h.cfg.Shards)

	// call request RPC
	if err := network.Request(h.ipt.Services[h.ipt.Endpoints[shard]], &database.RequestMsg{
		Transaction: &database.TransactionMsg{ // create a new transaction
			Sender:    sender,
			Receiver:  receiver,
			Amount:    int64(amount),
			SessionId: int64(session),
		},
		ReturnAddress: localAddress(h.cfg.Port), // set the return address
	}); err != nil {
		return fmt.Errorf("rpc failed: %v", err)
	}

	return nil
}
