package server

import (
	"database/sql"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/chirst/go-todo/database"
)

func setupTest(t *testing.T) (*httptest.Server, func()) {
	var db *sql.DB
	if os.Getenv("TEST_POSTGRES") != "" {
		db = database.OpenTestDB(t)
	}
	router := GetRouter(db)
	testServer := httptest.NewServer(router)
	teardown := func() {
		if db != nil {
			db.Close()
		}
		testServer.Close()
	}
	return testServer, teardown
}
