// Package repository contains functions to manipulate data in DB
package repository

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/url"
	"strings"
	"time"

	"github.com/oluwagbenga-joloko/articles-app/models"
)

// InputError represents an input error condition
type InputError struct {
	Message string
}

func (e *InputError) Error() string {
	return e.Message
}

// GetOrCreatePublisher gets publisher from DB if it exists and creates it if not.
// It returns error
func GetOrCreatePublisher(db *sql.DB, p *models.Publisher, name string) error {
	err := db.QueryRow("SELECT id FROM publishers p WHERE p.name = $1 ", name).Scan(&p.ID)
	if err == sql.ErrNoRows {
		err = db.QueryRow("INSERT into publishers(name) VALUES($1) RETURNING id", name).Scan(&p.ID)
	}
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// GetOrCreateCategory gets category from DB if it exists and creates it if not.
// It returns any error that occurs
func GetOrCreateCategory(db *sql.DB, c *models.Category, name string) error {
	err := db.QueryRow("SELECT id FROM categories p WHERE p.name = $1 ", name).Scan(&c.ID)
	if err == sql.ErrNoRows {
		err = db.QueryRow("INSERT into categories(name) VALUES($1) RETURNING id", name).Scan(&c.ID)
	}
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

// GetArticle gets Article from DB.
// It returns any error that occurs
func GetArticle(db *sql.DB, article *models.Article, id int) error {
	qString := `SELECT articles.id, articles.title, articles.body, categories.name, publishers.name, articles.created_at, articles.updated_at, articles.published_at
		FROM articles, categories, publishers
		WHERE publishers.id = articles.publisher_id AND categories.id = articles.category_id AND articles.id = $1`

	err := db.QueryRow(qString, id).Scan(&article.ID, &article.Title, &article.Body, &article.Category, &article.Publisher, &article.CreatedAt, &article.UpdatedAt, &article.PublishedAt)
	if err != nil {
		return err
	}
	return nil
}

// DeleteArticle deletes Article from DB.
// It returns any error that occurs
func DeleteArticle(db *sql.DB, id int) error {
	res, err := db.Exec("Delete from articles WHERE articles.id = $1", id)
	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil

}

// UpdateArticle updates Article in DB with data.
// It returns any error that occurs
func UpdateArticle(db *sql.DB, article *models.Article, data map[string]interface{}) error {
	updateableFields := []string{"title", "body", "published_at", "publisher", "category"}

	var val []interface{}
	var col []string

	num := 1
	for _, field := range updateableFields {
		if field == "publisher" {
			if v, ok := data["publisher"]; ok == true {
				if v == "" {
					return &InputError{"publisher body cannot be empty"}

				}

				var publisher models.Publisher
				err := GetOrCreatePublisher(db, &publisher, fmt.Sprintf("%v", data["publisher"]))
				if err != nil {
					return err
				}
				val = append(val, publisher.ID)
				col = append(col, fmt.Sprintf("%s = $%v", "publisher_id", num))
				article.Publisher = publisher.Name
				num++
			}

		} else if field == "category" {
			if v, ok := data["category"]; ok == true {
				if v == "" {
					return &InputError{"category body cannot be empty"}
				}
				var category models.Category
				err := GetOrCreateCategory(db, &category, fmt.Sprintf("%v", data["category"]))
				if err != nil {
					return err
				}
				val = append(val, category.ID)
				col = append(col, fmt.Sprintf("%s = $%v", "category_id", num))
				article.Category = category.Name
				num++
			}
		} else {
			if v, ok := data[field]; ok == true {
				if v == "" {
					return &InputError{fmt.Sprintf("%s cannot be empty", field)}
				}
				if v != "" {
					val = append(val, data[field])
					col = append(col, fmt.Sprintf("%s = $%v", field, num))
					num++
				}
			}

		}
	}
	bytes, err := json.Marshal(data)
	if err != nil {
		return err
	}
	json.Unmarshal(bytes, &article)
	qString := "UPDATE articles SET " + strings.Join(col, ", ") + fmt.Sprintf(" WHERE id = $%v", num)
	val = append(val, article.ID)

	res, err := db.Exec(qString, val...)

	if err != nil {
		return err
	}
	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

// CreateArticle creates new Article and saves in DB
// It returns any error that occurs
func CreateArticle(db *sql.DB, article *models.Article) error {
	if article.Body == "" {
		return &InputError{"article body is required and cannot be empty"}
	}
	if article.Title == "" {
		return &InputError{"article title is required and cannot be empty"}
	}

	if article.Publisher == "" {
		return &InputError{"article publisher is required and cannot be empty"}
	}
	if article.Category == "" {
		return &InputError{"article category is required and cannot be empty"}
	}

	var publisher models.Publisher
	err := GetOrCreatePublisher(db, &publisher, article.Publisher)
	if err != nil {
		return err
	}
	var category models.Category
	err = GetOrCreateCategory(db, &category, article.Category)
	if err != nil {
		return err
	}

	row := db.QueryRow("INSERT INTO articles(title, body, created_at, updated_at, publisher_id, category_id, published_at) VALUES($1, $2, $3, $4, $5, $6, $7) RETURNING id",
		article.Title, article.Body, time.Now(), time.Now(), publisher.ID, category.ID, article.PublishedAt)
	err = row.Scan(&article.ID)
	if err != nil {
		return err
	}
	return nil
}

// GetArticles Retrives article from DB
// It returns any error that occurs
func GetArticles(db *sql.DB, articles *[]models.Article, q url.Values) error {
	queryParams := map[string]string{"category": "categories.name", "publisher": "publishers.name", "created_at": "created_at", "published_at": "published_at"}
	var whereVal []interface{}
	var whereCol []string = []string{""}

	// move to util
	num := 1

	for key, dbKey := range queryParams {
		paramValue := q.Get(key)
		if paramValue != "" {
			whereVal = append(whereVal, paramValue)
			whereCol = append(whereCol, fmt.Sprintf(" %s = $%v", dbKey, num))
			num++
		}
	}
	var article models.Article
	qString := `SELECT articles.id, articles.title, articles.body, categories.name, publishers.name, articles.created_at, articles.updated_at, articles.published_at
		FROM articles, categories, publishers
		WHERE publishers.id = articles.publisher_id AND categories.id = articles.category_id`

	rows, err := db.Query(qString+strings.Join(whereCol, " AND "), whereVal...)
	defer rows.Close()
	if err != nil {
		return err
	}
	for rows.Next() {
		err = rows.Scan(&article.ID, &article.Title, &article.Body, &article.Category, &article.Publisher, &article.CreatedAt, &article.UpdatedAt, &article.PublishedAt)
		if err != nil {
			return err
		}
		*articles = append(*articles, article)
	}
	return nil
}
