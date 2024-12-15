package handler

import (
	"github.com/F24-CSE535/2pcbyz/cluster/internal/models"
	"github.com/F24-CSE535/2pcbyz/cluster/internal/network"
	"github.com/F24-CSE535/2pcbyz/cluster/pkg/rpc/database"

	"go.uber.org/zap"
	"golang.org/x/net/context"
)

func (h *Handler) begin(payload interface{}) {
	// get transaction
	trx := payload.(*database.RequestMsg)

	// get both shards
	sshard := findClientShard(trx.GetTransaction().GetSender(), h.Cfg.Shards)
	rshard := findClientShard(trx.GetTransaction().GetReceiver(), h.Cfg.Shards)

	// check transaction type
	var ctx context.Context
	if sshard == rshard {
		// call inter-shard
		ctx = context.WithValue(
			context.WithValue(
				context.Background(),
				"method",
				"intershard",
			),
			"request",
			trx,
		)
	} else {
		// call cross-shard
		ctx = context.WithValue(
			context.WithValue(
				context.Background(),
				"method",
				"crossshard",
			),
			"request",
			trx,
		)
	}

	// reconcile the context again
	h.Queue <- ctx
}

func (h *Handler) intershard(payload interface{}) {
	// get transaction
	trx := payload.(*database.RequestMsg)

	// get sessionId
	sessionId := int(trx.GetTransaction().GetSessionId())

	// insert locks
	if err := h.Storage.InsertLock(trx.GetTransaction().GetSender()); err != nil {
		h.Logger.Warn("failed to stores locks", zap.Int("session id", sessionId))
	}
	if err := h.Storage.InsertLock(trx.GetTransaction().GetReceiver()); err != nil {
		h.Logger.Warn("failed to stores locks", zap.Int("session id", sessionId))
	}

	// insert logs
	wals := make([]*models.Log, 0)
	wals = append(wals,
		&models.Log{
			SessionId: sessionId,
			Message:   models.WALStart,
		},
		&models.Log{
			SessionId: sessionId,
			Message:   models.WALUpdate,
			Record:    trx.GetTransaction().GetSender(),
			NewValue:  -1 * int(trx.GetTransaction().GetAmount()),
		},
		&models.Log{
			SessionId: sessionId,
			Message:   models.WALUpdate,
			Record:    trx.GetTransaction().GetReceiver(),
			NewValue:  int(trx.GetTransaction().GetAmount()),
		},
	)

	// store the logs
	if err := h.Storage.InsertBatchLogs(wals); err != nil {
		h.Logger.Warn("failed to store logs", zap.Error(err))
		return
	}

	// insert transaction
	if err := h.Storage.InsertTransaction(&models.Transaction{
		Sender:    trx.GetTransaction().GetSender(),
		Receiver:  trx.GetTransaction().GetReceiver(),
		Amount:    int(trx.GetTransaction().GetAmount()),
		SessionId: sessionId,
	}); err != nil {
		h.Logger.Warn("failed to store transaction", zap.Error(err))
	}

	// get the sender balance
	balance, err := h.Storage.GetClientBalance(trx.GetTransaction().GetSender())
	if err != nil {
		h.Logger.Warn("failed to get client balance", zap.Error(err))
		return
	}

	// call commit or abort
	var ctx context.Context
	if trx.GetTransaction().GetAmount() <= int64(balance) {
		ctx = context.WithValue(
			context.WithValue(
				context.Background(),
				"method",
				"commit",
			),
			"request",
			&database.CommitMsg{
				SessionId:     trx.GetTransaction().GetSessionId(),
				ReturnAddress: trx.GetReturnAddress(),
			},
		)
	} else {
		ctx = context.WithValue(
			context.WithValue(
				context.Background(),
				"method",
				"abort",
			),
			"request",
			&database.AbortMsg{
				SessionId:     trx.GetTransaction().GetSessionId(),
				ReturnAddress: trx.GetReturnAddress(),
			},
		)
	}

	// reconcile the context again
	h.Queue <- ctx
}

func (h *Handler) crossshard(payload interface{}) {

}

func (h *Handler) abort(payload interface{}) {
	// get transaction
	trx := payload.(*database.AbortMsg)

	// get sessionId
	sessionId := int(trx.GetSessionId())

	// insert log
	if err := h.Storage.InsertLog(&models.Log{
		SessionId: sessionId,
		Message:   models.WALAbort,
	}); err != nil {
		h.Logger.Warn("failed to store log", zap.Error(err))
	}

	// update transaction
	if err := h.Storage.UpdateTransactionStatus(sessionId, models.StsAbort); err != nil {
		h.Logger.Warn("failed to update transaction", zap.Error(err))
	}

	// reply back to the user
	if err := network.Reply(trx.GetReturnAddress(), "FAIL", sessionId); err != nil {
		h.Logger.Warn("failed to send reply", zap.Error(err))
	}
}

func (h *Handler) commit(payload interface{}) {
	// get transaction
	trx := payload.(*database.CommitMsg)

	// get sessionId
	sessionId := int(trx.GetSessionId())

	// insert log
	if err := h.Storage.InsertLog(&models.Log{
		SessionId: sessionId,
		Message:   models.WALCommit,
	}); err != nil {
		h.Logger.Warn("failed to store log", zap.Error(err))
	}

	// get all update logs
	wals, err := h.Storage.GetLogsBySessionId(sessionId)
	if err != nil {
		h.Logger.Warn("failed to get logs", zap.Error(err))
		return
	}

	// update clients
	for _, wal := range wals {
		// update the client balance
		if err := h.Storage.UpdateClientBalance(wal.Record, wal.NewValue, false); err != nil {
			h.Logger.Warn("failed to update balance", zap.Error(err), zap.String("client", wal.Record))
			return
		}
	}

	// update transaction
	if err := h.Storage.UpdateTransactionStatus(sessionId, models.StsCommit); err != nil {
		h.Logger.Warn("failed to update transaction", zap.Error(err))
	}

	// reply back to the user
	if err := network.Reply(trx.GetReturnAddress(), "OK", sessionId); err != nil {
		h.Logger.Warn("failed to send reply", zap.Error(err))
	}
}
