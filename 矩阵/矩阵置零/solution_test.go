package main

import "testing"

func Test_setZeroes1(t *testing.T) {
	matrix := [][]int{{1, 1, 1}, {1, 0, 1}, {1, 1, 1}}
	setZeroes1(matrix)
	t.Logf("setZeroes1 done: %v", matrix)
}
