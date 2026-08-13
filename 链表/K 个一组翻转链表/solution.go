package main

func reverseKGroup(head *ListNode, k int) *ListNode {
	prehead := new(ListNode)
	prehead.Next = head
	cur := head
	head = prehead
	if k == 0 || k == 1 {
		return prehead.Next
	}
	nextHead := head
	for cur != nil {
		for i := 0; i < k; i++ {
			if nextHead.Next == nil {
				return prehead.Next
			} else {
				nextHead = nextHead.Next
			}
		}
		for i := 1; i < k; i++ {
			temp := head.Next
			head.Next = cur.Next
			cur.Next = cur.Next.Next
			head.Next.Next = temp
		}
		head = cur
		nextHead = cur
		cur = cur.Next
	}
	return prehead.Next
}

/*
	给你链表的头节点 head ，每 k 个节点一组进行翻转，请你返回修改后的链表。

k 是一个正整数，它的值小于或等于链表的长度。如果节点总数不是 k 的整数倍，那么请将最后剩余的节点保持原有顺序。

你不能只是单纯的改变节点内部的值，而是需要实际进行节点交换。

解法: 1. 先判断链表长度是否小于k，如果小于k，则直接返回链表
2. 如果链表长度大于等于k，则将链表分成k个一组，每组进行翻转
3. 翻转每组链表后，将翻转后的链表连接起来
4. 返回翻转后的链表
翻转链表的方法:
一个代翻转的子链表，先存头（如果没有头则创建一个）并以该链表的第一个节点作为 cur
每一次头连接 cur 的下一个节点，cur 连接下两个节点,cur 的下一个节点连接头原本连接的节点
这样就完成了 头123->头 213 的翻转
继续 头 213 ->头 321 翻转
如果链表第一个节点作为 cur，那么就凭空创建一个头，最终返回头.next 就是了
*/
