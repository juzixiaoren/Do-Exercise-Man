package main

import "testing"

func Test_merge(t *testing.T) {
	r := merge([][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}})
	t.Logf("result: %v", r)
}
