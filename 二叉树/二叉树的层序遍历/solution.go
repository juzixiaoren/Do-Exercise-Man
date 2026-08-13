package main

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
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
		ans = append(ans, levelAns)
		queue = queue[num:]
	}
	return ans
}
