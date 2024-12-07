package consensus

import "github.com/f24-cse535/2pcbyz/internal/timer"

type handler struct {
	timer *timer.Timer
}

func (h *handler) hdPreprepare() {}

func (h *handler) hdPrepare() {}

func (h *handler) hdCommit() {}
