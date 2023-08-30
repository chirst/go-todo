package rest

import (
	"log"
	"net/http"

	"github.com/chirst/go-todo/auth"
)

func getUserID(r *http.Request) int {
	uid, ok := r.Context().Value(auth.UIDKey).(int)
	if !ok {
		log.Print("unable to assert user id is an int")
	}
	return uid
}
