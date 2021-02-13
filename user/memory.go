package user

// MemoryRepository persists users
type MemoryRepository struct {
	users []User
}

func (s *MemoryRepository) addUser(u User) *User {
	s.users = append(s.users, u)
	fakeID := int64(len(s.users))
	return &User{ID: fakeID, Username: u.Username, Password: u.Password}
}

func (s *MemoryRepository) getUserByName(n string) *User {
	for _, u := range s.users {
		if u.Username == n {
			return &u
		}
	}
	return nil
}
