package main

type Node struct {
	children map[byte]*Node
	isWord   bool
}
type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{
			children: make(map[byte]*Node),
			isWord:   false,
		},
	}

}

func (this *Trie) Insert(word string) {
	words := []byte(word)
	cur := this.root
	for i := 0; i < len(words); i++ {
		prefix := words[i]
		chiNode, ok := cur.children[prefix]
		if ok {
			cur = chiNode
		} else {
			newNode := &Node{children: make(map[byte]*Node)}
			cur.children[prefix] = newNode
			cur = newNode
		}
	}
	cur.isWord = true
}

func (this *Trie) Search(word string) bool {
	words := []byte(word)
	cur := this.root
	for i := 0; i < len(words); i++ {
		prefix := words[i]
		chiNode, ok := cur.children[prefix]
		if ok {
			cur = chiNode
		} else {
			return false
		}
	}
	if cur.isWord {
		return true
	}
	return false
}

func (this *Trie) StartsWith(prefix string) bool {
	words := []byte(prefix)
	cur := this.root
	for i := 0; i < len(words); i++ {
		prefix := words[i]
		chiNode, ok := cur.children[prefix]
		if ok {
			cur = chiNode
		} else {
			return false
		}
	}
	return true

}

/*
更经典的字典树解法，Node 保存：map[byte]*Node

构建的树为

root
 ↓
 a
 ↓
 p
 ↓
 p
 ↓
 l
 ↓
 e

 在最终节点isWord=true 即可


*/
