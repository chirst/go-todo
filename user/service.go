package user

import (
	"errors"
	"log"
	"todo/auth"
)

var errUsernameRequired error = errors.New("username required")
var errPasswordRequired error = errors.New("password required")
var errPasswordNotMatching error = errors.New("password not matching")
var errUserNotFound error = errors.New("user not found")
var errTokenGeneration error = errors.New("token generation error")
var errUserExists error = errors.New("user exists")

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
func (s *Service) AddUser(username, p string) (*User, error) {
	// TODO: stricter validation on name and pass
	if username == "" {
		return nil, errUsernameRequired
	}
	if p == "" {
		return nil, errPasswordRequired
	}
	u, err := s.r.getUserByName(username)
	if err != nil {
		return nil, err
	}
	if u != nil { // TODO: read this and make sure it doesn't suck
		return nil, errUserExists
	}
	hashedPassword, err := auth.GenerateFromPassword(p)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return s.r.addUser(User{ID: 0, Username: username, Password: *hashedPassword})
}

// GetUserTokenString returns an auth token string for the first user matching the
// given username and password. It returns nil for anything invalid.
func (s *Service) GetUserTokenString(username, password string) (*string, error) {
	u, err := s.r.getUserByName(username)
	if err != nil {
		return nil, err
	}
	if u == nil { // TODO: this should probably just be an error from the repo
		return nil, errUserNotFound
	}
	if auth.CompareHashAndPassword(u.Password, password) != nil {
		return nil, errPasswordNotMatching
	}
	_, tokenString, err := auth.GetTokenForUser(u.ID)
	if err != nil {
		log.Print(err)
		return nil, errTokenGeneration
	}
	return &tokenString, nil
}
