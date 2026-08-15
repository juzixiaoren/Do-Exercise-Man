package main

type Node struct {
	value    string
	children []*Node
	isWord   bool
}
type Trie struct {
	root *Node
}

func Constructor() Trie {
	return Trie{
		root: &Node{
			value:    "",
			children: []*Node{},
			isWord:   false,
		},
	}

}

func (this *Trie) Insert(word string) {
	words := []byte(word)
	cur := this.root
	flag := false
	for i := 0; i < len(words); i++ {
		flag = false
		prefix := words[0 : i+1]
		for _, chiNode := range cur.children {
			if string(prefix) == chiNode.value {
				cur = chiNode
				flag = true
				break
			} else {
				continue
			}
		}
		if !flag {
			newNode := &Node{value: string(prefix), children: []*Node{}}
			cur.children = append(cur.children, newNode)
			cur = newNode
		}
	}
	cur.isWord = true
}

func (this *Trie) Search(word string) bool {
	words := []byte(word)
	cur := this.root
	flag := false
	for i := 0; i < len(words); i++ {
		flag = false
		prefix := words[0 : i+1]
		for _, chiNode := range cur.children {
			if string(prefix) == chiNode.value {
				cur = chiNode
				flag = true
				break
			} else {
				continue
			}
		}
		if !flag {
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
	flag := false
	for i := 0; i < len(words); i++ {
		flag = false
		prefix := words[0 : i+1]
		for _, chiNode := range cur.children {
			if string(prefix) == chiNode.value {
				cur = chiNode
				flag = true
				break
			} else {
				continue
			}
		}
		if !flag {
			return false
		}
	}
	return true

}

/**
 * Your Trie object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Insert(word);
 * param_2 := obj.Search(word);
 * param_3 := obj.StartsWith(prefix);

 Trie（发音类似 "try"）或者说 前缀树 是一种树形数据结构，用于高效地存储和检索字符串数据集中的键。这一数据结构有相当多的应用情景，例如自动补全和拼写检查。

请你实现 Trie 类：

Trie() 初始化前缀树对象。
void insert(String word) 向前缀树中插入字符串 word 。
boolean search(String word) 如果字符串 word 在前缀树中，返回 true（即，在检索之前已经插入）；否则，返回 false 。
boolean startsWith(String prefix) 如果之前已经插入的字符串 word 的前缀之一为 prefix ，返回 true ；否则，返回 false 。


示例：

输入
["Trie", "insert", "search", "search", "startsWith", "insert", "search"]
[[], ["apple"], ["apple"], ["app"], ["app"], ["app"], ["app"]]
输出
[null, null, true, false, true, null, true]

解释
Trie trie = new Trie();
trie.insert("apple");
trie.search("apple");   // 返回 True
trie.search("app");     // 返回 False
trie.startsWith("app"); // 返回 True
trie.insert("app");
trie.search("app");     // 返回 True

解法：
"a"
"ap"
"app"
"appl"
"apple"
例如：
root
 ↓
"a"
 ↓
"ap"
 ↓
"app"
 ↓
"appl"
 ↓
"apple"
一层一层追下来，且如果是 word 保存 isword，方便查找。
更优方案：一个节点只保存一个 byte
*/
