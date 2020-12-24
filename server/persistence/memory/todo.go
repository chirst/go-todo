package memory

import "todo/domain"

// Todo ...
type Todo struct {
	Name string
}

// Storage ...
type TodoStorage struct {
	todos []Todo
}

// GetTodos ...
func (s *TodoStorage) GetTodos() []domain.Todo {
	ret := []domain.Todo{}
	for _, todo := range s.todos {
		ret = append(ret, domain.Todo{Name: todo.Name})
	}
	return ret
}

// AddTodo ...
func (s *TodoStorage) AddTodo(todo domain.Todo) *domain.Todo {
	newTodo := Todo{Name: todo.Name}
	s.todos = append(s.todos, newTodo)
	return &domain.Todo{Name: newTodo.Name}
}
