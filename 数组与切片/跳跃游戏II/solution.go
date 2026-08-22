package main

func jump(nums []int) int {
	canMove := 0
	pos := 0
	step := 0
	newPos := pos
	for canMove < len(nums)-1 {
		maxMove := 0
		for i := pos; i <= canMove; i++ {
			if i+nums[i] > maxMove {
				maxMove = i + nums[i]
				newPos = i
			} else {
				continue
			}
		}
		canMove = maxMove
		pos = newPos
		step++
	}
	return step
}
