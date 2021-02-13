package todo

import (
	"testing"
)

func TestAddTodo(t *testing.T) {
	r := new(TodoStorage)
	s := NewService(r)
	todo := Todo{
		Name: "do stuff",
	}

	addedTodo, err := s.AddTodo(todo)

	if addedTodo == nil {
		t.Errorf("got nil want not nil")
	}
	if err != nil {
		t.Errorf("got nil want not nil")
	}
}

func TestAddBlankName(t *testing.T) {
	r := new(TodoStorage)
	s := NewService(r)
	todo := Todo{
		Name: "",
	}

	addedTodo, err := s.AddTodo(todo)

	if err != ErrNameRequired {
		t.Errorf("got %v want %v", err, ErrNameRequired)
	}
	if addedTodo != nil {
		t.Errorf("got %v want nil", addedTodo)
	}
}

func TestGetTodos(t *testing.T) {
	r := new(TodoStorage)
	s := NewService(r)
	todo := Todo{
		Name: "do stuff",
	}
	s.AddTodo(todo)
	s.AddTodo(todo)
	s.AddTodo(todo)
	todos := s.GetTodos()
	if len(todos) != 3 {
		t.Errorf("got %v want %v", len(todos), 3)
	}
}
