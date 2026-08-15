package main

import "testing"

func Test_cloneGraph(t *testing.T) {
	n1 := &Node{Val: 1}
	n2 := &Node{Val: 2}
	n1.Neighbors = []*Node{n2}
	n2.Neighbors = []*Node{n1}
	r := cloneGraph(n1)
	t.Logf("result: %v", r != nil)
}
