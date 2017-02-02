package main

import (
	"log"
	"net/http"

	"github.com/divyag9/goqueues/packages/handler"
	"github.com/divyag9/goqueues/packages/storage"
	"github.com/gorilla/context"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	dbsession := storage.GetMongoDBSession()
	defer dbsession.Close()
	// Handler for incoming request
	http.HandleFunc("/", DBHandler(handle, dbsession))
	// Start the server
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}

// DBHandler is a wrapper for the underlying http handler
func DBHandler(fn http.HandlerFunc, dbsession interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		switch sessionType := dbsession.(type) {
		case *mgo.Session:
			dbsessionCopy := sessionType.Copy()
			defer dbsessionCopy.Close()
			context.Set(r, "dbsession", dbsessionCopy)
			fn(w, r)
		default:
			log.Fatalln("unknown session")
		}
	}
}

func handle(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		handler.HandleRead(w, r)
	case "POST":
		handler.HandleInsert(w, r)
	default:
		http.Error(w, "Not supported", http.StatusMethodNotAllowed)
	}
}
