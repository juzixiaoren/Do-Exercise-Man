package main

import "testing"

func Test_isHappy(t *testing.T) {
	r := isHappy(19)
	t.Logf("result: %v", r)
}
