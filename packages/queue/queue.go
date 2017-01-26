package queue

import (
	"time"
)

// Details represents the information about a queue
type Details struct {
	Name          string    `json:"name" bson:"name"`
	Type          string    `json:"type" bson:"type"`
	Depth         int64     `json:"depth" bson:"depth"`
	Rate          int64     `json:"rate" bson:"rate"`
	LastProcessed time.Time `json:"lastProcessed" bson:"last_processed"`
	LastReported  time.Time `bson:"last_reported"`
}
