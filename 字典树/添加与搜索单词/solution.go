package main

type WordDictionary struct {
	root *Node
	ans  map[string]struct{}
}
type Node struct {
	next   map[byte]*Node
	isWord bool
	word   string
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &Node{
			next: map[byte]*Node{},
		},
		ans: map[string]struct{}{},
	}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	for i := 0; i < len(word); i++ {
		_, ok := cur.next[word[i]]
		if !ok {
			cur.next[word[i]] = &Node{
				next: map[byte]*Node{},
				word: cur.word + string(word[i]),
			}
			cur = cur.next[word[i]]
		} else {
			cur = cur.next[word[i]]
		}
		if i == len(word)-1 {
			cur.isWord = true
		}
	}
}

type pos struct {
	x int
	y int
}

func (this *WordDictionary) Search(board [][]byte, start pos) {

	type Edge struct {
		posi pos
		node *Node
		dir  int
	}

	direct := [][]int{
		{0, 1},
		{1, 0},
		{0, -1},
		{-1, 0},
	}

	node, ok := this.root.next[board[start.y][start.x]]

	if !ok {
		return
	}

	stack := []Edge{
		{
			posi: start,
			node: node,
			dir:  0,
		},
	}

	visit := map[pos]bool{}
	visit[start] = true

	for len(stack) > 0 {

		top := &stack[len(stack)-1]

		// 找到单词
		if top.node.isWord {
			this.ans[top.node.word] = struct{}{}
		}

		// 四个方向都搜索完
		if top.dir == 4 {

			// 回溯
			visit[top.posi] = false

			stack = stack[:len(stack)-1]

			continue
		}

		// 当前方向
		d := direct[top.dir]

		top.dir++

		next := pos{
			x: top.posi.x + d[0],
			y: top.posi.y + d[1],
		}

		if next.x < 0 ||
			next.x >= len(board[0]) ||
			next.y < 0 ||
			next.y >= len(board) {
			continue
		}

		if visit[next] {
			continue
		}

		nextNode, ok := top.node.next[board[next.y][next.x]]

		if !ok {
			continue
		}

		visit[next] = true

		stack = append(stack, Edge{
			posi: next,
			node: nextNode,
			dir:  0,
		})

	}
}

func findWords(board [][]byte, words []string) []string {
	wordDict := Constructor()
	for _, word := range words {
		wordDict.AddWord(word)
	}
	for i := 0; i < len(board[0]); i++ {
		for j := 0; j < len(board); j++ {
			wordDict.Search(board, pos{i, j})
		}
	}
	res := []string{}
	for word := range wordDict.ans {
		res = append(res, word)
	}
	return res
}
