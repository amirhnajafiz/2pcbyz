package models

// list of system's write-ahead log types
const (
	WALStart  = "start"
	WALUpdate = "update"
	WALCommit = "commit"
	WALAbort  = "abort"
)

// Log is an entity used in our system's write-ahead logging process.
type Log struct {
	Message   string `bson:"message"`
	Record    string `bson:"record"`
	SessionId int    `bson:"session_id"`
	NewValue  int    `bson:"new_value"`
	Sequence  int    `bson:"sequence"`
}
