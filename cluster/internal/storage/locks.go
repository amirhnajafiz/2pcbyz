package storage

import (
	"context"
	"time"

	"github.com/F24-CSE535/2pcbyz/cluster/internal/models"
)

// InsertLock gets a record and inserts a lock record.
func (s *Storage) InsertLock(record string) error {
	_, err := s.locks.InsertOne(context.TODO(), &models.Lock{
		Record:    record,
		DeletedAt: time.Now().String(),
	})

	return err
}

// ReleaseLock unlocks a captured lock.
func (s *Storage) ReleaseLock(record string) error {
	_, err := s.locks.InsertOne(context.TODO(), &models.Lock{
		Record:    record,
		DeletedAt: time.Now().String(),
	})

	return err
}
