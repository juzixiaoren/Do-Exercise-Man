package main

import "slices"

func groupAnagrams(strs []string) [][]string {
	ans := [][]string{}
	m := map[string]int{}
	for i := 0; i < len(strs); i++ {
		copyStr := []byte(strs[i])
		slices.Sort(copyStr)

		index, ok := m[string(copyStr)]
		if ok {
			ans[index] = append(ans[index], strs[i])
		} else {
			ans = append(ans, []string{strs[i]})
			m[string(copyStr)] = len(ans) - 1
		}
	}
	return ans
}
