package rest

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"
	"todo/user"
)

func TestAddUser(t *testing.T) {
	r := new(user.MemoryRepository)
	s := user.NewService(r)

	userBody := struct {
		Username string
		Password string
	}{
		"guduser",
		"1234",
	}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(userBody)
	req, err := http.NewRequest("POST", "/users", buffer)
	if err != nil {
		log.Fatal(err)
	}
	w := httptest.NewRecorder()

	AddUser(s)(w, req)

	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected %#v, got: %#v", http.StatusOK, resp.StatusCode)
	}
}
