package main

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test_partition(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}}
	r := partition(head, 3)
	t.Logf("result: %v", r != nil)
}
