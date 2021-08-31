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

func (s *mockTodoService) GetTodos(userID int) (todo.Todos, error) {
	ts := todo.Todos{}
	return ts, nil
}

func (s *mockTodoService) AddTodo(t todo.Todo) (*todo.Todo, error) {
	return todo.NewTodo(1, "gud todo", nil, 1)
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

func TestGetTodos(t *testing.T) {
	s := &mockTodoService{}
	token, _, _ := auth.GetTokenForUser(1)
	ctx := context.WithValue(context.Background(), jwtauth.TokenCtxKey, token)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("GET", "/todos", nil)
	r = r.WithContext(ctx)

	GetTodos(s)(w, r)
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
	r := httptest.NewRequest("GET", "/todos", buffer)
	r = r.WithContext(ctx)

	AddTodo(s)(w, r)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected %#v, got: %#v", http.StatusOK, resp.StatusCode)
	}
}

func TestCompleteTodo(t *testing.T) {
	s := &mockTodoService{}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PUT", "/todos/1/complete", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("todoID", "1")
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	token, _, _ := auth.GetTokenForUser(1)
	r = r.WithContext(context.WithValue(r.Context(), jwtauth.TokenCtxKey, token))

	CompleteTodo(s)(w, r)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected %#v, got: %#v", http.StatusOK, resp.StatusCode)
	}
}

func TestIncompleteTodo(t *testing.T) {
	s := &mockTodoService{}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("PATCH", "/todos/1/incomplete", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("todoID", "1")
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	token, _, _ := auth.GetTokenForUser(1)
	r = r.WithContext(context.WithValue(r.Context(), jwtauth.TokenCtxKey, token))

	IncompleteTodo(s)(w, r)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Errorf("expected %#v, got: %#v", http.StatusOK, resp.StatusCode)
	}
}

func TestDeleteTodo(t *testing.T) {
	s := &mockTodoService{}
	w := httptest.NewRecorder()
	r := httptest.NewRequest("DELETE", "/todos/1", nil)
	rctx := chi.NewRouteContext()
	rctx.URLParams.Add("todoID", "1")
	r = r.WithContext(context.WithValue(r.Context(), chi.RouteCtxKey, rctx))
	token, _, _ := auth.GetTokenForUser(1)
	r = r.WithContext(context.WithValue(r.Context(), jwtauth.TokenCtxKey, token))

	DeleteTodo(s)(w, r)
	resp := w.Result()

	if resp.StatusCode != http.StatusNoContent {
		t.Errorf("expected %#v, got: %#v", http.StatusNoContent, resp.StatusCode)
	}
}
