package main

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func setupTest() (*httptest.Server, func()) {
	router := getRouter(nil)
	testServer := httptest.NewServer(router)
	teardown := func() {
		testServer.Close()
	}
	return testServer, teardown
}

func TestAddUser(t *testing.T) {
	testServer, teardown := setupTest()
	defer teardown()

	addUser(t, testServer)
	bearer := loginUser(t, testServer)
	addTodo(t, testServer, bearer)
}

func addUser(t *testing.T, testServer *httptest.Server) {
	var body = []byte(`{
		"username":"gud",
		"password":"wordpass"
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
	gotBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read body into bytes slice")
	}
	expectedBody := []byte(`{"id":1,"username":"gud"}`)
	if !bytes.Equal(gotBody, expectedBody) {
		t.Fatalf("got %s, want %s", gotBody, expectedBody)
	}
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
	buf := bytes.Buffer{}
	buf.ReadFrom(resp.Body)
	return "Bearer " + buf.String()[1:126]
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
	gotBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read body into bytes slice")
	}
	expectedBody := []byte(`{"id":1,"name":"todo1","completed":null,"userId":1,"priority":{"id":2,"name":"Normal","weight":2}}`)
	if !bytes.Equal(gotBody, expectedBody) {
		t.Fatalf("got %s, want %s", gotBody, expectedBody)
	}
}
