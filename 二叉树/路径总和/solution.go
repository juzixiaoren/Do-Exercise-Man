package main

func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	Sum := 0
	return dfs(root, targetSum, Sum)

}

func dfs(root *TreeNode, targetSum int, Sum int) bool {
	if root == nil {
		return false
	}
	Sum += root.Val
	if root.Right == nil && root.Left == nil {
		if Sum == targetSum {
			return true
		}
		return false
	} else {
		return dfs(root.Left, targetSum, Sum) || dfs(root.Right, targetSum, Sum)
	}

}

/*
题目：
给你二叉树的根节点 root 和一个表示目标和的整数 targetSum 。判断该树中是否存在 根节点到叶子节点 的路径，这条路径上所有节点值相加等于目标和 targetSum 。如果存在，返回 true ；否则，返回 false 。

叶子节点 是指没有子节点的节点。

解法:使用 dfs(递归，因为显式栈无法保存中间变量，需要二维显式栈才能保存中间变量)，递归过程中 Sum 不断传递（传递的是中间值)

解法 2，一次递归
*/

func hasPathSu2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	targetSum -= root.Val
	if root.Left == nil && root.Right == nil {
		if targetSum == 0 {
			return true
		}
		return false
	}
	return hasPathSu2(root.Left, targetSum) || hasPathSu2(root.Right, targetSum)

}
