package user

import (
	"todo/domain"
)

type Repository interface {
	AddUser(domain.User) *domain.User
}

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) AddUser(name, password string) *domain.User {
	u := domain.User{ID: 1, Name: name, Password: password}
	return s.r.AddUser(u)
}
