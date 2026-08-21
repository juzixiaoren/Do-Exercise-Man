package main

type WordDictionary struct {
	root *Node
}
type Node struct {
	next   map[byte]*Node
	isWord bool
}

func Constructor() WordDictionary {
	return WordDictionary{
		root: &Node{
			next: map[byte]*Node{},
		},
	}
}

func (this *WordDictionary) AddWord(word string) {
	cur := this.root
	for i := 0; i < len(word); i++ {
		_, ok := cur.next[word[i]]
		if !ok {
			cur.next[word[i]] = &Node{
				next: map[byte]*Node{},
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

func (this *WordDictionary) Search(word string) bool {
	type Edge struct {
		level int
		node  *Node
	}
	stack := []Edge{}

	_, ok := this.root.next[word[0]]
	if !ok {
		return false
	} else {
		stack = append(stack, Edge{level: 0, node: this.root.next[word[0]]})
	}
	for len(stack) != 0 {
		edge := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if edge.level+1 == len(word) {
			if edge.node.isWord == true {
				return true
			} else {
				continue
			}
		}
		if word[edge.level+1] != '.' {
			nNode, ok := edge.node.next[word[edge.level+1]]
			if !ok {
				continue
			} else {
				stack = append(stack, Edge{level: edge.level + 1, node: nNode})
			}
		} else {
			if len(edge.node.next) == 0 {
				continue
			} else {
				for nodeIndex := range edge.node.next {
					stack = append(stack, Edge{edge.level + 1, edge.node.next[nodeIndex]})
				}
			}
		}
	}
	return false
}
