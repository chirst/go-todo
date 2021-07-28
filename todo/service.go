package todo

import "errors"

var errNameRequired = errors.New("name is required")

// Repository for todos
type Repository interface {
	addTodo(Todo) (*Todo, error)
	getTodos(int64) ([]*Todo, error)
	completeTodo(todoID int64) error
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
func (s *Service) GetTodos(userID int64) (Todos, error) {
	return s.r.getTodos(userID)
}

// CompleteTodo marks a todo as complete
func (s *Service) CompleteTodo(todoID int64) error {
	return s.r.completeTodo(todoID)
}

// IncompleteTodo marks a todo as incomplete
func (s *Service) IncompleteTodo(todoId int64) {
	// TODO:
}

// DeleteTodo marks a todo as deleted where it will remain but not be accessed
func (s *Service) DeleteTodo(todoId int64) {
	// TODO:
}

// ChangeTodoName changes the name of a todo
func (s *Service) ChangeTodoName(todoId int64, name string) {
	// TODO:
}

// TODO: grouping and sorting
