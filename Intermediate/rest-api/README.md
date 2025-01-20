# REST API with CRUD Operations

A RESTful API to manage items, demonstrating basic CRUD operations in Go using the Gorilla Mux library.

## Features
- **Create**: Add a new item.
- **Read**: Retrieve all items or a specific item by ID.
- **Update**: Modify an existing item.
- **Delete**: Remove an item by ID.

## Requirements
- Go 1.20 or higher.
- Gorilla Mux library.

## Setup
1. Clone the repository:
   ```bash
   git clone https://github.com/arxel2468/Go-Projects.git
   cd Intermediate/rest-api
   go mod tidy
   go run main.go


## API Endpoints
GET /items: Retrieve all items.
GET /items/{id}: Retrieve a single item by ID.
POST /items: Create a new item.
PUT /items/{id}: Update an existing item.
DELETE /items/{id}: Delete an item by ID.


Example Usage
Add an item:
   ```bash
   curl -X POST -H "Content-Type: application/json" -d '{"name": "Laptop", "price": 1200}' http://localhost:8080/items
OR
   ```bash
   curl -X POST -H "Content-Type: application/json" -d "{\"name\": \"Laptop\", \"price\": 1200}" http://localhost:8080/items
View all items:
   ```bash
   curl http://localhost:8080/items

## Notes
The in-memory data store is reset every time the server restarts.
For production, connect to a database like PostgreSQL or MongoDB.