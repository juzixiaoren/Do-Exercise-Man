package main

func generateParenthesis(n int) []string {
	ans := []string{}
	dfs(n, 0, 0, "", &ans)
	return ans
}
func dfs(n, leftNum, rightNum int, path string, ans *[]string) {
	if leftNum == n && rightNum == n {
		*ans = append(*ans, path)
	} else {
		if leftNum < n {
			dfs(n, leftNum+1, rightNum, path+"(", ans)
		}
		if rightNum < leftNum {
			dfs(n, leftNum, rightNum+1, path+")", ans)
		}
	}
}
