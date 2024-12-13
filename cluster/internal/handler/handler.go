package handler

import (
	"context"

	"github.com/F24-CSE535/2pcbyz/cluster/internal/storage"

	"go.uber.org/zap"
)

// Handler is a process that gets requests from gRPC module and executes
// sub-handlers based on the input request.
type Handler struct {
	Logger  *zap.Logger
	Storage *storage.Storage
	Queue   chan context.Context
}

func (h *Handler) Start() {
	for {
		<-h.Queue
	}
}
