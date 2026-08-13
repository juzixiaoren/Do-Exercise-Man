package main

import "testing"

func Test_copyRandomList2(t *testing.T) {
	head := &Node{Val: 1}
	head.Next = &Node{Val: 2}
	head.Random = head.Next
	head.Next.Random = head
	r := copyRandomList2(head)
	t.Logf("result: %v", r != nil)
}
