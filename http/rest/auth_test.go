package rest

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestLogin(t *testing.T) {
	s := &mockUserService{}
	loginBody := loginBody{
		"guduser",
		"1234",
	}
	buffer := new(bytes.Buffer)
	json.NewEncoder(buffer).Encode(loginBody)
	w := httptest.NewRecorder()
	r := httptest.NewRequest("POST", "/login", buffer)

	Login(s)(w, r)
	resp := w.Result()

	if resp.StatusCode != http.StatusOK {
		t.Fatalf("expected %#v, got: %#v", http.StatusOK, resp.StatusCode)
	}
}
