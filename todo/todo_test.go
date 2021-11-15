package todo

import (
	"testing"
	"time"
)

func TestNewTodo(t *testing.T) {
	tests := map[string]struct {
		id        int
		name      string
		completed *time.Time
		userID    int
		priority  Priority
		want      error
	}{
		"blank name": {
			id:        0,
			name:      "",
			completed: nil,
			userID:    0,
			priority:  DefaultPriority(),
			want:      errNameRequired,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, got := NewTodo(
				tc.id,
				tc.name,
				tc.completed,
				tc.userID,
				tc.priority,
			)
			if tc.want != got {
				t.Fatalf("expected %#v, got %#v", tc.want, got)
			}
		})
	}
}

func TestTodoToJSON(t *testing.T) {
	nt, err := NewTodo(1, "gud todo", nil, 2, DefaultPriority())
	if err != nil {
		t.Errorf("got err: %#v want no err", err.Error())
	}

	j, err := nt.ToJSON()

	if err != nil {
		t.Errorf("got err: %#v want no err", err.Error())
	}
	// TODO: long line
	expectedJSON := `{"id":1,"name":"gud todo","completed":null,"userId":2,"priority":{"id":2,"name":"Normal","weight":2}}`
	if string(j) != expectedJSON {
		t.Errorf("got %s\nwant %s", string(j), expectedJSON)
	}
}

func TestTodosToJSON(t *testing.T) {
	nt, err := NewTodo(1, "gud todo", nil, 2, DefaultPriority())
	if err != nil {
		t.Errorf("got err: %#v want no err", err.Error())
	}
	var ts Todos = Todos{nt}

	j, err := ts.ToJSON()

	if err != nil {
		t.Errorf("got err: %#v want no err", err.Error())
	}
	// TODO: long line
	expectedJSON := `[{"id":1,"name":"gud todo","completed":null,"userId":2,"priority":{"id":2,"name":"Normal","weight":2}}]`
	if string(j) != expectedJSON {
		t.Errorf("got %s\nwant %s", string(j), expectedJSON)
	}
}
