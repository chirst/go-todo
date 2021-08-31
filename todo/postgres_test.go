package todo

import (
	"database/sql"
	"testing"
	"time"

	"github.com/chirst/go-todo/database"
)

func TestPostgresGetTodos(t *testing.T) {

	t.Run("gets for user", func(t *testing.T) {
		db := database.OpenTestDB(t)
		defer db.Close()

		r := &PostgresRepository{DB: db}

		firstUserID := insertUser(db, "u1")
		secondUserID := insertUser(db, "u2")
		insertTodo(t, r, firstUserID)
		insertTodo(t, r, secondUserID)
		insertTodo(t, r, secondUserID)

		todos, err := r.getTodos(firstUserID)

		if err != nil {
			t.Errorf("getTodos(%v) returned err: %v", firstUserID, err)
		}
		if len(todos) != 1 {
			t.Errorf("getTodos(%v) returned %v todos, want 1 todo", firstUserID, len(todos))
		}
	})

	t.Run("excludes deleted", func(t *testing.T) {
		db := database.OpenTestDB(t)
		defer db.Close()

		r := &PostgresRepository{DB: db}

		userID := insertUser(db, "u1")
		insertTodo(t, r, userID)
		insertTodo(t, r, userID)
		toBeDeletedID := insertTodo(t, r, userID)
		r.deleteTodo(userID, toBeDeletedID)

		todos, err := r.getTodos(userID)
		if err != nil {
			t.Errorf("getTodos(%v) returned err: %v", userID, err)
		}
		if len(todos) != 2 {
			t.Errorf("getTodos(%v) returned %v todos, want 2 todos", userID, len(todos))
		}
	})
}

func TestPostgresAddTodo(t *testing.T) {
	db := database.OpenTestDB(t)
	defer db.Close()

	r := &PostgresRepository{DB: db}

	firstUserID := insertUser(db, "u1")
	todo, err := NewTodo(0, "gud name", nil, firstUserID)
	if err != nil {
		t.Errorf("unable to create todo")
	}

	_, err = r.addTodo(*todo)

	if err != nil {
		t.Errorf("addTodo err: %s", err.Error())
	}
}

func TestPostgresCompleteTodo(t *testing.T) {

	t.Run("completes", func(t *testing.T) {
		db := database.OpenTestDB(t)
		defer db.Close()

		r := &PostgresRepository{DB: db}

		userID := insertUser(db, "u1")
		todoID := insertTodo(t, r, userID)

		err := r.completeTodo(userID, todoID)

		if err != nil {
			t.Errorf("completeTodo(%v, %v) err: %s", userID, todoID, err.Error())
		}

		time := getTodoTime(db, todoID)
		if time == nil {
			t.Errorf(
				"completeTodo(%v, %v) expected completed not to be nil",
				userID,
				todoID,
			)
		}
	})

	t.Run("errs for wrong user", func(t *testing.T) {
		db := database.OpenTestDB(t)
		defer db.Close()

		r := &PostgresRepository{DB: db}

		firstUserID := insertUser(db, "u1")
		secondUserID := insertUser(db, "u2")
		todoID := insertTodo(t, r, firstUserID)

		err := r.completeTodo(secondUserID, todoID)

		if err == nil {
			t.Errorf(
				"completeTodo(%v, %v) expected err todo id: %v owned by user id: %v",
				secondUserID,
				todoID,
				todoID,
				firstUserID,
			)
		}
	})
}

func TestPostgresIncompleteTodo(t *testing.T) {

	t.Run("incompletes", func(t *testing.T) {
		db := database.OpenTestDB(t)
		defer db.Close()

		r := &PostgresRepository{DB: db}

		userID := insertUser(db, "u1")
		todoID := insertCompleteTodo(t, r, userID)

		err := r.incompleteTodo(userID, todoID)

		if err != nil {
			t.Errorf("incompleteTodo(%v, %v) err: %s", userID, todoID, err.Error())
		}

		time := getTodoTime(db, todoID)
		if time != nil {
			t.Errorf(
				"incompleteTodo(%v, %v) expected complete to be nil",
				userID,
				todoID,
			)
		}
	})

	t.Run("errs for wrong user", func(t *testing.T) {
		db := database.OpenTestDB(t)
		defer db.Close()

		r := &PostgresRepository{DB: db}

		firstUserID := insertUser(db, "u1")
		secondUserID := insertUser(db, "u2")
		todoID := insertCompleteTodo(t, r, firstUserID)

		err := r.incompleteTodo(secondUserID, todoID)

		if err == nil {
			t.Errorf(
				"incompleteTodo(%v, %v) expected err todo id: %v owned by user id: %v",
				secondUserID,
				todoID,
				todoID,
				firstUserID,
			)
		}
	})
}

func TestPostgresDeleteTodo(t *testing.T) {

	t.Run("deletes", func(t *testing.T) {
		db := database.OpenTestDB(t)
		defer db.Close()

		r := &PostgresRepository{DB: db}

		userID := insertUser(db, "u1")
		todoID := insertTodo(t, r, userID)

		err := r.deleteTodo(userID, todoID)

		if err != nil {
			t.Errorf("deleteTodo(%v, %v) err: %s", userID, todoID, err.Error())
		}

		time := getDeletedTime(db, todoID)
		if time == nil {
			t.Errorf(
				"deleteTodo(%v, %v) expected deleted not to be nil",
				userID,
				todoID,
			)
		}
	})

	t.Run("errs for wrong user", func(t *testing.T) {
		db := database.OpenTestDB(t)
		defer db.Close()

		r := &PostgresRepository{DB: db}

		firstUserID := insertUser(db, "u1")
		secondUserID := insertUser(db, "u2")
		todoID := insertTodo(t, r, firstUserID)

		err := r.deleteTodo(secondUserID, todoID)

		if err == nil {
			t.Errorf(
				"deleteTodo(%v, %v) expected err todo id: %v owned by user id: %v",
				secondUserID,
				todoID,
				todoID,
				firstUserID,
			)
		}
	})
}

func insertUser(db *sql.DB, name string) int {
	result := db.QueryRow(`
		INSERT INTO public.user (username, password)
		VALUES ($1, '12345')
		RETURNING id
	`, name)
	var id int
	result.Scan(&id)
	return id
}

func insertTodo(t *testing.T, r *PostgresRepository, userID int) int {
	todo, err := NewTodo(0, "gud todo", nil, userID)
	if err != nil {
		t.Fatalf("unable to create todo")
	}
	addedTodo, err := r.addTodo(*todo)
	if err != nil {
		t.Fatalf("unable to add todo")
	}
	return addedTodo.id
}

func insertCompleteTodo(t *testing.T, r *PostgresRepository, userID int) int {
	now := time.Now()
	todo, err := NewTodo(0, "complete todo", &now, userID)
	if err != nil {
		t.Fatalf("unable to create todo")
	}
	addedTodo, err := r.addTodo(*todo)
	if err != nil {
		t.Fatalf("unable to add todo")
	}
	return addedTodo.id
}

func getTodoTime(db *sql.DB, todoID int) *time.Time {
	result := db.QueryRow("SELECT completed FROM todo WHERE id = $1", todoID)
	var completed *time.Time
	result.Scan(&completed)
	return completed
}

func getDeletedTime(db *sql.DB, todoID int) *time.Time {
	result := db.QueryRow("SELECT deleted FROM todo WHERE id = $1", todoID)
	var deleted *time.Time
	result.Scan(&deleted)
	return deleted
}
