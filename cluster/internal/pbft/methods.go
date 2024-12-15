package pbft

func (sm *StateMachine) request(payload interface{}) error {
	return nil
}

func (sm *StateMachine) prePrepare(payload interface{}) error {
	return nil
}

func (sm *StateMachine) ackPrePrepare(payload interface{}) error {
	return nil
}

func (sm *StateMachine) prepare(payload interface{}) error {
	return nil
}

func (sm *StateMachine) ackPrepare(payload interface{}) error {
	return nil
}

func (sm *StateMachine) commit(payload interface{}) error {
	return nil
}

func (sm *StateMachine) block() error {
	return nil
}

func (sm *StateMachine) unblock() error {
	return nil
}

func (sm *StateMachine) byzantine() error {
	return nil
}

func (sm *StateMachine) nonByzantine() error {
	return nil
}
