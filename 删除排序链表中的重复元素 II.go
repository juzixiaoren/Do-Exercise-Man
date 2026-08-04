package main

func deleteDuplicates(head *ListNode) *ListNode {
	preHead := new(ListNode)
	preHead.Next = head
	cur := preHead
	for cur != nil {
		if cur.Next == nil || cur.Next.Next == nil {
			return preHead.Next
		} else if cur.Next.Val != cur.Next.Next.Val {
			cur = cur.Next
			continue
		}
		for cur.Next.Next != nil && cur.Next.Val == cur.Next.Next.Val {
			cur.Next.Next = cur.Next.Next.Next
		}
		cur.Next = cur.Next.Next
	}
	return preHead.Next
}
