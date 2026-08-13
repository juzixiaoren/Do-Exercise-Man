package main

func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	stack := []*TreeNode{}
	cur := root
	var temp *TreeNode
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if temp != nil {
			if cur.Val <= temp.Val {
				return false
			}
		}
		temp = cur
		cur = cur.Right
	}
	return true
}
