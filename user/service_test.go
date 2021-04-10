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
			want:   &User{ID: 1, Username: "gud name"},
			want2:  nil,
		},
		"add existing": {
			input:  "gud name",
			input2: "1234",
			want:   nil,
			want2:  ErrUserExists,
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

			if got != nil {
				if tc.want.ID != got.ID || tc.want.Username != got.Username {
					t.Fatalf("expected %#v, got: %#v", tc.want, got)
				}
				if tc.input2 == got.Password {
					t.Fatalf("expected hashed password, got %#v", got.Password)
				}
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

	s.AddUser("gud name", "1234")

	tests := map[string]struct {
		input  string
		input2 string
		want   *string
		want2  error
	}{
		"happy path": {
			input:  "gud name",
			input2: "1234",
			want:   new(string),
			want2:  nil,
		},
		"password not matching": {
			input:  "gud name",
			input2: "123",
			want:   nil,
			want2:  ErrPasswordNotMatching,
		},
		"user not found": {
			input:  "wut name",
			input2: "1234",
			want:   nil,
			want2:  ErrUserNotFound,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, got2 := s.GetUserTokenString(tc.input, tc.input2)
			if reflect.TypeOf(tc.want) != reflect.TypeOf(got) {
				t.Fatalf("expected %#v, got %#v", tc.want, got)
			}
			if tc.want2 != got2 {
				t.Fatalf("expected %#v, got %#v", tc.want2, got2)
			}
		})
	}
}
