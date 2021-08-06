package database

import (
	"database/sql"
	"io/ioutil"
	"log"
	"path"
	"path/filepath"
	"runtime"

	_ "github.com/lib/pq"

	"github.com/chirst/go-todo/config"
)

// InitDB opens a db connection, runs migrations, and exits if anything goes wrong
func InitDB() *sql.DB {
	db, err := sql.Open("postgres", config.PostgresSourceName())
	if err != nil {
		log.Fatalf("failed to connect to db %s", err.Error())
	}
	_, b, _, ok := runtime.Caller(0)
	if !ok {
		log.Fatalf("unable to get root directory %s", err.Error())
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
