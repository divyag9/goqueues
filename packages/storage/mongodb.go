package storage

import (
	"log"
	"time"

	"github.com/divyag9/goqueues/packages/config"
	"github.com/divyag9/goqueues/packages/queue"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// Mongo contains the mongo database session
type Mongo struct {
	Session *mgo.Session
}

// GetSession returns the mongoDB session to pass to the handler
func (m *Mongo) GetSession(config *config.Details) (interface{}, error) {
	// Get dial information for mongodb
	host := config.DBHost
	username := config.DBUsername
	password := config.DBPassword
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{host},
		Timeout:  60 * time.Second,
		Username: username,
		Password: password,
	}
	// Dial mongoDB
	dbsession, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		return nil, err
	}
	log.Println("connected to mongodb")
	return dbsession, nil
}

// Get returns the details of a particular queue
func (m *Mongo) Get(id int) (*queue.Details, error) {
	// implement logic
	return nil, nil
}

// GetAll returns the details of all queues
func (m *Mongo) GetAll() ([]*queue.Details, error) {
	var queueDetails []*queue.Details
	err := m.Session.DB("queues").C("details").
		Find(nil).All(&queueDetails)
	if err != nil {
		return nil, err
	}
	return queueDetails, nil
}

// Save saves the queue details
func (m *Mongo) Save(qd *queue.Details) error {
	changeInfo, err := m.Session.DB("queues").C("details").
		Upsert(bson.M{"name": qd.Name, "type": qd.Type}, qd)
	if err != nil {
		return err
	}
	log.Println("Number of documents updated:", changeInfo.Updated)
	return nil
}
