package user

// Repository for users
type Repository interface {
	AddUser(User) *User
}

// Service for users
type Service struct {
	r Repository
}

// NewService creates an instance of this service
func NewService(r Repository) *Service {
	return &Service{r}
}

// AddUser validates, creates, and adds the user to persistence
func (s *Service) AddUser(name, password string) (*User, error) {
	u, err := CreateUser(name, password)
	if err != nil {
		return nil, err
	}
	return s.r.AddUser(*u), nil
}
