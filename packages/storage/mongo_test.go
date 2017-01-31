package storage

import (
	"testing"
	"time"

	"github.com/divyag9/goqueues/packages/queue"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

func TestGetAll(t *testing.T) {
	mongoDB := &Mongo{}
	dbsession, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		t.Fatalf("cannot dial mongo %v", err)
	}
	mongoDB.Session = dbsession
	defer dbsession.Close()

	_, err = mongoDB.GetAll()
	if err != nil {
		t.Errorf("Expected GetAll to return no error, returned error: %v ", err)
	}
}

func TestSave(t *testing.T) {
	mongoDB := &Mongo{}
	dbsession, err := mgo.Dial("mongodb://localhost")
	if err != nil {
		t.Fatalf("cannot dial mongo %v", err)
	}
	mongoDB.Session = dbsession
	queueDetails := queue.Details{ID: bson.NewObjectId(), Name: "foo", Type: "bar", Depth: 1000, Rate: 10, LastProcessed: time.Now(), LastReported: time.Now()}
	defer dbsession.Close()

	err = mongoDB.Save(&queueDetails)
	if err != nil {
		t.Errorf("Expected Save to return no error, returned error: %v", err)
	}
}
