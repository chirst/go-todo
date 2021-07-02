package todo

import (
	"encoding/json"
	"time"
)

// Todos is a list of Todo
type Todos []*Todo

// ToJSON converts Todos to json
func (ts *Todos) ToJSON() ([]byte, error) {
	return json.Marshal(ts)
}

// Todo models a single todo
// name is a short description of the todo
// completed is the time the todo has been completed or nil if the todo is incomplete
// userID is the user this todo belongs to
type Todo struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Completed *time.Time `json:"completed"`
	UserID    int64      `json:"userId"`
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

func (t *Todo) ToJSON() ([]byte, error) {
	return json.Marshal(t)
}
