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
	Name      string
	Leader    bool
	Cfg       *config.Config
	Ipt       *config.IPTable
	Logger    *zap.Logger
	Storage   *storage.Storage
	Consensus chan context.Context
	Queue     chan context.Context

	dispatcher chan context.Context
	notify     chan context.Context

	states map[int]string
}

func (h *Handler) dispatch() {
	for {
		ctx := <-h.dispatcher

		method := ctx.Value("method").(string)
		payload := ctx.Value("request")

		h.Queue <- context.WithValue(context.WithValue(context.Background(), "method", "process-"+method), "request", payload)

		<-h.notify
	}
}

// Start consuming messages.
func (h *Handler) Start() {
	h.states = make(map[int]string)
	h.dispatcher = make(chan context.Context, 20)
	h.notify = make(chan context.Context)

	go h.dispatch()

	for {
		// get context messages from queue
		ctx := <-h.Queue
		payload := ctx.Value("request")

		h.Logger.Debug("input request", zap.String("method", ctx.Value("method").(string)))

		// map of method to handler
		switch ctx.Value("method").(string) {
		case "begin":
			h.begin(payload)
		case "intershard":
			h.dispatcher <- ctx
		case "process-intershard":
			h.intershard(payload)
		case "prepare":
			h.dispatcher <- ctx
		case "process-prepare":
			h.prepare(payload)
		case "reply":
			h.reply(payload)
		case "abort":
			h.abort(payload)
		case "commit":
			h.commit(payload)
		}

		h.Logger.Debug("done", zap.String("method", ctx.Value("method").(string)))
	}
}
