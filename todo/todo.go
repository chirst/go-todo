package todo

import (
	"encoding/json"
	"time"
)

// Todo models a single todo
type Todo struct {
	id int

	// name is a short description of the todo
	name string

	// completed is the time the todo has been completed or nil when incomplete
	completed *time.Time

	// userID is the user this todo belongs to
	userID int
}

// NewTodo provides a consistent way of creating a valid Todo
func NewTodo(
	id int,
	name string,
	completed *time.Time,
	userID int,
) (*Todo, error) {
	t := &Todo{
		id:        id,
		completed: completed,
		userID:    userID,
	}
	err := t.setName(name)
	if err != nil {
		return nil, err
	}
	return t, nil
}

func (t *Todo) setName(n string) error {
	if n == "" {
		return errNameRequired
	}
	t.name = n
	return nil
}

// Todos is a list of Todo
type Todos []*Todo

type todoJSON struct {
	ID        int        `json:"id"`
	Name      string     `json:"name"`
	Completed *time.Time `json:"completed"`
	UserID    int        `json:"userId"`
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

type Priorities []*Priority

type Priority struct {
	id     int
	name   string
	weight int
}

type priorityJSON struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Weight int    `json:"weight"`
}

func (p *Priorities) ToJSON() ([]byte, error) {
	ps := []*priorityJSON{}
	for _, priority := range *p {
		ps = append(ps, &priorityJSON{
			ID:     priority.id,
			Name:   priority.name,
			Weight: priority.weight,
		})
	}
	return json.Marshal(ps)
}
