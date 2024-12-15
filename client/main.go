package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"github.com/F24-CSE535/2pcbyz/client/internal/config"
	"github.com/F24-CSE535/2pcbyz/client/internal/handler"
	"github.com/F24-CSE535/2pcbyz/client/internal/server"
	"github.com/F24-CSE535/2pcbyz/client/internal/utils"
)

func main() {
	args := os.Args
	if len(args) < 3 {
		panic("at least two arguments are needed (./main <config-path> <iptable> <testcase>)")
	}

	// load config file
	cfg := config.New(args[1])

	// load iptable file
	ipt := config.NewIPTable(args[2])

	// create a new handler
	hd := handler.NewHandler(&cfg, &ipt)

	// load the input tests
	if len(args) == 4 {
		if val, err := utils.CSVParseTestcaseFile(args[3]); err == nil {
			hd.SetTests(val)
		}
	}

	// run a gRPC server
	go server.ListenAndServe(cfg.Port)

	// in a for loop, read user commands
	reader := bufio.NewReader(os.Stdin)
	for {
		fmt.Print("\n$ ")

		input, _ := reader.ReadString('\n') // read input until newline
		input = strings.TrimSpace(input)

		// no input
		if len(input) == 0 {
			continue
		}

		// split into parts
		parts := strings.Split(input, " ")

		// create args for the client handlers
		cargs := parts[1:]
		cargsc := len(cargs)

		// call exec on handler
		if msg, err := hd.Exec(parts[0], cargsc, cargs); err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(msg)
		}
	}
}
