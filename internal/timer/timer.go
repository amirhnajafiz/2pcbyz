package timer

import (
	"sync"
	"time"

	"github.com/f24-cse535/2pcbyz/pkg/types"
)

type Timer struct {
	events    map[string]*event
	output    chan *types.Packet
	period    time.Duration
	tableLock sync.Mutex
}

func NewTimer(p time.Duration, oc chan *types.Packet) *Timer {
	return &Timer{
		events:    make(map[string]*event),
		output:    oc,
		period:    p,
		tableLock: sync.Mutex{},
	}
}

func (t *Timer) NewEvent(id string, label int, expiresAt time.Time) {
	t.tableLock.Lock()
	t.events[id] = &event{label: label, expiresAt: expiresAt}
	t.tableLock.Unlock()
}

func (t *Timer) Finish(id string) {
	t.tableLock.Lock()
	delete(t.events, id)
	t.tableLock.Unlock()
}

func (t *Timer) Start() {
	for {
		ts := time.Now()

		t.tableLock.Lock()
		for key, value := range t.events {
			if value.expiresAt.Before(ts) {
				t.output <- &types.Packet{
					Header:  value.label,
					Payload: key,
				}
			}
		}
		t.tableLock.Unlock()

		time.Sleep(t.period)
	}
}
