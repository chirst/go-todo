// TODO: integration tests
// TODO: return errors
// TODO: filter gets by user
package todo

import (
	"database/sql"
	"log"
)

// PostgresRepository persists todos
type PostgresRepository struct {
	DB *sql.DB
}

// GetTodos gets all todos in storage for a user
func (s *PostgresRepository) getTodos(userID int64) []*Todo {
	// TODO: handle this error!
	rows, _ := s.DB.Query(`
		SELECT id, name, completed
		FROM todo
	`)
	userTodos := make([]*Todo, 0)
	for rows.Next() {
		userTodo := new(Todo)
		err := rows.Scan(&userTodo.ID, &userTodo.Name, &userTodo.Completed)
		if err != nil {
			log.Print(err)
			return nil
		}
		userTodos = append(userTodos, userTodo)
	}
	return userTodos
}

// AddTodo adds a single todo to storage
func (s *PostgresRepository) addTodo(t Todo) *Todo {
	row := s.DB.QueryRow(`
		INSERT INTO todo (name, completed)
		VALUES ($1, $2)
		RETURNING id, name, completed
	`, t.Name, t.Completed)
	insertedTodo := new(Todo)
	err := row.Scan(&insertedTodo.ID, &insertedTodo.Name, &insertedTodo.Completed)
	if err != nil {
		log.Print(err)
		return nil
	}
	return insertedTodo
}
