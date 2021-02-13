package user

import (
	"testing"
)

func TestAddUser(t *testing.T) {
	userStorage := new(MemoryRepository)
	s := NewService(userStorage)
	newUser, err := s.AddUser("gud name", "1234")
	if err != nil {
		t.Errorf("got %v want no error", err)
	}
	if newUser.Username != "gud name" {
		t.Errorf("got %v want %v", newUser.Username, "gud name")
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
