package manager

import (
	"time"

	"github.com/F24-CSE535/2pc/client/pkg/rpc/database"
)

// handleReply accepts a reply message for an active session.
func (m *Manager) handleReply(msg *database.ReplyMsg) {
	if session, ok := m.cache[int(msg.GetSessionId())]; ok {
		// append the reply to the list
		session.Replys = append(session.Replys, msg)

		// check for the number of replys
		if len(session.Replys) == len(session.Participants) {
			fn := time.Now()

			// return the message to client
			session.Text = msg.GetText()
			m.output <- session

			// update performance metrics
			du := fn.Sub(session.StartedAt).Nanoseconds() / 1000000
			m.throughput = append(m.throughput, float64(1000/du))
			m.latency = append(m.latency, float64(du))
		}
	}
}

// handleTimeouts unblocks the sessions that are not finished and they hit timeout.
func (m *Manager) handleTimeouts() {
	for {
		for _, value := range m.cache {
			if len(value.Text) == 0 && time.Since(value.StartedAt) >= 5*time.Second {
				// reset the timer
				value.StartedAt = time.Now()
			}
		}

		time.Sleep(5 * time.Second)
	}
}
