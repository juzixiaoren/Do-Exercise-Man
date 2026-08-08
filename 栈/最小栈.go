package main

type MinStack struct {
	stack      []int
	minimstack []int
}

func Constructor() MinStack {
	return MinStack{}

}

func (this *MinStack) Push(value int) {
	if len(this.minimstack) == 0 ||
		value <= this.minimstack[len(this.minimstack)-1] {
		this.minimstack = append(this.minimstack, value)
	}
	this.stack = append(this.stack, value)
}

func (this *MinStack) Pop() {
	if len(this.stack) == 0 {
		return
	}
	value := this.stack[len(this.stack)-1]
	if value == this.minimstack[len(this.minimstack)-1] {
		this.minimstack = this.minimstack[:len(this.minimstack)-1]
	}
	this.stack = this.stack[:len(this.stack)-1]
}

func (this *MinStack) Top() int {
	if len(this.stack) == 0 {
		return 0
	}
	return this.stack[len(this.stack)-1]

}

func (this *MinStack) GetMin() int {
	if len(this.minimstack) == 0 {
		return 0
	}
	return this.minimstack[len(this.minimstack)-1]

}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(value);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
/*
最小栈，思路：使用两个栈，一个栈存储数据，一个栈存储最小值，当数据入栈时，如果数据小于等于最小值栈的栈顶元素，则将数据入最小值栈，当数据出栈时，如果数据等于最小值栈的栈顶元素，则将最小值栈的栈顶元素出栈，当获取最小值时，返回最小值栈的栈顶元素
核心：保存上一个最小值，因此用栈
不用辅助栈的方法：使用一个最小值变量
当存入的值为新的最小值是，比如原来最小值是 5，新的为 3，那么存入为 2*3-5=1
当出栈发现小于最小值时（1<3)说明这个是最小值，出栈后，最小值为 2*最小值-当前值即2*3-1=5恢复复成 5
*/
