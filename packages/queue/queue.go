package queue

import (
	"time"
)

// Details represents the information about a queue
type Details struct {
	Name          string
	Type          string
	Depth         int64
	Rate          int64
	LastProcessed time.Time
	LastReported  time.Time
}

// Saver interface has the save method for saving the queue information sent
type Saver interface {
	Save() error
}

// SaveDetails calls the QueueInfoSaver interface method to save the queue information
func SaveDetails(s Saver) error {
	return s.Save()
}
