package server

import (
	"bytes"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func makePost(
	t *testing.T,
	ts *httptest.Server,
	bearer string,
	endpoint string,
	body string,
) string {
	r, err := http.NewRequest(
		"POST",
		ts.URL+endpoint,
		bytes.NewBuffer([]byte(body)),
	)
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
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read body into bytes slice: %s", err.Error())
	}
	return string(b)
}

func makeGet(
	t *testing.T,
	ts *httptest.Server,
	bearer string,
	endpoint string,
) string {
	r, err := http.NewRequest("GET", ts.URL+endpoint, nil)
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
	b, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("unable to read body into bytes slice")
	}
	return string(b)
}
