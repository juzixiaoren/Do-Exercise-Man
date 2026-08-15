package main

import "fmt"

func solution(m int64) int64 {
	exp := qpow(2, m, 998244353)
	return famous(2, exp-1, 998244353)

}
func famous(base int64, exp int64, mod int64) int64 {
	exp = exp % (mod - 1)
	return qpow(base, exp, mod)
}
func qpow(base int64, exp int64, mod int64) int64 {
	var ans int64
	ans = 1
	for exp > 0 {
		if exp%2 == 1 {
			exp--
			ans = (ans * base) % mod
		} else {
			base = (base * base) % mod
			exp = exp / 2
		}
	}
	return ans

}
func main() {
	fmt.Println(solution(3))
}

/*
某团队有 2^m 个候选方案参与逐轮两两对比淘汰评审，方案编号依次为 1 到 2^m。每个方案的优先级评分就是其编号，评分越高的方案越优。

在任一轮对比中，若编号为 i 的方案与编号为 j 的方案被分到一起，则方案 i 胜出的概率为：

1
P(i 胜 j) = i / (i + j)
具体评审流程如下：

首轮，所有方案按初始编排顺序两两配对：位置 1 与位置 2 对比，位置 3 与位置 4 对比，依此类推，位置 2^m - 1 与位置 2^m 对比。位置上的方案编号不一定等于位置编号。
后续各轮，将上一轮胜出的方案按胜出先后排成一行，再依次每两个配对，胜者进入下一轮，直到产生唯一胜者。
初始编排是一个长度为 2^m 的排列，每个方案恰好出现一次，且顺序决定首轮配对关系。

请计算：共有多少种不同的初始编排顺序，能使编号为 1 的方案最终胜出的概率达到最大值。答案对 998244353 取模后输出。

名词解释：长度为 n 的排列是由 1, 2, ..., n 这 n 个整数按任意顺序组成的数组，每个整数恰好出现一次。例如 {3, 1, 4, 2} 是长度为 4 的排列，而 {2, 3, 2, 4} 和 {1, 5, 3, 2} 都不是排列，前者有重复元素，后者包含超出范围的数。

2^m-1
2^2^m-1
*/
