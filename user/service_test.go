package user

import (
	"reflect"
	"testing"
)

func TestAddUser(t *testing.T) {
	userStorage := new(MemoryRepository)
	s := NewService(userStorage)

	tests := map[string]struct {
		input  string
		input2 string
		want   *User
		want2  error
	}{
		"adds": {
			input:  "gud name",
			input2: "1234",
			want:   &User{ID: 1, Username: "gud name", Password: "1234"},
			want2:  nil,
		},
		"no username": {
			input:  "",
			input2: "1234",
			want:   nil,
			want2:  ErrUsernameRequired,
		},
		"no password": {
			input:  "gud name",
			input2: "",
			want:   nil,
			want2:  ErrPasswordRequired,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, got2 := s.AddUser(tc.input, tc.input2)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected %#v, got: %#v", tc.want, got)
			}
			if tc.want2 != got2 {
				t.Fatalf("expected %#v, got %#v", tc.want2, got2)
			}
		})
	}
}

func TestGetUserTokenString(t *testing.T) {
	userStorage := new(MemoryRepository)
	s := NewService(userStorage)
	s.r.addUser(User{0, "gud", "1234"})
	tokenString, err := s.GetUserTokenString("gud", "1234")
	if err != nil {
		t.Errorf("got %v want no error", err)
	}
	if tokenString == nil {
		t.Errorf("got %v want token string", tokenString)
	}
}
