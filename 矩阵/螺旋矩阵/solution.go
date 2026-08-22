package main

func spiralOrder(matrix [][]int) []int {
	acts := [][]int{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	left := 0
	act := 0
	right := len(matrix[0]) - 1
	up := len(matrix) - 1
	down := 0
	pos := []int{0, 0}
	ans := []int{}
	count := len(matrix) * len(matrix[0])
	for count != 0 {
		ans = append(ans, matrix[pos[1]][pos[0]])
		pos[0] = pos[0] + acts[act][0]
		pos[1] = pos[1] + acts[act][1]
		if pos[0] > right {
			down++
			act = (act + 1) % 4
			pos[0]--
			pos[1] = pos[1] + acts[act][1]
		} else if pos[0] < left {
			up--
			act = (act + 1) % 4
			pos[0]++
			pos[1] = pos[1] + acts[act][1]
		} else if pos[1] > up {
			right--
			act = (act + 1) % 4
			pos[1]--
			pos[0] = pos[0] + acts[act][0]
		} else if pos[1] < down {
			left++
			pos[1]++
			act = (act + 1) % 4
			pos[0] = pos[0] + acts[act][0]
		}
		count--
	}
	return ans
}
