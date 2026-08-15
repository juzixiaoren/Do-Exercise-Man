package main

func numIslands(grid [][]byte) int {
	ans := 0
	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			if grid[j][i] == '1' {
				ans++
				infect(grid, i, j)
			}
		}
	}
	return ans

}
func infect(grid [][]byte, i, j int) {
	if i < 0 || j < 0 || i > len(grid[0])-1 || j > len(grid)-1 || grid[j][i] != '1' {
		return
	}
	grid[j][i] = '2'
	infect(grid, i, j-1)
	infect(grid, i, j+1)
	infect(grid, i-1, j)
	infect(grid, i+1, j)
}

/*
给你一个由 '1'（陆地）和 '0'（水）组成的的二维网格，请你计算网格中岛屿的数量。

岛屿总是被水包围，并且每座岛屿只能由水平方向和/或竖直方向上相邻的陆地连接形成。

此外，你可以假设该网格的四条边均被水包围。



示例 1：

输入：grid = [
  ['1','1','1','1','0'],
  ['1','1','0','1','0'],
  ['1','1','0','0','0'],
  ['0','0','0','0','0']
]
输出：1
示例 2：

输入：grid = [
  ['1','1','0','0','0'],
  ['1','1','0','0','0'],
  ['0','0','1','0','0'],
  ['0','0','0','1','1']
]
输出：3

思路:扩散，遇到 1 的时候标记岛屿数+1
且使用递归将 1 变成 2，包括 1的上下左右都变成 2

*/
