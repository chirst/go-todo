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
	id         int
	name       string
	completed  *time.Time
	userID     int
	priorityID int
}

// GetTodos gets all todos in storage for a user
func (s *PostgresRepository) getTodos(userID int) ([]*Todo, error) {
	rows, err := s.DB.Query(`
		SELECT id, name, completed, user_id
		FROM todo
		WHERE user_id = $1 AND deleted IS NULL
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
		t, err := NewTodo(pgt.id, pgt.name, pgt.completed, pgt.userID, &pgt.priorityID)
		if err != nil {
			return nil, err
		}
		userTodos = append(userTodos, t)
	}
	return userTodos, nil
}

func (r *PostgresRepository) getTodo(userID, todoID int) (*Todo, error) {
	row := r.DB.QueryRow(`
		SELECT id, name, completed, user_id
		FROM todo
		WHERE user_id = $1 AND id = $2
	`, userID, todoID)
	pgt := postgresTodo{}
	err := row.Scan(&pgt.id, &pgt.name, &pgt.completed, &pgt.userID)
	if err != nil {
		return nil, err
	}
	return NewTodo(pgt.id, pgt.name, pgt.completed, pgt.userID, &pgt.priorityID)
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
	return NewTodo(pgt.id, pgt.name, pgt.completed, pgt.userID, &pgt.priorityID)
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

func (s *PostgresRepository) incompleteTodo(userID, todoID int) error {
	result, err := s.DB.Exec(`
		UPDATE todo
		SET completed = NULL
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

func (r *PostgresRepository) updateName(userID, todoID int, name string) error {
	result, err := r.DB.Exec(`
		UPDATE todo
		SET name = $3
		WHERE id = $2 AND user_id = $1
	`, userID, todoID, name)
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

func (s *PostgresRepository) deleteTodo(userID, todoID int) error {
	result, err := s.DB.Exec(`
		UPDATE todo
		SET deleted = timezone('utc', now())
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

func (s *PostgresRepository) getPriorities() (Priorities, error) {
	rows, err := s.DB.Query("SELECT id, name, weight FROM priority")
	if err != nil {
		return nil, err
	}
	ps := Priorities{}
	for rows.Next() {
		p := &Priority{}
		err := rows.Scan(&p.id, &p.name, &p.weight)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	return ps, nil
}
