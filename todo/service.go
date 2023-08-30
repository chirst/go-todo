package todo

import (
	"errors"
	"time"
)

var errNameRequired = errors.New("name is required")

// Repository defines a way to persist todos.
type Repository interface {
	addTodo(
		name string,
		completed *time.Time,
		userID int,
		priority priorityModel,
	) (*Todo, error)
	getTodos(int) ([]*Todo, error)
	getTodo(userID, todoID int) (*Todo, error)
	completeTodo(userID, todoID int) error
	incompleteTodo(userID, todoID int) error
	updateName(userID, todoID int, name string) error
	deleteTodo(userID, todoID int) error
	getPriorities() (Priorities, error)
	getPriority(priorityID int) (*priorityModel, error)
	updatePriority(userID, todoID, priorityID int) error
}

type service struct {
	r Repository
}

// NewService creates an instance of the TodoService.
func NewService(r Repository) *service {
	return &service{r}
}

// AddTodo is for creating, validating, and adding a new todo to persistence.
func (s *service) AddTodo(
	name string,
	completed *time.Time,
	userID int,
	priorityID *int,
) (*Todo, error) {
	priority := defaultPriority()
	if priorityID != nil {
		p, err := s.r.getPriority(*priorityID)
		if err != nil {
			return nil, err
		}
		priority = *p
	}
	return s.r.addTodo(name, completed, userID, priority)
}

// GetTodos gets all todos for the given user from persistence.
func (s *service) GetTodos(userID int) (Todos, error) {
	return s.r.getTodos(userID)
}

// CompleteTodo marks the given todo as complete.
func (s *service) CompleteTodo(userID, todoID int) error {
	return s.r.completeTodo(userID, todoID)
}

// IncompleteTodo marks the given todo as incomplete.
func (s *service) IncompleteTodo(userID, todoID int) error {
	return s.r.incompleteTodo(userID, todoID)
}

// ChangeTodoName changes the name of the given todo to the given name.
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

// DeleteTodo marks a todo as deleted where it will remain in persistence, but
// not be able to be accessed.
func (s *service) DeleteTodo(userID, todoID int) error {
	return s.r.deleteTodo(userID, todoID)
}

// GetPriorities gets all priorities.
func (s *service) GetPriorities() (Priorities, error) {
	return s.r.getPriorities()
}

// UpdatePriority changes the priority of the given todo.
func (s *service) UpdatePriority(userID, todoID, priorityID int) error {
	return s.r.updatePriority(userID, todoID, priorityID)
}
