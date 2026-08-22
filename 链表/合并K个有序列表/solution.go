package main

import "container/heap"

type ListNode struct {
	Val  int
	Next *ListNode
}
type Minheap []*ListNode

func (h Minheap) Less(i, j int) bool {
	return h[i].Val < h[j].Val
}
func (h Minheap) Len() int {
	return len(h)
}
func (h Minheap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *Minheap) Pop() any {
	value := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return value
}
func (h *Minheap) Push(node any) {
	*h = append(*h, node.(*ListNode))
}
func mergeKLists(lists []*ListNode) *ListNode {
	h := Minheap{}
	heap.Init(&h)
	for i := 0; i < len(lists); i++ {
		if lists[i] == nil {
			continue
		}
		heap.Push(&h, lists[i])

	}
	preHead := &ListNode{Next: nil}
	head := preHead
	for h.Len() != 0 {
		head.Next = heap.Pop(&h).(*ListNode)
		head = head.Next
		if head != nil && head.Next != nil {
			heap.Push(&h, head.Next)
		}
	}
	return preHead.Next
}
