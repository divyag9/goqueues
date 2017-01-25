package database

import (
	"github.com/divyag9/goqueues/packages/queue"
)

// DataAccessor contains methods for accessing the queue information from the database
type DataAccessor interface {
	Get(int) (*queue.Details, error)
	GetAll() ([]*queue.Details, error)
}

// DataSaver contains methods for saving the queue information to the database
type DataSaver interface {
	Save(*queue.Details) error
}

// GetQueueDetailsByID returns the details of a particular queue
func GetQueueDetailsByID(da DataAccessor, id int) (*queue.Details, error) {
	return da.Get(id)
}

// GetAllQueueDetails returns the details of all queue
func GetAllQueueDetails(da DataAccessor) ([]*queue.Details, error) {
	return da.GetAll()
}

// SaveQueueDetails saves the queue details
func SaveQueueDetails(ds DataSaver, qd *queue.Details) error {
	return ds.Save(qd)
}
