package main

func setZeroes(matrix [][]int) {
	var flag int
	flag = -876
	m := len(matrix)
	n := len(matrix[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == 0 {
				for k := 0; k < n; k++ {
					if matrix[i][k] != 0 {
						matrix[i][k] = flag
					}
				}
				for k := 0; k < m; k++ {
					if matrix[k][j] != 0 {
						matrix[k][j] = flag
					}
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == flag {
				matrix[i][j] = 0
			}
		}
	}
}
//使用魔法数字当哨兵，但是时间复杂度高