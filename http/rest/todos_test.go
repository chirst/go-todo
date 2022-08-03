package rest

import (
	"bytes"
	"context"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/chirst/go-todo/auth"
	"github.com/chirst/go-todo/todo"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

type mockTodoService struct{}

func (s *mockTodoService) GetTodos(userID int) (todo.Todos, error) {
	ts := todo.Todos{}
	return ts, nil
}

func (s *mockTodoService) AddTodo(
	name string,
	completed *time.Time,
	userID int,
	priorityID *int,
) (*todo.Todo, error) {
	return todo.MustNewTodo(1, "gud todo", nil, 1), nil
}

func (s *mockTodoService) CompleteTodo(userID int, todoID int) error {
	return nil
}

func (s *mockTodoService) IncompleteTodo(userID int, todoID int) error {
	return nil
}

func (s *mockTodoService) DeleteTodo(userID, todoID int) error {
	return nil
}

func (s *mockTodoService) ChangeTodoName(
	userID,
	todoID int,
	name string,
) error {
	return nil
}

func (s *mockTodoService) GetPriorities() (todo.Priorities, error) {
	return todo.Priorities{}, nil
}

func (s *mockTodoService) UpdatePriority(userID, todoID, priorityID int) error {
	return nil
}

func TestGetTodos(t *testing.T) {
	token, _, _ := auth.GetTokenForUser(1)
	ctx := context.WithValue(context.Background(), jwtauth.TokenCtxKey, token)
	r := httptest.NewRequest("GET", "/todos", nil).WithContext(ctx)
	w := httptest.NewRecorder()
	s := &mockTodoService{}

	GetTodos(s)(w, r)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("got %d, expected: %d", resp.StatusCode, http.StatusOK)
	}
}

func TestAddTodo(t *testing.T) {
	token, _, _ := auth.GetTokenForUser(1)
	ctx := context.WithValue(context.Background(), jwtauth.TokenCtxKey, token)
	todoBody := []byte(`{
		"name": "gud name"
	}`)
	r := httptest.NewRequest(
		"GET",
		"/todos",
		bytes.NewBuffer(todoBody),
	).WithContext(ctx)
	w := httptest.NewRecorder()
	s := &mockTodoService{}

	AddTodo(s)(w, r)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("got %d, expected: %d", resp.StatusCode, http.StatusOK)
	}
}

func TestPatchTodo(t *testing.T) {
	t.Run("all", func(t *testing.T) {
		s := &mockTodoService{}
		w := httptest.NewRecorder()
		body := []byte(`{
			"complete": true,
			"name": "new name",
			"priorityId": 2
		}`)
		r := httptest.NewRequest("PATCH", "/todos/1", bytes.NewBuffer(body))
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("todoID", "1")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
		token, _, _ := auth.GetTokenForUser(1)
		r = r.WithContext(
			context.WithValue(r.Context(), jwtauth.TokenCtxKey, token),
		)

		PatchTodo(s)(w, r)
		resp := w.Result()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("got %d, expected: %d", resp.StatusCode, http.StatusOK)
		}
	})

	t.Run("without a key", func(t *testing.T) {
		s := &mockTodoService{}
		w := httptest.NewRecorder()
		body := []byte(`{
			"name": "new name"
		}`)
		r := httptest.NewRequest(
			"PATCH",
			"/todos/1/complete",
			bytes.NewBuffer(body),
		)
		rctx := chi.NewRouteContext()
		rctx.URLParams.Add("todoID", "1")
		r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
		token, _, _ := auth.GetTokenForUser(1)
		r = r.WithContext(
			context.WithValue(r.Context(), jwtauth.TokenCtxKey, token),
		)

		PatchTodo(s)(w, r)
		resp := w.Result()

		if resp.StatusCode != http.StatusOK {
			t.Fatalf("got %d, expected: %d", resp.StatusCode, http.StatusOK)
		}
	})
}

func TestDeleteTodo(t *testing.T) {
	s := &mockTodoService{}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/todos/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("todoID", "1")
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	token, _, _ := auth.GetTokenForUser(1)
	r = r.WithContext(
		context.WithValue(r.Context(), jwtauth.TokenCtxKey, token),
	)

	DeleteTodo(s)(w, r)
	resp := w.Result()

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("got %d, expected: %d", resp.StatusCode, http.StatusNoContent)
	}
}

func TestGetPriorities(t *testing.T) {
	s := &mockTodoService{}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/priorities", nil)

	GetPriorities(s)(w, r)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("got %d, expected: %d", resp.StatusCode, http.StatusOK)
	}
}
