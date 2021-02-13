package todo

import (
	"errors"
)

// ErrNameRequired for an empty todo name
var ErrNameRequired = errors.New("Name is required")

// Repository for todos
type Repository interface {
	addTodo(Todo) *Todo
	getTodos() []Todo
}

// Service for todos
type Service struct {
	r Repository
}

// NewService creates an instance of this service
func NewService(r Repository) *Service {
	return &Service{r}
}

// AddTodo is for creating, validating and adding a new todo to persistence
func (s *Service) AddTodo(t Todo) (*Todo, error) {
	if t.Name == "" {
		return nil, ErrNameRequired
	}
	return s.r.addTodo(t), nil
}

// GetTodos gets all todos from persistence
func (s *Service) GetTodos() []Todo {
	return s.r.getTodos()
}
