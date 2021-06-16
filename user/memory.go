package user

// MemoryRepository persists users
type MemoryRepository struct {
	users []User
}

func (s *MemoryRepository) addUser(u User) (*User, error) {
	u.ID = int64(len(s.users)) + 1
	s.users = append(s.users, u)
	return &u, nil
}

func (s *MemoryRepository) getUserByName(n string) (*User, error) {
	for _, u := range s.users {
		if u.Username == n {
			return &u, nil
		}
	}
	return nil, nil
}
