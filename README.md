# Go Web API with Docker & SQLite

A simple Go-based REST API with SQLite, containerized using Docker and set up with GitHub Actions for CI/CD.

---

## Features

- REST API in Go
- SQLite database support
- Dockerized for easy deployment
- CI/CD pipeline via GitHub Actions
- Docker Compose for orchestration
- Unit tested with Go's `testing` package [2 features]

---

## Run Locally

### 1. Clone the repo

```bash
git clone https://github.com/AymanMagdy/books-store-api.git
cd books-store-api
```

### 2. Run with Docker Compose

```bash
docker-compose up --build
```

API will be running at [http://localhost:8080](http://localhost:8080)

---

## Run Tests

```bash
go test
```

---

## CI/CD (GitHub Actions)

This repo uses GitHub Actions to:

- Run Go tests
- Build the Docker image on push/PR

Workflow file: `.github/workflows/go-docker-ci.yml`

---

## Build Docker Image Manually

```bash
docker build -t go-docker-app .
docker run -p 8080:8080 go-docker-app
```

---

## Sample API Endpoints

| Method | Endpoint         | Description        |
|--------|------------------|--------------------|
| GET    | `/books`         | List all books     |
| GET    | `/books/{id}`    | Get book by ID     |
| POST   | `/books`         | Create new book    |
| PUT   | `/books/{id}`     | Edit a book        |
| DELETE | `/books/{id}`    | Delete book        |

---

## Run Script Example

```bash
#!/bin/bash

set -e

echo "Building Docker image..."
docker-compose build

echo "Starting containers..."
docker-compose up -d

echo "Waiting for the server to start..."
sleep 3

echo "Sending test request to http://localhost:8080 ..."
response=$(curl -s -o /dev/null -w "%{http_code}" http://localhost:8080)

if [ "$response" -eq 200 ]; then
  echo "running!"
else
  echo "failed. HTTP status: $response"
fi
```