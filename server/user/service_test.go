package user

import (
	"testing"
	"todo/persistence/memory"
)

func TestAddUser(t *testing.T) {
	userStorage := new(memory.UserStorage)
	s := NewService(userStorage)
	newUser := s.AddUser("gud name", "1234")
	if newUser.Name != "gud name" {
		t.Errorf("got %v want %v", newUser.Name, "gud name")
	}
}
