package main

import "testing"

func Test_evalRPN(t *testing.T) {
	r := evalRPN([]string{"2", "1", "+", "3", "*"})
	t.Logf("result: %d", r)
}
