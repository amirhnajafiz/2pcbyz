package handler

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/F24-CSE535/2pcbyz/client/internal/network"
	"github.com/F24-CSE535/2pcbyz/client/pkg/rpc/database"
)

// request accepts all parameters for a transaction and calls a Request RPC.
func (h *Handler) request(argc int, argv []string) (string, error) {
	if argc < 3 {
		return "", errors.New("not enough arguments (sender receiver amount)")
	}

	// parse input arguments
	sender := argv[0]
	receiver := argv[1]
	amount, _ := strconv.Atoi(argv[2])

	// set a sessionId for the transaction
	session := h.session
	h.session++

	// find the clinet shard
	sshard := findClientShard(sender, h.cfg.Shards)
	cshard := findClientShard(receiver, h.cfg.Shards)

	// check availability
	if h.lives[sshard]-h.byzantines[sshard] < 3 || h.lives[cshard]-h.byzantines[cshard] < 3 {
		return fmt.Sprintf("%d: not enough servers to process the transaction", session), nil
	}

	// call request RPC
	if err := network.Request(h.ipt.Services[h.ipt.Endpoints[sshard]], &database.RequestMsg{
		Transaction: &database.TransactionMsg{ // create a new transaction
			Sender:    sender,
			Receiver:  receiver,
			Amount:    int64(amount),
			SessionId: int64(session),
		},
		ReturnAddress: localAddress(h.cfg.Port), // set the return address
	}); err != nil {
		return "", fmt.Errorf("rpc failed: %v", err)
	}

	return fmt.Sprintf("transaction %d (%s %s %d) submitted", session, sender, receiver, amount), nil
}

// printBalance accepts a client id and gets its balance over all nodes of a cluster.
func (h *Handler) printBalance(argc int, argv []string) (string, error) {
	if argc < 1 {
		return "", errors.New("not enough arguments (client)")
	}

	client := argv[0]

	// get client shard
	shard := findClientShard(client, h.cfg.Shards)

	// loop over all nodes inside a cluster and call printbalance
	output := fmt.Sprintf("client: %s\n", client)
	for _, svc := range strings.Split(h.ipt.Endpoints[fmt.Sprintf("E%s", shard)], ":") {
		if amount, err := network.PrintBalance(h.ipt.Services[svc], client); err != nil {
			return "", err
		} else {
			output = fmt.Sprintf("%s\t- %s: %d\n", output, svc, amount)
		}
	}

	return output, nil
}

// printDatastore loops over all nodes and call PrintDatastore RPC.
func (h *Handler) printDatastore(_ int, _ []string) (string, error) {
	list := []string{"C1", "C2", "C3"}

	// loop over all nodes and call print datastore
	output := "datastores:\n"
	for _, cluster := range list {
		for _, svc := range strings.Split(h.ipt.Endpoints[fmt.Sprintf("E%s", cluster)], ":") {
			if datastore, err := network.PrintDatastore(h.ipt.Services[svc]); err != nil {
				return "", err
			} else {
				output = fmt.Sprintf("%s\t- %s:\n", output, svc)
				for _, item := range datastore {
					output = fmt.Sprintf("%s\t\t- %s\n", output, item)
				}
			}
		}
	}

	return output, nil
}

// next runs the next testcase.
func (h *Handler) next(_ int, _ []string) (string, error) {
	// check the index variable
	if h.index == len(h.tests) {
		return "end of tests", nil
	}

	list := []string{"C1", "C2", "C3"}
	for _, item := range list {
		h.lives[item] = 0
		h.byzantines[item] = 0
	}

	// block all servers
	for _, svc := range strings.Split(h.ipt.Endpoints["all"], ":") {
		network.Block(h.ipt.Services[svc])
		network.NonByzantine(h.ipt.Services[svc])
	}

	output := fmt.Sprintf("test set %d:\n", h.index+1)

	// update servers status
	for _, svc := range h.tests[h.index]["servers"].([]string) {
		network.Unblock(h.ipt.Services[svc])
		for _, cluster := range list {
			for _, tmp := range strings.Split(h.ipt.Endpoints[fmt.Sprintf("E%s", cluster)], ":") {
				if svc == tmp {
					h.lives[cluster]++
				}
			}
		}
	}

	for _, svc := range h.tests[h.index]["byzantines"].([]string) {
		network.Byzantine(h.ipt.Services[svc])
		for _, cluster := range list {
			for _, tmp := range strings.Split(h.ipt.Endpoints[fmt.Sprintf("E%s", cluster)], ":") {
				if svc == tmp {
					h.byzantines[cluster]++
				}
			}
		}
	}

	// make transactions by calling the handler request
	for _, trx := range h.tests[h.index]["transactions"].([][]string) {
		if msg, err := h.request(3, trx); err != nil {
			return "", err
		} else {
			output = fmt.Sprintf("%s\t- %s\n", output, msg)
		}
	}

	// increase index
	h.index++

	return output, nil
}

// exit terminates the client program.
func (h *Handler) exit(_ int, _ []string) (string, error) {
	os.Exit(0)

	return "", nil
}
