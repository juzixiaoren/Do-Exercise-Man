package main

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
