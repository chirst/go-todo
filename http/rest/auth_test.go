package rest

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

type mockTokenGetter struct{}

func (s *mockTokenGetter) GetUserTokenString(
	username,
	password string,
) (*string, error) {
	ts := "asdf33890fjxl;aksd"
	return &ts, nil
}

func TestLogin(t *testing.T) {
	w := httptest.NewRecorder()
	body := []byte(`{
		"username": "guduser",
		"password": "1234"
	}`)
	r := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	s := &mockTokenGetter{}

	Login(s)(w, r)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("got %d, expected: %d", resp.StatusCode, http.StatusOK)
	}
}
