package main

func combine(n int, k int) [][]int {
	if n < 1 {
		return [][]int{}
	}
	ans := [][]int{}
	dfs(n, k, 0, &ans, []int{})
	return ans
}

func dfs(n, k, num int, ans *[][]int, path []int) {
	if len(path) == k {
		temp := make([]int, len(path))
		copy(temp, path)
		*ans = append(*ans, temp)
		return
	} else {
		for i := num + 1; i <= n; i++ {
			dfs(n, k, i, ans, append(path, i))
		}
	}
}

/**/

func combine2(n int, k int) [][]int {

	ans := [][]int{}

	dfs2(1, n, k, []int{}, &ans)

	return ans
}

func dfs2(start, n, k int, path []int, ans *[][]int) {

	if len(path) == k {

		temp := append([]int{}, path...)
		*ans = append(*ans, temp)

		return
	}

	for i := start; i <= n; i++ {

		path = append(path, i)

		dfs2(i+1, n, k, path, ans)

		path = path[:len(path)-1]
	}
}
