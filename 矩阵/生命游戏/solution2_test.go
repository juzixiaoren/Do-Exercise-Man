package main

import "testing"

func Test_gameOfLife1(t *testing.T) {
	board := [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}
	gameOfLife1(board)
	t.Logf("gameOfLife1 done: %v", board)
}
