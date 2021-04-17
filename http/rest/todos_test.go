package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"todo/auth"
	"todo/todo"

	"github.com/go-chi/jwtauth"
)

func TestGetTodos(t *testing.T) {
	r := new(todo.MemoryRepository)
	s := todo.NewService(r)

	token, _, _ := auth.GetTokenForUser(1)
	ctx := context.WithValue(context.Background(), jwtauth.TokenCtxKey, token)
	req, err := http.NewRequestWithContext(ctx, "GET", "/todos", nil)
	if err != nil {
		log.Fatal(err)
	}
	w := httptest.NewRecorder()

	GetTodos(s)(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected %#v, got: %#v", http.StatusOK, resp.StatusCode)
	}
}

func TestAddTodo(t *testing.T) {
	r := new(todo.MemoryRepository)
	s := todo.NewService(r)

	token, _, _ := auth.GetTokenForUser(1)
	ctx := context.WithValue(context.Background(), jwtauth.TokenCtxKey, token)
	todoBody := todo.Todo{
		ID:     1,
		Name:   "gud name",
		UserID: 1,
	}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(todoBody)
	req, err := http.NewRequestWithContext(ctx, "GET", "/todos", buffer)
	if err != nil {
		log.Fatal(err)
	}
	w := httptest.NewRecorder()

	AddTodo(s)(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected %#v, got: %#v", http.StatusOK, resp.StatusCode)
	}
}
