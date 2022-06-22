package todo

import (
	"database/sql"
	"testing"
	"time"

	"github.com/chirst/go-todo/database"
)

func TestPostgresGetTodos(t *testing.T) {

	t.Run("gets for user", func(t *testing.T) {
		db, teardown := database.OpenTestDB(t)
		defer teardown()

		r := &PostgresRepository{DB: db}

		firstUserID := insertUser(t, db, "u1")
		secondUserID := insertUser(t, db, "u2")
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
		db, teardown := database.OpenTestDB(t)
		defer teardown()

		r := &PostgresRepository{DB: db}

		userID := insertUser(t, db, "u1")
		insertTodo(t, r, userID)
		insertTodo(t, r, userID)
		toBeDeletedID := insertTodo(t, r, userID)
		err := r.deleteTodo(userID, toBeDeletedID)
		if err != nil {
			t.Errorf("deleteTodo want no error got: %s", err.Error())
		}

		todos, err := r.getTodos(userID)
		if err != nil {
			t.Errorf("getTodos(%v) returned err: %v", userID, err)
		}
		if len(todos) != 2 {
			t.Errorf("getTodos(%v) returned %v todos, want 2 todos", userID, len(todos))
		}
	})
}

func TestPostgresGetTodo(t *testing.T) {
	db, teardown := database.OpenTestDB(t)
	defer teardown()

	r := &PostgresRepository{DB: db}

	userID := insertUser(t, db, "u1")
	todoID := insertTodo(t, r, userID)

	todo, err := r.getTodo(userID, todoID)

	if err != nil {
		t.Fatalf("expected no error got error %s", err.Error())
	}
	if todo == nil {
		t.Fatalf("expected todo got nil")
	}
	if todo.id != todoID {
		t.Fatalf("expected todo with id: %v, got todo with id: %v", todoID, todo.id)
	}
	if todo.priority.id != 2 {
		t.Fatalf("expected todo with default priorityID of 2, got priorityID: %v", todo.priority.id)
	}
}

func TestPostgresAddTodo(t *testing.T) {
	db, teardown := database.OpenTestDB(t)
	defer teardown()

	r := &PostgresRepository{DB: db}

	firstUserID := insertUser(t, db, "u1")

	_, err := r.addTodo("gud name", nil, firstUserID, defaultPriority())

	if err != nil {
		t.Errorf("addTodo err: %s", err.Error())
	}
}

func TestPostgresCompleteTodo(t *testing.T) {

	t.Run("completes", func(t *testing.T) {
		db, teardown := database.OpenTestDB(t)
		defer teardown()

		r := &PostgresRepository{DB: db}

		userID := insertUser(t, db, "u1")
		todoID := insertTodo(t, r, userID)

		err := r.completeTodo(userID, todoID)

		if err != nil {
			t.Errorf("completeTodo(%v, %v) err: %s", userID, todoID, err.Error())
		}

		time := getTodoTime(t, db, todoID)
		if time == nil {
			t.Errorf(
				"completeTodo(%v, %v) expected completed not to be nil",
				userID,
				todoID,
			)
		}
	})

	t.Run("errs for wrong user", func(t *testing.T) {
		db, teardown := database.OpenTestDB(t)
		defer teardown()

		r := &PostgresRepository{DB: db}

		firstUserID := insertUser(t, db, "u1")
		secondUserID := insertUser(t, db, "u2")
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
		db, teardown := database.OpenTestDB(t)
		defer teardown()

		r := &PostgresRepository{DB: db}

		userID := insertUser(t, db, "u1")
		todoID := insertCompleteTodo(t, r, userID)

		err := r.incompleteTodo(userID, todoID)

		if err != nil {
			t.Errorf("incompleteTodo(%v, %v) err: %s", userID, todoID, err.Error())
		}

		time := getTodoTime(t, db, todoID)
		if time != nil {
			t.Errorf(
				"incompleteTodo(%v, %v) expected complete to be nil",
				userID,
				todoID,
			)
		}
	})

	t.Run("errs for wrong user", func(t *testing.T) {
		db, teardown := database.OpenTestDB(t)
		defer teardown()

		r := &PostgresRepository{DB: db}

		firstUserID := insertUser(t, db, "u1")
		secondUserID := insertUser(t, db, "u2")
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

func TestPostgresUpdateName(t *testing.T) {
	db, teardown := database.OpenTestDB(t)
	defer teardown()

	r := &PostgresRepository{DB: db}

	userID := insertUser(t, db, "u1")
	todoID := insertTodo(t, r, userID)
	newName := "guddest new name"

	err := r.updateName(userID, todoID, newName)

	if err != nil {
		t.Fatalf("expected err to be nil got err: %s", err.Error())
	}

	updatedTodo, err := r.getTodo(userID, todoID)
	if err != nil {
		t.Fatalf("failed to lookup todo with userID: %v, todoID %v", userID, todoID)
	}

	if updatedTodo.name != newName {
		t.Fatalf("expected todo to have name: %s, got: %s", newName, updatedTodo.name)
	}
}

func TestPostgresUpdatePriority(t *testing.T) {
	db, teardown := database.OpenTestDB(t)
	defer teardown()

	r := &PostgresRepository{DB: db}

	userID := insertUser(t, db, "u1")
	todoID := insertTodo(t, r, userID)
	newPriorityID := 1

	err := r.updatePriority(userID, todoID, newPriorityID)

	if err != nil {
		t.Fatalf("expected err to be nil got err: %s", err.Error())
	}

	updatedTodo, err := r.getTodo(userID, todoID)
	if err != nil {
		t.Fatalf("failed to lookup todo with userID: %v, todoID %v", userID, todoID)
	}

	if updatedTodo.priority.id != newPriorityID {
		t.Fatalf(
			"expected todo to have priorityID: %v, got: %v",
			newPriorityID,
			updatedTodo.priority.id,
		)
	}
}

func TestPostgresDeleteTodo(t *testing.T) {

	t.Run("deletes", func(t *testing.T) {
		db, teardown := database.OpenTestDB(t)
		defer teardown()

		r := &PostgresRepository{DB: db}

		userID := insertUser(t, db, "u1")
		todoID := insertTodo(t, r, userID)

		err := r.deleteTodo(userID, todoID)

		if err != nil {
			t.Errorf("deleteTodo(%v, %v) err: %s", userID, todoID, err.Error())
		}

		time := getDeletedTime(t, db, todoID)
		if time == nil {
			t.Errorf(
				"deleteTodo(%v, %v) expected deleted not to be nil",
				userID,
				todoID,
			)
		}
	})

	t.Run("errs for wrong user", func(t *testing.T) {
		db, teardown := database.OpenTestDB(t)
		defer teardown()

		r := &PostgresRepository{DB: db}

		firstUserID := insertUser(t, db, "u1")
		secondUserID := insertUser(t, db, "u2")
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

func TestPostgresGetPriorities(t *testing.T) {
	db, teardown := database.OpenTestDB(t)
	defer teardown()

	r := &PostgresRepository{DB: db}

	ps, err := r.getPriorities()

	if err != nil {
		t.Errorf("got error: %#v, want no error", err.Error())
	}
	if psLen := len(ps); psLen < 1 {
		t.Errorf("got %#v priorities, want more than 1", psLen)
	}
}

func TestPostgresGetPriority(t *testing.T) {
	db, teardown := database.OpenTestDB(t)
	defer teardown()

	r := &PostgresRepository{DB: db}

	queryID := 1
	p, err := r.getPriority(queryID)

	if err != nil {
		t.Errorf("got err: %#v, want no error", err.Error())
	}
	if p.id != queryID {
		t.Errorf("got id: %#v, want %#v", p.id, queryID)
	}
}

func insertUser(t *testing.T, db *sql.DB, name string) int {
	result := db.QueryRow(`
		INSERT INTO public.user (username, password)
		VALUES ($1, '12345')
		RETURNING id
	`, name)
	var id int
	err := result.Scan(&id)
	if err != nil {
		t.Fatalf(err.Error())
	}
	return id
}

func insertTodo(t *testing.T, r *PostgresRepository, userID int) int {
	addedTodo, err := r.addTodo("gud todo", nil, userID, defaultPriority())
	if err != nil {
		t.Fatalf("unable to add todo")
	}
	return addedTodo.id
}

func insertCompleteTodo(t *testing.T, r *PostgresRepository, userID int) int {
	now := time.Now()
	addedTodo, err := r.addTodo("complete todo", &now, userID, defaultPriority())
	if err != nil {
		t.Fatalf("unable to add todo")
	}
	return addedTodo.id
}

func getTodoTime(t *testing.T, db *sql.DB, todoID int) *time.Time {
	result := db.QueryRow("SELECT completed FROM todo WHERE id = $1", todoID)
	var completed *time.Time
	err := result.Scan(&completed)
	if err != nil {
		t.Fatalf(err.Error())
	}
	return completed
}

func getDeletedTime(t *testing.T, db *sql.DB, todoID int) *time.Time {
	result := db.QueryRow("SELECT deleted FROM todo WHERE id = $1", todoID)
	var deleted *time.Time
	err := result.Scan(&deleted)
	if err != nil {
		t.Fatalf(err.Error())
	}
	return deleted
}
