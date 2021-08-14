package todo

import (
	"fmt"
	"time"
)

// MemoryRepository persists todos
type MemoryRepository struct {
	todos []memoryTodo
}

type memoryTodo struct {
	id        int
	name      string
	completed *time.Time
	deleted   *time.Time
	userID    int
}

// GetTodos gets all todos in storage
func (r *MemoryRepository) getTodos(userID int) ([]*Todo, error) {
	var userTodos []*Todo
	userTodos = []*Todo{}
	for i := range r.todos {
		if r.todos[i].userID == userID {
			userTodo, err := NewTodo(
				r.todos[i].id,
				r.todos[i].name,
				r.todos[i].completed,
				r.todos[i].userID,
			)
			if err != nil {
				return nil, err
			}
			userTodos = append(userTodos, userTodo)
		}
	}
	return userTodos, nil
}

// AddTodo adds a single todo to storage
func (r *MemoryRepository) addTodo(t Todo) (*Todo, error) {
	id := int(len(r.todos)) + 1
	mt := memoryTodo{
		id:        id,
		name:      t.name,
		completed: t.completed,
		deleted:   nil,
		userID:    t.userID,
	}
	r.todos = append(r.todos, mt)
	newTodo, err := NewTodo(id, mt.name, mt.completed, mt.userID)
	if err != nil {
		return nil, err
	}
	return newTodo, nil
}

// Complete todo marks todo with the given id as complete and returns no error
// on success
func (r *MemoryRepository) completeTodo(userID, todoID int) error {
	t, err := r.getTodo(userID, todoID)
	if err != nil {
		return err
	}
	now := time.Now()
	t.completed = &now
	return nil
}

func (r *MemoryRepository) deleteTodo(userID, todoID int) error {
	t, err := r.getTodo(userID, todoID)
	if err != nil {
		return err
	}
	now := time.Now()
	t.deleted = &now
	return nil
}

func (r *MemoryRepository) getTodo(userID, id int) (*memoryTodo, error) {
	for i := range r.todos {
		if r.todos[i].id == id && r.todos[i].userID == userID {
			return &r.todos[i], nil
		}
	}
	return nil, fmt.Errorf("no todo found with id: %v and userID: %v", id, userID)
}
