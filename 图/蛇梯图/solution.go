package main

type Edge struct {
	value int
	step  int
}

func snakesAndLadders(board [][]int) int {
	expandboard := []int{0}

	for i := len(board) - 1; i >= 0; i-- {

		row := board[i]

		if (len(board)-1-i)%2 == 1 {
			// 奇数行反转
			for j := len(row) - 1; j >= 0; j-- {
				expandboard = append(expandboard, row[j])
			}
		} else {
			for j := 0; j < len(row); j++ {
				expandboard = append(expandboard, row[j])
			}
		}
	}
	visit := make([]bool, len(board[0])*len(board)+1)
	queue := []Edge{
		{
			value: 1,
			step:  0,
		},
	}
	visit[1] = true
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		for i := 1; i < 7; i++ {
			if expandboard[cur.value+i] != -1 {
				next := expandboard[cur.value+i]
				if visit[next] {
					continue
				} else {
					if next == len(expandboard)-1 {
						return cur.step + 1
					}
					visit[next] = true
					queue = append(queue, Edge{value: next, step: cur.step + 1})
				}
			} else if cur.value+i == len(expandboard)-1 {
				return cur.step + 1
			} else {
				if visit[cur.value+i] {
					continue
				} else {
					visit[cur.value+i] = true
					queue = append(queue, Edge{value: cur.value + i, step: cur.step + 1})
				}
			}
		}
	}
	return -1

}

/*
意思是从左下角开始，到右上角，按顺序走，一次能走 1 步到 6 步，即从 0 开始可以走 1 2 3 4 5 6 任意一个节点，其中会有黑洞，遇到黑洞就会传送走，求最少步数
*/
func snakesAndLadders_fix(board [][]int) int {

	n := len(board)

	// 展开棋盘
	arr := []int{0}

	for i := n - 1; i >= 0; i-- {

		if (n-1-i)%2 == 1 {
			for j := n - 1; j >= 0; j-- {
				arr = append(arr, board[i][j])
			}
		} else {
			for j := 0; j < n; j++ {
				arr = append(arr, board[i][j])
			}
		}
	}

	visit := make([]bool, n*n+1)

	queue := []Edge{
		{
			value: 1,
			step:  0,
		},
	}

	visit[1] = true

	for len(queue) > 0 {

		cur := queue[0]
		queue = queue[1:]

		for i := 1; i <= 6; i++ {

			next := cur.value + i

			if next > n*n {
				continue
			}

			if arr[next] != -1 {
				next = arr[next]
			}

			if next == n*n {
				return cur.step + 1
			}

			if !visit[next] {
				visit[next] = true
				queue = append(queue, Edge{
					value: next,
					step:  cur.step + 1,
				})
			}
		}
	}

	return -1
}
