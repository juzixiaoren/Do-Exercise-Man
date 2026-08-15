package main

import "testing"

func Test_mergeTwoLists(t *testing.T) {
	l1 := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 4}}}
	l2 := &ListNode{Val: 1, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4}}}
	r := mergeTwoLists(l1, l2)
	t.Logf("result: %v", r != nil)
}
