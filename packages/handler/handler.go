package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/divyag9/goqueues/packages/queue"
	"github.com/divyag9/goqueues/packages/storage"
	"github.com/gorilla/context"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

// HandleInsert saves the queue details
func HandleInsert(w http.ResponseWriter, r *http.Request) {
	mongoSession := context.Get(r, "database").(*mgo.Session)
	mongoDB := &storage.Mongo{}
	mongoDB.Session = mongoSession

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
	// Give the queue details a unique ID
	queueDetails.ID = bson.NewObjectId()
	queueDetails.LastReported = time.Now()

	// Insert it into the database
	if err := storage.SaveQueueDetails(mongoDB, queueDetails); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Saved queue details")

	// Send response
	w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(`{"result":"success"}`))
}

// HandleRead retrieves the queue details
func HandleRead(w http.ResponseWriter, r *http.Request) {
	mongoSession := context.Get(r, "database").(*mgo.Session)
	mongoDB := &storage.Mongo{}
	mongoDB.Session = mongoSession

	// Retrieve queue details
	queueDetails, err := storage.GetAllQueueDetails(mongoDB)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	log.Println("Retrieved queue details")

	// Send encoded response
	if err := json.NewEncoder(w).Encode(queueDetails); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
