package handler

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/divyag9/goqueues/packages/storage"
	"github.com/gorilla/context"

	"bytes"
)

func TestHandleInsert(t *testing.T) {
	var json = `{"name": "foo", "type": "bar", "depth": 1000, "rate": 10,"lastprocessed": "2008-09-17T20:04:26Z"}`
	// Create a request to pass to our handler.
	req, err := http.NewRequest("POST", "/", bytes.NewBufferString(json))
	if err != nil {
		t.Fatal(err)
	}
	dbsession := storage.GetMongoDBSession()
	defer dbsession.Close()
	context.Set(req, "dbsession", dbsession)
	// Create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleInsert)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}

	// Check the response body is what we expect.
	expected := `{"result":"success"}`
	if rr.Body.String() != expected {
		t.Errorf("handler returned unexpected body: got %v want %v",
			rr.Body.String(), expected)
	}
}

func TestHandleRead(t *testing.T) {
	// Create a request to pass to our handler.
	req, err := http.NewRequest("GET", "/", nil)
	if err != nil {
		t.Fatal(err)
	}
	dbsession := storage.GetMongoDBSession()
	defer dbsession.Close()
	context.Set(req, "dbsession", dbsession)
	// Create a ResponseRecorder to record the response.
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HandleRead)
	handler.ServeHTTP(rr, req)

	// Check the status code is what we expect.
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
