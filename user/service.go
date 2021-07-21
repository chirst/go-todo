package user

import (
	"errors"

	"github.com/chirst/go-todo/auth"
)

var (
	errUsernameRequired    = errors.New("username required")
	errPasswordRequired    = errors.New("password required")
	errPasswordNotMatching = errors.New("password not matching")
	errTokenGeneration     = errors.New("token generation error")
)

// Repository for users
type Repository interface {
	addUser(User) (*User, error)
	getUserByName(string) (*User, error)
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
func (s *Service) AddUser(u *User) (*User, error) {
	// TODO: test password != input password
	hashedPassword, err := auth.GenerateFromPassword(u.password)
	if err != nil {
		return nil, err
	}
	u.password = *hashedPassword
	return s.r.addUser(*u)
}

// GetUserTokenString returns an auth token string for the first user matching
// the given username and password. It returns nil for anything invalid.
func (s *Service) GetUserTokenString(username, password string) (*string, error) {
	u, err := s.r.getUserByName(username)
	if err != nil {
		return nil, err
	}
	if auth.CompareHashAndPassword(u.password, password) != nil {
		return nil, errPasswordNotMatching
	}
	_, tokenString, err := auth.GetTokenForUser(u.id)
	if err != nil {
		return nil, errTokenGeneration
	}
	return &tokenString, nil
}
