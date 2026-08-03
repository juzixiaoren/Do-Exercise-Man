package main

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	cur := head
	for i := 1; i < left-1; i++ {
		cur = cur.Next
	}
	if left == 1 {
		cur = new(ListNode)
		cur.Next = head
	}
	leftNode := cur.Next
	for i := left; i < right; i++ {
		temp := cur.Next
		temp2 := leftNode.Next.Next
		cur.Next = leftNode.Next
		cur.Next.Next = temp
		leftNode.Next = temp2
	}
	if left == 1 {
		head = cur.Next
	}
	return head
}

/*
给你单链表的头指针 head 和两个整数 left 和 right ，其中 left <= right 。
请你反转从位置 left 到位置 right 的链表节点，返回 反转后的链表 。
解法:
1. 找到left位置的节点
2. 找到right位置的节点
3. 反转left到right之间的节点
4. 将left位置的节点的Next指向right位置的节点
5. 将right位置的节点的Next指向left位置的节点的Next
6. 返回head
*/
