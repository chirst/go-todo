package main

import (
	"bytes"
	"database/sql"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/chirst/go-todo/database"
	"github.com/kinbiko/jsonassert"
)

func setupTest(t *testing.T) (*httptest.Server, func()) {
	var db *sql.DB
	if os.Getenv("TEST_POSTGRES") != "" {
		db = database.OpenTestDB(t)
	}
	router := getRouter(db)
	testServer := httptest.NewServer(router)
	teardown := func() {
		if db != nil {
			db.Close()
		}
		testServer.Close()
	}
	return testServer, teardown
}

func TestBehavior(t *testing.T) {
	testServer, teardown := setupTest(t)
	defer teardown()

	addUser(t, testServer)
	bearer := loginUser(t, testServer)
	addTodo(t, testServer, bearer)
	getPriorities(t, testServer, bearer)
	getTodos(t, testServer, bearer)
}

func addUser(t *testing.T, testServer *httptest.Server) {
	var body = []byte(`{
		"username": "gud",
		"password": "wordpass"
	}`)
	r, err := http.NewRequest("POST", testServer.URL+"/users", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("err creating new request: %s", err.Error())
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatalf("err doing request: %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("got status code: %d, want %d", resp.StatusCode, http.StatusOK)
	}
	assertJSONEqual(t, resp.Body, `{
		"id": "<<PRESENCE>>",
		"username": "gud"
	}`)
}

func loginUser(t *testing.T, testServer *httptest.Server) string {
	var body = []byte(`{
		"username": "gud",
		"password": "wordpass"
	}`)
	r, err := http.NewRequest("POST", testServer.URL+"/login", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("err creating new request: %s", err.Error())
	}
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatalf("err doing request: %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("got status code: %d, want %d", resp.StatusCode, http.StatusOK)
	}
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("failed to read login response body to bytes")
	}
	bs := string(b)
	return "Bearer " + bs[1:len(bs)-2]
}

func addTodo(t *testing.T, testServer *httptest.Server, bearer string) {
	var body = []byte(`{
		"name": "todo1"
	}`)
	r, err := http.NewRequest("POST", testServer.URL+"/todos", bytes.NewBuffer(body))
	if err != nil {
		t.Fatalf("err creating new request: %s", err.Error())
	}
	r.Header.Set("Authorization", bearer)
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatalf("err doing request: %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("got status code: %d, want %d", resp.StatusCode, http.StatusOK)
	}
	assertJSONEqual(t, resp.Body, `{
		"id": "<<PRESENCE>>",
		"name": "todo1",
		"completed": null,
		"userId": "<<PRESENCE>>",
		"priority": {
			"id": 2,
			"name": "Normal",
			"weight": 2
		}
	}`)
}

func getPriorities(t *testing.T, testServer *httptest.Server, bearer string) {
	r, err := http.NewRequest("GET", testServer.URL+"/priorities", nil)
	if err != nil {
		t.Fatalf("err creating new request: %s", err.Error())
	}
	r.Header.Set("Authorization", bearer)
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatalf("err doing request: %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("got status code: %d, want %d", resp.StatusCode, http.StatusOK)
	}
	assertJSONEqual(t, resp.Body, `[
		{
			"id": 1,
			"name": "Low",
			"weight": 1
		},
		{
			"id": 2,
			"name": "Normal",
			"weight": 2
		},
		{
			"id": 3,
			"name": "High",
			"weight": 3
		}
	]`)
}

func getTodos(t *testing.T, testServer *httptest.Server, bearer string) {
	r, err := http.NewRequest("GET", testServer.URL+"/todos", nil)
	if err != nil {
		t.Fatalf("err creating new request: %s", err.Error())
	}
	r.Header.Set("Authorization", bearer)
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatalf("err doing request: %s", err.Error())
	}
	if resp.StatusCode != http.StatusOK {
		t.Fatalf("got status code: %d, want %d", resp.StatusCode, http.StatusOK)
	}
	assertJSONEqual(t, resp.Body, `[
		{
			"id": "<<PRESENCE>>",
			"completed": null,
			"name": "todo1",
			"userId": "<<PRESENCE>>",
			"priority": {
				"id": 2,
				"name": "Normal",
				"weight": 2
			}
		}
	]`)
}

func assertJSONEqual(t *testing.T, respBody io.ReadCloser, expected string) {
	got, err := io.ReadAll(respBody)
	if err != nil {
		t.Fatalf("unable to read body into bytes slice")
	}
	ja := jsonassert.New(t)
	ja.Assertf(string(got), expected)
}
