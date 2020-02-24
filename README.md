# Articles App

## Description

Server side implementation of the **Articles App**

Find the Database schema [here](https://www.lucidchart.com/invitations/accept/2c64a572-095d-421e-939d-7105c9c98e03)

Find Code documentation [here](https://godoc.org/github.com/oluwagbenga-joloko/articles-app)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Follow [here](https://golang.org/doc/install") to install Golang

### Installing

- clone the repository and navigate into the project directory

```bash
git clone git@github.com:oluwagbenga-joloko/articles-app.git && cd articles-app
```

- Copy the content of the `.env_sample` file into a `.env` file. Edit the file to relect your local settings.

```bash
cp .env.sample .env
```

- source the `.env` file

```bash
source .env
```

- build the application

```bash
make build
```

- run the applcation. The app should now be available from your browser at `http://127.0.0.1:8080`

```bash
make run
```

- run the tests

```bash
make test
```

## Endpoints

- GET /article/:id - Get a specific article using its id.

- GET /article - Get all articles. filter results by `category`, `publisher`, `created_at`, `published_at` using url query params

- PUT /article - Updates title, body, category, publisher, and published_at fields.

- POST /article - Creates a new article

- DELETE /article/:id - Delete an article

Below is an article example:

```json
{
  "title": "Lorem ipsum dolor sit amet",
  "body": "Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
  "category": "Lorem ipsum",
  "publisher": "John Doe",
  "created_at": "2020-02-22T22:39:22.716611Z",
  "updated_at": "2020-02-22T22:39:22.716611Z",
  "published_at": "2018-09-22T22:39:23.716611Z"
}
```

> Note: If a publisher does not exist, a new one is created and assigned to the article. The same applies to the category field.
