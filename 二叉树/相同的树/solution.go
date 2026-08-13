package main

func isSameTree(p *TreeNode, q *TreeNode) bool {
	if p == nil && q == nil {
		return true
	} else if p == nil || q == nil {
		return false
	} else {
		if p.Val == q.Val {
			return true && isSameTree(p.Left, q.Left) && isSameTree(p.Right, q.Right)
		} else {
			return false
		}
	}
}

/*
递归，空间复杂度 O(logn)栈的长度实际上是
*/
func isSameTreeUseStack(p *TreeNode, q *TreeNode) bool {
	stack := [][]*TreeNode{}
	stack = append(stack, []*TreeNode{p, q})
	for len(stack) != 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node[0] == nil && node[1] == nil {
			continue
		} else if node[0] == nil || node[1] == nil {
			return false
		} else {
			if node[0].Val == node[1].Val {
				stack = append(stack, []*TreeNode{node[0].Left, node[1].Left})
				stack = append(stack, []*TreeNode{node[0].Right, node[1].Right})
			} else {
				return false
			}
		}
	}
	return true
}
