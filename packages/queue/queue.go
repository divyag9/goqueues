package queue

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Details represents the information about a queue
type Details struct {
	ID            bson.ObjectId `bson:"_id"`
	Name          string        `bson:"name"`
	Type          string        `bson:"type"`
	Depth         int64         `bson:"depth"`
	Rate          int64         `bson:"rate"`
	LastProcessed time.Time     `bson:"last_processed"`
	LastReported  time.Time     `bson:"last_reported"`
}

// RequestDetails represents incoming request information about a queue
type RequestDetails struct {
	Name          string    `json:"name"`
	Type          string    `json:"type"`
	Depth         int64     `json:"depth"`
	Rate          int64     `json:"rate"`
	LastProcessed time.Time `json:"lastprocessed"`
}
