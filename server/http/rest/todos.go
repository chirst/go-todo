package rest

import (
	"encoding/json"
	"net/http"
	"todo/adding"
	"todo/domain"
	"todo/listing"

	"github.com/go-chi/chi"
)

// Todos ...
func Todos(router chi.Router, listingService *listing.Service, addingService *adding.Service) {
	router.Get("/todos", getTodos(listingService))
	router.Post("/todos", addTodo(addingService))
}

func getTodos(service *listing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		todos := service.GetTodos()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	}
}

func addTodo(service *adding.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var bodyTodo domain.Todo
		err := decoder.Decode(&bodyTodo)
		if err != nil {
			http.Error(w, http.StatusText(400), 400)
			return
		}

		newTodo, err := service.AddTodo(bodyTodo)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newTodo)
	}
}
