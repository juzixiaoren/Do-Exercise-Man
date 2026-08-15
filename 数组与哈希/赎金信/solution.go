package main

func canConstruct(ransomNote string, magazine string) bool {
	num := make([]int, 26)
	for i := 0; i < len(magazine); i++ {
		num[magazine[i]-'a']++
	}
	for i := 0; i < len(ransomNote); i++ {
		num[ransomNote[i]-'a']--
		if num[ransomNote[i]-'a'] < 0 {
			return false
		}
	}
	return true
}
