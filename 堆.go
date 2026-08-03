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
func main() {
	testheap()

}
