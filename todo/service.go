package todo

import "errors"

var errNameRequired error = errors.New("name is required")

// Repository for todos
type Repository interface {
	addTodo(Todo) (*Todo, error)
	getTodos(int64) ([]*Todo, error)
}

// Service for todos
type Service struct {
	r Repository
}

// NewService creates an instance of the todo service
func NewService(r Repository) *Service {
	return &Service{r}
}

// AddTodo is for creating, validating, and adding a new todo to persistence
func (s *Service) AddTodo(t Todo) (*Todo, error) {
	return s.r.addTodo(t)
}

// GetTodos gets all todos for user from persistence
func (s *Service) GetTodos(userID int64) ([]*Todo, error) {
	return s.r.getTodos(userID)
}
