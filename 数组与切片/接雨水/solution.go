package main

func trap(height []int) int {
	leftMax := 0
	rightMax := 0
	left := 0
	right := len(height) - 1
	ans := 0
	for left < right {
		leftMax = max(height[left], leftMax)
		rightMax = max(height[right], rightMax)
		if height[left] < height[right] {
			ans += min(leftMax, rightMax) - height[left]
			left++
		} else {
			ans += min(leftMax, rightMax) - height[right]
			right--
		}
	}
	return ans
}
