package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = preorder[0]
	if len(preorder) == 1 && len(inorder) == 1 {
		return root
	}
	i := 0
	for {
		if root.Val != inorder[i] {
			i++
			continue
		}
		break
	}
	leftInorder := inorder[0:i]
	RightInorder := inorder[i+1:]
	leftPreorder := preorder[1 : 1+i]
	RightPreorder := preorder[1+i:]
	root.Left = buildTree(leftPreorder, leftInorder)
	root.Right = buildTree(RightPreorder, RightInorder)
	return root
}

func buildTreeFix(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}

	root := &TreeNode{
		Val: preorder[0],
	}

	rootIndex := 0
	for inorder[rootIndex] != root.Val {
		rootIndex++
	}

	root.Left = buildTree(
		preorder[1:rootIndex+1],
		inorder[:rootIndex],
	)

	root.Right = buildTree(
		preorder[rootIndex+1:],
		inorder[rootIndex+1:],
	)

	return root
}

/*
根据前序遍历和中序遍历建树
前序遍历：
根 -> 左子树 -> 右子树
所以当前前序区间的第一个元素，一定是根节点。
中序遍历：
左子树 -> 根 -> 右子树
所以左子树的中序遍历就是中序遍历中根左边的部分。同理其他也是

所以关键点在于找出中序遍历中根的位置
比如根位置是 3，说明左子树长为 3

所以前序遍历为根—>333->右子树
中序遍历为 333->根->左子树
即实际上子树的遍历数组的长度是一样的

*/
