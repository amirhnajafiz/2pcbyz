package handler

import (
	"fmt"
	"time"

	"github.com/F24-CSE535/2pcbyz/client/internal/config"
)

// Handler is map that binds the user input commands to an executable function.
type Handler struct {
	session  int
	cfg      *config.Config
	ipt      *config.IPTable
	handlers map[string]func(int, []string) error
}

// NewHandler returns a handler instance.
func NewHandler(cfg *config.Config, ipt *config.IPTable) *Handler {
	instance := &Handler{
		session:  int(time.Now().Unix()),
		cfg:      cfg,
		ipt:      ipt,
		handlers: make(map[string]func(int, []string) error),
	}

	// define a handles map to callback function
	instance.handlers["request"] = instance.request

	return instance
}

// Exec accepts all user inputs and calls a command in handler's handlers map.
func (h *Handler) Exec(cmd string, argc int, argv []string) error {
	if callback, ok := h.handlers[cmd]; ok {
		return callback(argc, argv)
	}

	return fmt.Errorf("no command: %s", cmd)
}
