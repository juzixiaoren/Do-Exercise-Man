package main

type TNode struct {
	Val   int
	Left  *TNode
	Right *TNode
	Next  *TNode
}

func connect(root *TNode) *TNode {
	if root == nil {
		return root
	}
	queue := []*TNode{}
	queue = append(queue, root)
	for len(queue) != 0 {
		nodeNum := len(queue)
		for i := 0; i < nodeNum; i++ {
			if i == nodeNum-1 {
				queue[i].Next = nil
			} else {
				queue[i].Next = queue[i+1]
			}
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[nodeNum:]
	}
	return root
}
