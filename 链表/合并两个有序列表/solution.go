package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	head := ListNode{}
	tail := &head
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			tail.Val = list2.Val
			list2 = list2.Next
		} else {
			tail.Val = list1.Val
			list1 = list1.Next
		}
		tail.Next = &ListNode{}
		tail = tail.Next
	}
	if list1 == nil {
		tail.Val = list2.Val
		tail.Next = list2.Next
	}
	if list2 == nil {
		tail.Val = list1.Val
		tail.Next = list1.Next
	}
	return &head
}

/*
这里需要优化，避免每次都创建新的节点，可以使用一个哨兵节点来简化代码。
*/

func mergeTwoLists_fix(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	head := new(ListNode)
	tail := head
	for list1 != nil && list2 != nil {
		if list1.Val >= list2.Val {
			tail.Next = list2
			list2 = list2.Next
		} else {
			tail.Next = list1
			list1 = list1.Next
		}
		tail = tail.Next
	}
	if list1 == nil {
		tail.Next = list2
	} else {
		tail.Next = list1
	}
	return head.Next
}
