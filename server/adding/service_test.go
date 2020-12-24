package adding

import (
	"testing"
	"todo/domain"
	"todo/persistence/memory"
)

func TestAddTodo(t *testing.T) {
	r := new(memory.TodoStorage)
	s := NewService(r)
	todo := domain.Todo{
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
	r := new(memory.TodoStorage)
	s := NewService(r)
	todo := domain.Todo{
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
