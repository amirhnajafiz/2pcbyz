package server

import (
	"context"
	"fmt"

	"github.com/F24-CSE535/2pcbyz/cluster/internal/storage"
	"github.com/F24-CSE535/2pcbyz/cluster/pkg/rpc/api"
	"google.golang.org/protobuf/types/known/emptypb"
)

type API struct {
	api.UnimplementedAPIServer

	storage *storage.Storage
}

// PrintBalance accepts a printbalance message and returns a printbalance response.
func (a *API) PrintBalance(_ context.Context, msg *api.PrintBalanceMsg) (*api.PrintBalanceRsp, error) {
	balance, err := a.storage.GetClientBalance(msg.GetClient())
	if err != nil {
		return nil, fmt.Errorf("database failed: %v", err)
	}

	return &api.PrintBalanceRsp{
		Balance: int64(balance),
	}, nil
}

// PrintDatastore returns all committed transactions inside this node.
func (a *API) PrintDatastore(_ *emptypb.Empty, stream api.API_PrintDatastoreServer) error {
	trxs, err := a.storage.GetCommittedTransactions()
	if err != nil {
		return fmt.Errorf("database failed: %v", err)
	}

	// send datastore one by one
	for _, trx := range trxs {
		if err := stream.Send(&api.DatastoreRsp{
			Sender:    trx.Sender,
			Receiver:  trx.Receiver,
			Amount:    int64(trx.Amount),
			SessionId: int64(trx.SessionId),
		}); err != nil {
			return err
		}
	}

	return nil
}
