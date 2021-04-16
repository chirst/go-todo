package rest

import (
	"encoding/json"
	"net/http"
	"todo/auth"
	"todo/todo"
)

// GetTodos returns all todos
func GetTodos(service *todo.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := auth.GetUIDClaim(r.Context())
		todos := service.GetTodos(uid)
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	}
}

// AddTodo adds a todo
func AddTodo(service *todo.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var bodyTodo todo.Todo
		err := decoder.Decode(&bodyTodo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		uid := auth.GetUIDClaim(r.Context())
		bodyTodo.UserID = uid
		newTodo, err := service.AddTodo(bodyTodo)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newTodo)
	}
}
