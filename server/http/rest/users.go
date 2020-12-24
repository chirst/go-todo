package rest

import (
	"encoding/json"
	"net/http"
	"todo/domain"
	"todo/user"

	"github.com/go-chi/chi"
)

// Users creates endpoints and handlers for users
func Users(router chi.Router, usersService *user.Service) {
	router.Post("/users", addUser(usersService))
}

func addUser(service *user.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		service.AddUser("gud", "1234")
		user := domain.User{ID: 1, Name: "gud", Password: "1234"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
