package todo

import "log"

// MemoryRepository persists todos
type MemoryRepository struct {
	todos []Todo
}

// GetTodos gets all todos in storage
func (s *MemoryRepository) getTodos(userID int64) ([]*Todo, error) {
	var userTodos []*Todo
	userTodos = []*Todo{}
	for i := range s.todos {
		if s.todos[i].userId == userID {
			userTodos = append(userTodos, &s.todos[i])
		}
	}
	return userTodos, nil
}

// AddTodo adds a single todo to storage
func (s *MemoryRepository) addTodo(t Todo) (*Todo, error) {
	id := int64(len(s.todos)) + 1
	nt, err := NewTodo(id, t.name, t.completed, t.userId)
	s.todos = append(s.todos, *nt)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return nt, nil
}
