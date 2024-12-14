package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Storage is a module that uses mongo-driver library to handle MongoDB queries.
type Storage struct {
	shards *mongo.Collection
}

// NewStorage opens a MongoDB connection and returns an instance of storage struct.
func NewStorage(uri string, database string) (*Storage, error) {
	// open a new connection to MongoDB cluster
	conn, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to open a MongoDB connection: %v", err)
	}

	// create a new instance
	instance := Storage{
		shards: conn.Database(database).Collection("shards"),
	}

	return &instance, nil
}
