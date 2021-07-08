package todo

import (
	"reflect"
	"testing"
	"time"
)

func TestNewTodo(t *testing.T) {
	tests := map[string]struct {
		id        int64
		name      string
		completed *time.Time
		userID    int64
		want      *Todo
		want2     error
	}{
		"blank name": {
			id:        0,
			name:      "",
			completed: nil,
			userID:    0,
			want:      nil,
			want2:     errNameRequired,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got, got2 := NewTodo(tc.id, tc.name, tc.completed, tc.userID)
			if !reflect.DeepEqual(tc.want, got) {
				t.Fatalf("expected %#v, got: %#v", tc.want, got)
			}
			if tc.want2 != got2 {
				t.Fatalf("expected %#v, got %#v", tc.want2, got2)
			}
		})
	}
}
