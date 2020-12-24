package memory

import (
	"todo/user"
)

type User struct {
	ID       int64
	Name     string
	Password string
}

type UserStorage struct {
	users []User
}

func (s *UserStorage) AddUser(u user.User) *user.User {
	newUser := User{u.ID, u.Name, u.Password}
	s.users = append(s.users, newUser)
	return &user.User{ID: u.ID, Name: u.Name, Password: u.Password}
}
