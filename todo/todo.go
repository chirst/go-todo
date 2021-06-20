package todo

import "time"

// Todo models a single todo
// Name is a short description of the todo
// Completed is the time the todo has been completed or nil if the todo is incomplete
// UserID is the user this todo belongs to
type Todo struct {
	id        int64      `json:"id"`
	name      string     `json:"name"`
	completed *time.Time `json:"completed"`
	userID    int64      `json:"userId"`
}

// NewTodo provides a consistent way of creating a valid Todo
func NewTodo(id int64, name string, completed *time.Time, userID int64) (*Todo, error) {
	// TODO: test no name
	if name == "" {
		return nil, errNameRequired
	}
	return &Todo{
		id,
		name,
		completed,
		userID,
	}, nil
}

// UserID is the related user id for this todo
// probably a bad idea...
func (t *Todo) UserID() int64 {
	return t.userID
}

// Name is the name of the todo
func (t *Todo) Name() string {
	return t.name
}

// Completed is the time the todo was completed
func (t *Todo) Completed() *time.Time {
	return t.completed
}
