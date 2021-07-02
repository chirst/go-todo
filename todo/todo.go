package todo

import "time"

// Todo models a single todo
// name is a short description of the todo
// completed is the time the todo has been completed or nil if the todo is incomplete
// userID is the user this todo belongs to
type Todo struct {
	id        int64
	name      string
	completed *time.Time
	userID    int64
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

func (t *Todo) ID() int64 {
	return t.id
}

// Name is the name of the todo
func (t *Todo) Name() string {
	return t.name
}

// Completed is the time the todo was completed
func (t *Todo) Completed() *time.Time {
	return t.completed
}

// UserID is the related user id for this todo
func (t *Todo) UserID() int64 {
	return t.userID
}
