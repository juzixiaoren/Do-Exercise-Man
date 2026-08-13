package main

import "testing"

func Test_simplifyPath2(t *testing.T) {
	r := simplifyPath2("/a/./b/../../c/")
	t.Logf("result: %s", r)
}
