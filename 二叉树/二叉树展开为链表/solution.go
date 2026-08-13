package main

func flatten(root *TreeNode) {
	if root == nil {
		return
	}
	temp := root.Right
	root.Right = root.Left
	root.Left = nil
	cur := root
	for cur.Right != nil {
		cur = cur.Right
	}
	cur.Right = temp
	flatten(root.Right)

}

/*
给你二叉树的根结点 root ，请你将它展开为一个单链表：

展开后的单链表应该同样使用 TreeNode ，其中 right 子指针指向链表中下一个结点，而左子指针始终为 null 。
展开后的单链表应该与二叉树 先序遍历 顺序相同。

上面是递归写法，思路是将右子树接到左子树最右边，将左子树接到根节点右边，左边清空
然后对子树继续执行这样的操作
下面不使用递归
*/
func flatten2(root *TreeNode) {
	for root != nil {
		if root.Left == nil {
			root = root.Right
			continue
		}
		tail := root.Left
		for tail.Right != nil {
			tail = tail.Right
		}
		tail.Right = root.Right
		root.Right = root.Left
		root.Left = nil
		root = root.Right
		continue

	}
}

/*
	解法相似:
	将右子树接到左子树最右边，将左子树接到根节点右边，左边清空
	然后对子树继续执行这样的操作

*/
