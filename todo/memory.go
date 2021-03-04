package todo

// MemoryRepository persists todos
type MemoryRepository struct {
	todos []Todo
}

// GetTodos gets all todos in storage
func (s *MemoryRepository) getTodos(userID int64) []Todo {
	var userTodos []Todo
	for _, t := range s.todos {
		if t.UserID == userID {
			userTodos = append(userTodos, t)
		}
	}
	return userTodos
}

// AddTodo adds a single todo to storage
func (s *MemoryRepository) addTodo(t Todo) *Todo {
	s.todos = append(s.todos, t)
	return &t
}
