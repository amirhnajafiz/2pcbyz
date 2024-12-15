package storage

import (
	"context"
	"fmt"

	"github.com/F24-CSE535/2pcbyz/cluster/internal/models"
	"go.mongodb.org/mongo-driver/bson"
)

// DeleteShards removes all existing shards of the node.
func (s *Storage) DeleteShards() error {
	// creating an empty filter
	filter := bson.D{}

	// delete all existing shards
	_, err := s.shards.DeleteMany(context.TODO(), filter)
	if err != nil {
		return err
	}

	return nil
}

// InsertShards creates client's records for the node.
func (s *Storage) InsertShards(start, end int) error {
	// create a list of records
	records := make([]interface{}, 0)
	for i := start; i <= end; i++ {
		records = append(records, &models.Client{
			Client:  fmt.Sprintf("%d", i),
			Balance: 10,
		})
	}

	// insert all shards
	_, err := s.shards.InsertMany(context.TODO(), records)

	return err
}
