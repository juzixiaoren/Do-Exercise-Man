package main

func containsNearbyDuplicate(nums []int, k int) bool {
	m := map[int]int{}
	for i := 0; i <= k && i < len(nums); i++ {
		_, ok := m[nums[i]]
		if ok {
			return true
		} else {
			m[nums[i]] = i
		}
	}
	if len(nums) <= k+1 {
		return false
	}
	ri := k + 1
	le := 0
	for ri < len(nums) {
		delete(m, nums[le])
		_, ok := m[nums[ri]]
		if ok {
			return true
		} else {
			m[nums[ri]] = ri
		}
		ri++
		le++
	}
	return false
}

/*
219. 存在重复元素 II
给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。



示例 1：

输入：nums = [1,2,3,1], k = 3
输出：true
示例 2：

输入：nums = [1,0,1,1], k = 1
输出：true
示例 3：

输入：nums = [1,2,3,1,2,3], k = 2
输出：false

思路：维护一个大小为k的滑动窗口，窗口内不能有重复元素，如果窗口内有重复元素，则返回true，否则返回false。使用hashmap 键不唯一的特性
补充，go 的 map 使用 m:map[int]int{} 会自动初始化.
获取元素 value,ok := m[nums[i]]  // ok 为 true 表示元素存在，false 表示元素不存在
删除元素 delete(m, nums[i])
go 没有 hash set,使用m := map[int]struct{}{}来弄 set
*/
