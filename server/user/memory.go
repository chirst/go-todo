package user

type user struct {
	ID       int64
	Name     string
	Password string
}

// UserStorage persists users
type UserStorage struct {
	users []user
}

// AddUser saves a single user
func (s *UserStorage) AddUser(u User) *User {
	newUser := user{u.ID, u.Name, u.Password}
	s.users = append(s.users, newUser)
	return &User{ID: u.ID, Name: u.Name, Password: u.Password}
}
