package main

func insert(intervals [][]int, newInterval []int) [][]int {
	ans := [][]int{}
	flag := false
	if len(intervals) == 0 {
		ans = append(ans, newInterval)
		return ans
	}
	for i := 0; i < len(intervals); i++ {
		if intervals[i][1] < newInterval[0] {
			ans = append(ans, intervals[i])
		} else {
			if intervals[i][0] > newInterval[1] {
				if !flag {
					ans = append(ans, newInterval)
					flag = true
				}
				ans = append(ans, intervals[i])
				continue
			} else {
				newInterval[0] = min(intervals[i][0], newInterval[0])
				newInterval[1] = max(intervals[i][1], newInterval[1])
			}
		}
	}
	if !flag {
		ans = append(ans, newInterval)
	}
	return ans
}

/*给你一个 无重叠的 ，按照区间起始端点排序的区间列表 intervals，其中 intervals[i] = [starti, endi] 表示第 i 个区间的开始和结束，并且 intervals 按照 starti 升序排列。同样给定一个区间 newInterval = [start, end] 表示另一个区间的开始和结束。

在 intervals 中插入区间 newInterval，使得 intervals 依然按照 starti 升序排列，且区间之间不重叠（如果有必要的话，可以合并区间）。

返回插入之后的 intervals。

注意 你不需要原地修改 intervals。你可以创建一个新数组然后返回它。



示例 1：

输入：intervals = [[1,3],[6,9]], newInterval = [2,5]
输出：[[1,5],[6,9]]
示例 2：

输入：intervals = [[1,2],[3,5],[6,7],[8,10],[12,16]], newInterval = [4,8]
输出：[[1,2],[3,10],[12,16]]
解释：这是因为新的区间 [4,8] 与 [3,5],[6,7],[8,10] 重叠。
思路：遍历 intervals，如果当前区间的结束时间小于 newInterval 的开始时间，则直接添加到结果中；如果当前区间的开始时间大于 newInterval 的结束时间，则直接添加到结果中；否则，合并当前区间和 newInterval。


*/
