package main

import (
	"fmt"
)

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	nts := []string{
		"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz",
	}
	type Edge struct {
		path  string
		level int
	}
	ans := []string{}
	stack := []Edge{}
	for _, word := range nts[digits[0]-'0'-2] {
		stack = append(stack, Edge{path: string(word),
			level: 0,
		})
	}
	for len(stack) != 0 {
		edge := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		numIdx := edge.level + 1
		if numIdx > len(digits)-1 {
			ans = append(ans, edge.path)
			continue
		}
		num := digits[numIdx] - '0'
		for _, word := range nts[num-2] {
			stack = append(stack, Edge{path: edge.path + string(word),
				level: numIdx,
			})
		}
	}
	return ans
}

func main() {
	ans := letterCombinations("23")
	fmt.Println(ans)
}

func letterCombinations_fix(digits string) []string {
	if digits == "" {
		return nil
	}

	mapping := []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	ans := make([]string, 0)
	var dfs func(int, []byte)
	dfs = func(index int, path []byte) {
		if index == len(digits) {
			ans = append(ans, string(path))
			return
		}

		d := digits[index] - '0'
		if d < 2 || d > 9 {
			return
		}

		for _, ch := range mapping[d] {
			dfs(index+1, append(path, byte(ch)))
		}
	}

	dfs(0, nil)
	return ans
}
