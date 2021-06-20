package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"todo/auth"
	"todo/todo"
)

// GetTodos returns all todos
func GetTodos(service *todo.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := auth.GetUIDClaim(r.Context())
		todos, err := service.GetTodos(uid)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(todos)
	}
}

// AddTodo adds a todo
func AddTodo(service *todo.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var bt bodyTodo
		err := decoder.Decode(&bt)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		uid := auth.GetUIDClaim(r.Context())
		t, err := todo.NewTodo(0, bt.name, bt.completed, uid)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
		}
		newTodo, err := service.AddTodo(*t)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(newTodo)
	}
}

type bodyTodo struct {
	name      string
	completed *time.Time
}
