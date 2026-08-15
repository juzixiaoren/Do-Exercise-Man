package main

import "slices"

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	flag := true
	queue := []*TreeNode{}
	ans := [][]int{}
	queue = append(queue, root)
	for len(queue) != 0 {
		levelAns := []int{}
		num := len(queue)
		for i := 0; i < num; i++ {
			levelAns = append(levelAns, queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		if !flag {
			slices.Reverse(levelAns)
		}
		flag = !flag
		ans = append(ans, levelAns)
		queue = queue[num:]
	}
	return ans
}
