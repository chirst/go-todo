package user

import "errors"

// User ...
type User struct {
	ID       int64  `json:"id"`
	Name     string `json:"name"`
	Password string `json:"-"`
}

func createUser(name, password string) (*User, error) {
	if name == "" {
		return nil, errors.New("name required")
	}
	if password == "" {
		return nil, errors.New("password required")
	}
	return &User{ID: 0, Name: name, Password: password}, nil
}
