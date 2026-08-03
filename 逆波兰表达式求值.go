package main

import "strconv"

func evalRPN(tokens []string) int {
	stack := []int{}
	for i := 0; i < len(tokens); i++ {
		num, err := strconv.Atoi(tokens[i])
		if err == nil {
			stack = append(stack, num)
		} else {
			nums1 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			nums2 := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			switch tokens[i] {
			case "+":
				{
					stack = append(stack, nums2+nums1)
				}
			case "-":
				{
					stack = append(stack, nums2-nums1)
				}
			case "*":
				{
					stack = append(stack, nums2*nums1)
				}
			case "/":
				{
					stack = append(stack, int(nums2/nums1))
				}
			}
		}
	}
	return stack[0]
}
