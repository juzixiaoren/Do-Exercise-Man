package main

import "testing"

func Test_twoSum(t *testing.T) {
	r := twoSum([]int{2, 7, 11, 15}, 9)
	t.Logf("result: %v", r)
}
