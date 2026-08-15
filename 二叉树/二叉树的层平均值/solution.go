package main

func averageOfLevels(root *TreeNode) []float64 {
	if root == nil {
		return []float64{}
	}
	queue := []*TreeNode{}
	ans := []float64{}
	queue = append(queue, root)
	average := 0.0
	for len(queue) != 0 {
		num := len(queue)
		average = 0.0
		for i := 0; i < num; i++ {
			average += float64(queue[i].Val)
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
		}
		average = average / float64(num)
		ans = append(ans, average)
		queue = queue[num:]
	}
	return ans
}
