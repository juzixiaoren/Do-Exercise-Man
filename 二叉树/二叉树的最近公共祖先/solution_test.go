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

func Test_lowestCommonAncestor(t *testing.T) {
	root := buildTree([]interface{}{3, 5, 1, 6, 2, 0, 8, nil, nil, 7, 4})
	p := root.Left
	q := root.Right
	r := lowestCommonAncestor(root, p, q)
	t.Logf("result: %v", r != nil)
}
