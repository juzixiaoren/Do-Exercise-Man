package main

func twoSum1(nums []int, target int) []int {
	var ans []int
	mp := map[int]int{}
	for i := 0; i < len(nums); i++ {
		tofind := target - nums[i]
		j, ok := mp[tofind]
		if ok {
			ans = append(ans, i)
			ans = append(ans, j)
			return ans
		} else {
			mp[nums[i]] = i
		}
	}
	return ans
}
