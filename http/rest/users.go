package rest

import (
	"encoding/json"
	"net/http"
	"todo/user"
)

// AddUser adds a user
func AddUser(service *user.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		b := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{}
		json.NewDecoder(r.Body).Decode(&b)
		user, err := service.AddUser(b.Username, b.Password)
		if err != nil {
			http.Error(w, err.Error(), 400)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}
