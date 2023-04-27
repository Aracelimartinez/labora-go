package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
		"fmt"
		"strconv"
		"strings"
)

type Item struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
}

var Items []Item

func getItems(w http.ResponseWriter, r *http.Request) {
	// Función para obtener todos los elementos
	w.Header().Set("Content-Type", "application/json")

	pageParam := r.URL.Query().Get("page")
	itemsPerPageParam := r.URL.Query().Get("itemsPerPage")
	page, err := strconv.Atoi(pageParam)
	if err != nil {
		page = 1
	}

	itemsPerPage, err := strconv.Atoi(itemsPerPageParam)
	if err != nil {
		itemsPerPage = 10
	}

	start := (page - 1) * itemsPerPage
	end := start + itemsPerPage
	if end > len(Items) {
		end = len(Items)
	}

	json.NewEncoder(w).Encode(Items[start:end])
}

func getItem(w http.ResponseWriter, r *http.Request) {
	// Función para obtener un elemento específico
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idItem := params["id"]

	for _, item := range Items {
		if item.ID == idItem {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "ID no encontrado", http.StatusNotFound)
}

func searchItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	paramName := r.URL.Query()
	name, err := paramName["name"]
	if !err {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	var itemsFound []Item
	for _, item := range Items {
		if strings.EqualFold(item.Name, name[0]) {
			itemsFound = append(itemsFound, item)
		}
	}
	json.NewEncoder(w).Encode(itemsFound)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	// Función para crear un nuevo elemento
	w.Header().Set("Content-Type", "application/json")
	var item Item
	err := json.NewDecoder(r.Body).Decode(&item)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	Items = append(Items, item)
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(Items)
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	// Función para actualizar un elemento existente
	w.Header().Set("Content-Type", "application/json")
	var itemUpdated Item
	err := json.NewDecoder(r.Body).Decode(&itemUpdated)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	params := mux.Vars(r)
	idItem := params["id"]
	for i, item := range Items {
		if item.ID == idItem {
			Items[i] = itemUpdated
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(Items)
			return
		}
	}
	http.Error(w, "Elemento no encontrado", http.StatusNotFound)
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	// Función para eliminar un elemento
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	idItem:= params["id"]

	for i, item := range Items {
		if item.ID == idItem {
			Items = append(Items[:i], Items[i+1:]...)
			w.WriteHeader(http.StatusOK)
			json.NewEncoder(w).Encode(Items)
			return
		}
	}
	http.Error(w, "Elemento no encontrado", http.StatusNotFound)
}

func main() {
	for i := 1; i <= 50; i++ {
		Items = append(Items, Item{ID: fmt.Sprintf("%d", i), Name: fmt.Sprintf("Item %d", i)})
	}

	router := mux.NewRouter()

	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItem).Methods("GET")
	router.HandleFunc("/item", searchItem).Methods("GET")
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}
