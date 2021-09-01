package user

import (
	"database/sql"
)

// PostgresRepository persists users
type PostgresRepository struct {
	DB *sql.DB
}

type postgresUser struct {
	ID       int
	Username string
	Password string
}

func (s *PostgresRepository) addUser(u User) (*User, error) {
	result := s.DB.QueryRow(`
		INSERT INTO public.user (username, password)
		VALUES ($1, $2)
		RETURNING id, username, password
	`, u.username, u.password)
	insertedUser := postgresUser{}
	err := result.Scan(
		&insertedUser.ID,
		&insertedUser.Username,
		&insertedUser.Password,
	)
	if err != nil {
		return nil, err
	}
	domainUser, err := NewUser(
		insertedUser.ID,
		insertedUser.Username,
		insertedUser.Password,
	)
	if err != nil {
		return nil, err
	}
	return domainUser, nil
}

func (s *PostgresRepository) getUserByName(n string) (*User, error) {
	row := s.DB.QueryRow(`
		SELECT id, username, password
		FROM public.user
		WHERE username = $1
	`, n)
	pgUser := postgresUser{}
	err := row.Scan(&pgUser.ID, &pgUser.Username, &pgUser.Password)
	if err != nil {
		return nil, err
	}
	domainUser, err := NewUser(pgUser.ID, pgUser.Username, pgUser.Password)
	if err != nil {
		return nil, err
	}
	return domainUser, nil
}
