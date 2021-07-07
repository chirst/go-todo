package todo

// TODO: integration tests

import (
	"database/sql"
	"log"
	"time"
)

// PostgresRepository persists todos
type PostgresRepository struct {
	DB *sql.DB
}

type postgresTodo struct {
	id        int64
	name      string
	completed *time.Time
	userID    int64
}

// GetTodos gets all todos in storage for a user
func (s *PostgresRepository) getTodos(userID int64) ([]*Todo, error) {
	rows, err := s.DB.Query(`
		SELECT id, name, completed, user_id
		FROM todo
		WHERE user_id = $1
	`, userID)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	userTodos := make([]*Todo, 0)
	for rows.Next() {
		pgt := postgresTodo{}
		err := rows.Scan(&pgt.id, &pgt.name, &pgt.completed, &pgt.userID)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		t, err := NewTodo(pgt.id, pgt.name, pgt.completed, pgt.userID)
		if err != nil {
			log.Print(err)
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
		log.Print(err)
		return nil, err
	}
	return NewTodo(pgt.id, pgt.name, pgt.completed, pgt.userID)
}
