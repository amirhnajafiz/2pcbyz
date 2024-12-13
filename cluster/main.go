package main

import (
	"os"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		panic("at least four arguments are needed (./main <config-path> <iptable-path>)")
	}
}
