package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/oluwagbenga-joloko/articles-app/models"
)

func TestCreateArticleSuccess(t *testing.T) {
	clearTable()
	payload, _ := json.Marshal(article1)
	req := httptest.NewRequest("POST", "/article", bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusCreated)
	}
	var resArticle models.Article
	json.Unmarshal(rr.Body.Bytes(), &resArticle)

	if article1.Title != resArticle.Title {
		t.Errorf("wrong article Title: got %v want %v",
			resArticle.Title, article1.Title)
	}
	if article1.Body != resArticle.Body {
		t.Errorf("wrong article body: got %v want %v",
			resArticle.Body, article1.Body)
	}
	if article1.Publisher != resArticle.Publisher {
		t.Errorf("wrong article publiser: got %v want %v",
			resArticle.Publisher, article1.Publisher)
	}
	if article1.Category != resArticle.Category {
		t.Errorf("wrong article publiser: got %v want %v",
			resArticle.Publisher, article1.Publisher)
	}
	if article1.PublishedAt != resArticle.PublishedAt {
		t.Errorf("wrong article publishedAT: got %v want %v",
			resArticle.PublishedAt, article1.PublishedAt)
	}
}

// copy maps only one level deep
func copy(m map[string]interface{}) map[string]interface{} {
	targetMap := make(map[string]interface{})
	for key, value := range m {
		targetMap[key] = value
	}
	return targetMap

}
func TestCreateArticleFaliure(t *testing.T) {
	clearTable()
	type Test struct {
		payload         []byte
		expectedStatus  int
		expectedMessage string
	}
	var tests = []Test{
		Test{[]byte("{"), http.StatusBadRequest, "invalid json request"},
	}
	article := map[string]interface{}{
		"title":     "give",
		"body":      "Give lady of they such",
		"category":  "comedy",
		"publisher": "Mary Jane",
	}
	for k := range article {
		invalidArticle := copy(article)
		delete(invalidArticle, k)
		payload, _ := json.Marshal(invalidArticle)
		test := Test{payload, http.StatusBadRequest, fmt.Sprintf("article %s is required and cannot be empty", k)}
		tests = append(tests, test)

	}

	for _, test := range tests {
		req := httptest.NewRequest("POST", "/article", bytes.NewBuffer(test.payload))
		rr := httptest.NewRecorder()
		app.Router.ServeHTTP(rr, req)
		if status := rr.Code; status != test.expectedStatus {
			t.Errorf("handler returned wrong status code: got %v want %v",
				status, test.expectedStatus)
		}
		var resData map[string]string
		json.Unmarshal(rr.Body.Bytes(), &resData)

		if actualMessage := resData["message"]; actualMessage != test.expectedMessage {
			t.Errorf("handler returned wrong message: got %v want %v",
				actualMessage, test.expectedMessage)
		}

	}

}

func TestGetArticleSuccess(t *testing.T) {
	clearTable()
	CreateArticle(&article1)
	req := httptest.NewRequest("GET", fmt.Sprintf("/article/%v", article1.ID), nil)
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var resArticle models.Article
	json.Unmarshal(rr.Body.Bytes(), &resArticle)

	if article1.Title != resArticle.Title {
		t.Errorf("wrong article name: got %v want %v",
			resArticle.Title, article1.Title)
	}
	if article1.Body != resArticle.Body {
		t.Errorf("wrong article body: got %v want %v",
			resArticle.Body, article1.Body)
	}
	if article1.Publisher != resArticle.Publisher {
		t.Errorf("wrong article publiser: got %v want %v",
			resArticle.Publisher, article1.Publisher)
	}
	if article1.Category != resArticle.Category {
		t.Errorf("wrong article publiser: got %v want %v",
			resArticle.Publisher, article1.Publisher)
	}
	if article1.PublishedAt != resArticle.PublishedAt {
		t.Errorf("wrong article publishedAT: got %v want %v",
			resArticle.PublishedAt, article1.PublishedAt)
	}
}

func TestGetArticlesSuccess(t *testing.T) {
	clearTable()
	articles := []models.Article{article1, article2}
	for _, article := range articles {
		CreateArticle(&article)

	}
	req := httptest.NewRequest("GET", fmt.Sprintf("/article"), nil)
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var resArticles []models.Article
	json.Unmarshal(rr.Body.Bytes(), &resArticles)
	if l := len(resArticles); l != 2 {
		t.Errorf("wrong number of articles: got %v want %v",
			l, 2)

	}

	for i, article := range articles {
		resArticle := resArticles[i]
		if article.Title != resArticle.Title {
			t.Errorf("wrong article name: got %v want %v",
				resArticle.Title, article1.Title)
		}
		if article.Body != resArticle.Body {
			t.Errorf("wrong article body: got %v want %v",
				resArticle.Body, article1.Body)
		}
		if article.Publisher != resArticle.Publisher {
			t.Errorf("wrong article publiser: got %v want %v",
				resArticle.Publisher, article1.Publisher)
		}
		if article.Category != resArticle.Category {
			t.Errorf("wrong article publiser: got %v want %v",
				resArticle.Publisher, article1.Publisher)
		}
		if article1.PublishedAt != resArticle.PublishedAt {
			t.Errorf("wrong article publishedAT: got %v want %v",
				resArticle.PublishedAt, article1.PublishedAt)
		}

	}

}

func TestUpdateArticleSuccess(t *testing.T) {
	clearTable()
	CreateArticle(&article1)
	updateArticle := article2
	updateArticle.PublishedAt = time.Now()
	payload, _ := json.Marshal(updateArticle)
	fmt.Printf(string(payload))
	req := httptest.NewRequest("PUT", fmt.Sprintf("/article/%v", article1.ID), bytes.NewBuffer(payload))
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var resArticle models.Article
	json.Unmarshal(rr.Body.Bytes(), &resArticle)

	if updateArticle.Title != resArticle.Title {
		t.Errorf("wrong article name: got %v want %v",
			resArticle.Title, article1.Title)
	}
	if updateArticle.Body != resArticle.Body {
		t.Errorf("wrong article body: got %v want %v",
			resArticle.Body, article1.Body)
	}
	if updateArticle.Publisher != resArticle.Publisher {
		t.Errorf("wrong article publiser: got %v want %v",
			resArticle.Publisher, article1.Publisher)
	}
	if updateArticle.Category != resArticle.Category {
		t.Errorf("wrong article category: got %v want %v",
			resArticle.Category, article1.Category)
	}
	if updateArticle.PublishedAt.Format(time.UnixDate) != resArticle.PublishedAt.Format(time.UnixDate) {
		t.Errorf("wrong article publishedAT: got %v want %v",
			resArticle.PublishedAt, updateArticle.PublishedAt)
	}
}

func TestDeleteArticleSuccess(t *testing.T) {
	clearTable()
	CreateArticle(&article1)
	req := httptest.NewRequest("DELETE", fmt.Sprintf("/article/%v", article1.ID), nil)
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code: got %v want %v",
			status, http.StatusOK)
	}
	var resData map[string]string
	json.Unmarshal(rr.Body.Bytes(), &resData)
	{
		expectedMessage := "article deleted"
		if actualMessage := resData["message"]; actualMessage != expectedMessage {
			t.Errorf("handler returned wrong status code: got %v want %v",
				actualMessage, expectedMessage)
		}
	}

}
