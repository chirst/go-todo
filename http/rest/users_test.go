package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chirst/go-todo/user"
)

type mockUserService struct{}

func (s *mockUserService) AddUser(_ *user.User) (*user.User, error) {
	u, err := user.NewUser(1, "guduser", "1234")
	if err != nil {
		panic(err.Error())
	}
	return u, nil
}

func (s *mockUserService) GetUserTokenString(username, password string) (*string, error) {
	ts := "asdf33890fjxl;aksd"
	return &ts, nil
}

func TestAddUser(t *testing.T) {
	s := &mockUserService{}
	userBody := addUserBody{
		"guduser",
		"1234",
	}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(userBody)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/users", buffer)

	AddUser(s)(w, r)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected %#v, got: %#v", http.StatusOK, resp.StatusCode)
	}
}
