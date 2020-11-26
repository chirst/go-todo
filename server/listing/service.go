package listing

type Repository interface {
	GetTodos() []Todo
}

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) GetTodos() []Todo {
	return s.r.GetTodos()
}
