package handlers

import (
	"encoding/json"
	"net/http"
	"rest-api/models"
	"strconv"

	"github.com/gorilla/mux"
)

// In-memory data store
var items = make(map[int]models.Item)
var idCounter = 1

// Get all items
func GetItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	itemList := make([]models.Item, 0, len(items))
	for _, item := range items {
		itemList = append(itemList, item)
	}
	json.NewEncoder(w).Encode(itemList)
}

// Get a single item by ID
func GetItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if item, exists := items[id]; exists {
		json.NewEncoder(w).Encode(item)
	} else {
		http.Error(w, "Item not found", http.StatusNotFound)
	}
}

// Create a new item
func CreateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var item models.Item
	_ = json.NewDecoder(r.Body).Decode(&item)

	item.ID = idCounter
	items[idCounter] = item
	idCounter++

	json.NewEncoder(w).Encode(item)
}

// Update an existing item
func UpdateItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if _, exists := items[id]; exists {
		var updatedItem models.Item
		_ = json.NewDecoder(r.Body).Decode(&updatedItem)
		updatedItem.ID = id
		items[id] = updatedItem
		json.NewEncoder(w).Encode(updatedItem)
	} else {
		http.Error(w, "Item not found", http.StatusNotFound)
	}
}

// Delete an item by ID
func DeleteItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	id, _ := strconv.Atoi(params["id"])

	if _, exists := items[id]; exists {
		delete(items, id)
		w.WriteHeader(http.StatusNoContent)
	} else {
		http.Error(w, "Item not found", http.StatusNotFound)
	}
}
