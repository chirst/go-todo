package user

// MemoryRepository persists users
type MemoryRepository struct {
	users []User
}

func (s *MemoryRepository) addUser(u User) *User {
	s.users = append(s.users, u)
	return &User{ID: u.ID, Name: u.Name, Password: u.Password}
}

func (s *MemoryRepository) getUserByName(n string) *User {
	for _, u := range s.users {
		if u.Name == n {
			return &u
		}
	}
	return nil
}