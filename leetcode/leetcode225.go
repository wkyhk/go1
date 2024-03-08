package leetcode

/*
225. 用队列实现栈
请你仅使用两个队列实现一个后入先出（LIFO）的栈，并支持普通栈的全部四种操作（push、top、pop 和 empty）。

实现 MyStack 类：

void push(int x) 将元素 x 压入栈顶。
int pop() 移除并返回栈顶元素。
int top() 返回栈顶元素。
boolean empty() 如果栈是空的，返回 true ；否则，返回 false 。
链接：https://leetcode.cn/problems/implement-stack-using-queues/
*/
type MyStack struct {
	stack []int
}

func Constructor1() MyStack {
	return MyStack{}
}

func (this *MyStack) Push(x int) {
	this.stack = append(this.stack, x)
}

func (this *MyStack) Pop() int {
	ret := this.stack[len(this.stack)-1]
	this.stack = this.stack[:len(this.stack)-1]
	return ret
}

func (this *MyStack) Top() int {
	return this.stack[len(this.stack)-1]
}

func (this *MyStack) Empty() bool {
	return len(this.stack) == 0
}
