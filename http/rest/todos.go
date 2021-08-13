package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/chirst/go-todo/auth"
	"github.com/chirst/go-todo/todo"
	"github.com/go-chi/chi"
)

// GetTodos returns all todos belonging to the current user
func GetTodos(s todo.TodoService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		uid := auth.GetUIDClaim(r.Context())
		todos, err := s.GetTodos(uid)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Unable to get todos", http.StatusInternalServerError)
			return
		}
		jsonTodos, err := todos.ToJSON()
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Unable to serialize todos", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonTodos)
	}
}

type addTodoBody struct {
	Name      string
	Completed *time.Time
}

// AddTodo adds a todo for the current user
func AddTodo(s todo.TodoService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bt := addTodoBody{}
		if err := json.NewDecoder(r.Body).Decode(&bt); err != nil {
			log.Print(err.Error())
			http.Error(w, "Unable to decode the request body", http.StatusBadRequest)
			return
		}
		uid := auth.GetUIDClaim(r.Context())
		t, err := todo.NewTodo(0, bt.Name, bt.Completed, uid)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Unable to create todo", http.StatusBadRequest)
			return
		}
		addedTodo, err := s.AddTodo(*t)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Unable to add todo", http.StatusInternalServerError)
			return
		}
		jsonTodo, err := addedTodo.ToJSON()
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Unable to serialize added todo", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(jsonTodo)
	}
}

func CompleteTodo(s todo.TodoService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		todoID := chi.URLParam(r, "todoID")
		id, err := strconv.Atoi(todoID)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Unable to complete todo", http.StatusBadRequest)
			return
		}
		uid := auth.GetUIDClaim(r.Context())
		err = s.CompleteTodo(uid, id)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Unable to complete todo", http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
	}
}
