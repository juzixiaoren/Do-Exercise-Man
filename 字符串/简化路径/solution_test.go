package main

import "testing"

func Test_simplifyPath(t *testing.T) {
	r := simplifyPath("/a/./b/../../c/")
	t.Logf("result: %s", r)
}
