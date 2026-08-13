package main

import "math"

func getMinimumDifference(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{}
	stack = append(stack, root)
	minAns := 0
	if root.Left != nil {
		minAns = int(math.Abs(float64(root.Val - root.Left.Val)))
	} else {
		minAns = int(math.Abs(float64(root.Val - root.Right.Val)))
	}
	cur := root.Left
	temp := root.Val
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		cur = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if temp != cur.Val {
			minAns = min(minAns, int(math.Abs(float64(cur.Val)-float64(temp))))
		}
		temp = cur.Val
		cur = cur.Right
	}
	return minAns
}

/*上面的写复杂了*/
func getMinimumDifferenceFix(root *TreeNode) int {
	if root == nil {
		return 0
	}
	stack := []*TreeNode{}
	minAns := 0
	if root.Left != nil {
		minAns = root.Val - root.Left.Val
	} else {
		minAns = root.Right.Val - root.Val
	}
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
			minAns = min(minAns, cur.Val-temp.Val)
		}
		temp = cur
		cur = cur.Right
	}
	return minAns
}

/*
二叉搜索树的最小绝对差

给你一个二叉搜索树的根节点 root ，返回 树中任意两不同节点值之间的最小差值 。

差值是一个正数，其数值等于两值之差的绝对值。

做法：二叉搜索树的排序使用的是中序排序，而中序排序可以用显示栈模拟

注意中序排序的显式栈需要先把最左边压进去，并且终止条件为 cur!=nil||len(stack)!=0((!!!!))
*/
