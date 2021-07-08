package todo

import (
	"testing"
	"time"
)

func TestNewTodo(t *testing.T) {
	tests := map[string]struct {
		id        int64
		name      string
		completed *time.Time
		userID    int64
		want      error
	}{
		"blank name": {
			id:        0,
			name:      "",
			completed: nil,
			userID:    0,
			want:      errNameRequired,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			_, got := NewTodo(tc.id, tc.name, tc.completed, tc.userID)
			if tc.want != got {
				t.Fatalf("expected %#v, got %#v", tc.want, got)
			}
		})
	}
}
