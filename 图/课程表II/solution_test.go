package main

import "testing"

func Test_findOrder(t *testing.T) {
	r := findOrder(2, [][]int{{1, 0}})
	t.Logf("result: %v", r)
}
