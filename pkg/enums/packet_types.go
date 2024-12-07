package enums

const (
	PacketViewChange int = iota + 1
	PacketNewView
	PacketTimeoutLeader
	PacketTimeoutViewChange
	PacketTimeoutNewView
)

const (
	PacketPreprepare int = iota + 10
	PacketPrepare
	PacketCommit
	PacketAckPreprepare
	PacketAckPrepare
	PacketTimeoutPreprepare
	PacketTimeoutPrepare
)

const (
	PacketRequest int = iota + 20
)
