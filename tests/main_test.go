package tests

import (
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
	"time"

	application "github.com/oluwagbenga-joloko/articles-app/app"
	"github.com/oluwagbenga-joloko/articles-app/models"
	"github.com/oluwagbenga-joloko/articles-app/repository"
)

var app application.App

var article1 = models.Article{
	Title:     "Uncommonly",
	Body:      "Full he none no side. Uncommonly surrounded considered for him are its. It we is read good soon",
	Category:  "drama",
	Publisher: "Mark Cane",
}
var article2 = models.Article{
	Title:     "Give lady",
	Body:      "Give lady of they such they sure it. Me contained explained my education. Vulgar as hearts by garret. ",
	Category:  "comedy",
	Publisher: "Mary Jane",
}
var article3 = models.Article{
	Title:       "favourable no",
	Body:        "Surrounded affronting favourable no mr. Lain knew like half she yet joy.",
	Category:    "comedy",
	Publisher:   "mr. Lain",
	PublishedAt: time.Now(),
}

func clearTable() {
	_, err := app.DB.Exec(models.ClearTables)
	if err != nil {
		log.Fatal(err)
	}

}

func CreateArticle(a *models.Article) {
	err := repository.CreateArticle(app.DB, a)
	if err != nil {
		log.Fatal(err)
	}
}

func TestMain(m *testing.M) {
	log.Println("Do stuff BEFORE the tests!")
	dbURL := os.Getenv("TEST_DB_URL")
	app.InitializeDb(dbURL)
	defer app.DB.Close()
	app.InitializeRoutes()
	exitVal := m.Run()
	_, err := app.DB.Exec(models.TearDown)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Do stuff AFTER the tests!")
	os.Exit(exitVal)
}

func TestHomeResponse(t *testing.T) {
	req := httptest.NewRequest("GET", "/", nil)
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
}
