package server

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/kinbiko/jsonassert"
)

func addUser(t *testing.T, ts *httptest.Server) {
	var body = []byte(`{
		"username": "gud",
		"password": "wordpass"
	}`)
	r, err := http.NewRequest("POST", ts.URL+"/users", bytes.NewBuffer(body))
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
	got, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read body into bytes slice")
	}
	jsonassert.New(t).Assertf(string(got), `{
		"id": "<<PRESENCE>>",
		"username": "gud"
	}`)
}

func loginUser(t *testing.T, ts *httptest.Server) string {
	var body = []byte(`{
		"username": "gud",
		"password": "wordpass"
	}`)
	r, err := http.NewRequest("POST", ts.URL+"/login", bytes.NewBuffer(body))
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

func addTodo(t *testing.T, ts *httptest.Server, bearer string) {
	respBody := makePost(t, ts, bearer, "/todos", `{
		"name": "todo1"
	}`)
	jsonassert.New(t).Assertf(respBody, `{
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

func getPriorities(t *testing.T, ts *httptest.Server, bearer string) {
	respBody := makeGet(t, ts, bearer, "/priorities")
	jsonassert.New(t).Assertf(respBody, `[
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

func getTodos(t *testing.T, ts *httptest.Server, bearer string) int {
	respBody := makeGet(t, ts, bearer, "/todos")
	jsonassert.New(t).Assertf(respBody, `[
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
	type todoList struct {
		ID int
	}
	td := []todoList{}
	err := json.Unmarshal([]byte(respBody), &td)
	if err != nil {
		t.Fatalf("err unmarshaling response: %s", err.Error())
	}
	return td[0].ID
}

func patchTodo(t *testing.T, ts *httptest.Server, bearer string, todoID int) {
	url := fmt.Sprintf("%s/todos/%d", ts.URL, todoID)
	body := []byte(`{
		"complete": true
	}`)
	r, err := http.NewRequest("PATCH", url, bytes.NewBuffer(body))
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
}

func deleteTodo(t *testing.T, ts *httptest.Server, bearer string, todoID int) {
	url := fmt.Sprintf("%s/todos/%d", ts.URL, todoID)
	r, err := http.NewRequest("DELETE", url, nil)
	if err != nil {
		t.Fatalf("err creating new request: %s", err.Error())
	}
	r.Header.Set("Authorization", bearer)
	resp, err := http.DefaultClient.Do(r)
	if err != nil {
		t.Fatalf("err doing request: %s", err.Error())
	}
	if resp.StatusCode != http.StatusNoContent {
		t.Fatalf(
			"got status code: %d, want %d",
			resp.StatusCode,
			http.StatusNoContent,
		)
	}
}
