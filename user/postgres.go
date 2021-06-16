// TODO: integration tests
package user

import (
	"database/sql"
	"log"
)

// PostgresRepository persists users
type PostgresRepository struct {
	DB *sql.DB
}

func (s *PostgresRepository) addUser(u User) (*User, error) {
	result := s.DB.QueryRow(`
		INSERT INTO public.user (username, password)
		VALUES ($1, $2)
		RETURNING id, username, password
	`, u.Username, u.Password)
	insertedUser := new(User)
	err := result.Scan(&insertedUser.ID, &insertedUser.Username, &insertedUser.Password)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return insertedUser, nil
}

func (s *PostgresRepository) getUserByName(n string) (*User, error) {
	row := s.DB.QueryRow(`
		SELECT id, username, password
		FROM public.user
	`)
	user := new(User)
	err := row.Scan(&user.ID, &user.Username, &user.Password)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return user, nil
}
