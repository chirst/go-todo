package user

import "errors"

// User ...
type User struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
	Password string `json:"-"`
}

func createUser(username, password string) (*User, error) {
	if username == "" {
		return nil, errors.New("username required")
	}
	if password == "" {
		return nil, errors.New("password required")
	}
	return &User{ID: 0, Username: username, Password: password}, nil
}
