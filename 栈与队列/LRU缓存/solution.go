package main

type DoublyLinkList struct {
	Prev *DoublyLinkList
	Next *DoublyLinkList
	val  int
	key  int
}

type LRUCache struct {
	Head     *DoublyLinkList
	Tail     *DoublyLinkList
	Capacity int
	Cache    map[int]*DoublyLinkList
}

func Constructor(capacity int) LRUCache {
	head := new(DoublyLinkList)
	tail := new(DoublyLinkList)
	head.Next = tail
	tail.Prev = head
	cache := make(map[int]*DoublyLinkList)
	return LRUCache{
		Head:     head,
		Tail:     tail,
		Capacity: capacity,
		Cache:    cache,
	}

}

func (this *LRUCache) Get(key int) int {
	node, ok := this.Cache[key]
	if ok {
		this.DeleteNode(node)
		this.ToHead(node)
		return node.val
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Cache[key]
	if ok {
		node.val = value
		this.DeleteNode(node)
		this.ToHead(node)
		return
	} else {
		node := new(DoublyLinkList)
		node.key = key
		node.val = value
		if len(this.Cache) >= this.Capacity {
			delete(this.Cache, this.Tail.Prev.key)
			this.DeleteNode(this.Tail.Prev)
			this.Cache[key] = node
			this.ToHead(node)
		} else {
			this.Cache[key] = node
			this.ToHead(node)
		}
	}
}

func (this *LRUCache) ToHead(node *DoublyLinkList) {
	node.Next = this.Head.Next
	this.Head.Next.Prev = node
	this.Head.Next = node
	node.Prev = this.Head
}
func (this *LRUCache) DeleteNode(node *DoublyLinkList) {
	node.Prev.Next = node.Next
	node.Next.Prev = node.Prev
}
