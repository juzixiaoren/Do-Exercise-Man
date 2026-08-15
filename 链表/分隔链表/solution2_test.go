package main

import "testing"

func Test_partition2(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 5, Next: &ListNode{Val: 2}}}}}}
	r := partition2(head, 3)
	t.Logf("result: %v", r != nil)
}
