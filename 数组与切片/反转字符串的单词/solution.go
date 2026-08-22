package main

import "slices"

func reverseWords(s string) string {
	words := []string{}
	word := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			if word != "" {
				words = append(words, word)
			}
			word = ""
			continue
		} else {
			word = word + string(s[i])
		}
	}
	if word != "" {
		words = append(words, word)
	}
	slices.Reverse(words)
	word = ""
	for i := 0; i < len(words); i++ {
		if i != len(words)-1 {
			word = word + words[i] + " "
		} else {
			word = word + words[i]
		}
	}
	return word
}
