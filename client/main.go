package main

import (
	"os"

	"github.com/F24-CSE535/2pcbyz/client/internal/config"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		panic("at least two arguments are needed (./main <config-path> <iptable>)")
	}

	// load config file
	_ = config.New(args[1])

	// load iptable
	_ = config.NewIPTable(args[2])
}
