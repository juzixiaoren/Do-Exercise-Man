package main

import "testing"

func Test_longestConsecutive2(t *testing.T) {
	r := longestConsecutive2([]int{100, 4, 200, 1, 3, 2})
	t.Logf("result: %d", r)
}
