# Articles App

## Description

Server side implementation of the **Articles App**

Find the Database schema [here](https://www.lucidchart.com/invitations/accept/2c64a572-095d-421e-939d-7105c9c98e03)

## Getting Started

These instructions will get you a copy of the project up and running on your local machine for development and testing purposes.

### Prerequisites

- Follow [here](https://golang.org/doc/install install go") to install Golang

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
