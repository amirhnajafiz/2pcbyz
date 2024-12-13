package manager

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/F24-CSE535/2pc/client/internal/utils"
	"github.com/F24-CSE535/2pc/client/pkg/models"
	"github.com/F24-CSE535/2pc/client/pkg/rpc/database"
)

func (m *Manager) PrintBalance(argc int, argv []string) string {
	// check the number of arguments
	if argc < 1 {
		return "not enough arguments"
	}

	// get the shard
	cluster, err := m.storage.GetClientShard(argv[0])
	if err != nil {
		return fmt.Errorf("database failed: %v", err).Error()
	}

	// get the cluster's services
	services := strings.Split(m.dialer.Nodes[fmt.Sprintf("E%s", cluster)], ":")

	// make RPC call
	output := fmt.Sprintf("--  server  -  %s  --\n", argv[0])
	for _, svc := range services {
		if balance, err := m.dialer.PrintBalance(svc, argv[0]); err != nil {
			return fmt.Errorf("server failed: %v", err).Error()
		} else {
			output = fmt.Sprintf("%s     %s     -  %d\n", output, svc, balance)
		}
	}

	return output
}

func (m *Manager) PrintDatastores(argc int, argv []string) ([]string, string) {
	// set a list of servers
	servers := strings.Split(m.dialer.Nodes["all"], ":")

	// create a list of records
	records := make([]string, 0)

	// loop over servers and get messages
	for _, svc := range servers {
		// make RPC call
		list, err := m.dialer.PrintDatastore(svc)
		if err != nil {
			return nil, err.Error()
		}

		// get sessions from the database
		output := fmt.Sprintf("%s:\n", svc)
		for _, msg := range list {
			if ses, err := m.storage.GetSessionById(int(msg.GetSessionId())); err == nil {
				output = fmt.Sprintf(
					"%s\t[<%d, %s>, (%s, %s, %d)]\n",
					output,
					msg.GetBallotNumberSequence(),
					msg.GetBallotNumberPid(),
					ses.Sender,
					ses.Receiver,
					ses.Amount,
				)
			} else {
				log.Printf("failed to get session %d: %v\n", msg.GetSessionId(), err)
			}
		}

		records = append(records, output)
	}

	return records, ""
}

func (m *Manager) Transaction(argc int, argv []string) (string, bool) {
	// check the number of arguments
	if argc < 3 {
		return "not enough arguments", false
	}

	// extract data from input command
	sender := argv[0]
	receiver := argv[1]
	amount, _ := strconv.Atoi(argv[2])
	sessionId := m.memory.GetSession()

	// create a new session
	session := models.Session{
		Id:       sessionId,
		Sender:   sender,
		Receiver: receiver,
		Amount:   amount,
		Replys:   make([]*database.ReplyMsg, 0),
	}

	// save the transaction into cache
	session.StartedAt = time.Now()
	m.cache[sessionId] = &session

	// store session for future optimizations
	if err := m.storage.InsertSession(&session); err != nil {
		log.Printf("failed to store cross-shard metric: %v\n", err)
	}

	return fmt.Sprintf("transaction %d (%s, %s, %d): sent", sessionId, sender, receiver, amount), true
}

func (m *Manager) Block(argc int, argv []string) string {
	// check the number of arguments
	if argc < 1 {
		return "not enough arguments"
	}

	// make RPC call
	if err := m.dialer.Block(argv[0]); err != nil {
		return fmt.Errorf("server failed: %v", err).Error()
	}

	return "blocked"
}

func (m *Manager) Unblock(argc int, argv []string) string {
	// check the number of arguments
	if argc < 1 {
		return "not enough arguments"
	}

	// make RPC call
	if err := m.dialer.Unblock(argv[0]); err != nil {
		return fmt.Errorf("server failed: %v", err).Error()
	}

	return "unblocked"
}

func (m *Manager) LoadTests(argc int, argv []string) string {
	// check the number of arguments
	if argc < 1 {
		return "not enough arguments"
	}

	// load tests
	tc, err := utils.CSVParseTestcaseFile(argv[0])
	if err != nil {
		return err.Error()
	}

	// set tests
	m.tests = tc
	m.index = 0

	return fmt.Sprintf("load %d testsets", len(tc))
}
