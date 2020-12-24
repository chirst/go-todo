package rest

import (
	"encoding/json"
	"net/http"
	"todo/domain"
	"todo/todo"

	"github.com/go-chi/chi"
)

// Todos ...
func Todos(router chi.Router, todoService *todo.Service) {
	router.Get("/todos", getTodos(todoService))
	router.Post("/todos", addTodo(todoService))
}

func getTodos(service *todo.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		todos := service.GetTodos()
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	}
}

func addTodo(service *todo.Service) func(w http.ResponseWriter, r *http.Request) {
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
