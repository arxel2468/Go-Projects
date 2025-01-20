package main

import (
	"log"
	"net/http"
	"rest-api/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Define routes
	r.HandleFunc("/items", handlers.GetItems).Methods("GET")
	r.HandleFunc("/items/{id}", handlers.GetItem).Methods("GET")
	r.HandleFunc("/items", handlers.CreateItem).Methods("POST")
	r.HandleFunc("/items/{id}", handlers.UpdateItem).Methods("PUT")
	r.HandleFunc("/items/{id}", handlers.DeleteItem).Methods("DELETE")

	// Start server
	log.Println("Starting server on :8080...")
	log.Fatal(http.ListenAndServe(":8080", r))
}
