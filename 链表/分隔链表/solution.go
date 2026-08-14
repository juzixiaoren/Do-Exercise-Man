package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func partition(head *ListNode, x int) *ListNode {
	preHead := new(ListNode)
	preHead.Next = head
	cur := preHead
	var hasLarge bool
	var largePre *ListNode
	for cur != nil && cur.Next != nil {
		if cur.Next.Val >= x && hasLarge {
		} else if cur.Next.Val >= x {
			largePre = cur
			hasLarge = true
		} else if hasLarge {
			temp := largePre.Next
			largePre.Next = cur.Next
			cur.Next = cur.Next.Next
			largePre.Next.Next = temp
			largePre = largePre.Next
			continue
		} else {
		}
		cur = cur.Next
	}
	return preHead.Next
}

/*
给定一个链表的头节点 head 和一个特定值 x，请对链表进行分隔，使得所有小于 x 的节点都出现在大于或等于 x 的节点之前。
你应当保留两个分区中每个节点的初始相对位置。

思路：维护一个大于 X 的区域，遇到小于 X 的插入到大于 X 的区域前即可

更好的思路是，维护两个链表，一个插大于 x 的，一个插小于 x 的，然后连接在一起
*/

func partition2(head *ListNode, x int) *ListNode {
	largeList := new(ListNode)
	smallList := new(ListNode)
	var curLarge *ListNode
	var curSmall *ListNode
	curLarge = largeList
	curSmall = smallList
	for head != nil {
		if head.Val >= x {
			curLarge.Next = head
			curLarge = curLarge.Next
		} else {
			curSmall.Next = head
			curSmall = curSmall.Next
		}
		head = head.Next
	}
	curLarge.Next = nil
	curSmall.Next = largeList.Next
	return smallList.Next
}
