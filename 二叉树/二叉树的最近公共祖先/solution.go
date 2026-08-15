package main

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}
	if root == p || root == q {
		return root
	}
	leftFind := lowestCommonAncestor(root.Left, p, q)
	rightFind := lowestCommonAncestor(root.Right, p, q)
	if leftFind == nil && rightFind == nil {
		return nil
	} else if leftFind == nil {
		return rightFind
	} else if rightFind == nil {
		return leftFind
	} else {
		return root
	}
}

/*
给定一个二叉树, 找到该树中两个指定节点的最近公共祖先。

对于有根树 T 的两个节点 p、q，最近公共祖先表示为一个节点 x，满足 x 是 p、q 的祖先且 x 的深度尽可能大（一个节点也可以是它自己的祖先）。

思路：这题最关键的不是“向上找父节点”，而是：
从下往上汇报：当前子树里有没有找到 p 或 q。

所以处理方法是：左子树找到了且右子树也找到了，返回自身
左子树和右子树都没找到，返回空
左子树找到了，右子树没找到，返回左子树(这里的左子树找到了即为情况 1 的返回自身，返回的就是公共节点)
递归实现


*/
