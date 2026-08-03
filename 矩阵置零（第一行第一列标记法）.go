package main

func setZeroes1(matrix [][]int) {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return
	}

	m := len(matrix)
	n := len(matrix[0])

	firstColZero := false
	firstRowZero := false

	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			firstColZero = true
			break
		}
	}

	for j := 0; j < n; j++ {
		if matrix[0][j] == 0 {
			firstRowZero = true
			break
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0] = 0
				matrix[0][j] = 0
			}
		}
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
	}

	if firstColZero {
		for i := 0; i < m; i++ {
			matrix[i][0] = 0
		}
	}

	if firstRowZero {
		for j := 0; j < n; j++ {
			matrix[0][j] = 0
		}
	}
}

/*
原题：矩阵置零
概述：给定一个 m x n 的矩阵，如果一个元素为 0，则将其所在行和列的所有元素都设为 0。请使用原地算法，即不创建额外空间。

解题思路：
1. 首先扫描第一行和第一列，记录是否有 0。
2. 然后扫描剩余部分，如果找到 0，则将对应的第一行和第一列的元素设为 0。
3. 再次扫描剩余部分，根据第一行和第一列的 0 来将对应行和列设为 0。
4. 最后根据第一步记录的是否有 0，来将第一行和第一列设为 0。
*/
