package adding

import (
	"errors"
	"todo/domain"
)

// ErrNameRequired ...
var ErrNameRequired = errors.New("Name is required")

// Repository ...
type Repository interface {
	AddTodo(domain.Todo) *domain.Todo
}

// Service ...
type Service struct {
	r Repository
}

// NewService ...
func NewService(r Repository) *Service {
	return &Service{r}
}

// AddTodo is for adding a new Todo
func (s *Service) AddTodo(todo domain.Todo) (*domain.Todo, error) {
	if todo.Name == "" {
		return nil, ErrNameRequired
	}
	return s.r.AddTodo(todo), nil
}
