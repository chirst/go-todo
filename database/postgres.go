package database

import (
	"database/sql"
	"io/ioutil"
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"testing"

	"github.com/DATA-DOG/go-txdb"
	_ "github.com/lib/pq" // Database driver for Postgres

	"github.com/chirst/go-todo/config"
)

// InitDB opens a db connection, runs migrations, and exits if anything goes
// wrong.
func InitDB() *sql.DB {
	db, err := sql.Open("postgres", config.PostgresSourceName())
	if err != nil {
		log.Fatalf("failed to connect to db %s", err.Error())
	}
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
	return db
}

var registered bool

// OpenTestDB returns a transactional database rolled back when it is closed.
func OpenTestDB(t *testing.T) *sql.DB {
	if os.Getenv("TEST_POSTGRES") == "" {
		t.Skip("Skipped Postgres test. Define TEST_POSTGRES env variable to run postgres tests")
	}
	if !registered {
		db := InitDB()
		db.Close()
		txdb.Register("txdb", "postgres", config.PostgresSourceName())
		registered = true
	}
	db, err := sql.Open("txdb", "identifier")
	if err != nil {
		log.Fatal(err)
	}
	return db
}
