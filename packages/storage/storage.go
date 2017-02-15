package storage

import (
	"github.com/divyag9/goqueues/packages/config"
	"github.com/divyag9/goqueues/packages/queue"
)

// Details contains storage information
type Details struct {
}

// Database contains methods for accessing the database information
type Database interface {
	GetSession(*config.Details) (interface{}, error)
}

// DataAccessor contains methods for accessing the queue information from the database
type DataAccessor interface {
	Get(int) (*queue.Details, error)
	GetAll() ([]*queue.Details, error)
}

// DataSaver contains methods for saving the queue information to the database
type DataSaver interface {
	Save(*queue.Details) error
}

// GetSession returns database session
func GetSession(db Database, config *config.Details) (interface{}, error) {
	return db.GetSession(config)
}

// GetQueueDetailsByID returns the details of a particular queue
func GetQueueDetailsByID(da DataAccessor, id int) (*queue.Details, error) {
	return da.Get(id)
}

// GetAllQueueDetails returns the details of all queues
func GetAllQueueDetails(da DataAccessor) ([]*queue.Details, error) {
	return da.GetAll()
}

// SaveQueueDetails saves the queue details
func SaveQueueDetails(ds DataSaver, qd *queue.Details) error {
	return ds.Save(qd)
}
