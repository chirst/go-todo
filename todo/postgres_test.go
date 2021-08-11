package todo

import (
	"database/sql"
	"testing"

	"github.com/chirst/go-todo/database"
)

func TestPostgresGetTodos(t *testing.T) {
	db := database.OpenTestDB(t)
	defer db.Close()

	r := &PostgresRepository{DB: db}

	firstUserID := insertUser(db, "u1")
	secondUserID := insertUser(db, "u2")
	insertTodo(t, r, firstUserID)
	insertTodo(t, r, secondUserID)

	todos, err := r.getTodos(firstUserID)

	if err != nil {
		t.Errorf("getTodos(1) returned err: %v", err)
	}
	if len(todos) != 1 {
		t.Errorf("getTodos(1) returned %v todos, want 1 todos", len(todos))
	}
}

func TestPostgresAddTodo(t *testing.T) {
	db := database.OpenTestDB(t)
	defer db.Close()

	r := &PostgresRepository{DB: db}

	firstUserID := insertUser(db, "u1")
	todo, _ := NewTodo(0, "gud name", nil, firstUserID)

	_, err := r.addTodo(*todo)

	if err != nil {
		t.Errorf("addTodo() err: %s", err.Error())
	}
}

func insertUser(db *sql.DB, name string) int64 {
	result := db.QueryRow(`
		INSERT INTO public.user (username, password)
		VALUES ($1, '12345')
		RETURNING id
	`, name)
	var id int64
	result.Scan(&id)
	return id
}

func insertTodo(t *testing.T, r *PostgresRepository, userID int64) {
	todo, err := NewTodo(0, "gud todo", nil, userID)
	if err != nil {
		t.Fatalf("unable to create todo")
	}
	r.addTodo(*todo)
}
