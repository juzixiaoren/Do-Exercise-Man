package main

import "testing"

type ListNode struct {
	Val  int
	Next *ListNode
}

func Test_addTwoNumbers(t *testing.T) {
	l1 := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	l2 := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	r := addTwoNumbers(l1, l2)
	t.Logf("result: %v", r != nil)
}
