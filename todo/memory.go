package todo

// MemoryRepository persists todos
type MemoryRepository struct {
	todos []Todo
}

// GetTodos gets all todos in storage
func (s *MemoryRepository) getTodos(userID int64) ([]*Todo, error) {
	var userTodos []*Todo
	userTodos = []*Todo{}
	for i := range s.todos {
		if s.todos[i].UserID == userID {
			userTodos = append(userTodos, &s.todos[i])
		}
	}
	return userTodos, nil
}

// AddTodo adds a single todo to storage
func (s *MemoryRepository) addTodo(t Todo) (*Todo, error) {
	t.ID = int64(len(s.todos)) + 1
	s.todos = append(s.todos, t)
	return &t, nil
}
