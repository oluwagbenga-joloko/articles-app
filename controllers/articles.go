package controllers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/oluwagbenga-joloko/articles-app/models"
	"github.com/oluwagbenga-joloko/articles-app/utils"
)

// CreateArticleHandler ...
func CreateArticleHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		var article models.Article
		err := json.NewDecoder(r.Body).Decode(&article)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "invalid json request")
			return
		}
		err = models.CreateArticle(db, &article)
		if err != nil {
			if inputE, ok := err.(*models.InputError); ok {
				utils.RespondWithError(w, http.StatusBadRequest, inputE.Message)
				return
			}
			utils.RespondWithError(w, http.StatusInternalServerError, "internal server error")
			return
		}
		utils.RespondWithJSON(w, http.StatusCreated, article)
	}

}

// ReadArticlesHandler ....
func ReadArticlesHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		q := r.URL.Query()
		var articles []models.Article
		err := models.GetArticles(db, &articles, q)
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, articles)
	}
}

//ReadArticleSingleArticleHandler ...
func ReadArticleSingleArticleHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {

		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusNotFound, "invalid article id")
			return
		}

		var article models.Article

		err = models.GetArticle(db, &article, id)
		if err == sql.ErrNoRows {
			utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("article with id %v not found", id))
			return
		}
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, article)
	}
}

// DeleteArticleSingleArticleHandler ...
func DeleteArticleSingleArticleHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusNotFound, "invalid article id")
			return
		}

		var article models.Article

		err = models.GetArticle(db, &article, id)
		if err == sql.ErrNoRows {
			utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("article with id %v not found", id))
			return
		}
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusInternalServerError, "internal server error")
			return
		}
		err = models.DeleteArticle(db, article.ID)
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusInternalServerError, "internal server error")
			return
		}
		utils.RespondWithJSON(w, http.StatusOK, map[string]string{"message": "article deleted"})
	}
}

//UpdateArticleHandler ...
func UpdateArticleHandler(db *sql.DB) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		id, err := strconv.Atoi(vars["id"])
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusNotFound, "invalid article id")
			return
		}

		var data map[string]interface{}
		err = json.NewDecoder(r.Body).Decode(&data)
		if err != nil {
			utils.RespondWithError(w, http.StatusBadRequest, "invalid json request")
			return
		}

		var article models.Article
		err = models.GetArticle(db, &article, id)
		if err == sql.ErrNoRows {
			utils.RespondWithError(w, http.StatusNotFound, fmt.Sprintf("article with id %v not found", id))
			return
		}
		if err != nil {
			fmt.Println(err)
			utils.RespondWithError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		err = models.UpdateArticle(db, &article, data)
		if err != nil {
			fmt.Println(err)
			if inputE, ok := err.(*models.InputError); ok {
				utils.RespondWithError(w, http.StatusBadRequest, inputE.Message)
				return
			}
			utils.RespondWithError(w, http.StatusInternalServerError, "internal server error")
			return
		}

		utils.RespondWithJSON(w, http.StatusOK, article)
	}

}
