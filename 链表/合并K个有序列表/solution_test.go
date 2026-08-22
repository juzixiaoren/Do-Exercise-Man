package main

import "testing"

func TestSolution(t *testing.T) {
	mergeKLists([]*ListNode{&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}}})
}
