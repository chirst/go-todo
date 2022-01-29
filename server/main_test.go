package server

import (
	"testing"
)

func TestBehavior(t *testing.T) {
	ts, teardown := setupTest(t)
	defer teardown()

	addUser(t, ts)
	bearer := loginUser(t, ts)
	addTodo(t, ts, bearer)
	getPriorities(t, ts, bearer)
	firstTodoId := getTodos(t, ts, bearer)
	patchTodo(t, ts, bearer, firstTodoId)
	deleteTodo(t, ts, bearer, firstTodoId)
}
