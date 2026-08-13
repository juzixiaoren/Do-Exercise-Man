package main

import "testing"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	root := &TreeNode{Val: preorder[0]}
	idx := 0
	for inorder[idx] != root.Val {
		idx++
	}
	root.Left = buildTree(preorder[1:1+idx], inorder[:idx])
	root.Right = buildTree(preorder[1+idx:], inorder[idx+1:])
	return root
}

func Test_buildTree1(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	r := buildTree1(preorder, inorder)
	t.Logf("result: %v", r != nil)
}
