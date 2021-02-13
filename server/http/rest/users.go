package rest

import (
	"encoding/json"
	"net/http"
	"todo/user"

	"github.com/go-chi/chi"
)

// Users creates endpoints and handlers for users
func Users(router chi.Router, usersService *user.Service) {
	router.Post("/users", addUser(usersService))
}

func addUser(service *user.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		user, err := service.AddUser("gud", "1234")
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
