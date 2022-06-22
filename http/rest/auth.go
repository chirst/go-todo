package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/chirst/go-todo/user"
)

type loginBody struct {
	Username string
	Password string
}

// Login returns an auth token for a valid login
func Login(s user.UserService) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		b := loginBody{}
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
			return
		}
		tokenString, err := s.GetUserTokenString(b.Username, b.Password)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
			return
		}
		w.Header().Set("Content-Type", "text/plain")
		err = json.NewEncoder(w).Encode(tokenString)
		if err != nil {
			log.Print(err.Error())
			http.Error(
				w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError,
			)
		}
	}
}
