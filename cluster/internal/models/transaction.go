package models

// list of system's transaction's status
const (
	StsCommit = "commit"
	StsAbort  = "abort"
)

// Transaction is a bson model to store the client requests.
type Transaction struct {
	Sender    string `bson:"sender"`
	Receiver  string `bson:"receiver"`
	Amount    int    `bson:"amount"`
	SessionId int    `bson:"session_id"`
	Status    string `bson:"status"`
}
