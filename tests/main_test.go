package test

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

var app App

func TestMain(m *testing.M) {
	dbURL := "postgres://vmpnnoyk:AF5w7RkOJ9lmTivXvgR8Bsy6tllNnUSR@kandula.db.elephantsql.com:5432/vmpnnoyk"
	log.Println("Do stuff BEFORE the tests!")
	app = App{}
	dbURL := os.Getenv("TEST_DB_URL")
	app.initializeDb(dbURL)
	app.initializeRoutes()
	exitVal := m.Run()
	log.Println("Do stuff AFTER the tests!")
	os.Exit(exitVal)
}

func TestHomeResponse(t *testing.T) {

	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	app.router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
