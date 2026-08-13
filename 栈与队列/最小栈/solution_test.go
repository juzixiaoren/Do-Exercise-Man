package main

import "testing"

func Test_Constructor1(t *testing.T) {
	s := Constructor1()
	s.Push(1)
	s.Push(2)
	t.Logf("top: %d, min: %d", s.Top(), s.GetMin())
}
