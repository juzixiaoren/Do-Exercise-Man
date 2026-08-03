package main

import (
	"strconv"
)

func calculate(s string) int {
	ans := 0
	num := 0
	sign := 1
	stack := []int{}
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			continue
		}
		t, err := strconv.Atoi(string(s[i]))
		if err == nil {
			num = num*10 + t
		}
		switch s[i] {
		case '+':
			ans += sign * num
			num = 0
			sign = 1
		case '-':
			ans += sign * num
			num = 0
			sign = -1
		case '(':
			stack = append(stack, ans)
			stack = append(stack, sign)
			sign = 1
			ans = 0
		case ')':
			ans += sign * num
			sign = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			ans = sign*ans + stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			num = 0
		}
	}
	return ans
}

/**
* @lc app=leetcode.cn id=224 lang=golang
*
* [224] 基本计算器

1. 遇到数字，直接累加到num中
2. 遇到+，将sign * num 加到ans中，然后sign = 1，num = 0
3. 遇到-，将sign * num 加到ans中，然后sign = -1，num = 0
4. 遇到(，将ans和sign压入栈中，然后sign = 1，num = 0
5. 遇到)，将sign * num 加到ans中，然后sign = 栈顶元素，ans = sign * ans + 栈顶元素，然后num = 0
*/
