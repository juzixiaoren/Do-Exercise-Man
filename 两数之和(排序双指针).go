package main

import "sort"

func twoSum(nums []int, target int) []int {
	copynums := make([]int, len(nums))
	copy(copynums, nums)
	var ans []int
	sort.Slice(copynums, func(i int, j int) bool {
		return copynums[i] < copynums[j]
	})
	var i int = 0
	var j int = len(nums) - 1
	for i < j {
		sum := copynums[i] + copynums[j]
		if sum == target {
			for k := 0; k < len(nums); k++ {
				if nums[k] == copynums[i] || nums[k] == copynums[j] {
					ans = append(ans, k)
					if len(ans) == 2 {
						return ans
					}
				}
			}
		} else if sum <= target {
			i++
		} else {
			j--
		}
	}
	return ans
}
