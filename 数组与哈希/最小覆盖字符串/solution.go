package main

func minWindow(s string, t string) string {
	if len(t) == 0 || len(s) < len(t) {
		return ""
	}

	need := map[byte]int{}
	window := map[byte]int{}

	for i := 0; i < len(t); i++ {
		need[t[i]]++
	}

	left := 0
	count := 0

	start := 0
	minLen := len(s) + 1

	for right := 0; right < len(s); right++ {
		word := s[right]

		// 当前字符是 t 需要的
		if need[word] > 0 {
			window[word]++

			// 只有没有超过需求数量，才算真正匹配一个
			if window[word] <= need[word] {
				count++
			}
		}

		// 当前窗口已经覆盖 t
		for count == len(t) {
			// 先记录当前合法窗口
			if right-left+1 < minLen {
				minLen = right - left + 1
				start = left
			}

			// 尝试移动左边界
			leftChar := s[left]

			if need[leftChar] > 0 {
				// 如果删掉的是一个"必要字符"
				if window[leftChar] <= need[leftChar] {
					count--
				}

				window[leftChar]--
			}

			left++
		}
	}

	if minLen == len(s)+1 {
		return ""
	}

	return s[start : start+minLen]
}
