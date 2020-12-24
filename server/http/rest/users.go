package rest

import (
	"encoding/json"
	"net/http"
	"todo/user"

	"github.com/go-chi/chi"
)

func Users(router chi.Router, usersService *user.Service) {
	router.Post("/users", addUser(usersService))
}

func addUser(service *user.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		service.AddUser("gud", "1234")
		user := user.User{ID: 1, Name: "gud", Password: "1234"}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
