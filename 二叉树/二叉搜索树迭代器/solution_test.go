package main

import "testing"

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

func Test_Constructor(t *testing.T) {
	it := Constructor(buildTree([]interface{}{7, 3, 15, nil, nil, 9, 20}))
	t.Logf("next: %d, hasNext: %v", it.Next(), it.HasNext())
}
