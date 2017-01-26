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
