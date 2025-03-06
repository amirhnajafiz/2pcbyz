package timers

import (
	"time"

	"github.com/F24-CSE535/2pcbyz/cluster/internal/pbft/memory"

	"go.uber.org/zap"
)

// NewLeaderTimer returns an instance of leader timer.
func NewLeaderTimer(
	logger *zap.Logger,
	memory *memory.SharedMemory,
	leaderTo,
	leaderPi int,
) *LeaderTimer {
	instance := LeaderTimer{
		leaderTimeout:      time.Duration(leaderTo) * time.Second,
		leaderPingInterval: time.Duration(leaderPi) * time.Second,
		logger:             logger,
		memory:             memory,
		leaderPingChan:     make(chan bool),
		leaderTimerChan:    make(chan bool),
	}

	// start the leader timer and leader pinger
	go instance.leaderTimer()
	go instance.leaderPinger()

	return &instance
}

// NewPBFTTimer returns an instance of PBFT timer.
func NewPBFTTimer(
	cto int,
	logger *zap.Logger,
	memory *memory.SharedMemory,
	dispatcherNotifyChan chan bool,
) *PBFTTimer {
	return &PBFTTimer{
		consensusTimeout:     time.Duration(cto) * time.Millisecond,
		logger:               logger,
		memory:               memory,
		consensusTimerChan:   make(chan bool),
		dispatcherNotifyChan: dispatcherNotifyChan,
	}
}
