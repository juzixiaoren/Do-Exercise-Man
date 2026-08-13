package main

func gameOfLife(board [][]int) {
	directions := [][2]int{{-1, -1}, {-1, 0}, {-1, 1}, {0, -1}, {0, 1}, {1, -1}, {1, 0}, {1, 1}}
	m := len(board)
	n := len(board[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			count := 0
			for k := 0; k < 8; k++ {
				x, y := i+directions[k][0], j+directions[k][1]
				if x >= 0 && x < m && y >= 0 && y < n {
					if board[x][y]&1 == 1 {
						count++
					}
				}
			}
			if board[i][j] == 1 {
				if count < 2 || count > 3 {
					continue
				}
				board[i][j] = 3
			}
			if board[i][j] == 0 {
				if count == 3 {
					board[i][j] = 2
				}
			}
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			board[i][j] = board[i][j] >> 1
		}
	}
}
