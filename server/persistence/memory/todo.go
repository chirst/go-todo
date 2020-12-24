package memory

import (
	"todo/adding"
	"todo/listing"
)

// Todo ...
type Todo struct {
	Name string
}

// Storage ...
type TodoStorage struct {
	todos []Todo
}

// GetTodos ...
func (s *TodoStorage) GetTodos() []listing.Todo {
	ret := []listing.Todo{}
	for _, todo := range s.todos {
		ret = append(ret, listing.Todo{Name: todo.Name})
	}
	return ret
}

// AddTodo ...
func (s *TodoStorage) AddTodo(todo adding.Todo) *adding.Todo {
	newTodo := Todo{Name: todo.Name}
	s.todos = append(s.todos, newTodo)
	return &adding.Todo{Name: newTodo.Name}
}
