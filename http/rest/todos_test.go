package rest

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chirst/go-todo/auth"
	"github.com/chirst/go-todo/todo"

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
	buffer := new(bytes.Buffer)
	todoBody := addTodoBody{
		Name: "gud name",
	}
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

// func TestCompleteTodo(t *testing.T) {
// 	r := new(todo.MemoryRepository)
// 	s := todo.NewService(r)

// 	todo, err := todo.NewTodo(0, "incomplete todo", nil, 1)
// 	if err != nil {
// 		t.Fatalf("failed to create new todo")
// 	}
// 	_, err = s.AddTodo(*todo)
// 	if err != nil {
// 		t.Fatalf("failed to add todo")
// 	}

// 	token, _, _ := auth.GetTokenForUser(1)
// 	ctx := context.WithValue(context.Background(), jwtauth.TokenCtxKey, token)
// 	buffer := new(bytes.Buffer)
// 	req, err := http.NewRequestWithContext(ctx, "PUT", "/todos/1/complete", buffer)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	w := httptest.NewRecorder()

// 	CompleteTodo(s)(w, req)

// 	resp := w.Result()

// 	if resp.StatusCode != http.StatusOK {
// 		t.Fatalf("expected %#v, got: %#v", http.StatusOK, resp.StatusCode)
// 	}
// }
