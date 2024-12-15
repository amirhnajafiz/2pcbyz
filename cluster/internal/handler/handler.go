package handler

import (
	"context"

	"github.com/F24-CSE535/2pcbyz/cluster/internal/config"
	"github.com/F24-CSE535/2pcbyz/cluster/internal/storage"

	"go.uber.org/zap"
)

// Handler is a process that gets requests from gRPC module and executes sub-handlers based on the input request.
type Handler struct {
	Sequence  int
	Port      int
	Cfg       *config.Config
	Ipt       *config.IPTable
	Logger    *zap.Logger
	Storage   *storage.Storage
	Consensus chan context.Context
	Queue     chan context.Context

	states map[int]string
}

// Start consuming messages.
func (h *Handler) Start() {
	h.states = make(map[int]string)

	for {
		// get context messages from queue
		ctx := <-h.Queue
		payload := ctx.Value("request")

		// map of method to handler
		switch ctx.Value("method").(string) {
		case "request":
			h.begin(payload)
		case "intershard":
			h.intershard(payload)
		case "prepare":
			h.prepare(payload)
		case "reply":
			h.reply(payload)
		case "abort":
			h.abort(payload)
		case "commit":
			h.commit(payload)
		}
	}
}
