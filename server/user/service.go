package user

import (
	"todo/domain"
)

// Repository for users
type Repository interface {
	AddUser(domain.User) *domain.User
}

// Service for users
type Service struct {
	r Repository
}

// NewService creates an instance of this service
func NewService(r Repository) *Service {
	return &Service{r}
}

// AddUser validates, creates, and adds the user to persistence
func (s *Service) AddUser(name, password string) *domain.User {
	u := domain.User{ID: 1, Name: name, Password: password}
	return s.r.AddUser(u)
}
