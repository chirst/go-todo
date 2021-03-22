package todo

import "errors"

var ErrNameRequired error = errors.New("Name is required")

// Repository for todos
type Repository interface {
	addTodo(Todo) *Todo
	getTodos(int64) []Todo
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
	if t.Name == "" {
		return nil, ErrNameRequired
	}
	return s.r.addTodo(t), nil
}

// GetTodos gets all todos for user from persistence
func (s *Service) GetTodos(userID int64) []Todo {
	// Todo: error for non existing UserID
	// Todo: if no todos return an empty array
	return s.r.getTodos(userID)
}
