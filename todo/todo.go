package todo

import (
	"encoding/json"
	"time"
)

// Todo models a single todo
type Todo struct {
	id int64

	// name is a short description of the todo
	name string

	// completed is the time the todo has been completed or nil when incomplete
	completed *time.Time

	// userID is the user this todo belongs to
	userID int64
}

// NewTodo provides a consistent way of creating a valid Todo
func NewTodo(
	id int64,
	name string,
	completed *time.Time,
	userID int64,
) (*Todo, error) {
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

// Todos is a list of Todo
type Todos []*Todo

type todoJSON struct {
	ID        int64      `json:"id"`
	Name      string     `json:"name"`
	Completed *time.Time `json:"completed"`
	UserID    int64      `json:"userId"`
	// TODO: HATEOAS to single todo
}

// ToJSON converts Todos to json
func (ts *Todos) ToJSON() ([]byte, error) {
	jts := []*todoJSON{}
	for _, t := range *ts {
		jts = append(jts, &todoJSON{
			t.id,
			t.name,
			t.completed,
			t.userID,
		})
	}
	return json.Marshal(jts)
}

// ToJSON converts a Todo to json
func (t *Todo) ToJSON() ([]byte, error) {
	return json.Marshal(&todoJSON{
		t.id,
		t.name,
		t.completed,
		t.userID,
	})
}
