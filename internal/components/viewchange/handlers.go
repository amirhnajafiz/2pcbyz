package viewchange

import "github.com/f24-cse535/2pcbyz/internal/timer"

type handler struct {
	timer *timer.Timer
}

func (h *handler) hdViewChange() {}

func (h *handler) hdNewView() {}

func (h *handler) hdDefault() {}
