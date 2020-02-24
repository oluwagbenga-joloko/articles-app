// Package models contains appliction models
package models

import (
	"time"
)

// Article represents an article
type Article struct {
	ID          int       `json:"id"`
	Title       string    `json:"title"`
	Body        string    `json:"body"`
	Category    string    `json:"category"`
	Publisher   string    `json:"publisher" `
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	PublishedAt time.Time `json:"published_at"`
}

// Publisher represents a publisher
type Publisher struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Category represents a category
type Category struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// Setup represents query string to create tables
var Setup = `
BEGIN;

CREATE TABLE IF NOT EXISTS publishers (
	id serial PRIMARY KEY,
	name VARCHAR(200) NOT NULL 
);
CREATE TABLE IF NOT EXISTS categories (
	id serial PRIMARY KEY,
	name VARCHAR(200) NOT NULL
);
CREATE TABLE IF NOT EXISTS articles (
	id serial PRIMARY KEY,
	title VARCHAR(200) NOT NULL,
	body TEXT NOT NULL,
	category_id integer NOT NULL REFERENCES categories(id) ON DELETE CASCADE,
	publisher_id integer NOT NULL REFERENCES publishers(id) ON DELETE CASCADE,
	created_at TIMESTAMP NOT NULL,
	updated_at TIMESTAMP NOT NULL,
	published_at TIMESTAMP
);
COMMIT;
`

// TearDown represents query string to drop tables
var TearDown = `
BEGIN;
DROP TABLE articles;
DROP TABLE categories;
DROP TABLE publishers;
COMMIT;
`

//ClearTables represents query string to clear data from tables
var ClearTables = `
BEGIN;
DELETE FROM articles;
DELETE FROM categories;
DELETE FROM publishers;
BEGIN;
`
