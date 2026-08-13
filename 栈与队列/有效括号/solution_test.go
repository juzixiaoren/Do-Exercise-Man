package main

import "testing"

func Test_isValid(t *testing.T) {
	r := isValid("()[]{}")
	t.Logf("result: %v", r)
}
