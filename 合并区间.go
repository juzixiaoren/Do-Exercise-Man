package main

import "sort"

func merge(intervals [][]int) [][]int {
	ans := [][]int{}
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	if len(intervals) < 1 {
		return intervals
	}
	union := []int{intervals[0][0], intervals[0][1]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= union[1] {
			union[1] = max(union[1], intervals[i][1])
		} else {
			ans = append(ans, union)
			union = intervals[i]
		}
	}
	ans = append(ans, union)
	return ans

}
