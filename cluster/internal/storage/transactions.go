package storage

import (
	"context"

	"github.com/F24-CSE535/2pcbyz/cluster/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

// InsertTransaction adds a new transaction to the node's transactions.
func (s *Storage) InsertTransaction(trx *models.Transaction) error {
	// insert transaction
	_, err := s.transactions.InsertOne(context.TODO(), trx)

	return err
}

// GetCommittedTransactions returns the list of committed transactions on the node.
func (s *Storage) GetCommittedTransactions() ([]*models.Transaction, error) {
	// create a filter to select committeds
	filter := bson.M{"status": models.StsCommit}

	// find all documents that match the filter
	cursor, err := s.transactions.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// decode the results into a slice of Transaction structs
	var results []*models.Transaction
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}

// UpdateTransactionStatus gets a trx sessionId and a status and updates the transaction status.
func (s *Storage) UpdateTransactionStatus(sessionId int, status string) error {
	// create filter and update option
	filter := bson.M{"session_id": sessionId}
	update := bson.D{{Key: "$set", Value: bson.D{{Key: "status", Value: status}}}}

	// perform the update query
	_, err := s.transactions.UpdateMany(context.TODO(), filter, update)

	return err
}
