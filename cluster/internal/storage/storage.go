package storage

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Storage is a module that uses mongo-driver library to handle MongoDB queries.
type Storage struct {
	clients      *mongo.Collection
	locks        *mongo.Collection
	logs         *mongo.Collection
	transactions *mongo.Collection
}

// NewStorage opens a MongoDB connection and returns an instance of storage struct.
func NewStorage(uri, database, partition string) (*Storage, error) {
	// open a new connection to MongoDB cluster
	conn, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		return nil, fmt.Errorf("failed to open a MongoDB connection: %v", err)
	}

	// create a new instance
	instance := Storage{
		locks:        conn.Database(database).Collection(fmt.Sprintf("%s_locks", partition)),
		logs:         conn.Database(database).Collection(fmt.Sprintf("%s_logs", partition)),
		clients:      conn.Database(database).Collection(fmt.Sprintf("%s_clients", partition)),
		transactions: conn.Database(database).Collection(fmt.Sprintf("%s_transactions", partition)),
	}

	return &instance, nil
}
