package main

import "testing"

func Test_summaryRanges(t *testing.T) {
	r := summaryRanges([]int{0, 1, 2, 4, 5, 7})
	t.Logf("result: %v", r)
}
