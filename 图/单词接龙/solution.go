package main

func ladderLength(beginWord string, endWord string, wordList []string) int {
	visit := map[string]struct{}{}
	hasWord := map[string]struct{}{}
	legalWordList := []byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	for _, word := range wordList {
		hasWord[word] = struct{}{}
	}
	if endWord == beginWord {
		return 1
	}
	_, ok := hasWord[endWord]
	if !ok {
		return 0
	}
	type Edge struct {
		word string
		step int
	}
	queue := []Edge{
		{
			word: beginWord,
			step: 1,
		},
	}
	visit[beginWord] = struct{}{}
	for len(queue) != 0 {
		cur := queue[0]
		queue = queue[1:]
		for i := 0; i < len([]byte(cur.word)); i++ {
			for _, legalWord := range legalWordList {
				if cur.word[i] == legalWord {
					continue
				} else {
					newWord := []byte(cur.word)
					newWord[i] = legalWord
					nextWord := string(newWord)
					if nextWord == endWord {
						return cur.step + 1
					}
					_, ok := visit[nextWord]
					if ok {
						continue
					} else {
						_, ok = hasWord[nextWord]
						if ok {
							queue = append(queue, Edge{
								word: nextWord,
								step: cur.step + 1,
							})
							visit[nextWord] = struct{}{}
						} else {
							continue
						}
					}
				}
			}
		}
	}
	return 0
}
