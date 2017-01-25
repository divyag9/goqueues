package main

import (
	"log"
	"net/http"

	"github.com/divyag9/goqueues/packages/database"
	"github.com/divyag9/goqueues/packages/queue"
	"github.com/gorilla/context"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	http.HandleFunc("/", DBHandler(handle))
	// start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// DBHandler is a wrapper for the underlying http handler
func DBHandler(fn http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		dbsession, err := mgo.Dial("localhost")
		if err != nil {
			log.Fatal("cannot dial mongo", err)
		}
		dbcopy := dbsession.Copy()
		defer dbcopy.Close()
		defer dbsession.Close()
		context.Set(r, "database", dbcopy)
		fn(w, r)
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handleRead(w, r)
	case "POST":
		handleInsert(w, r)
	default:
		http.Error(w, "Not supported", http.StatusMethodNotAllowed)
	}
}

func handleInsert(w http.ResponseWriter, r *http.Request) {
	mongoSession := context.Get(r, "database").(*mgo.Session)
	mongoDB := &database.Mongo{}
	mongoDB.Session = mongoSession
	testQueue := &queue.Details{}

	database.SaveQueueDetails(mongoDB, testQueue)
	// implement logic
}

func handleRead(w http.ResponseWriter, r *http.Request) {
	// implement logic
}
