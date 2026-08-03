package main

func copyRandomList2(head *Node) *Node {
	if head == nil {
		return nil
	}
	cur := head
	for cur != nil {
		temp := cur.Next
		cur.Next = new(Node)
		cur.Next.Val = cur.Val
		cur.Next.Next = temp
		cur = cur.Next.Next
	}
	cur = head
	for cur != nil {
		if cur.Random == nil {
			cur.Next.Random = nil
		} else {
			cur.Next.Random = cur.Random.Next
		}
		cur = cur.Next.Next
	}
	cur = head
	ans := cur.Next
	anscur := ans
	for anscur != nil {
		if anscur.Next != nil {
			cur.Next = cur.Next.Next
			anscur.Next = anscur.Next.Next
			anscur = anscur.Next
			cur = cur.Next
		} else {
			cur.Next = nil
			break
		}
	}
	return ans
}

/*
更聪明的写法
先在每一个节点后面增加一个复制节点
然后设置随机指针
然后分离链表
*/
