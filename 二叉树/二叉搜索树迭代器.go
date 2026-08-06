package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}
type BSTIterator struct {
	stack []*TreeNode
	root  *TreeNode
	cur   *TreeNode
}

func Constructor(root *TreeNode) BSTIterator {
	cur := root
	stack := []*TreeNode{}
	for cur != nil {
		stack = append(stack, cur)
		cur = cur.Left
	}
	return BSTIterator{
		root:  root,
		stack: stack,
		cur:   cur,
	}

}

func (this *BSTIterator) Next() int {
	for this.cur != nil {
		this.stack = append(this.stack, this.cur)
		this.cur = this.cur.Left
	}
	node := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	this.cur = node.Right
	return node.Val

}

func (this *BSTIterator) HasNext() bool {
	if len(this.stack) == 0 && this.cur == nil {
		return false
	}
	return true
}

/*
实现一个二叉搜索树迭代器类BSTIterator ，表示一个按中序遍历二叉搜索树（BST）的迭代器：
BSTIterator(TreeNode root) 初始化 BSTIterator 类的一个对象。BST 的根节点 root 会作为构造函数的一部分给出。指针应初始化为一个不存在于 BST 中的数字，且该数字小于 BST 中的任何元素。
boolean hasNext() 如果向指针右侧遍历存在数字，则返回 true ；否则返回 false 。
int next()将指针向右移动，然后返回指针处的数字。

即实现中序遍历，但是中间可以暂停下一步（一步一步遍历)

显式栈中序遍历的核心在于
步骤 1:左边走到头,入栈
步骤 2:栈顶出栈，处理
步骤 3:转向右子树
这里的关键点在于步骤1 为 cur!=nil
步骤三这里转向右子树如果右子树不存在，那么直接进入步骤 2，这个出的栈即是根节点，所以实现的左根右中序遍历
判断是否有下一个的条件为栈为空且cur=nil
所以一步一步操作需要额外保存 cur
*/
