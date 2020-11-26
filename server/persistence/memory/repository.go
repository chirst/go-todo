package memory

import "todo/listing"

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
