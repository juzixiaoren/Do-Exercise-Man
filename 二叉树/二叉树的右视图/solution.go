package main

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue := []*TreeNode{}
	ans := []int{}
	queue = append(queue, root)
	for len(queue) != 0 {
		index := len(queue)
		ans = append(ans, queue[index-1].Val)
		for i := 0; i < index; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		queue = queue[index:]
	}
	return ans
}
