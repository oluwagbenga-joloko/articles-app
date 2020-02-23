package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	application "github.com/oluwagbenga-joloko/articles-app/app"
)

func main() {
	dbURL := os.Getenv("DB_URL")
	app := application.App{}
	app.InitializeDb(dbURL)
	app.InitializeRoutes()
	fmt.Println("server runing on port 8080")
	defer app.DB.Close()
	log.Fatal(http.ListenAndServe(":8080", app.Router))
}
