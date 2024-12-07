package viewchange

import (
	"time"

	"github.com/f24-cse535/2pcbyz/internal/timer"
	"github.com/f24-cse535/2pcbyz/pkg/enums"
	"github.com/f24-cse535/2pcbyz/pkg/types"
)

type ViewchangeDispatcher struct {
	inputChannel   chan *types.Packet
	forwardChannel chan *types.Packet

	hd handler
}

func NewDispatcher(fc chan *types.Packet) *ViewchangeDispatcher {
	ic := make(chan *types.Packet)

	return &ViewchangeDispatcher{
		inputChannel:   ic,
		forwardChannel: fc,
		hd: handler{
			timer: timer.NewTimer(1*time.Second, ic),
		},
	}
}

func (v *ViewchangeDispatcher) Start() {
	for {
		pkt := <-v.inputChannel

		switch pkt.Header {
		case enums.PacketViewChange:
			v.hd.hdViewChange()
		case enums.PacketNewView:
			v.hd.hdNewView()
		default:
			v.hd.hdDefault()
		}

		if v.forwardChannel != nil {
			v.forwardChannel <- pkt
		}
	}
}

func (v *ViewchangeDispatcher) C() chan *types.Packet {
	return v.inputChannel
}
