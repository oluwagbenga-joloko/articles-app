package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"github.com/oluwagbenga-joloko/articles-app/controllers"
	"github.com/oluwagbenga-joloko/articles-app/models"
)

// App ...
type App struct {
	db     *sql.DB
	router *mux.Router
}

func (app *App) initializeDb(connStr string) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(models.Setup)
	if err != nil {
		log.Fatal(err)
	}
	app.db = db
}

func (app *App) initializeRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/article", controllers.CreateArticleHandler(app.db)).Methods("POST")
	router.HandleFunc("/article", controllers.ReadArticlesHandler(app.db)).Methods("GET")
	router.HandleFunc("/article/{id}", controllers.ReadArticleSingleArticleHandler(app.db)).Methods("GET")
	router.HandleFunc("/article/{id}", controllers.DeleteArticleSingleArticleHandler(app.db)).Methods("DELETE")
	router.HandleFunc("/article/{id}", controllers.UpdateArticleHandler(app.db)).Methods("PUT")
	app.router = router

}

func main() {
	dbURL := os.Getenv("DB_URL")
	app := App{}
	app.initializeDb(dbURL)
	app.initializeRoutes()
	fmt.Println("server runing on port 8080")
	defer app.db.Close()
	log.Fatal(http.ListenAndServe(":8080", app.router))

}

// HomeHandler ...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "welcome to articles app"})

}
