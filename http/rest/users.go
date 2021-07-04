package rest

import (
	"encoding/json"
	"log"
	"net/http"
	"todo/user"
)

type addUserBody struct {
	Username string
	Password string
}

// AddUser adds a user
func AddUser(service *user.Service) func(w http.ResponseWriter, r *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		body := addUserBody{}
		if err := json.NewDecoder(r.Body).Decode(&body); err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		newUser, err := user.NewUser(0, body.Username, body.Password)
		addedUser, err := service.AddUser(newUser)
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		decodedUser, err := addedUser.ToJSON()
		if err != nil {
			log.Print(err)
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Header().Set("Content-Type", "application/json")
		w.Write(decodedUser)
	}
}
