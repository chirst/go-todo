package user

type Repository interface {
	AddUser(User) *User
}

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) AddUser(name, password string) *User {
	u := User{1, name, password}
	return s.r.AddUser(u)
}
