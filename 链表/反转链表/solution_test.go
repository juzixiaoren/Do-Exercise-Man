package main

import "testing"

func Test_reverseBetween(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3}}}
	r := reverseBetween(head, 1, 2)
	t.Logf("result: %v", r != nil)
}
