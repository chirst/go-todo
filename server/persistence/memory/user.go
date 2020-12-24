package memory

import (
	"todo/domain"
)

type User struct {
	ID       int64
	Name     string
	Password string
}

type UserStorage struct {
	users []User
}

func (s *UserStorage) AddUser(u domain.User) *domain.User {
	newUser := User{u.ID, u.Name, u.Password}
	s.users = append(s.users, newUser)
	return &domain.User{ID: u.ID, Name: u.Name, Password: u.Password}
}
