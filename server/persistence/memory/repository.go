package memory

import (
	"todo/adding"
	"todo/listing"
)

type Storage struct {
	todos []Todo
}

func (s *Storage) GetTodos() []listing.Todo {
	ret := []listing.Todo{}
	for _, todo := range s.todos {
		ret = append(ret, listing.Todo{Name: todo.Name})
	}
	return ret
}

func (s *Storage) AddTodo(todo adding.Todo) *adding.Todo {
	newTodo := Todo{Name: todo.Name}
	s.todos = append(s.todos, newTodo)
	return &adding.Todo{Name: newTodo.Name}
}
