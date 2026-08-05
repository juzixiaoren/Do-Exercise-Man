package main

func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 0
	dfs_(root, 0, &ans)
	return ans
}

func dfs_(root *TreeNode, sum int, num *int) {
	if root == nil {
		return
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*num += sum
	} else {
		dfs_(root.Left, sum, num)
		dfs_(root.Right, sum, num)
	}

}
