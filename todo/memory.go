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
	id         int
	name       string
	completed  *time.Time
	deleted    *time.Time
	userID     int
	priorityID int
}

// GetTodos gets all todos in storage
func (r *MemoryRepository) getTodos(userID int) ([]*Todo, error) {
	var userTodos []*Todo
	userTodos = []*Todo{}
	for i := range r.todos {
		if r.todos[i].userID == userID {
			priority, err := r.getPriority(r.todos[i].priorityID)
			if err != nil {
				return nil, err
			}
			userTodo, err := newTodo(
				r.todos[i].id,
				r.todos[i].name,
				r.todos[i].completed,
				r.todos[i].userID,
				*priority,
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
func (r *MemoryRepository) addTodo(
	name string,
	completed *time.Time,
	userID int,
	priority priorityModel,
) (*Todo, error) {
	id := int(len(r.todos)) + 1
	mt := memoryTodo{
		id:         id,
		name:       name,
		completed:  completed,
		deleted:    nil,
		userID:     userID,
		priorityID: priority.id,
	}
	r.todos = append(r.todos, mt)
	newTodo, err := newTodo(id, mt.name, mt.completed, mt.userID, priority)
	if err != nil {
		return nil, err
	}
	return newTodo, nil
}

// Complete todo marks todo with the given id as complete and returns no error
// on success
func (r *MemoryRepository) completeTodo(userID, todoID int) error {
	t, err := r.getMemoryTodo(userID, todoID)
	if err != nil {
		return err
	}
	now := time.Now()
	t.completed = &now
	return nil
}

func (r *MemoryRepository) incompleteTodo(userID, todoID int) error {
	t, err := r.getMemoryTodo(userID, todoID)
	if err != nil {
		return err
	}
	t.completed = nil
	return nil
}

func (r *MemoryRepository) deleteTodo(userID, todoID int) error {
	t, err := r.getMemoryTodo(userID, todoID)
	if err != nil {
		return err
	}
	now := time.Now()
	t.deleted = &now
	return nil
}

func (r *MemoryRepository) getPriorities() (Priorities, error) {
	return []*priorityModel{
		{
			id:     1,
			name:   "Low",
			weight: 1,
		},
		{
			id:     2,
			name:   "Normal",
			weight: 2,
		},
		{
			id:     3,
			name:   "High",
			weight: 3,
		},
	}, nil
}

func (r *MemoryRepository) getPriority(priorityID int) (*priorityModel, error) {
	switch priorityID {
	case 1:
		return &priorityModel{
			id:     1,
			name:   "Low",
			weight: 1,
		}, nil
	case 2:
		return &priorityModel{
			id:     2,
			name:   "Normal",
			weight: 2,
		}, nil
	case 3:
		return &priorityModel{
			id:     3,
			name:   "High",
			weight: 3,
		}, nil
	default:
		return nil, fmt.Errorf("priority with id %#v not found", priorityID)
	}
}

func (r *MemoryRepository) updateName(userID, todoID int, name string) error {
	t, err := r.getMemoryTodo(userID, todoID)
	if err != nil {
		return err
	}
	t.name = name
	return nil
}

func (r *MemoryRepository) updatePriority(userID, todoID, priorityID int) error {
	t, err := r.getMemoryTodo(userID, todoID)
	if err != nil {
		return err
	}
	t.priorityID = priorityID
	return nil
}

func (r *MemoryRepository) getTodo(userID, id int) (*Todo, error) {
	t, err := r.getMemoryTodo(userID, id)
	if err != nil {
		return nil, err
	}
	priority, err := r.getPriority(t.priorityID)
	if err != nil {
		return nil, err
	}
	return newTodo(t.id, t.name, t.completed, t.userID, *priority)
}

func (r *MemoryRepository) getMemoryTodo(userID, id int) (*memoryTodo, error) {
	for i := range r.todos {
		if r.todos[i].id == id && r.todos[i].userID == userID {
			return &r.todos[i], nil
		}
	}
	return nil, fmt.Errorf("no todo found with id: %v and userID: %v", id, userID)
}
