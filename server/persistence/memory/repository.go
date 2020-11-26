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

func (s *Storage) AddTodo() adding.Todo {
	newTodo := Todo{Name: "do stuff"}
	s.todos = append(s.todos, newTodo)
	return adding.Todo{Name: newTodo.Name}
}
