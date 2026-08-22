package main

func lengthOfLongestSubstring(s string) int {
	m := map[byte]struct{}{}
	left := 0
	ans := 0
	for right := 0; right < len(s); right++ {
		if len(m) == 0 {
			m[s[right]] = struct{}{}
			continue
		}
		_, ok := m[s[right]]
		if ok {
			ans = max(ans, right-left)
			for left < right && s[left] != s[right] {
				delete(m, s[left])
				left++
			}
			left++
		} else {
			m[s[right]] = struct{}{}
		}
	}
	ans = max(ans, len(s)-left)
	return ans
}
