package user

import (
	"errors"
	"todo/auth"
)

var ErrUsernameRequired error = errors.New("username required")
var ErrPasswordRequired error = errors.New("password required")
var ErrPasswordNotMatching error = errors.New("password not matching")
var ErrUserNotFound error = errors.New("user not found")
var ErrTokenGeneration error = errors.New("token generation error")
var ErrUserExists error = errors.New("user exists")

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
func (s *Service) AddUser(username, p string) (*User, error) {
	// Todo: check if user already exists
	// Todo: stricter validation on name and pass
	if username == "" {
		return nil, ErrUsernameRequired
	}
	if p == "" {
		return nil, ErrPasswordRequired
	}
	if s.r.getUserByName(username) != nil {
		return nil, ErrUserExists
	}
	hashedPassword, err := auth.GenerateFromPassword(p)
	if err != nil {
		return nil, err
	}
	u := User{ID: 0, Username: username, Password: *hashedPassword}
	return s.r.addUser(u), nil
}

// GetUserTokenString returns an auth token string for the first user matching the
// given username and password. It returns nil for anything invalid.
func (s *Service) GetUserTokenString(username, password string) (*string, error) {
	u := s.r.getUserByName(username)
	if u == nil {
		return nil, ErrUserNotFound
	}
	if auth.CompareHashAndPassword(u.Password, password) != nil {
		return nil, ErrPasswordNotMatching
	}
	_, tokenString, err := auth.GetTokenForUser(u.ID)
	if err != nil {
		return nil, ErrTokenGeneration
	}
	return &tokenString, nil
}
