package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
	"todo/auth"
	"todo/todo"
)

type bodyTodo struct {
	Name      string
	Completed *time.Time
}

type responseTodo struct {
	Id        int64      `json:"id"`
	Name      string     `json:"name"`
	Completed *time.Time `json:"completed"`
	UserID    int64      `json:"userId"`
}

func toResponse(t *todo.Todo) *responseTodo {
	return &responseTodo{
		Id:        t.ID(),
		Name:      t.Name(),
		Completed: t.Completed(),
		UserID:    t.UserID(),
	}
}

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
		var todosResponse []responseTodo
		for _, todo := range todos {
			todosResponse = append(todosResponse, *toResponse(todo))
		}
		json.NewEncoder(w).Encode(todosResponse)
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
		t, err := todo.NewTodo(0, bt.Name, bt.Completed, uid)
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
		json.NewEncoder(w).Encode(toResponse(newTodo))
	}
}
