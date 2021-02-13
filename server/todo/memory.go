package todo

type todo struct {
	Name string
}

// TodoStorage persists todos
type TodoStorage struct {
	todos []todo
}

// GetTodos gets all todos in storage
func (s *TodoStorage) GetTodos() []Todo {
	ret := []Todo{}
	for _, t := range s.todos {
		ret = append(ret, Todo{Name: t.Name})
	}
	return ret
}

// AddTodo adds a single todo to storage
func (s *TodoStorage) AddTodo(t Todo) *Todo {
	newTodo := todo{Name: t.Name}
	s.todos = append(s.todos, newTodo)
	return &Todo{Name: newTodo.Name}
}
