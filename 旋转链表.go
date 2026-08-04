package main

func rotateRight(head *ListNode, k int) *ListNode {
	cur := head
	if cur == nil {
		return head
	}
	n := 0
	for cur.Next != nil {
		n++
		cur = cur.Next
	}
	n++
	cur.Next = head
	move := n - k%n
	for i := 0; i < move; i++ {
		cur = cur.Next
	}
	head = cur.Next
	cur.Next = nil
	return head
}

/*
给你一个链表的头节点 head ，旋转链表，将链表每个节点向右移动 k 个位置。
思路
1->2->3->4->5 旋转 2，不难看出倒数第二个将会成为头节点，那么倒数第二个实际上就是正数的第 n-k%n+1 个节点，即 5-2+1=4
那么先指针直到尾部，将尾部和头相连，不难看出尾部继续移动 1 就到头部第一个节点，移动 n 就到头部的第 n 个节点
那么只需要再移动n-k%n+1步就到旋转后的头节点了。
这个时候只需要再做一步，在头节点之前断开，那么移动 n-k%n 即可到断开处
复杂度 O(n)只需要一次遍历再多 n-k%n步即可
*/
