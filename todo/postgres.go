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

type postgresPriority struct {
	id     int
	name   string
	weight int
}

func (r *PostgresRepository) getTodos(userID int) ([]*Todo, error) {
	rows, err := r.DB.Query(`
		SELECT
			t.id,
			t.name,
			t.completed,
			t.user_id,
			p.id,
			p.name,
			p.weight
		FROM todo t
			JOIN priority p ON p.id = t.priority_id
		WHERE t.user_id = $1 AND t.deleted IS NULL
	`, userID)
	defer func() {
		err = rows.Close()
	}()
	if err != nil {
		return nil, err
	}
	userTodos := make([]*Todo, 0)
	for rows.Next() {
		pgt := postgresTodo{}
		pgp := postgresPriority{}
		err := rows.Scan(
			&pgt.id,
			&pgt.name,
			&pgt.completed,
			&pgt.userID,
			&pgp.id,
			&pgp.name,
			&pgp.weight,
		)
		if err != nil {
			return nil, err
		}
		priority := priorityModel(pgp)
		t, err := newTodo(pgt.id, pgt.name, pgt.completed, pgt.userID, priority)
		if err != nil {
			return nil, err
		}
		userTodos = append(userTodos, t)
	}
	if err = rows.Err(); rows.Err() != nil {
		return nil, err
	}
	return userTodos, err
}

func (r *PostgresRepository) getTodo(userID, todoID int) (*Todo, error) {
	row := r.DB.QueryRow(`
		SELECT
			t.id,
			t.name,
			t.completed,
			t.user_id,
			t.priority_id,
			p.id,
			p.name,
			p.weight
		FROM todo t
			JOIN priority p ON p.id = t.priority_id
		WHERE t.user_id = $1 AND t.id = $2
	`, userID, todoID)
	pgt := postgresTodo{}
	pgp := postgresPriority{}
	err := row.Scan(
		&pgt.id,
		&pgt.name,
		&pgt.completed,
		&pgt.userID,
		&pgt.priorityID,
		&pgp.id,
		&pgp.name,
		&pgp.weight,
	)
	if err != nil {
		return nil, err
	}
	priority := priorityModel(pgp)
	return newTodo(pgt.id, pgt.name, pgt.completed, pgt.userID, priority)
}

func (r *PostgresRepository) addTodo(
	name string,
	completed *time.Time,
	userID int,
	priority priorityModel,
) (*Todo, error) {
	row := r.DB.QueryRow(`
		INSERT INTO todo (name, completed, user_id)
		VALUES ($1, $2, $3)
		RETURNING id, name, completed, user_id
	`, name, completed, userID)
	pgt := postgresTodo{}
	err := row.Scan(&pgt.id, &pgt.name, &pgt.completed, &pgt.userID)
	if err != nil {
		return nil, err
	}
	return newTodo(pgt.id, pgt.name, pgt.completed, pgt.userID, priority)
}

func (r *PostgresRepository) completeTodo(userID, todoID int) error {
	result, err := r.DB.Exec(`
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

func (r *PostgresRepository) incompleteTodo(userID, todoID int) error {
	result, err := r.DB.Exec(`
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

func (r *PostgresRepository) updatePriority(
	userID,
	todoID,
	priorityID int,
) error {
	result, err := r.DB.Exec(`
		UPDATE todo
		SET priority_id = $3
		WHERE id = $2 AND user_id = $1
	`, userID, todoID, priorityID)
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

func (r *PostgresRepository) deleteTodo(userID, todoID int) error {
	result, err := r.DB.Exec(`
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

func (r *PostgresRepository) getPriorities() (Priorities, error) {
	rows, err := r.DB.Query("SELECT id, name, weight FROM priority")
	defer func() {
		err = rows.Close()
	}()
	if err != nil {
		return nil, err
	}
	ps := Priorities{}
	for rows.Next() {
		p := &priorityModel{}
		err := rows.Scan(&p.id, &p.name, &p.weight)
		if err != nil {
			return nil, err
		}
		ps = append(ps, p)
	}
	if err = rows.Err(); err != nil {
		return nil, err
	}
	return ps, err
}

func (r *PostgresRepository) getPriority(
	priorityID int,
) (*priorityModel, error) {
	row := r.DB.QueryRow(`
		SELECT id, name, weight
		FROM priority
		WHERE id = $1
	`, priorityID)
	pgp := &postgresPriority{}
	err := row.Scan(&pgp.id, &pgp.name, &pgp.weight)
	if err != nil {
		return nil, err
	}
	return &priorityModel{
		id:     pgp.id,
		name:   pgp.name,
		weight: pgp.weight,
	}, nil
}
