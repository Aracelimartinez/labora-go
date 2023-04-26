package main

import (
    "encoding/json"
    "net/http"
    "github.com/gorilla/mux"
)

type Item struct {
	ID      string `json:"id"`
	Name    string `json:"name"`
}

var items = []Item{
	{
		ID:   "1",
		Name: "Item 1",
	},
	{
		ID:   "2",
		Name: "Item 2",
	},
	{
		ID:   "3",
		Name: "Item 3",
	},
	{
		ID:   "4",
		Name: "Item 4",
	},
	{
		ID:   "5",
		Name: "Item 5",
	},
}

func getItems(w http.ResponseWriter, r *http.Request) {
	// Función para obtener todos los elementos
	json.NewEncoder(w).Encode(items)
}

func getItem(w http.ResponseWriter, r *http.Request) {
	// Función para obtener un elemento específico
	params := mux.Vars(r)
	idItem, err := params["id"]

	for _, item := range items {
		if item.ID == idItem {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	http.Error(w, "ID no encontrado", http.StatusNotFound)
}

func createItem(w http.ResponseWriter, r *http.Request) {
	// Función para crear un nuevo elemento
}

func updateItem(w http.ResponseWriter, r *http.Request) {
	// Función para actualizar un elemento existente
}

func deleteItem(w http.ResponseWriter, r *http.Request) {
	// Función para eliminar un elemento
}

func main() {
	router := mux.NewRouter()
	//Tarea.
	router.HandleFunc("/items", getItems).Methods("GET")
	router.HandleFunc("/items/{id}", getItem).Methods("GET")

	//Más adelante.
	router.HandleFunc("/items", createItem).Methods("POST")
	router.HandleFunc("/items/{id}", updateItem).Methods("PUT")
	router.HandleFunc("/items/{id}", deleteItem).Methods("DELETE")

	http.ListenAndServe(":8000", router)
}
