package adding

import "errors"

var ErrNameRequired = errors.New("Name is required")

type Repository interface {
	AddTodo(Todo) Todo
}

type Service struct {
	r Repository
}

func NewService(r Repository) *Service {
	return &Service{r}
}

func (s *Service) AddTodo(todo Todo) (Todo, error) {
	if todo.Name == "" {
		return Todo{}, ErrNameRequired
	}
	return s.r.AddTodo(todo), nil
}
