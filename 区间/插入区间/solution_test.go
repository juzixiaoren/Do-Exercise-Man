package main

import "testing"

func Test_insert(t *testing.T) {
	r := insert([][]int{{1, 3}, {6, 9}}, []int{2, 5})
	t.Logf("result: %v", r)
}
