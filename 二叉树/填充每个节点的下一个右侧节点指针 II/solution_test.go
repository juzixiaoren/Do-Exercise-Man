package main

import "testing"

func buildTNode() *TNode {
	return &TNode{
		Val:  1,
		Left: &TNode{Val: 2, Left: &TNode{Val: 4}, Right: &TNode{Val: 5}},
		Right: &TNode{Val: 3, Right: &TNode{Val: 7}},
	}
}

func Test_connect(t *testing.T) {
	r := connect(buildTNode())
	t.Logf("connect done: %v", r != nil)
}
