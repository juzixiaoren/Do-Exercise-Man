package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		current.Next = &ListNode{
			Val: sum % 10,
		}

		carry = sum / 10
		current = current.Next
	}

	return dummy.Next
}

/*
题目：两数相加
给你两个 非空 的链表，表示两个非负整数。它们每位数字都是按照 逆序 的方式存储的，并且每个节点只能存储 一位 数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0 开头。
思路：使用两个指针分别指向两个链表的头节点，然后依次遍历两个链表，将两个链表的节点值相加
如果相加结果大于等于10，则将结果减去10，并将进位标记为1，否则将进位标记为0。
然后将相加结果的个位数作为新链表的节点值，将进位标记作为新链表的节点值。最后返回新链表的头节点。

*/
