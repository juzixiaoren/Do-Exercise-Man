package main

import "testing"

func Test_containsNearbyDuplicate(t *testing.T) {
	r := containsNearbyDuplicate([]int{1, 2, 3, 1}, 3)
	t.Logf("result: %v", r)
}
