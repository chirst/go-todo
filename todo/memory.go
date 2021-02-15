package todo

// MemoryRepository persists todos
type MemoryRepository struct {
	todos []Todo
}

// GetTodos gets all todos in storage
func (s *MemoryRepository) getTodos() []Todo {
	return s.todos
}

// AddTodo adds a single todo to storage
func (s *MemoryRepository) addTodo(t Todo) *Todo {
	s.todos = append(s.todos, t)
	return &Todo{Name: t.Name, Completed: t.Completed}
}
