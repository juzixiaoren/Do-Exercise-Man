package main

func permute(nums []int) [][]int {
	ans := [][]int{}
	visit := make([]bool, len(nums))
	dfs(visit, 0, nums, &ans, []int{})
	return ans
}
func dfs(visit []bool, level int, nums []int, ans *[][]int, path []int) {
	if level == len(nums) {
		*ans = append(*ans, path)
		return
	} else {
		for i := 0; i < len(visit); i++ {
			if visit[i] == false {
				visit[i] = true
				dfs(visit, level+1, nums, ans, append(path, nums[i]))
				visit[i] = false
			}
		}
	}
}
