package main

import "fmt"

func solution(n int, k int, S int, R int) []int {
	ans := []int{}
	leftNum := n - k
	leftLevel := R
	rightNum := k
	rightLevel := S - R
	posibleLevel := []int{
		1, 2, 3, 4, 5, 6,
	}
	for _, gap := range posibleLevel {
		if leftNum < leftLevel || leftNum*gap < leftLevel { //leftNum*gap < leftLevel 说明 gap 太小，无法满足 leftLevel
			continue
		} else if rightNum*gap > rightLevel || rightNum*6 < rightLevel { //rightNum*gap > rightLevel 说明 gap 太大，无法满足 rightLevel
			continue
		} else {
			extra := leftLevel - leftNum //extra 为 leftLevel 与 leftNum 的差值，即需要额外分配的等级
			for i := 0; i < leftNum; i++ {
				newLevel := 1
				for extra != 0 && newLevel < gap { //newLevel < gap 说明 newLevel 还可以继续增加
					newLevel++
					extra--
				}
				ans = append(ans, newLevel)
			}
			extra = rightLevel - rightNum*gap
			for j := 0; j < rightNum; j++ {
				newLevel := gap
				for extra != 0 && newLevel < 6 {
					extra--
					newLevel++
				}
				ans = append(ans, newLevel)
			}

			break
		}
	}
	return ans
}
func main() {
	ans := solution(4, 2, 15, 5)
	if len(ans) == 0 {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}

/*
题目 1：等级序列还原
题目内容
产线质检得到 n 件工件的质量评级序列 {v1, v2, ..., vn}，每件评级为 1 至 6 级之间的整数，评级总和记为 S。

随后质检流程升级，评级最高的 k 件工件被召回重测。若存在多件最高评级，可以从相同评级的工件中任意选择。召回后剩余 n - k 件的评级总和记为 R。已知 1 <= k < n。

现在只知道四个整数 (n, k, S, R)，请还原出一组可能的原始评级序列 {v1, v2, ..., vn}，满足：

每个 vi 属于 {1, 2, 3, 4, 5, 6}；
所有评级的总和等于 S；
剔除评级最高的 k 件后，剩余 n - k 件的评级和等于 R。
若不存在满足条件的序列，输出 -1。

输入描述
在一行上输入四个整数 n, k, S, R。

数据范围：

2 <= n <= 200000
1 <= k < n
1 <= R < S <= 1.2 * 10^6
输出描述
如果有解，输出一行 n 个整数，代表任意一个合法评级序列，顺序任意。

无解则输出 -1。

思路，一定会有 n-k 个工件<某个值
k 个工件大于某个值
这个值只能从 1，2 ，3，4，5，6 里面取，直接枚举即可
*/
