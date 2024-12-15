package memory

// *SharedMemory is a local storage for processes and handlers.
type SharedMemory struct {
	SeqLT  int
	Leader string
	Name   string
	AcMsgs []string
	Inputs []string
}

// NewSharedMemory returns an instance of shared memory.
func NewSharedMemory() *SharedMemory {
	return &SharedMemory{}
}

func (s *SharedMemory) GetClientLastTimestamp() int {
	return s.SeqLT
}

func (s *SharedMemory) GetLeader() string {
	return s.Leader
}

func (s *SharedMemory) GetNodeName() string {
	return s.Name
}

func (s *SharedMemory) SetLeader(input string) {
	s.Leader = input
}

func (s *SharedMemory) AppendInput(in string) {
	s.Inputs = append(s.Inputs, in)
}

func (s *SharedMemory) GetInputMessages() []string {
	return s.Inputs
}

func (s *SharedMemory) ResetAcceptedMessages() {
	s.AcMsgs = make([]string, 0)
}
