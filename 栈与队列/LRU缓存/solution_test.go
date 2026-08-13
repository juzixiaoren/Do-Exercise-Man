package main

import "testing"

func Test_Constructor(t *testing.T) {
	c := Constructor(2)
	c.Put(1, 1)
	c.Put(2, 2)
	t.Logf("get 1: %d", c.Get(1))
	c.Put(3, 3)
	t.Logf("get 2: %d", c.Get(2))
}
