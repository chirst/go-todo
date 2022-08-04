package todo

import (
	"encoding/json"
	"time"
)

// Todo models a single todo
type Todo todoModel

// Todos is a list of Todo
type Todos []*Todo

// todoModel should only be created by calling `newTodo`
type todoModel struct {
	id int

	// name is a short description of the todo
	name string

	// completed is the time the todo has been completed or nil when incomplete
	completed *time.Time

	// userID is the user this todo belongs to
	userID int

	// priority is the priority this todo has assigned to it
	priority priorityModel
}

// newTodo provides a consistent way of creating a valid Todo.
func newTodo(
	id int,
	name string,
	completed *time.Time,
	userID int,
	priority priorityModel,
) (*Todo, error) {
	t := &Todo{
		id:        id,
		completed: completed,
		userID:    userID,
		priority:  priority,
	}
	err := t.setName(name)
	if err != nil {
		return nil, err
	}
	return t, nil
}

// MustNewTodo provides a way to create a todo for testing purposes. If an error
// occurs creating a Todo there will be a panic. This function is not meant to
// be called by application code.
func MustNewTodo(
	id int,
	name string,
	completed *time.Time,
	userID int,
) *Todo {
	t, err := newTodo(id, name, completed, userID, defaultPriority())
	if err != nil {
		panic(err)
	}
	return t
}

func (t *Todo) setName(n string) error {
	if n == "" {
		return errNameRequired
	}
	t.name = n
	return nil
}

type todoJSON struct {
	ID        int          `json:"id"`
	Name      string       `json:"name"`
	Completed *time.Time   `json:"completed"`
	UserID    int          `json:"userId"`
	Priority  priorityJSON `json:"priority"`
}

// ToJSON serializes a list of todos to JSON
func (ts *Todos) ToJSON() ([]byte, error) {
	jts := []*todoJSON{}
	for _, t := range *ts {
		p := priorityJSON{
			ID:     t.priority.id,
			Name:   t.priority.name,
			Weight: t.priority.weight,
		}
		jts = append(jts, &todoJSON{
			t.id,
			t.name,
			t.completed,
			t.userID,
			p,
		})
	}
	return json.Marshal(jts)
}

// ToJSON serializes a single todo to JSON
func (t *Todo) ToJSON() ([]byte, error) {
	p := priorityJSON{
		ID:     t.priority.id,
		Name:   t.priority.name,
		Weight: t.priority.weight,
	}
	return json.Marshal(&todoJSON{
		ID:        t.id,
		Name:      t.name,
		Completed: t.completed,
		UserID:    t.userID,
		Priority:  p,
	})
}

// Priority is the priority of a todo
type Priority priorityModel

// Priorities is a list of priority
type Priorities []*priorityModel

type priorityModel struct {
	id     int
	name   string
	weight int
}

// defaultPriority skips the DB and returns the default priority for a Todo.
func defaultPriority() priorityModel {
	return priorityModel{
		id:     2,
		name:   "Normal",
		weight: 2,
	}
}

type priorityJSON struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Weight int    `json:"weight"`
}

// ToJSON serializes priorities to JSON
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
