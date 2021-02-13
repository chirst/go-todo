package rest

import (
	"encoding/json"
	"net/http"
	"todo/user"
)

// AddUser adds a user
func AddUser(service *user.Service) func(w http.ResponseWriter, r *http.Request) {
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
