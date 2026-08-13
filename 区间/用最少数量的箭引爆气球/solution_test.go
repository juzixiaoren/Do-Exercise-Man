package main

import "testing"

func Test_findMinArrowShots(t *testing.T) {
	r := findMinArrowShots([][]int{{10, 16}, {2, 8}, {1, 6}, {7, 12}})
	t.Logf("result: %d", r)
}
