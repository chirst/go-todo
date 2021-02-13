package user

// MemoryRepository persists users
type MemoryRepository struct {
	users []User
}

// AddUser saves a single user
func (s *MemoryRepository) addUser(u User) *User {
	s.users = append(s.users, u)
	return &User{ID: u.ID, Name: u.Name, Password: u.Password}
}
