package main

import (
	"log"
	"net/http"

	"github.com/divyag9/goqueues/packages/handler"
	"github.com/gorilla/context"
	mgo "gopkg.in/mgo.v2"
)

func main() {
	// Get the mongodb session and pass to the handler
	dbsession, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		log.Fatalln("cannot dial mongo ", err)
	}
	log.Println("connected to mongodb")
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
		dbcopy := dbsession.(*mgo.Session).Copy()
		defer dbcopy.Close()
		context.Set(r, "database", dbcopy)
		fn(w, r)
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
