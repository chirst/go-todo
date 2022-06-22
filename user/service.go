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

// Repository defines a way of persisting users.
type Repository interface {
	addUser(User) (*User, error)
	getUserByName(string) (*User, error)
}

// Service defines a way of managing users.
type Service interface {
	AddUser(u *User) (*User, error)
	GetUserTokenString(username, password string) (*string, error)
}

type service struct {
	r Repository
}

// NewService creates an instance of the users service.
func NewService(r Repository) Service {
	return &service{r}
}

// AddUser validates, creates, and adds the user to persistence.
func (s *service) AddUser(u *User) (*User, error) {
	hashedPassword, err := auth.GenerateFromPassword(u.password)
	if err != nil {
		return nil, err
	}
	u.password = *hashedPassword
	return s.r.addUser(*u)
}

// GetUserTokenString returns an auth token string for the first user matching
// the given username and password. It returns nil and an error for anything
// invalid.
func (s *service) GetUserTokenString(username, password string) (*string, error) {
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
