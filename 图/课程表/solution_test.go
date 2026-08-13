package main

import "testing"

func Test_canFinish(t *testing.T) {
	r := canFinish(2, [][]int{{1, 0}})
	t.Logf("result: %v", r)
}
