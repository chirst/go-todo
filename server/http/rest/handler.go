package rest

import (
	"encoding/json"
	"net/http"
	"todo/adding"
	"todo/listing"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

// Handler ...
func Handler(listingService *listing.Service, addingService *adding.Service) http.Handler {
	router := chi.NewRouter()
	router.Use(middleware.Logger)
	router.Get("/todos", getTodos(listingService))
	router.Post("/todos", addTodo(addingService))
	return router
}

func getTodos(service *listing.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		todos := service.GetTodos()
		json.NewEncoder(w).Encode(todos)
	}
}

func addTodo(service *adding.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		newTodo := service.AddTodo()
		json.NewEncoder(w).Encode(newTodo)
	}
}
