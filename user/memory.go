package user

// MemoryRepository persists users
type MemoryRepository struct {
	users []User
}

func (s *MemoryRepository) addUser(u User) (*User, error) {
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
	// TODO: not the greatest thing to return double nil
	return nil, nil
}
