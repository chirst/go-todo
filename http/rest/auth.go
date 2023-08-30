package rest

import (
	"encoding/json"
	"log"
	"net/http"
)

type tokenGetter interface {
	GetUserTokenString(username, password string) (*string, error)
}

type loginBody struct {
	Username string
	Password string
}

type loginResponse struct {
	Jwt string `json:"jwt"`
}

// Login returns an auth token for a valid login
func Login(s tokenGetter) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		b := loginBody{}
		err := json.NewDecoder(r.Body).Decode(&b)
		if err != nil {
			log.Print(err.Error())
			http.Error(
				w,
				http.StatusText(http.StatusBadRequest),
				http.StatusBadRequest,
			)
			return
		}
		tokenString, err := s.GetUserTokenString(b.Username, b.Password)
		if err != nil {
			log.Print(err.Error())
			http.Error(
				w,
				http.StatusText(http.StatusUnauthorized),
				http.StatusUnauthorized,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		err = json.NewEncoder(w).Encode(loginResponse{
			Jwt: *tokenString,
		})
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
