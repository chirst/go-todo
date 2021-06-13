package todo

import "time"

// Todo models a single todo
// Name is a short description of the todo
// Completed is the time the todo has been completed or nil if the todo is incomplete
// UserID is the user this todo belongs to
type Todo struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Completed *time.Time `json:"completed"`
	UserID    int64      `json:"userId"`
}
