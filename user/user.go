package user

import "encoding/json"

// User ...
type User struct {
	id       int64
	username string
	password string
}

func NewUser(id int64, username string, password string) (*User, error) {
	// TODO: stricter validation on name and pass
	// TODO: test no username
	// TODO: test no password
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

type UserJSON struct {
	ID       int64  `json:"id"`
	Username string `json:"username"`
}

func (u *User) ToJSON() ([]byte, error) {
	uj := UserJSON{
		ID:       u.id,
		Username: u.username,
	}
	return json.Marshal(uj)
}
