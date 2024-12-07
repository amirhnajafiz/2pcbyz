package statemachine

import (
	"github.com/f24-cse535/2pcbyz/pkg/enums"
	"github.com/f24-cse535/2pcbyz/pkg/types"
)

type StatemachineDispatcher struct {
	inputChannel chan *types.Packet

	hd handler
}

func NewDispatcher() *StatemachineDispatcher {
	return &StatemachineDispatcher{
		inputChannel: make(chan *types.Packet),
		hd:           handler{},
	}
}

func (s *StatemachineDispatcher) Start() {
	for {
		pkt := <-s.inputChannel

		switch pkt.Header {
		case enums.PacketRequest:
			s.hd.hdRequest()
		}
	}
}

func (s *StatemachineDispatcher) C() chan *types.Packet {
	return s.inputChannel
}
