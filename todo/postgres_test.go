package todo

import (
	"testing"

	"github.com/chirst/go-todo/config"
	"github.com/chirst/go-todo/database"
)

func TestPostgresGetTodos(t *testing.T) {
	if config.SkipPostgres() {
		t.Skip("skipped TestPostgresGetTodos")
	}
	db := database.InitDB()
	defer db.Close()
	r := &PostgresRepository{DB: db}

	todos, err := r.getTodos(1)

	if err != nil {
		t.Errorf("getTodos(1) returned err: %v", err)
	}
	if len(todos) != 2 {
		t.Errorf("getTodos(1) returned %v todos, want 1 todos", len(todos))
	}
}
