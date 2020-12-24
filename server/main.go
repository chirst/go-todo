package main

import (
	"fmt"
	"net/http"
	"todo/adding"
	"todo/auth"
	"todo/http/rest"
	"todo/listing"
	"todo/persistence/memory"
	"todo/user"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	todosRepo := new(memory.TodoStorage)
	usersRepo := new(memory.UserStorage)
	listingService := listing.NewService(todosRepo)
	addingService := adding.NewService(todosRepo)
	usersService := user.NewService(usersRepo)

	router := chi.NewRouter()
	router.Use(middleware.Logger)

	// protected routes
	router.Group(func(r chi.Router) {
		r.Use(auth.Verifier)
		r.Use(auth.Authenticator)

		rest.Todos(r, listingService, addingService)
	})

	// unprotected routes
	router.Group(func(r chi.Router) {
		rest.Users(r, usersService)
	})

	address := "localhost:3000"
	fmt.Printf("starting server on %s\n", address)
	http.ListenAndServe(address, router)
}
