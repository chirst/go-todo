package todo

import (
	"errors"
	"time"
)

var errNameRequired = errors.New("name is required")

// Repository for todos
type Repository interface {
	addTodo(
		name string,
		completed *time.Time,
		userID int,
		priority Priority,
	) (*Todo, error)
	getTodos(int) ([]*Todo, error)
	getTodo(userID, todoID int) (*Todo, error)
	completeTodo(userID, todoID int) error
	incompleteTodo(userID, todoID int) error
	updateName(userID, todoID int, name string) error
	deleteTodo(userID, todoID int) error
	getPriorities() (Priorities, error)
	getPriority(priorityID int) (*Priority, error)
	updatePriority(userID, todoID, priorityID int) error
}

// TodoService for todos
type TodoService interface {
	AddTodo(
		name string,
		completed *time.Time,
		userID int,
		priorityID *int,
	) (*Todo, error)
	GetTodos(userID int) (Todos, error)
	CompleteTodo(userID, todoID int) error
	IncompleteTodo(userID, todoID int) error
	ChangeTodoName(userID, todoID int, name string) error
	DeleteTodo(userID, todoID int) error
	GetPriorities() (Priorities, error)
	UpdatePriority(userID, todoID, priorityID int) error
}

type service struct {
	r Repository
}

// NewService creates an instance of the todo service
func NewService(r Repository) TodoService {
	return &service{r}
}

// AddTodo is for creating, validating, and adding a new todo to persistence
func (s *service) AddTodo(
	name string,
	completed *time.Time,
	userID int,
	priorityID *int,
) (*Todo, error) {
	priority := DefaultPriority()
	if priorityID != nil {
		p, err := s.r.getPriority(*priorityID)
		if err != nil {
			return nil, err
		}
		priority = *p
	}
	return s.r.addTodo(name, completed, userID, priority)
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

// GetPriorities gets all priorities
func (s *service) GetPriorities() (Priorities, error) {
	return s.r.getPriorities()
}

// UpdatePriority changes the given todos priority
func (s *service) UpdatePriority(userID, todoID, priorityID int) error {
	return s.r.updatePriority(userID, todoID, priorityID)
}
