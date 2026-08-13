package main

import "testing"

func Test_longestConsecutive(t *testing.T) {
	r := longestConsecutive([]int{100, 4, 200, 1, 3, 2})
	t.Logf("result: %d", r)
}
