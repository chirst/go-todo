package user

import (
	"testing"
)

func TestNewUser(t *testing.T) {
	tests := map[string]struct {
		id       int64
		username string
		password string
		want     error
	}{
		"blank username": {
			id:       0,
			username: "",
			password: "1234",
			want:     errUsernameRequired,
		},
		"blank password": {
			id:       0,
			username: "name",
			password: "",
			want:     errPasswordRequired,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, got := NewUser(tc.id, tc.username, tc.password)
			if tc.want != got {
				t.Fatalf("expected %#v, got %#v", tc.want, got)
			}
		})
	}
}
