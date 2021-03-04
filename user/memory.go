package user

// MemoryRepository persists users
type MemoryRepository struct {
	users []User
}

func (s *MemoryRepository) addUser(u User) *User {
	u.ID = int64(len(s.users)) + 1
	s.users = append(s.users, u)
	return &u
}

func (s *MemoryRepository) getUserByName(n string) *User {
	for _, u := range s.users {
		if u.Username == n {
			return &u
		}
	}
	return nil
}
