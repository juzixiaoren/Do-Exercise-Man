package main

import "testing"

func Test_numIslands(t *testing.T) {
	grid := [][]byte{
		{'1', '1', '0', '0', '0'},
		{'1', '0', '0', '1', '0'},
		{'0', '0', '0', '0', '0'},
		{'0', '0', '0', '1', '1'},
	}
	r := numIslands(grid)
	t.Logf("result: %d", r)
}
