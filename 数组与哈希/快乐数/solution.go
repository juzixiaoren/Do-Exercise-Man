package main

func isHappy(n int) bool {
	var fast int
	var slow int
	fast = n
	slow = n
	for true {
		fast = docul(fast)
		fast = docul(fast)
		slow = docul(slow)
		if fast == 1 || slow == 1 {
			return true
		}
		if fast == slow {
			return false
		}
	}
	return true
}
func docul(n int) int {
	var digit []int
	sum := 0
	for n > 0 {
		digit = append(digit, n%10)
		n = n / 10
	}
	for i := 0; i < len(digit); i++ {
		sum += digit[i] * digit[i]
	}
	return sum
}

/*
编写一个算法来判断一个数 n 是不是快乐数。

「快乐数」 定义为：

对于一个正整数，每一次将该数替换为它每个位置上的数字的平方和。
然后重复这个过程直到这个数变为 1，也可能是 无限循环 但始终变不到 1。
如果这个过程 结果为 1，那么这个数就是快乐数。
如果 n 是 快乐数 就返回 true ；不是，则返回 false 。



示例 1：

输入：n = 19
输出：true
解释：
12 + 92 = 82
82 + 22 = 68
62 + 82 = 100
12 + 02 + 02 = 1
示例 2：

输入：n = 2
输出：false

解法: 使用快慢指针，快指针每次走两步，慢指针每次走一步，如果快指针等于慢指针，则说明有环，返回 false，如果快指针等于 1，则说明是快乐数，返回 true。
说明，所有的无限循环的题目，都会进入一个环，所以可以使用快慢指针来判断，而进入环的都可以用快慢指针来判断。
当你看到循环二字之时，请无理由使用快慢指针
*/
