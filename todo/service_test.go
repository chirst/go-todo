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

	tests := map[string]struct {
		input Todo
		want  *Todo
		want2 error
	}{
		"adds": {
			input: Todo{
				Name:      "do stuff",
				Completed: now,
				UserID:    1,
			},
			want: &Todo{
				Name:      "do stuff",
				Completed: now,
				UserID:    1,
			},
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
			got := s.GetTodos(tc.userID)
			if len(got) != tc.wantTodoLength {
				t.Errorf("expected %#v, got: %#v", tc.wantTodoLength, len(got))
			}
		})
	}
}
