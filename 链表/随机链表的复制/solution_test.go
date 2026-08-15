package main

import "testing"

func Test_copyRandomList(t *testing.T) {
	head := &Node{Val: 1}
	head.Next = &Node{Val: 2}
	head.Random = head.Next
	head.Next.Random = head
	r := copyRandomList(head)
	t.Logf("result: %v", r != nil)
}
