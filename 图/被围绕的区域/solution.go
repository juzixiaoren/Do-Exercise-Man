package main

func solve(board [][]byte) {
	for i := 0; i < len(board[0]); i++ {
		j := 0
		j_ := len(board) - 1
		if board[j][i] == 'O' {
			infect_unT(board, i, j)
		}
		if board[j_][i] == 'O' {
			infect_unT(board, i, j_)
		}
	}
	for j := 0; j < len(board); j++ {
		i := 0
		i_ := len(board[0]) - 1
		if board[j][i] == 'O' {
			infect_unT(board, i, j)
		}
		if board[j][i_] == 'O' {
			infect_unT(board, i_, j)
		}
	}
	for i := 0; i < len(board[0]); i++ {
		for j := 0; j < len(board); j++ {
			switch board[j][i] {
			case 'O':
				board[j][i] = 'X'
			case 'T':
				board[j][i] = 'O'
			}
		}
	}
}

func infect_unT(board [][]byte, i, j int) {
	if i < 0 || j < 0 || i > len(board[0])-1 || j > len(board)-1 || board[j][i] != 'O' {
		return
	}
	board[j][i] = 'T'
	infect_unT(board, i, j-1)
	infect_unT(board, i, j+1)
	infect_unT(board, i-1, j)
	infect_unT(board, i+1, j)
}
