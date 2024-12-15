package handler

import (
	"fmt"
	"time"

	"github.com/F24-CSE535/2pcbyz/client/internal/config"
)

// Handler is map that binds the user input commands to an executable function.
type Handler struct {
	session  int
	index    int
	cfg      *config.Config
	ipt      *config.IPTable
	handlers map[string]func(int, []string) (string, error)
	tests    []map[string]interface{}

	lives      map[string]int
	byzantines map[string]int
}

// NewHandler returns a handler instance.
func NewHandler(cfg *config.Config, ipt *config.IPTable) *Handler {
	instance := &Handler{
		session:    int(time.Now().Unix()),
		index:      0,
		cfg:        cfg,
		ipt:        ipt,
		handlers:   make(map[string]func(int, []string) (string, error)),
		lives:      make(map[string]int),
		byzantines: make(map[string]int),
	}

	// define a handles map to callback function
	instance.handlers["request"] = instance.request
	instance.handlers["printbalance"] = instance.printBalance
	instance.handlers["printdatastore"] = instance.printDatastore
	instance.handlers["next"] = instance.next
	instance.handlers["exit"] = instance.exit

	return instance
}

// SetTests updates the value of tests variable.
func (h *Handler) SetTests(ts []map[string]interface{}) {
	h.index = 0
	h.tests = ts
}

// Exec accepts all user inputs and calls a command in handler's handlers map.
func (h *Handler) Exec(cmd string, argc int, argv []string) (string, error) {
	if callback, ok := h.handlers[cmd]; ok {
		return callback(argc, argv)
	}

	return "", fmt.Errorf("no command: %s", cmd)
}
