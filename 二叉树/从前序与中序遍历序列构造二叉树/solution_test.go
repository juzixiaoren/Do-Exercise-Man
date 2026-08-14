package main

import "testing"

func Test_buildTree1(t *testing.T) {
	preorder := []int{3, 9, 20, 15, 7}
	inorder := []int{9, 3, 15, 20, 7}
	r := buildTree(preorder, inorder)
	t.Logf("result: %v", r != nil)
}
