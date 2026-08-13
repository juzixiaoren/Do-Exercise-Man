package main

type PathPair struct {
	node *TreeNode
	sum  int
}

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	sumMax := root.Val
	Postorder(root, &sumMax)
	return sumMax
}
func Postorder(root *TreeNode, sumMax *int) int {
	if root == nil {
		return 0
	}
	sum1 := Postorder(root.Left, sumMax)
	sum2 := Postorder(root.Right, sumMax)
	ans := max(sum1+root.Val, sum2+root.Val, root.Val)
	*sumMax = max(ans, *sumMax, sum1+sum2+root.Val)
	return ans
}
