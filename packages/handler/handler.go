package handler

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	mgo "gopkg.in/mgo.v2"

	"github.com/divyag9/goqueues/packages/queue"
	"github.com/divyag9/goqueues/packages/storage"
	"github.com/gorilla/context"
)

// HandleInsert saves the queue details
func HandleInsert(w http.ResponseWriter, r *http.Request) {

	// Decode the request body into RequestDetails
	requestDetails := &queue.RequestDetails{}
	if err := json.NewDecoder(r.Body).Decode(requestDetails); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Set the queueDetails
	queueDetails := &queue.Details{}
	queueDetails.Name = requestDetails.Name
	queueDetails.Type = requestDetails.Type
	queueDetails.Depth = requestDetails.Depth
	queueDetails.Rate = requestDetails.Rate
	queueDetails.LastProcessed = requestDetails.LastProcessed
	queueDetails.LastReported = time.Now()

	// Get the dbsession and insert into the database
	dbsession := context.Get(r, "dbsession")
	insertFunction := insertQueueDetails(queueDetails)
	if err := executeOperation(dbsession, insertFunction); err != nil {
		http.Error(w, fmt.Sprintf("Error occured while saving queue details: %q", err.Error()), 100)
		return
	}

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"result":"success"}`))
}

// HandleRead retrieves the queue details
func HandleRead(w http.ResponseWriter, r *http.Request) {
	// Get the storage db and retrieve queue details
	dbsession := context.Get(r, "dbsession")
	var queueDetails []*queue.Details
	retrieveFunction := retrieveQueueDetails(&queueDetails)
	err := executeOperation(dbsession, retrieveFunction)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error occured while retrieving details: %q", err.Error()), 101)
		return
	}

	// Send encoded response
	if err := json.NewEncoder(w).Encode(queueDetails); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func insertQueueDetails(queueDetails *queue.Details) func(mongo *storage.Mongo) error {
	return func(mongo *storage.Mongo) error {
		// Insert it into the database
		err := storage.SaveQueueDetails(mongo, queueDetails)
		if err != nil {
			return err
		}
		return nil
	}
}

func retrieveQueueDetails(queueDetails *[]*queue.Details) func(mongo *storage.Mongo) error {
	return func(mongo *storage.Mongo) error {
		var detailsReturned []*queue.Details
		var err error
		// Retrieve queue details
		detailsReturned, err = storage.GetAllQueueDetails(mongo)
		if err != nil {
			return err
		}
		(*queueDetails) = detailsReturned
		return nil
	}
}

func executeOperation(dbsession interface{}, operation func(*storage.Mongo) error) error {
	switch sessionType := dbsession.(type) {
	case *mgo.Session:
		dbsessionCopy := sessionType.Copy()
		mongo := &storage.Mongo{}
		mongo.Session = dbsessionCopy
		defer dbsessionCopy.Close()
		err := operation(mongo)
		if err != nil {
			return err
		}
	default:
		return errors.New("Unknown database session type")
	}
	return nil
}
