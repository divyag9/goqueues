package queue

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

// Details represents the information about a queue
type Details struct {
	ID            bson.ObjectId `json:"id" bson:"_id"`
	Name          string        `json:"name" bson:"name"`
	Type          string        `json:"type" bson:"type"`
	Depth         int64         `json:"depth" bson:"depth"`
	Rate          int64         `json:"rate" bson:"rate"`
	LastProcessed time.Time     `json:"lastprocessed" bson:"last_processed"`
	LastReported  time.Time     `json:"lastreported" bson:"last_reported"`
}

// RequestDetails represents incoming request information about a queue
type RequestDetails struct {
	Name          string    `json:"name" bson:"name"`
	Type          string    `json:"type" bson:"type"`
	Depth         int64     `json:"depth" bson:"depth"`
	Rate          int64     `json:"rate" bson:"rate"`
	LastProcessed time.Time `json:"lastprocessed" bson:"last_processed"`
}
