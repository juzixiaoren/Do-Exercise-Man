package main

import "testing"

func Test_removeNthFromEnd2(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}
	r := removeNthFromEnd2(head, 2)
	t.Logf("result: %v", r != nil)
}
