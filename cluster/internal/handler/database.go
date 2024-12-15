package handler

import "github.com/F24-CSE535/2pcbyz/cluster/pkg/rpc/database"

func (h *Handler) request(payload interface{}) {
	// get transaction
	trx := payload.(*database.RequestMsg)

	// insert locks

	// insert logs

	// insert transaction

	// call commit
}

func (h *Handler) abort(payload interface{}) {
	// get transaction
	trx := payload.(*database.AbortMsg)

	// insert log

	// update transaction
}

func (h *Handler) commit(payload interface{}) {
	// get transaction
	trx := payload.(*database.CommitMsg)

	// insert log

	// update client

	// update transaction
}
