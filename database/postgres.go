package database

import (
	"database/sql"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"runtime"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq"

	"github.com/chirst/go-todo/config"
)

// InitDB opens a db connection, runs migrations, and exits if anything goes wrong
func InitDB() *sql.DB {
	db, err := sql.Open("postgres", config.PostgresSourceName())
	if err != nil {
		log.Fatalf("failed to connect to db %s", err.Error())
	}
	runMigrations(db)
	return db
}

// InitTestDB registers a transactional database used by OpenTestDB
func InitTestDB() {
	txdb.Register("txdb", "postgres", config.PostgresSourceName())
}

// OpenTestDB returns a transactional database rolled back when it is closed
func OpenTestDB() *sql.DB {
	db, err := sql.Open("txdb", "identifier")
	if err != nil {
		log.Fatal(err)
	}
	runMigrations(db)
	return db
}

func runMigrations(db *sql.DB) {
	_, b, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalf("unable to get root directory")
	}
	d := path.Join(path.Dir(b))
	rootPath := filepath.Dir(d)
	files, err := ioutil.ReadDir(path.Join(
		rootPath,
		"database",
		"migrations",
	))
	if err != nil {
		log.Fatalf("failed to read migration directory %s", err.Error())
	}
	for _, f := range files {
		log.Printf("starting migration: %s\n", f.Name())
		q, err := ioutil.ReadFile(path.Join(
			rootPath,
			"database",
			"migrations",
			f.Name(),
		))
		if err != nil {
			log.Fatalf("failed to read migration file %s", err.Error())
		}
		if _, err = db.Exec(string(q)); err != nil {
			log.Fatalf("failed to run migration %s", err.Error())
		}
		log.Printf("finished migration: %s\n", f.Name())
	}
}
