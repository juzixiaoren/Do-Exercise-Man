package main

import "testing"

func Test_setZeroes(t *testing.T) {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes(matrix)
	t.Logf("setZeroes done: %v", matrix)
}
