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

type todoGetter interface {
	GetTodos(userID int) (todo.Todos, error)
}

// GetTodos returns all todos belonging to the current user
func GetTodos(s todoGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		uid, err := auth.GetUIDClaim(r.Context())
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Unable to get user", http.StatusUnauthorized)
			return
		}
		todos, err := s.GetTodos(*uid)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Unable to get todos", http.StatusInternalServerError)
			return
		}
		jsonTodos, err := todos.ToJSON()
		if err != nil {
			log.Print(err.Error())
			http.Error(
				w,
				"Unable to serialize todos",
				http.StatusInternalServerError,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonTodos)
		if err != nil {
			log.Print(err.Error())
			http.Error(
				w,
				"Unable to write response",
				http.StatusInternalServerError,
			)
		}
	}
}

type todoAdder interface {
	AddTodo(
		name string,
		completed *time.Time,
		userID int,
		priorityID *int,
	) (*todo.Todo, error)
}

type addTodoBody struct {
	Name      string
	Completed *time.Time
	Priority  *int
}

// AddTodo adds a todo for the current user
func AddTodo(s todoAdder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		bt := addTodoBody{}
		if err := json.NewDecoder(r.Body).Decode(&bt); err != nil {
			log.Print(err.Error())
			http.Error(
				w,
				"Unable to decode the request body",
				http.StatusBadRequest,
			)
			return
		}
		uid, err := auth.GetUIDClaim(r.Context())
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Unable to get user", http.StatusUnauthorized)
			return
		}
		addedTodo, err := s.AddTodo(bt.Name, bt.Completed, *uid, bt.Priority)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Unable to add todo", http.StatusInternalServerError)
			return
		}
		jsonTodo, err := addedTodo.ToJSON()
		if err != nil {
			log.Print(err.Error())
			http.Error(
				w,
				"Unable to serialize added todo",
				http.StatusInternalServerError,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonTodo)
		if err != nil {
			log.Print(err.Error())
			http.Error(
				w,
				"Unable to write response",
				http.StatusInternalServerError,
			)
		}
	}
}

type todoPatcher interface {
	UpdatePriority(userID, todoID, priorityID int) error
	CompleteTodo(userID, todoID int) error
	IncompleteTodo(userID, todoID int) error
	ChangeTodoName(userID, todoID int, name string) error
}

type patchTodoBody struct {
	Complete   *bool
	Name       *string
	PriorityID *int
}

// PatchTodo updates the given optional fields of patchTodoBody
func PatchTodo(s todoPatcher) func( // nolint:cyclop // TODO
	w http.ResponseWriter,
	r *http.Request,
) {
	return func(w http.ResponseWriter, r *http.Request) {
		id := chi.URLParam(r, "todoID")
		todoID, err := strconv.Atoi(id)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Unable to read todo id", http.StatusBadRequest)
			return
		}

		uid, err := auth.GetUIDClaim(r.Context())
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Unable to get user", http.StatusUnauthorized)
			return
		}

		body := patchTodoBody{}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			log.Print(err.Error())
			http.Error(w, "Unable to read body", http.StatusBadRequest)
			return
		}

		if body.Complete != nil {
			if *body.Complete {
				err = s.CompleteTodo(*uid, todoID)
			} else {
				err = s.IncompleteTodo(*uid, todoID)
			}
			if err != nil {
				log.Print(err.Error())
				http.Error(
					w,
					"Unable to change completion status",
					http.StatusInternalServerError,
				)
			}
		}

		if body.Name != nil {
			err = s.ChangeTodoName(*uid, todoID, *body.Name)
			if err != nil {
				log.Print(err.Error())
				http.Error(
					w,
					"Unable to change todo name",
					http.StatusInternalServerError,
				)
			}
		}

		if body.PriorityID != nil {
			err = s.UpdatePriority(*uid, todoID, *body.PriorityID)
			if err != nil {
				log.Print(err.Error())
				http.Error(
					w,
					"Unable to change todo priority",
					http.StatusInternalServerError,
				)
			}
		}
	}
}

type todoDeleter interface {
	DeleteTodo(userID, todoID int) error
}

// DeleteTodo deletes todo with the given id
func DeleteTodo(s todoDeleter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		todoID := chi.URLParam(r, "todoID")
		id, err := strconv.Atoi(todoID)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "unable to delete todo", http.StatusBadRequest)
			return
		}
		uid, err := auth.GetUIDClaim(r.Context())
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Unable to get user", http.StatusUnauthorized)
			return
		}
		err = s.DeleteTodo(*uid, id)
		if err != nil {
			log.Print(err.Error())
			http.Error(
				w,
				"unable to delete todo",
				http.StatusInternalServerError,
			)
			return
		}
		w.WriteHeader(http.StatusNoContent)
		w.Header().Set("Content-Type", "text/plain")
	}
}

type priorityGetter interface {
	GetPriorities() (todo.Priorities, error)
}

// GetPriorities returns all possible priorities.
func GetPriorities(s priorityGetter) func(
	w http.ResponseWriter,
	r *http.Request,
) {
	return func(w http.ResponseWriter, r *http.Request) {
		ps, err := s.GetPriorities()
		if err != nil {
			log.Print(err.Error())
			http.Error(
				w,
				"unable to get priorities",
				http.StatusInternalServerError,
			)
			return
		}
		jsonPriorities, err := ps.ToJSON()
		if err != nil {
			log.Print(err.Error())
			http.Error(
				w,
				"unable to serialize priorities",
				http.StatusInternalServerError,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonPriorities)
		if err != nil {
			log.Print(err.Error())
			http.Error(
				w,
				"Unable to write response",
				http.StatusInternalServerError,
			)
		}
	}
}
