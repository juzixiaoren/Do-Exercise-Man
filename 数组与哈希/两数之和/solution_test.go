package main

import "testing"

func Test_twoSum1(t *testing.T) {
	r := twoSum1([]int{2, 7, 11, 15}, 9)
	t.Logf("result: %v", r)
}
