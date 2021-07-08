package user

import "encoding/json"

// User ...
type User struct {
	id       int64
	username string
	password string
}

// NewUser is a way to create a valid User
func NewUser(id int64, username string, password string) (*User, error) {
	// TODO: stricter validation on name and pass
	if username == "" {
		return nil, errUsernameRequired
	}
	if password == "" {
		return nil, errPasswordRequired
	}
	return &User{
		id,
		username,
		password,
	}, nil
}

type userJSON struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

// ToJSON converts a User to JSON
func (u *User) ToJSON() ([]byte, error) {
	uj := userJSON{
		ID:       u.id,
		Username: u.username,
	}
	return json.Marshal(uj)
}
