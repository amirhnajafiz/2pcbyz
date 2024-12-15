package memory

// SharedMemory is a local storage for processes and handlers.
type SharedMemory struct{}

// NewSharedMemory returns an instance of shared memory.
func NewSharedMemory() *SharedMemory {
	return &SharedMemory{}
}

func (s SharedMemory) GetLeader() string {
	return ""
}

func (s SharedMemory) GetNodeName() string {
	return ""
}

func (s SharedMemory) SetLeader(input string) {
}

func (s SharedMemory) ResetAcceptedMessages() {
}
