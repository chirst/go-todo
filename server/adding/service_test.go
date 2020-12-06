package adding

import (
	"testing"
)

func TestAddTodo(t *testing.T) {
	mockRepo := new(mockStorage)
	s := NewService(mockRepo)
	todo := Todo{
		Name: "do stuff",
	}

	addedTodo, err := s.AddTodo(todo)

	if len(mockRepo.todos) != 1 {
		t.Errorf("got %d want %d", len(mockRepo.todos), 1)
	}
	if addedTodo == nil {
		t.Errorf("got nil want not nil")
	}
	if err != nil {
		t.Errorf("got nil want not nil")
	}
}

func TestAddBlankName(t *testing.T) {
	mockRepo := new(mockStorage)
	s := NewService(mockRepo)
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

type mockStorage struct {
	todos []Todo
}

func (s *mockStorage) AddTodo(todo Todo) *Todo {
	newTodo := Todo{Name: todo.Name}
	s.todos = append(s.todos, newTodo)
	return &Todo{Name: newTodo.Name}
}
