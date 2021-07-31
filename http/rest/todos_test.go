package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chirst/go-todo/auth"
	"github.com/chirst/go-todo/todo"

	"github.com/go-chi/chi"
	"github.com/go-chi/jwtauth"
)

type mockTodoService struct{}

func (s *mockTodoService) CompleteTodo(todoID int64) error {
	return nil
}

func (s *mockTodoService) GetTodos(userID int64) (todo.Todos, error) {
	ts := todo.Todos{}
	return ts, nil
}

func (s *mockTodoService) AddTodo(t todo.Todo) (*todo.Todo, error) {
	retTodo, err := todo.NewTodo(1, "gud todo", nil, 1)
	if err != nil {
		panic(err.Error())
	}
	return retTodo, nil
}

func TestGetTodos(t *testing.T) {
	s := &mockTodoService{}

	token, _, _ := auth.GetTokenForUser(1)
	ctx := context.WithValue(context.Background(), jwtauth.TokenCtxKey, token)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/todos", nil)
	req = req.WithContext(ctx)

	GetTodos(s)(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected %#v, got: %#v", http.StatusOK, resp.StatusCode)
	}
}

func TestAddTodo(t *testing.T) {
	s := &mockTodoService{}

	token, _, _ := auth.GetTokenForUser(1)
	ctx := context.WithValue(context.Background(), jwtauth.TokenCtxKey, token)

	buffer := new(bytes.Buffer)
	todoBody := addTodoBody{
		Name: "gud name",
	}
	json.NewEncoder(buffer).Encode(todoBody)

	w := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/todos", buffer)
	req = req.WithContext(ctx)

	AddTodo(s)(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected %#v, got: %#v", http.StatusOK, resp.StatusCode)
	}
}

func TestCompleteTodo(t *testing.T) {
	s := &mockTodoService{}

	w := httptest.NewRecorder()
	req := httptest.NewRequest("PUT", "/todos/1/complete", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("todoID", "1")
	req = req.WithContext(context.WithValue(req.Context(), chi.RouteCtxKey, rctx))

	CompleteTodo(s)(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected %#v, got: %#v", http.StatusOK, resp.StatusCode)
	}
}
