package main

import "strconv"

func summaryRanges(nums []int) []string {
	ans := []string{}
	var le int
	var ri int
	if len(nums) == 0 {
		return ans
	}
	if len(nums) == 1 {
		ans = append(ans, strconv.Itoa(nums[0]))
		return ans
	}
	le = nums[0]
	ri = nums[0]

	for i := 1; i < len(nums); i++ {
		if nums[i] == ri+1 {
			ri = nums[i]
		} else {
			if le == ri {
				ans = append(ans, strconv.Itoa(le))
			} else {
				ans = append(ans, strconv.Itoa(le)+"->"+strconv.Itoa(ri))
			}
			le = nums[i]
			ri = nums[i]
		}
	}
	if le == ri {
		ans = append(ans, strconv.Itoa(le))
	} else {
		ans = append(ans, strconv.Itoa(le)+"->"+strconv.Itoa(ri))
	}
	return ans
}

/*
给定一个  无重复元素 的 有序 整数数组 nums 。

区间 [a,b] 是从 a 到 b（包含）的所有整数的集合。

返回 恰好覆盖数组中所有数字 的 最小有序 区间范围列表 。也就是说，nums 的每个元素都恰好被某个区间范围所覆盖，并且不存在属于某个区间但不属于 nums 的数字 x 。

列表中的每个区间范围 [a,b] 应该按如下格式输出：

"a->b" ，如果 a != b
"a" ，如果 a == b


示例 1：

输入：nums = [0,1,2,4,5,7]
输出：["0->2","4->5","7"]
解释：区间范围是：
[0,2] --> "0->2"
[4,5] --> "4->5"
[7,7] --> "7"
示例 2：

输入：nums = [0,2,3,4,6,8,9]
输出：["0","2->4","6","8->9"]
解释：区间范围是：
[0,0] --> "0"
[2,4] --> "2->4"
[6,6] --> "6"
[8,9] --> "8->9"

思路：使用双指针，一个指针指向区间的起始位置，一个指针指向区间的结束位置。如果两个指针指向的元素相等，则说明区间只有一个元素，否则说明区间有多个元素。
遍历一次即可
*/
