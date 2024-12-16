package handler

import (
	"fmt"
	"strings"
	"time"

	"github.com/F24-CSE535/2pcbyz/cluster/internal/models"
	"github.com/F24-CSE535/2pcbyz/cluster/internal/network"
	"github.com/F24-CSE535/2pcbyz/cluster/pkg/rpc/database"

	"go.uber.org/zap"
	"golang.org/x/net/context"
)

func (h *Handler) begin(payload interface{}) {
	// get transaction
	trx := payload.(*database.RequestMsg)

	h.Logger.Debug("session id is here", zap.Int64("id", trx.GetTransaction().GetSessionId()))

	// insert transaction
	if err := h.Storage.InsertTransaction(&models.Transaction{
		Sender:    trx.GetTransaction().GetSender(),
		Receiver:  trx.GetTransaction().GetReceiver(),
		Amount:    int(trx.GetTransaction().GetAmount()),
		SessionId: int(trx.GetTransaction().GetSessionId()),
		Sequence:  int(trx.GetTransaction().GetSequence()),
	}); err != nil {
		h.Logger.Warn("failed to store transaction", zap.Error(err))
	}

	// get both shards
	sshard := findClientShard(trx.GetTransaction().GetSender(), h.Cfg.Shards)
	rshard := findClientShard(trx.GetTransaction().GetReceiver(), h.Cfg.Shards)

	// check transaction type
	if sshard == rshard {
		// call inter-shard
		h.Queue <- context.WithValue(
			context.WithValue(
				context.Background(),
				"method",
				"intershard",
			),
			"request",
			trx,
		)
	} else {
		trx.CoordinatorAddress = localAddress(h.Port)

		h.Logger.Info("cross-shard transaction", zap.Int64("session_id", trx.GetTransaction().GetSessionId()))

		// call cross-shard
		if h.Leader {
			time.Sleep(1 * time.Second)
			for _, svc := range strings.Split(h.Ipt.Endpoints[fmt.Sprintf("E%s", rshard)], ":") {
				if err := network.Prepare(h.Ipt.Services[svc], trx); err != nil {
					h.Logger.Error("failed to call participant", zap.Error(err))
				}
			}
		}
	}
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

func (h *Handler) prepare(payload interface{}) {
	// get transaction
	trx := payload.(*database.RequestMsg)

	// get sessionId
	sessionId := int(trx.GetTransaction().GetSessionId())
	h.Logger.Debug("session id is here", zap.Int("id", sessionId))

	// insert locks
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
		Sequence:  int(trx.GetTransaction().GetSequence()),
	}); err != nil {
		h.Logger.Warn("failed to store transaction", zap.Error(err))
	}

	// get the receiver balance
	balance, err := h.Storage.GetClientBalance(trx.GetTransaction().GetReceiver())
	if err != nil {
		h.Logger.Warn("failed to get client balance", zap.Error(err))
		return
	}

	// set commit or abort
	var msg string
	if trx.GetTransaction().GetAmount() <= int64(balance) {
		msg = "commit"
	} else {
		msg = "abort"
	}

	// callback the coordinator
	if h.Leader {
		if err := network.Reply(trx.CoordinatorAddress, trx.ReturnAddress, localAddress(h.Port), msg, sessionId); err != nil {
			h.Logger.Warn("failed to call the coordinator", zap.Error(err))
		}
	}
}

func (h *Handler) reply(payload interface{}) {
	// get reply message
	msg := payload.(*database.ReplyMsg)

	// get sessionId
	sessionId := int(msg.GetSessionId())
	h.Logger.Debug("session id is here", zap.Int("id", sessionId))

	// call reply on all other nodes
	if h.Leader {
		for _, svc := range strings.Split(h.Ipt.Endpoints[fmt.Sprintf("E%s", h.Cfg.Name)], ":") {
			if svc != h.Name {
				if err := network.Reply(h.Ipt.Services[svc], msg.GetReturnAddress(), msg.GetParticipantAddress(), msg.GetText(), sessionId); err != nil {
					h.Logger.Warn("failed to call node", zap.Error(err))
				}
			}
		}
	}

	// check for commit or abort
	var ctx context.Context
	if msg.GetText() == "abort" {
		ctx = context.WithValue(
			context.WithValue(
				context.Background(),
				"method",
				"abort",
			),
			"request",
			&database.AbortMsg{
				SessionId:     msg.SessionId,
				ReturnAddress: msg.GetReturnAddress(),
			},
		)
	} else if msg.GetText() == "commit" {
		// get transaction
		trx, err := h.Storage.GetTransaction(sessionId)
		if err != nil {
			h.Logger.Error("failed to get transaction", zap.Error(err))
			return
		}

		// insert locks
		if err := h.Storage.InsertLock(trx.Sender); err != nil {
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
				Record:    trx.Sender,
				NewValue:  -1 * trx.Amount,
			},
		)

		// store the logs
		if err := h.Storage.InsertBatchLogs(wals); err != nil {
			h.Logger.Warn("failed to store logs", zap.Error(err))
			return
		}

		// get the sender balance
		balance, err := h.Storage.GetClientBalance(trx.Sender)
		if err != nil {
			h.Logger.Warn("failed to get client balance", zap.Error(err))
			return
		}

		// call commit or abort
		if int64(trx.Amount) <= int64(balance) {
			ctx = context.WithValue(
				context.WithValue(
					context.Background(),
					"method",
					"commit",
				),
				"request",
				&database.CommitMsg{
					SessionId:     int64(sessionId),
					ReturnAddress: msg.GetReturnAddress(),
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
					SessionId:     int64(sessionId),
					ReturnAddress: msg.GetReturnAddress(),
				},
			)
		}
	}

	// callback the cluster
	if h.Leader {
		if ctx.Value("method").(string) == "abort" {
			// call abort on participant
			if err := network.Abort(msg.GetParticipantAddress(), msg.GetReturnAddress(), sessionId); err != nil {
				h.Logger.Warn("failed to call the participant", zap.Error(err))
			}
		} else if ctx.Value("method").(string) == "commit" {
			// call commit on participant
			if err := network.Commit(msg.GetParticipantAddress(), msg.GetReturnAddress(), sessionId); err != nil {
				h.Logger.Warn("failed to call the participant", zap.Error(err))
			}
		}
	}

	// reconcile the context again
	h.Queue <- ctx
}

func (h *Handler) abort(payload interface{}) {
	defer func() {
		go func() {
			h.notify <- nil
		}()

		time.Sleep(1 * time.Second)
	}()

	// get transaction
	trx := payload.(*database.AbortMsg)

	// get sessionId
	sessionId := int(trx.GetSessionId())
	h.Logger.Debug("session id is here", zap.Int("id", sessionId))

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
	if err := network.Reply(trx.GetReturnAddress(), "", "", "FAIL", sessionId); err != nil {
		h.Logger.Warn("failed to send reply", zap.Error(err))
	}
}

func (h *Handler) commit(payload interface{}) {
	defer func() {
		go func() {
			h.notify <- nil
		}()

		time.Sleep(1 * time.Second)
	}()

	// get transaction
	trx := payload.(*database.CommitMsg)

	// get sessionId
	sessionId := int(trx.GetSessionId())
	h.Logger.Debug("session id is here", zap.Int("id", sessionId))

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
	if err := network.Reply(trx.GetReturnAddress(), "", "", "OK", sessionId); err != nil {
		h.Logger.Warn("failed to send reply", zap.Error(err))
	}
}
