package rest

import (
	"encoding/json"
	"net/http"
	"todo/user"
)

type loginBody struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Login returns an auth token for a valid login
func Login(userService *user.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		b := struct {
			Username string `json:"username"`
			Password string `json:"password"`
		}{"", ""}
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		tokenString, err := userService.GetUserTokenString(b.Username, b.Password)
		if err != nil {
			http.Error(w, http.StatusText(http.StatusForbidden), http.StatusForbidden)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tokenString)
	}
}
