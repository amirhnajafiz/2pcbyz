package consensus

import (
	"time"

	"github.com/f24-cse535/2pcbyz/internal/timer"
	"github.com/f24-cse535/2pcbyz/pkg/enums"
	"github.com/f24-cse535/2pcbyz/pkg/types"
)

type ConsensusDispatcher struct {
	inputChannel   chan *types.Packet
	forwardChannel chan *types.Packet

	hd handler
}

func NewDispatcher(fc chan *types.Packet) *ConsensusDispatcher {
	ic := make(chan *types.Packet)

	return &ConsensusDispatcher{
		inputChannel:   ic,
		forwardChannel: fc,
		hd: handler{
			timer: timer.NewTimer(1*time.Second, ic),
		},
	}
}

func (c *ConsensusDispatcher) Start() {
	for {
		pkt := <-c.inputChannel

		switch pkt.Header {
		case enums.PacketPreprepare:
			c.hd.hdPreprepare()
		case enums.PacketPrepare:
			c.hd.hdPrepare()
		case enums.PacketCommit:
			c.hd.hdCommit()
		}

		if c.forwardChannel != nil {
			c.forwardChannel <- pkt
		}
	}
}

func (c *ConsensusDispatcher) C() chan *types.Packet {
	return c.inputChannel
}
