package main

import (
	"log"
	"net/http"
	"os"

	mgo "gopkg.in/mgo.v2"

	"github.com/divyag9/goqueues/packages/config"
	"github.com/divyag9/goqueues/packages/handler"
	"github.com/divyag9/goqueues/packages/storage"
	"github.com/gorilla/context"
)

func main() {
	mongoDB := &storage.Mongo{}
	configDetails := validateAndSetConfigDetails()
	dbSession, err := storage.GetSession(mongoDB, configDetails)
	if err != nil {
		log.Fatal("Cannot connect to database")
	}
	switch sessionType := dbSession.(type) {
	case *mgo.Session:
		defer sessionType.Close()
	default:
		log.Fatalln("Failed to close database session")
	}
	// Handler for incoming request
	http.HandleFunc("/queues", DBHandler(handle, dbSession))
	// Start the server
	if err := http.ListenAndServeTLS(":443", configDetails.CertPath, configDetails.KeyPath, nil); err != nil {
		log.Fatal(err)
	}
}

// DBHandler is a wrapper for the underlying http handler
func DBHandler(fn http.HandlerFunc, dbsession interface{}) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		context.Set(r, "dbsession", dbsession)
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

func validateAndSetConfigDetails() *config.Details {
	certPath := os.Getenv("CERT_PATH")
	if certPath == "" {
		log.Fatal("Missing environment variable: CERT_PATH. Required environment variables :  CERT_PATH, KEY_PATH, DB_HOST, DB_USERNAME, DB_PASSWORD")
	}
	keyPath := os.Getenv("KEY_PATH")
	if keyPath == "" {
		log.Fatal("Missing environment variable: KEY_PATH. Required environment variables :  CERT_PATH, KEY_PATH, DB_HOST, DB_USERNAME, DB_PASSWORD")
	}
	dbHost := os.Getenv("DB_HOST")
	if dbHost == "" {
		log.Fatal("Missing environment variable: DB_HOST. Required environment variables :  CERT_PATH, KEY_PATH, DB_HOST, DB_USERNAME, DB_PASSWORD")
	}
	dbUsername := os.Getenv("DB_USERNAME")
	if dbUsername == "" {
		log.Fatal("Missing environment variable: DB_USERNAME. Required environment variables :  CERT_PATH, KEY_PATH, DB_HOST, DB_USERNAME, DB_PASSWORD")
	}
	dbPassword := os.Getenv("DB_PASSWORD")
	if dbPassword == "" {
		log.Fatal("Missing environment variable: DB_PASSWORD. Required environment variables :  CERT_PATH, KEY_PATH, DB_HOST, DB_USERNAME, DB_PASSWORD")
	}

	configDetails := &config.Details{}
	configDetails.CertPath = certPath
	configDetails.KeyPath = keyPath
	configDetails.DBHost = dbHost
	configDetails.DBUsername = dbUsername
	configDetails.DBPassword = dbPassword

	return configDetails
}
