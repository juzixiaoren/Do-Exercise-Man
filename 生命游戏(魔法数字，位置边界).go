package main

func gameOfLife1(board [][]int) {
	m := len(board)
	n := len(board[0])
	if m == 0 || n == 0 {
		return
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			check_alive(board, i, j, m, n)
		}
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			switch board[i][j] {
			case -1:
				board[i][j] = 1
			case 2:
				board[i][j] = 0
			}
		}
	}
}
func check_alive(board [][]int, i int, j int, m int, n int) {
	count := 0
	if i == 0 && j == 0 && m >= 2 && n >= 2 {
		if board[i+1][j] >= 1 {
			count++
		}
		if board[i][j+1] >= 1 {
			count++
		}
		if board[i+1][j+1] >= 1 {
			count++
		}
		board[i][j] = change_live(count, board[i][j])
	} else if i == 0 && j == n-1 && m >= 2 && n >= 2 {
		if board[i+1][j] >= 1 {
			count++
		}
		if board[i][j-1] >= 1 {
			count++
		}
		if board[i+1][j-1] >= 1 {
			count++
		}
		board[i][j] = change_live(count, board[i][j])
	} else if i == m-1 && j == 0 && m >= 2 && n >= 2 {
		if board[i-1][j] >= 1 {
			count++
		}
		if board[i][j+1] >= 1 {
			count++
		}
		if board[i-1][j+1] >= 1 {
			count++
		}
		board[i][j] = change_live(count, board[i][j])
	} else if i == m-1 && j == n-1 && m >= 2 && n >= 2 {
		if board[i-1][j] >= 1 {
			count++
		}
		if board[i][j-1] >= 1 {
			count++
		}
		if board[i-1][j-1] >= 1 {
			count++
		}
		board[i][j] = change_live(count, board[i][j])
	} else if i == 0 && m >= 2 && n >= 2 {
		if board[i+1][j] >= 1 {
			count++
		}
		if board[i][j+1] >= 1 {
			count++
		}
		if board[i][j-1] >= 1 {
			count++
		}
		if board[i+1][j-1] >= 1 {
			count++
		}
		if board[i+1][j+1] >= 1 {
			count++
		}
		board[i][j] = change_live(count, board[i][j])
	} else if i == m-1 && m >= 2 && n >= 2 {
		if board[i-1][j] >= 1 {
			count++
		}
		if board[i][j+1] >= 1 {
			count++
		}
		if board[i][j-1] >= 1 {
			count++
		}
		if board[i-1][j-1] >= 1 {
			count++
		}
		if board[i-1][j+1] >= 1 {
			count++
		}
		board[i][j] = change_live(count, board[i][j])
	} else if j == 0 && m >= 2 && n >= 2 {
		if board[i-1][j] >= 1 {
			count++
		}
		if board[i+1][j] >= 1 {
			count++
		}
		if board[i][j+1] >= 1 {
			count++
		}
		if board[i-1][j+1] >= 1 {
			count++
		}
		if board[i+1][j+1] >= 1 {
			count++
		}
		board[i][j] = change_live(count, board[i][j])
	} else if j == n-1 && m >= 2 && n >= 2 {
		if board[i-1][j] >= 1 {
			count++
		}
		if board[i+1][j] >= 1 {
			count++
		}
		if board[i][j-1] >= 1 {
			count++
		}
		if board[i-1][j-1] >= 1 {
			count++
		}
		if board[i+1][j-1] >= 1 {
			count++
		}
		board[i][j] = change_live(count, board[i][j])
	} else if m >= 2 && n >= 2 {
		if board[i-1][j] >= 1 {
			count++
		}
		if board[i+1][j] >= 1 {
			count++
		}
		if board[i][j+1] >= 1 {
			count++
		}
		if board[i][j-1] >= 1 {
			count++
		}
		if board[i-1][j-1] >= 1 {
			count++
		}
		if board[i-1][j+1] >= 1 {
			count++
		}
		if board[i+1][j-1] >= 1 {
			count++
		}
		if board[i+1][j+1] >= 1 {
			count++
		}
		board[i][j] = change_live(count, board[i][j])
	} else if m == 1 && n == 1 {
		board[i][j] = 0
	} else if m == 1 {
		if j == 0 {
			board[i][j] = 0
			return
		}
		if j == n-1 {
			board[i][j] = 0
			return
		}
		if board[i][j-1] >= 1 {
			count++
		}
		if board[i][j+1] >= 1 {
			count++
		}
		board[i][j] = change_live(count, board[i][j])
	} else if n == 1 {
		if i == 0 {
			board[i][j] = 0
			return
		}
		if i == m-1 {
			board[i][j] = 0
			return
		}
		if board[i-1][j] >= 1 {
			count++
		}
		if board[i+1][j] >= 1 {
			count++
		}
		board[i][j] = change_live(count, board[i][j])
	}
}

func change_live(count int, live int) int {
	if live == 0 && count == 3 {
		return -1
	}
	if count < 2 {
		if live == 1 {
			return 2
		}
		return 0
	}
	if count == 2 || count == 3 {
		if live == 0 {
			return 0
		}
		return 1
	}
	if count > 3 {
		if live == 1 {
			return 2
		}
		return 0
	}
	return 0
}
