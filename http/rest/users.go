package rest

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/chirst/go-todo/user"
)

type userAdder interface {
	AddUser(u *user.User) (*user.User, error)
}

type addUserBody struct {
	Username string
	Password string
}

// AddUser adds a user
func AddUser(s userAdder) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := addUserBody{}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			log.Print(err.Error())
			http.Error(
				w,
				"Unable to decode request body",
				http.StatusBadRequest,
			)
			return
		}
		newUser, err := user.NewUser(0, body.Username, body.Password)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Unable to create user", http.StatusBadRequest)
			return
		}
		addedUser, err := s.AddUser(newUser)
		if err != nil {
			log.Print(err.Error())
			http.Error(w, "Unable to add user", http.StatusInternalServerError)
			return
		}
		jsonUser, err := addedUser.ToJSON()
		if err != nil {
			log.Print(err.Error())
			http.Error(
				w,
				"Unable to serialize added user",
				http.StatusInternalServerError,
			)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		_, err = w.Write(jsonUser)
		if err != nil {
			log.Print(err.Error())
			http.Error(
				w,
				"Unable to write response",
				http.StatusInternalServerError,
			)
		}
	}
}
