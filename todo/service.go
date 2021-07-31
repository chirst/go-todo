package todo

import "errors"

var errNameRequired = errors.New("name is required")

// Repository for todos
type Repository interface {
	addTodo(Todo) (*Todo, error)
	getTodos(int64) ([]*Todo, error)
	completeTodo(todoID int64) error
}

// TodoService for todos
type TodoService interface {
	AddTodo(t Todo) (*Todo, error)
	GetTodos(userID int64) (Todos, error)
	CompleteTodo(todoID int64) error
}

type service struct {
	r Repository
}

// NewService creates an instance of the todo service
func NewService(r Repository) TodoService {
	return &service{r}
}

// AddTodo is for creating, validating, and adding a new todo to persistence
func (s *service) AddTodo(t Todo) (*Todo, error) {
	return s.r.addTodo(t)
}

// GetTodos gets all todos for user from persistence
func (s *service) GetTodos(userID int64) (Todos, error) {
	return s.r.getTodos(userID)
}

// CompleteTodo marks a todo as complete
func (s *service) CompleteTodo(todoID int64) error {
	return s.r.completeTodo(todoID)
}

// IncompleteTodo marks a todo as incomplete
func (s *service) IncompleteTodo(todoId int64) {
	// TODO:
}

// DeleteTodo marks a todo as deleted where it will remain but not be accessed
func (s *service) DeleteTodo(todoId int64) {
	// TODO:
}

// ChangeTodoName changes the name of a todo
func (s *service) ChangeTodoName(todoId int64, name string) {
	// TODO:
}

// TODO: grouping and sorting
