package adding

type Repository interface {
	AddTodo() Todo
}

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) AddTodo() Todo {
	return s.r.AddTodo()
}
