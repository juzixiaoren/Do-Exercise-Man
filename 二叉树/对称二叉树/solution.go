package main

func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	} else {
		reverse(root.Left)
		return isSame(root.Left, root.Right)
	}

}
func isSame(q *TreeNode, p *TreeNode) bool {
	if q == nil && p == nil {
		return true
	} else if q == nil || p == nil {
		return false
	} else {
		return p.Val == q.Val && isSame(q.Left, p.Left) && isSame(q.Right, p.Right)
	}
}
func reverse(root *TreeNode) {
	if root == nil {
		return
	} else {
		reverse(root.Left)
		reverse(root.Right)
		root.Left, root.Right = root.Right, root.Left
	}
}

/*
题目：给你一个二叉树的根节点 root ， 检查它是否轴对称。
将左子树 or 右子树先翻转
判断两个树是否相等
参考翻转二叉树.go和相同的树.go

*/
