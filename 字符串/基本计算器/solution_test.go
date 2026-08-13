package main

import "testing"

func Test_calculate(t *testing.T) {
	r := calculate("1 + 2")
	t.Logf("result: %d", r)
}
