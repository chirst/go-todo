package todo

import (
	"reflect"
	"testing"
)

func TestAddTodo(t *testing.T) {
	r := new(MemoryRepository)
	s := NewService(r)

	tests := map[string]struct {
		input Todo
		want  *Todo
		want2 error
	}{
		"happy path": {
			input: Todo{Name: "do stuff"},
			want:  &Todo{Name: "do stuff"},
			want2: nil,
		},
		"no name": {
			input: Todo{Name: ""},
			want:  nil,
			want2: ErrNameRequired,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, got2 := s.AddTodo(tc.input)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected %#v, got: %#v", tc.want, got)
			}
			if tc.want2 != got2 {
				t.Fatalf("expected %#v, got %#v", tc.want2, got2)
			}
		})
	}
}

func TestGetTodos(t *testing.T) {
	r := new(MemoryRepository)
	s := NewService(r)
	todo := Todo{
		UserID: 1,
		Name:   "do stuff",
	}
	nonUserTodo := Todo{
		UserID: 2,
		Name:   "gud todo",
	}
	s.AddTodo(todo)
	s.AddTodo(todo)
	s.AddTodo(todo)
	s.AddTodo(nonUserTodo)
	todos := s.GetTodos(1)
	if len(todos) != 3 {
		t.Errorf("got %v want %v", len(todos), 3)
	}
}
