package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"todo/user"
)

type loginBody struct {
	Username string
	Password string
}

// Login returns an auth token for a valid login
func Login(userService *user.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		b := loginBody{}
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		tokenString, err := userService.GetUserTokenString(b.Username, b.Password)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusForbidden)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		json.NewEncoder(w).Encode(tokenString)
	}
}
