package main

import "testing"

func Test_rotate(t *testing.T) {
	matrix := [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}
	rotate(matrix)
	t.Logf("rotate done: %v", matrix)
}
