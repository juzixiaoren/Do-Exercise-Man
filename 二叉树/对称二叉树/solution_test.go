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

func Test_isSymmetric(t *testing.T) {
	r := isSymmetric(buildTree([]interface{}{1, 2, 2, nil, 3, nil, 3}))
	t.Logf("result: %v", r)
}
