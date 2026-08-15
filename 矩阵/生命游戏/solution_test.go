package main

import "testing"

func Test_gameOfLife(t *testing.T) {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife(board)
	t.Logf("gameOfLife done: %v", board)
}
