package handler

import (
	"context"

	"github.com/F24-CSE535/2pcbyz/cluster/internal/storage"

	"go.uber.org/zap"
)

// Handler is a process that gets requests from gRPC module and executes sub-handlers based on the input request.
type Handler struct {
	Logger  *zap.Logger
	Storage *storage.Storage
	Queue   chan context.Context
}

// Start consuming messages.
func (h *Handler) Start() {
	for {
		// get context messages from queue
		ctx := <-h.Queue
		payload := ctx.Value("request")

		// map of method to handler
		switch ctx.Value("method").(string) {
		case "request":
			h.request(payload)
		case "abort":
			h.abort(payload)
		case "commit":
			h.commit(payload)
		}
	}
}
