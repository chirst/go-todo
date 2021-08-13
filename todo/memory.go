package todo

import (
	"fmt"
	"time"
)

// MemoryRepository persists todos
type MemoryRepository struct {
	todos []Todo
}

// GetTodos gets all todos in storage
func (r *MemoryRepository) getTodos(userID int) ([]*Todo, error) {
	var userTodos []*Todo
	userTodos = []*Todo{}
	for i := range r.todos {
		if r.todos[i].userID == userID {
			userTodos = append(userTodos, &r.todos[i])
		}
	}
	return userTodos, nil
}

// AddTodo adds a single todo to storage
func (r *MemoryRepository) addTodo(t Todo) (*Todo, error) {
	id := int(len(r.todos)) + 1
	nt, err := NewTodo(id, t.name, t.completed, t.userID)
	r.todos = append(r.todos, *nt)
	if err != nil {
		return nil, err
	}
	return nt, nil
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

func (r *MemoryRepository) getTodo(userID, id int) (*Todo, error) {
	for i := range r.todos {
		if r.todos[i].id == id && r.todos[i].userID == userID {
			return &r.todos[i], nil
		}
	}
	return nil, fmt.Errorf("no todo found with id: %v and userID: %v", id, userID)
}
