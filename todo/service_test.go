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

	exampleTodo, err := NewTodo(1, "do stuff", &now, 1, DefaultPriority())
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
			got, got2 := s.AddTodo(
				tc.input.name,
				tc.input.completed,
				tc.input.userID,
				&tc.input.priority.id,
			)
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

	defaultPriorityID := DefaultPriority().id
	s.AddTodo("do stuff", nil, 1, &defaultPriorityID)
	s.AddTodo("do stuff", nil, 1, &defaultPriorityID)
	s.AddTodo("do stuff", nil, 1, &defaultPriorityID)
	s.AddTodo("gud todo", nil, 2, &defaultPriorityID)

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
	addedTodo, err := r.addTodo("todo1", nil, userID, DefaultPriority())
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

func TestIncompleteTodo(t *testing.T) {
	r := &MemoryRepository{}
	s := NewService(r)

	userID := 1
	now := time.Now()
	addedTodo, err := r.addTodo("complete todo", &now, userID, DefaultPriority())
	if err != nil {
		t.Fatalf("failed to add todo")
	}

	err = s.IncompleteTodo(userID, addedTodo.id)
	if err != nil {
		t.Fatalf(err.Error())
	}

	incompletedTodo, err := r.getTodo(userID, addedTodo.id)
	if err != nil {
		t.Fatalf(err.Error())
	}

	if incompletedTodo.completed != nil {
		t.Fatalf("expected incomplete todo, got completed todo")
	}
}

func TestChangeTodoName(t *testing.T) {

	t.Run("invalid name", func(t *testing.T) {
		r := &MemoryRepository{}
		s := NewService(r)

		todo, err := r.addTodo("gud name", nil, 1, DefaultPriority())
		if err != nil {
			t.Fatalf("error adding todo")
		}

		err = s.ChangeTodoName(todo.userID, todo.id, "")
		if err == nil {
			t.Fatalf("got nil error expected err")
		}
	})

	t.Run("valid name", func(t *testing.T) {
		r := &MemoryRepository{}
		s := NewService(r)

		todo, err := r.addTodo("gud name", nil, 1, DefaultPriority())
		if err != nil {
			t.Fatalf("error adding todo")
		}

		newName := "gudder name"
		err = s.ChangeTodoName(todo.userID, todo.id, newName)
		if err != nil {
			t.Fatalf("got error expected no err")
		}

		memoryTodo, err := r.getMemoryTodo(todo.userID, todo.id)
		if err != nil {
			t.Fatalf("unable to get memory todo")
		}
		if memoryTodo.name != newName {
			t.Fatalf("got wrong name: %s, want: %s", memoryTodo.name, newName)
		}
	})
}

func TestDeleteTodo(t *testing.T) {
	r := &MemoryRepository{}
	s := NewService(r)

	var userID int = 1
	addedTodo, err := r.addTodo("todo1", nil, userID, DefaultPriority())
	if err != nil {
		t.Errorf("failed to add todo")
	}

	err = s.DeleteTodo(userID, addedTodo.id)
	if err != nil {
		t.Errorf("failed to delete todo with err: %v", err.Error())
	}

	deletedTodo, err := r.getMemoryTodo(userID, addedTodo.id)
	if err != nil {
		t.Errorf("failed to get deleted todo with err: %v", err.Error())
	}

	if deletedTodo.deleted == nil {
		t.Errorf("expected deleted todo, got nil for deleted")
	}
}

func TestGetPriorities(t *testing.T) {
	r := &MemoryRepository{}
	s := NewService(r)

	ps, err := s.GetPriorities()

	if err != nil {
		t.Errorf("got error: %#v, want no error", err.Error())
	}
	if psLen := len(ps); psLen != 3 {
		t.Errorf("expected %#v priorities, got %#v", 3, psLen)
	}
}

func TestUpdatePriority(t *testing.T) {
	r := &MemoryRepository{}
	s := NewService(r)

	uid := 1
	newPriorityId := 1
	addedTodo, err := r.addTodo("todo1", nil, uid, DefaultPriority())
	if err != nil {
		t.Errorf("failed to add todo")
	}

	err = s.UpdatePriority(uid, addedTodo.id, newPriorityId)

	if err != nil {
		t.Errorf("got error: %#v, want no error", err.Error())
	}
	savedTd, err := r.getTodo(uid, addedTodo.id)
	if err != nil {
		t.Errorf("error getting saved todo")
	}
	if savedTd.priority.id != newPriorityId {
		t.Errorf(
			"saved priorityID %#v, want priorityID %#v",
			savedTd.priority.id,
			newPriorityId,
		)
	}
}
