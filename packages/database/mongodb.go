package database

import (
	"github.com/divyag9/goqueues/packages/queue"
	"gopkg.in/mgo.v2"
)

// Mongo contains the mongo database session
type Mongo struct {
	Session *mgo.Session
}

// Get returns the details of a particular queue
func (m *Mongo) Get(id int) (*queue.Details, error) {
	// implement logic
	return nil, nil
}

// GetAll returns the details of all queue
func (m *Mongo) GetAll() ([]*queue.Details, error) {
	// implement logic
	return nil, nil
}

// Save saves the queue details
func (m *Mongo) Save(*queue.Details) error {
	// implement logic
	return nil
}
