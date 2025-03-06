package storage

import (
	"context"

	"github.com/F24-CSE535/2pcbyz/cluster/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

// InsertLog adds a new log to the node's logs.
func (s *Storage) InsertLog(log *models.Log) error {
	// insert log
	_, err := s.logs.InsertOne(context.TODO(), log)

	return err
}

// InsertBatchLogs into the database.
func (s *Storage) InsertBatchLogs(logs []*models.Log) error {
	// convert log model to interface
	records := make([]interface{}, 0)
	for _, item := range logs {
		records = append(records, item)
	}

	_, err := s.logs.InsertMany(context.TODO(), records)

	return err
}

// GetLogsBySessionId gets a sessionId and returns the logs for that session.
func (s *Storage) GetLogsBySessionId(sessionId int) ([]*models.Log, error) {
	// create a filter for the specified log
	filter := bson.M{
		"session_id": sessionId,
		"message":    models.WALUpdate,
	}

	// find all documents that match the filter
	cursor, err := s.logs.Find(context.TODO(), filter)
	if err != nil {
		return nil, err
	}
	defer cursor.Close(context.TODO())

	// decode the results into a slice of Logs structs
	var results []*models.Log
	if err = cursor.All(context.TODO(), &results); err != nil {
		return nil, err
	}

	return results, nil
}
