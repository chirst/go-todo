package main

import (
	"fmt"
	"net/http"
	"todo/auth"
	"todo/http/rest"
	"todo/todo"
	"todo/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	todosRepo := new(todo.MemoryRepository)
	usersRepo := new(user.MemoryRepository)
	todoService := todo.NewService(todosRepo)
	usersService := user.NewService(usersRepo)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// protected routes
	router.Group(func(r chi.Router) {
		r.Use(auth.Verifier)
		r.Use(auth.Authenticator)

		r.Get("/todos", rest.GetTodos(todoService))
		r.Post("/todos", rest.AddTodo(todoService))
	})

	// unprotected routes
	router.Group(func(r chi.Router) {
		r.Post("/users", rest.AddUser(usersService))
		r.Post("/login", rest.Login(usersService))
	})

	address := "localhost:3000"
	fmt.Printf("starting server on %s\n", address)
	http.ListenAndServe(address, router)
}
