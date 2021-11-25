package main

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	main()
	exitVal := m.Run()
	os.Exit(exitVal)
}

func TestAddTodo(t *testing.T) {

}
