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

	exampleTodo, err := NewTodo(1, "do stuff", &now, 1)
	if err != nil {
		t.Errorf("unable to create todo")
	}

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

	todo, err := NewTodo(0, "do stuff", nil, 1)
	if err != nil {
		t.Errorf("unable to create todo")
	}
	nonUserTodo, err := NewTodo(0, "gud todo", nil, 2)
	if err != nil {
		t.Errorf("unable to create todo")
	}
	s.AddTodo(*todo)
	s.AddTodo(*todo)
	s.AddTodo(*todo)
	s.AddTodo(*nonUserTodo)

	tests := map[string]struct {
		userID         int
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

	var userID int = 1
	incompleteTodo, err := NewTodo(0, "todo1", nil, userID)
	if err != nil {
		t.Fatalf("error creating todo")
	}
	addedTodo, err := r.addTodo(*incompleteTodo)
	if err != nil {
		t.Fatalf("failed to add todo")
	}

	err = s.CompleteTodo(userID, addedTodo.id)
	if err != nil {
		t.Fatalf(err.Error())
	}

	completedTodo, err := r.getTodo(userID, addedTodo.id)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if completedTodo.completed == nil {
		t.Fatalf("expected completed todo, got incomplete todo")
	}
}

func TestDeleteTodo(t *testing.T) {
	r := &MemoryRepository{}
	s := NewService(r)

	var userID int = 1
	td, err := NewTodo(0, "todo1", nil, userID)
	if err != nil {
		t.Errorf("error creating todo")
	}
	addedTodo, err := r.addTodo(*td)
	if err != nil {
		t.Errorf("failed to add todo")
	}

	err = s.DeleteTodo(userID, addedTodo.id)
	if err != nil {
		t.Errorf("failed to delete todo with err: %v", err.Error())
	}

	deletedTodo, err := r.getTodo(userID, addedTodo.id)
	if err != nil {
		t.Errorf("failed to get deleted todo with err: %v", err.Error())
	}

	if deletedTodo.deleted == nil {
		t.Errorf("expected deleted todo, got nil for deleted")
	}
}
