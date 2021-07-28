package todo

import (
	"reflect"
	"testing"
	"time"
)

func TestAddTodo(t *testing.T) {
	r := new(MemoryRepository)
	s := NewService(r)

	now := time.Now()

	exampleTodo, _ := NewTodo(1, "do stuff", &now, 1)

	tests := map[string]struct {
		input Todo
		want  *Todo
		want2 error
	}{
		"adds": {
			input: *exampleTodo,
			want:  exampleTodo,
			want2: nil,
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

	todo, _ := NewTodo(0, "do stuff", nil, 1)
	nonUserTodo, _ := NewTodo(0, "gud todo", nil, 2)
	s.AddTodo(*todo)
	s.AddTodo(*todo)
	s.AddTodo(*todo)
	s.AddTodo(*nonUserTodo)

	tests := map[string]struct {
		userID         int64
		wantTodoLength int
	}{
		"gets": {
			1,
			3,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, _ := s.GetTodos(tc.userID)
			if len(got) != tc.wantTodoLength {
				t.Fatalf("expected %#v, got: %#v", tc.wantTodoLength, len(got))
			}
		})
	}
}

func TestCompleteTodo(t *testing.T) {
	r := new(MemoryRepository)
	s := NewService(r)

	incompleteTodo, err := NewTodo(0, "todo1", nil, 1)
	if err != nil {
		t.Fatalf("error creating todo")
	}
	addedTodo, err := r.addTodo(*incompleteTodo)
	if err != nil {
		t.Fatalf("failed to add todo")
	}

	err = s.CompleteTodo(addedTodo.id)
	if err != nil {
		t.Fatalf(err.Error())
	}

	completedTodo, err := r.getTodo(addedTodo.id)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if completedTodo.completed == nil {
		t.Fatalf("expected completed todo, got incomplete todo")
	}
}
