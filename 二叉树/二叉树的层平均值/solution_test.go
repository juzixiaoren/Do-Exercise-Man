package main

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(vals []interface{}) *TreeNode {
	if len(vals) == 0 || vals[0] == nil {
		return nil
	}
	root := &TreeNode{Val: vals[0].(int)}
	queue := []*TreeNode{root}
	i := 1
	for len(queue) > 0 && i < len(vals) {
		cur := queue[0]
		queue = queue[1:]
		if i < len(vals) && vals[i] != nil {
			cur.Left = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, cur.Left)
		}
		i++
		if i < len(vals) && vals[i] != nil {
			cur.Right = &TreeNode{Val: vals[i].(int)}
			queue = append(queue, cur.Right)
		}
		i++
	}
	return root
}

func Test_averageOfLevels(t *testing.T) {
	r := averageOfLevels(buildTree([]interface{}{3, 9, 20, nil, nil, 15, 7}))
	t.Logf("result: %v", r)
}
