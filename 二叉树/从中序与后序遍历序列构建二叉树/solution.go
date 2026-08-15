package main

func buildTree1(inorder []int, postorder []int) *TreeNode {
	if len(inorder) == 0 {
		return nil
	}
	root := new(TreeNode)
	root.Val = postorder[len(postorder)-1]
	rootIndex := 0
	for {
		if inorder[rootIndex] == root.Val {
			break
		}
		rootIndex++
	}
	leftInorder := inorder[:rootIndex]
	rightInorder := inorder[rootIndex+1:]
	leftPostorder := postorder[:rootIndex]
	rightPostorder := postorder[rootIndex : len(postorder)-1]
	root.Left = buildTree(leftInorder, leftPostorder)
	root.Right = buildTree(rightInorder, rightPostorder)
	return root
}
