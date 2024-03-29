package rest

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/chirst/go-todo/user"
)

type mockUserAdder struct{}

func (s *mockUserAdder) AddUser(_ *user.User) (*user.User, error) {
	u, err := user.NewUser(1, "guduser", "1234")
	if err != nil {
		panic(err.Error())
	}
	return u, nil
}

func TestAddUser(t *testing.T) {
	s := &mockUserAdder{}
	w := httptest.NewRecorder()
	body := []byte(`{
		"username": "guduser",
		"password": "1234"
	}`)
	r := httptest.NewRequest("POST", "/users", bytes.NewBuffer(body))

	AddUser(s)(w, r)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("got %d, expected: %d", resp.StatusCode, http.StatusOK)
	}
}
