package main

import "math"

func countNodes(root *TreeNode) int {
	leftHeight := 0
	rightHeight := 0
	cur := root
	if cur == nil {
		return 0
	}
	for cur != nil {
		cur = cur.Left
		leftHeight++
	}
	cur = root
	for cur != nil {
		cur = cur.Right
		rightHeight++
	}
	if leftHeight == rightHeight {
		return int(math.Pow(2, float64(leftHeight))) - 1
	} else {
		return countNodes(root.Left) + countNodes(root.Right) + 1
	}
}

/*
题目：完全二叉树的节点个数
解题思路：如果左子树的高度等于右子树的高度，则左子树是满二叉树，右子树是完全二叉树，否则左子树是完全二叉树，右子树是满二叉树。
所以要么自己是满二叉树，要么左子树是满二叉树，右子树是完全二叉树，否则左子树是完全二叉树，右子树是满二叉树。因此
复杂度是时间复杂度：O(log² n)
空间复杂度：O(log n)

*/
