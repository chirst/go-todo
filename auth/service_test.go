package auth

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestAuthMiddleware(t *testing.T) {
	next := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uid, ok := r.Context().Value(UIDKey).(int)
		if !ok {
			t.Errorf("type assertion is not ok")
		}
		if uid != 1 {
			t.Errorf("expected uid to be 1, got %d", uid)
		}
	})
	w := httptest.NewRecorder()
	_, bearer, _ := GetTokenForUser(1)
	r := httptest.NewRequest("GET", "/resource", nil)
	r.Header.Set("Authorization", "BEARER "+bearer)

	ah := Authenticator(next)
	Verifier(ah).ServeHTTP(w, r)

	if w.Code != http.StatusOK {
		t.Errorf("expected %d, got %d", http.StatusOK, w.Code)
	}
}
