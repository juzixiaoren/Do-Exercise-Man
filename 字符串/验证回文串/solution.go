package main

func isPalindrome(s string) bool {
	fix_s := []byte{}
	for i := 0; i < len(s); i++ {
		word := s[i]
		if word >= 'A' && word <= 'Z' {
			word = word + 32
			fix_s = append(fix_s, word)
		} else if word >= 'a' && word <= 'z' {
			fix_s = append(fix_s, word)
		} else if word >= '0' && word <= '9' {
			fix_s = append(fix_s, word)
		}
	}
	s = string(fix_s)
	left := 0
	right := len(s) - 1
	for left < right {
		if s[left] != s[right] {
			return false
		} else {
			left++
			right--
		}
	}
	return true
}
