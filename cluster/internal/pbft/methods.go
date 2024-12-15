package pbft

import (
	"context"
	"fmt"
	"strings"

	"github.com/F24-CSE535/2pcbyz/cluster/internal/network"
	"github.com/F24-CSE535/2pcbyz/cluster/pkg/rpc/database"
	"github.com/F24-CSE535/2pcbyz/cluster/pkg/rpc/pbft"

	"go.uber.org/zap"
)

func (sm *StateMachine) request(payload interface{}) error {
	// get transaction
	trx := payload.(*database.RequestMsg)

	// set sequence number
	trx.GetTransaction().Sequence = int64(sm.sequence)
	sm.sequence++

	// call preprepare on all
	for _, svc := range strings.Split(sm.Ipt.Endpoints[fmt.Sprintf("E%s", sm.Cluster)], ":") {
		if err := network.PrePrepare(sm.Ipt.Services[svc], &pbft.PrePrepareMsg{
			Transaction: &pbft.TransactionMsg{
				Sender:        trx.GetTransaction().GetSender(),
				Receiver:      trx.GetTransaction().GetReceiver(),
				Amount:        trx.GetTransaction().GetAmount(),
				ReturnAddress: trx.GetReturnAddress(),
				SessionId:     trx.GetTransaction().GetSessionId(),
				Sequence:      trx.GetTransaction().GetSequence(),
			},
		}); err != nil {
			sm.Logger.Warn("failed to send preprepare", zap.Error(err))
		}
	}

	return nil
}

func (sm *StateMachine) prePrepare(payload interface{}) error {
	// get preprepare message
	msg := payload.(*pbft.PrePrepareMsg)

	// convert it to a transaction
	trx := &database.TransactionMsg{
		Sender:    msg.GetTransaction().GetSender(),
		Receiver:  msg.GetTransaction().GetReceiver(),
		Amount:    msg.GetTransaction().GetAmount(),
		SessionId: msg.GetTransaction().GetSessionId(),
		Sequence:  msg.GetTransaction().GetSequence(),
	}

	// create a request message
	req := &database.RequestMsg{
		Transaction:   trx,
		ReturnAddress: msg.Transaction.GetReturnAddress(),
	}

	if sm.block || sm.byzantine {
		return fmt.Errorf("node cannot process this message: block %b, byzantine %b", sm.block, sm.byzantine)
	}

	// send to handler
	sm.Queue <- context.WithValue(context.WithValue(context.Background(), "method", "begin"), "request", req)

	return nil
}

func (sm *StateMachine) ackPrePrepare(payload interface{}) error {
	return nil
}

func (sm *StateMachine) prepare(payload interface{}) error {
	return nil
}

func (sm *StateMachine) ackPrepare(payload interface{}) error {
	return nil
}

func (sm *StateMachine) commit(payload interface{}) error {
	return nil
}
