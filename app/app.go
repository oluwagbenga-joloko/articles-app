package application

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	// Add Postgres DB Driver
	_ "github.com/lib/pq"
	"github.com/oluwagbenga-joloko/articles-app/controllers"
	"github.com/oluwagbenga-joloko/articles-app/models"
)

// App ...
type App struct {
	DB     *sql.DB
	Router *mux.Router
}

// InitializeDb ...
func (app *App) InitializeDb(connStr string) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal(err)
	}
	_, err = db.Exec(models.Setup)
	if err != nil {
		log.Fatal(err)
	}
	app.DB = db
}

// InitializeRoutes ...
func (app *App) InitializeRoutes() {
	router := mux.NewRouter()
	router.HandleFunc("/", HomeHandler).Methods("GET")
	router.HandleFunc("/article", controllers.CreateArticleHandler(app.DB)).Methods("POST")
	router.HandleFunc("/article", controllers.ReadArticlesHandler(app.DB)).Methods("GET")
	router.HandleFunc("/article/{id}", controllers.ReadArticleSingleArticleHandler(app.DB)).Methods("GET")
	router.HandleFunc("/article/{id}", controllers.DeleteArticleSingleArticleHandler(app.DB)).Methods("DELETE")
	router.HandleFunc("/article/{id}", controllers.UpdateArticleHandler(app.DB)).Methods("PUT")
	app.Router = router
}

// HomeHandler ...
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{"message": "welcome to articles app"})

}
