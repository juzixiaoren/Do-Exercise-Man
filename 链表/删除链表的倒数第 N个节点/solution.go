package main

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	stack := []*ListNode{}
	cur := head
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Next
	}
	if len(stack) == n {
		return head.Next
	}
	if len(stack) < 3 {
		switch n {
		case 1:
			head.Next = nil
			return head
		case 2:
			return stack[1]
		default:
			return nil
		}
	}
	stack = stack[:len(stack)-n+1]
	stack[len(stack)-2].Next = stack[len(stack)-1].Next
	return head
}

/*
给你一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。
用栈全部压进去，然后冒出直到 top 为倒数第 n 个节点，然后删除即可
唯一要注意的是边界情况如只有两个节点，或者删除的是头节点
*/

func removeNthFromEnd2(head *ListNode, n int) *ListNode {
	preHead := new(ListNode)
	preHead.Next = head
	fast := preHead
	slow := preHead
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return preHead.Next
}

/*
然而方法二更为精妙啊
*/
