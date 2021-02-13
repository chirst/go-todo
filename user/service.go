package user

import (
	"errors"
	"todo/auth"
)

// Repository for users
type Repository interface {
	addUser(User) *User
	getUserByName(string) *User
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
func (s *Service) AddUser(username, password string) (*User, error) {
	u, err := createUser(username, password)
	if err != nil {
		return nil, err
	}
	return s.r.addUser(*u), nil
}

// GetUserTokenString returns an auth token string for the first user matching the
// given username and password. It returns nil for anything invalid.
func (s *Service) GetUserTokenString(username, password string) (*string, error) {
	u := s.r.getUserByName(username)
	if u == nil {
		return nil, errors.New("user not found")
	}
	if u.Password != password {
		return nil, errors.New("password not matching")
	}
	_, tokenString, err := auth.GetTokenForUser(u.ID)
	if err != nil {
		return nil, errors.New("unable to generate token")
	}
	return &tokenString, nil
}
