package todo

import "time"

// Todo models a single todo
// Name is a short description of the todo
// Completed is the time the todo has been completed or nil if the todo is incomplete
type Todo struct {
	Name      string    `json:"name"`
	Completed time.Time `json:"completed"`
}
