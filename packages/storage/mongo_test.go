package storage

import (
	"os"
	"testing"
	"time"

	"github.com/divyag9/goqueues/packages/config"
	"github.com/divyag9/goqueues/packages/queue"
	"gopkg.in/mgo.v2"
)

var configDetails = config.Details{CertPath: os.Getenv("CERT_PATH"), KeyPath: os.Getenv("KEY_PATH"), DBHost: os.Getenv("DB_HOST"), DBUsername: os.Getenv("DB_USERNAME"), DBPassword: os.Getenv("DB_PASSWORD")}

func TestGetSession(t *testing.T) {
	mongoDB := &Mongo{}
	_, err := mongoDB.GetSession(&configDetails)
	if err != nil {
		t.Errorf("Expected GetSession to return no error, returned error: %v ", err)
	}
}

func TestGetAll(t *testing.T) {
	mongoDB := &Mongo{}
	dbsession, _ := mongoDB.GetSession(&configDetails)
	mongoSession := dbsession.(*mgo.Session)
	mongoDB.Session = mongoSession
	defer mongoSession.Close()

	_, err := mongoDB.GetAll()
	if err != nil {
		t.Errorf("Expected GetAll to return no error, returned error: %v ", err)
	}
}

func TestSave(t *testing.T) {
	queueDetails := queue.Details{Name: "foo", Type: "bar", Depth: 1000, Rate: 10, LastProcessed: time.Now(), LastReported: time.Now()}
	mongoDB := &Mongo{}
	dbsession, _ := mongoDB.GetSession(&configDetails)
	mongoSession := dbsession.(*mgo.Session)
	mongoDB.Session = mongoSession
	defer mongoSession.Close()

	err := mongoDB.Save(&queueDetails)
	if err != nil {
		t.Errorf("Expected Save to return no error, returned error: %v", err)
	}
}
