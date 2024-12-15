package pbft

import (
	"context"

	"go.uber.org/zap"
)

type StateMachine struct {
	Consensus chan context.Context
	Queue     chan context.Context
	Logger    *zap.Logger
}

func (sm *StateMachine) Start() {
	for {
		// get context messages from queue
		ctx := <-sm.Consensus
		payload := ctx.Value("request")

		// create an error variable for handlers result
		var err error

		// map of method to handler
		switch ctx.Value("method").(string) {
		case "request":
			err = sm.request(payload)
		case "preprepare":
			err = sm.prePrepare(payload)
		case "ackpreprepare":
			err = sm.ackPrePrepare(payload)
		case "prepare":
			err = sm.prepare(payload)
		case "ackprepare":
			err = sm.ackPrepare(payload)
		case "commit":
			err = sm.commit(payload)
		default:
			sm.Queue <- ctx
		}

		// check error
		if err != nil {
			sm.Logger.Warn("state-machine error", zap.Error(err))
		}
	}
}
