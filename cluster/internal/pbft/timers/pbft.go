package timers

import (
	"time"

	"github.com/F24-CSE535/2pcbyz/cluster/internal/pbft/memory"

	"go.uber.org/zap"
)

// PBFTTimer will be used in PBFT handler to start requests' timer.
type PBFTTimer struct {
	logger *zap.Logger
	memory *memory.SharedMemory

	consensusTimeout time.Duration

	consensusTimerChan   chan bool
	dispatcherNotifyChan chan bool
}

// consensusTimers starts a consensus timer and if it hits timeout it will generate a timeout message.
func (p *PBFTTimer) StartConsensusTimer(ra string, sessionId int) {
	// create a new timer and start it
	timer := time.NewTimer(p.consensusTimeout)

	select {
	case <-p.consensusTimerChan:
		timer.Stop()
	case <-timer.C:
		// close everything and reply with timeout
		p.memory.ResetAcceptedMessages()
		p.logger.Info("consensus timeout", zap.Int("session id", sessionId))

		// accept next request
		p.dispatcherNotifyChan <- true
	}
}

// FinishConsensusTimer sends a notify signal to the consensus timer channel.
func (p *PBFTTimer) FinishConsensusTimer() {
	p.consensusTimerChan <- true
}
