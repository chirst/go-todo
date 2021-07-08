package user

import "errors"

var errUniqueUsername error = errors.New("user exists")
var errUserNotFound error = errors.New("user not found")

// MemoryRepository persists users
type MemoryRepository struct {
	users []User
}

func (s *MemoryRepository) addUser(u User) (*User, error) {
	for _, user := range s.users {
		if user.username == u.username {
			return nil, errUniqueUsername
		}
	}
	u.id = int64(len(s.users)) + 1
	s.users = append(s.users, u)
	return &u, nil
}

func (s *MemoryRepository) getUserByName(n string) (*User, error) {
	for _, u := range s.users {
		if u.username == n {
			return &u, nil
		}
	}
	return nil, errUserNotFound
}
