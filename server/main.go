package main

import (
	"fmt"
	"net/http"
	"todo/adding"
	"todo/auth"
	"todo/http/rest"
	"todo/listing"
	"todo/persistence/memory"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

func main() {
	todosRepo := new(memory.Storage)
	listingService := listing.NewService(todosRepo)
	addingService := adding.NewService(todosRepo)

	router := chi.NewRouter()

	// protected routes
	router.Group(func(r chi.Router) {
		r.Use(middleware.Logger)
		r.Use(auth.Verifier)
		r.Use(auth.Authenticator)

		rest.Todos(r, listingService, addingService)
	})

	// unprotected routes
	router.Group(func(r chi.Router) {
		rest.Users(r)
	})

	fmt.Println("started server on localhost:3000")
	http.ListenAndServe("localhost:3000", router)
}
