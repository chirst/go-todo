package user

import (
	"errors"
	"reflect"
	"testing"
)

func TestAddUser(t *testing.T) {
	userStorage := new(MemoryRepository)
	s := NewService(userStorage)

	u, err := NewUser(0, "gud name", "1234")
	if err != nil {
		t.Fatalf("error creating user")
	}

	tests := map[string]struct {
		input *User
		want  *User
		want2 error
	}{
		"adds": {
			input: u,
			want:  &User{id: 1, username: "gud name"},
			want2: nil,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, got2 := s.AddUser(tc.input)
			if got != nil && tc.want != nil {
				if tc.want.id != got.id || tc.want.username != got.username {
					t.Fatalf("expected %#v, got: %#v", tc.want, got)
				}
			}
			if !errors.Is(tc.want2, got2) {
				t.Fatalf("expected %#v, got %#v", tc.want2, got2)
			}
		})
	}
}

func TestAddUserHashesPassword(t *testing.T) {
	userStorage := new(MemoryRepository)
	s := NewService(userStorage)

	username := "guduser"
	password := "wordpass"
	nu, err := NewUser(0, username, password)
	if err != nil {
		t.Errorf("unable to create user")
	}
	_, err = s.AddUser(nu)
	if err != nil {
		t.Errorf(err.Error())
	}

	u, err := userStorage.getUserByName(username)
	if err != nil {
		t.Errorf("unable to retrieve user by name")
	}
	if u.password == password {
		t.Errorf("expected %s to be hashed and not match %s", u.password, password)
	}
}

func TestGetUserTokenString(t *testing.T) {
	userStorage := new(MemoryRepository)
	s := NewService(userStorage)

	u, err := NewUser(0, "gud name", "1234")
	if err != nil {
		t.Fatalf("error creating user")
	}
	_, err = s.AddUser(u)
	if err != nil {
		t.Fatalf("error adding user")
	}

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
			want2:  errPasswordNotMatching,
		},
		"user not found": {
			input:  "wut name",
			input2: "1234",
			want:   nil,
			want2:  errUserNotFound,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, got2 := s.GetUserTokenString(tc.input, tc.input2)
			if reflect.TypeOf(tc.want) != reflect.TypeOf(got) {
				t.Fatalf("expected %#v, got %#v", tc.want, got)
			}
			if !errors.Is(tc.want2, got2) {
				t.Fatalf("expected %#v, got %#v", tc.want2, got2)
			}
		})
	}
}
