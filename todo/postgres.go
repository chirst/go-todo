package todo

import (
	"database/sql"
	"errors"
	"time"
)

// PostgresRepository persists todos
type PostgresRepository struct {
	DB *sql.DB
}

type postgresTodo struct {
	id        int
	name      string
	completed *time.Time
	userID    int
}

// GetTodos gets all todos in storage for a user
func (s *PostgresRepository) getTodos(userID int) ([]*Todo, error) {
	rows, err := s.DB.Query(`
		SELECT id, name, completed, user_id
		FROM todo
		WHERE user_id = $1
	`, userID)
	if err != nil {
		return nil, err
	}
	userTodos := make([]*Todo, 0)
	for rows.Next() {
		pgt := postgresTodo{}
		err := rows.Scan(&pgt.id, &pgt.name, &pgt.completed, &pgt.userID)
		if err != nil {
			return nil, err
		}
		t, err := NewTodo(pgt.id, pgt.name, pgt.completed, pgt.userID)
		if err != nil {
			return nil, err
		}
		userTodos = append(userTodos, t)
	}
	return userTodos, nil
}

// AddTodo adds a single todo to storage
func (s *PostgresRepository) addTodo(t Todo) (*Todo, error) {
	row := s.DB.QueryRow(`
		INSERT INTO todo (name, completed, user_id)
		VALUES ($1, $2, $3)
		RETURNING id, name, completed, user_id
	`, t.name, t.completed, t.userID)
	pgt := postgresTodo{}
	err := row.Scan(&pgt.id, &pgt.name, &pgt.completed, &pgt.userID)
	if err != nil {
		return nil, err
	}
	return NewTodo(pgt.id, pgt.name, pgt.completed, pgt.userID)
}

// Complete todo marks todo with the given id as complete and returns no error
// on success
func (s *PostgresRepository) completeTodo(userID, todoID int) error {
	result, err := s.DB.Exec(`
		UPDATE todo
		SET completed = timezone('utc', now())
		WHERE id = $1 AND user_id = $2
	`, todoID, userID)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected != 1 {
		return errors.New("no rows affected")
	}
	return nil
}
