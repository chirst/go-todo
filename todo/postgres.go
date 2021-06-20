// TODO: integration tests
// TODO: filter gets by user
// TODO: fix listing columns in query and scan
package todo

import (
	"database/sql"
	"log"
	"time"
)

// PostgresRepository persists todos
type PostgresRepository struct {
	DB *sql.DB
}

type dbtodo struct {
	id        int64
	name      string
	completed *time.Time
}

// GetTodos gets all todos in storage for a user
func (s *PostgresRepository) getTodos(userID int64) ([]*Todo, error) {
	rows, err := s.DB.Query(`
		SELECT id, name, completed
		FROM todo
	`)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	userTodos := make([]*Todo, 0)
	for rows.Next() {
		dbt := new(dbtodo)
		err := rows.Scan(&dbt.id, &dbt.name, &dbt.completed)
		if err != nil {
			log.Print(err)
			return nil, err
		}
		t, err := NewTodo(dbt.id, dbt.name, dbt.completed, 0)
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
		INSERT INTO todo (name, completed)
		VALUES ($1, $2)
		RETURNING id, name, completed
	`, t.Name(), t.Completed())
	dbt := new(dbtodo)
	err := row.Scan(&dbt.id, &dbt.name, &dbt.completed)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return NewTodo(dbt.id, dbt.name, dbt.completed, 0)
}
