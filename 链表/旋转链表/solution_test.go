package main

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test_rotateRight(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	r := rotateRight(head, 2)
	t.Logf("result: %v", r != nil)
}
