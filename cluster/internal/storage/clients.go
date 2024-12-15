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
	_, err := s.clients.DeleteMany(context.TODO(), filter)
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
	_, err := s.clients.InsertMany(context.TODO(), records)

	return err
}

// GetClientBalance returns a balance value by accepting a client.
func (s *Storage) GetClientBalance(client string) (int, error) {
	// create a filter for the specified cluster
	filter := bson.M{"client": client}

	// decode the response
	var clientInstance models.Client
	if err := s.clients.FindOne(context.TODO(), filter).Decode(&clientInstance); err != nil {
		return 0, err
	}

	return int(clientInstance.Balance), nil
}

// UpdateClientBalance gets a client and new balance to update the balance value.
func (s *Storage) UpdateClientBalance(client string, balance int, set bool) error {
	// create a filter for the specified cluster
	filter := bson.M{"client": client}

	// define the update operation
	var update bson.D
	if set {
		update = bson.D{{Key: "$set", Value: bson.D{{Key: "balance", Value: balance}}}}
	} else {
		update = bson.D{{Key: "$inc", Value: bson.D{{Key: "balance", Value: balance}}}}
	}

	// perform the update query
	_, err := s.clients.UpdateMany(context.TODO(), filter, update)

	return err
}
