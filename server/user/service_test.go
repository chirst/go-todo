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
	if newUser.Name != "gud name" {
		t.Errorf("got %v want %v", newUser.Name, "gud name")
	}
}
