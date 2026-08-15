package main

import "testing"

func Test_canConstruct(t *testing.T) {
	r := canConstruct("a", "b")
	t.Logf("result: %v", r)
}
