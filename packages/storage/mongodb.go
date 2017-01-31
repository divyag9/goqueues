package storage

import (
	"github.com/divyag9/goqueues/packages/queue"
	mgo "gopkg.in/mgo.v2"
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

// GetAll returns the details of all queues
func (m *Mongo) GetAll() ([]*queue.Details, error) {
	var queueDetails []*queue.Details
	if err := m.Session.DB("queues").C("details").
		Find(nil).Sort("-when").Limit(100).All(&queueDetails); err != nil {

		return nil, err
	}
	return queueDetails, nil
}

// Save saves the queue details
func (m *Mongo) Save(qd *queue.Details) error {
	if err := m.Session.DB("queues").C("details").Insert(qd); err != nil {
		return err
	}
	return nil
}
