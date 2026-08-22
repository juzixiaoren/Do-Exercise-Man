package main

func lengthOfLastWord(s string) int {
	ans := 0
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			s = s[:i]
		} else {
			break
		}
	}
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			break
		} else {
			ans++
		}
	}
	return ans
}
