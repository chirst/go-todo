package todo

import "errors"

var errNameRequired = errors.New("name is required")

// Repository for todos
type Repository interface {
	addTodo(Todo) (*Todo, error)
	getTodos(int) ([]*Todo, error)
	getTodo(userID, todoID int) (*Todo, error)
	completeTodo(userID, todoID int) error
	incompleteTodo(userID, todoID int) error
	updateName(userID, todoID int, name string) error
	deleteTodo(userID, todoID int) error
}

// TodoService for todos
type TodoService interface {
	AddTodo(t Todo) (*Todo, error)
	GetTodos(userID int) (Todos, error)
	CompleteTodo(userID, todoID int) error
	IncompleteTodo(userID, todoID int) error
	ChangeTodoName(userID, todoID int, name string) error
	DeleteTodo(userID, todoID int) error
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
func (s *service) GetTodos(userID int) (Todos, error) {
	return s.r.getTodos(userID)
}

// CompleteTodo marks a todo as complete
func (s *service) CompleteTodo(userID, todoID int) error {
	return s.r.completeTodo(userID, todoID)
}

// IncompleteTodo marks a todo as incomplete
func (s *service) IncompleteTodo(userID, todoID int) error {
	return s.r.incompleteTodo(userID, todoID)
}

// ChangeTodoName changes the name of a todo
func (s *service) ChangeTodoName(userID int, todoID int, name string) error {
	t, err := s.r.getTodo(userID, todoID)
	if err != nil {
		return err
	}
	err = t.setName(name)
	if err != nil {
		return err
	}
	return s.r.updateName(userID, todoID, name)
}

// DeleteTodo marks a todo as deleted where it will remain but not be accessed
func (s *service) DeleteTodo(userID, todoID int) error {
	return s.r.deleteTodo(userID, todoID)
}
