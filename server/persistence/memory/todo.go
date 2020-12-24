package memory

import "todo/domain"

type todo struct {
	Name string
}

// TodoStorage persists todos
type TodoStorage struct {
	todos []todo
}

// GetTodos gets all todos in storage
func (s *TodoStorage) GetTodos() []domain.Todo {
	ret := []domain.Todo{}
	for _, t := range s.todos {
		ret = append(ret, domain.Todo{Name: t.Name})
	}
	return ret
}

// AddTodo adds a single todo to storage
func (s *TodoStorage) AddTodo(t domain.Todo) *domain.Todo {
	newTodo := todo{Name: t.Name}
	s.todos = append(s.todos, newTodo)
	return &domain.Todo{Name: newTodo.Name}
}
