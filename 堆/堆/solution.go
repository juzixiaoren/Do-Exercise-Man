package main

import (
	"container/heap"
	"fmt"
)

type MaxHeap []int

func (h MaxHeap) Len() int {
	return len(h)
}
func (h MaxHeap) Less(i, j int) bool {
	return (h[i] > h[j])
}

func (h *MaxHeap) Pop() any {
	n := len(*h)
	value := (*h)[n-1]
	*h = (*h)[:n-1]
	return value
}
func (h MaxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}
func (h *MaxHeap) Push(value any) {
	*h = append(*h, value.(int))
}

func testheap() {
	h := &MaxHeap{1, 2, 3, 4}
	heap.Init(h)
	heap.Push(h, 5)
	fmt.Println((*h)[0])
}
func ma2in() {
	testheap()

}

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] } // 小根堆
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x any) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

/*
go 的堆需要自己实现方法
Len
Less
Swap
Push
Pop

*/
