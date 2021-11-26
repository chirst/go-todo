package rest

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	w := httptest.NewRecorder()
	body := []byte(`{
		"username": "guduser",
		"password": "1234"
	}`)
	r := httptest.NewRequest("POST", "/login", bytes.NewBuffer(body))
	s := &mockUserService{}

	Login(s)(w, r)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("got %d, expected: %d", resp.StatusCode, http.StatusOK)
	}
}
