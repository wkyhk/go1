package leetcode

import "fmt"

/*
1670. 设计前中后队列
请你设计一个队列，支持在前，中，后三个位置的 push 和 pop 操作。

请你完成 FrontMiddleBack 类：

FrontMiddleBack() 初始化队列。
void pushFront(int val) 将 val 添加到队列的 最前面 。
void pushMiddle(int val) 将 val 添加到队列的 正中间 。
void pushBack(int val) 将 val 添加到队里的 最后面 。
int popFront() 将 最前面 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
int popMiddle() 将 正中间 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
int popBack() 将 最后面 的元素从队列中删除并返回值，如果删除之前队列为空，那么返回 -1 。
请注意当有 两个 中间位置的时候，选择靠前面的位置进行操作。比方说：
将 6 添加到 [1, 2, 3, 4, 5] 的中间位置，结果数组为 [1, 2, 6, 3, 4, 5] 。
从 [1, 2, 3, 4, 5, 6] 的中间位置弹出元素，返回 3 ，数组变为 [1, 2, 4, 5, 6] 。
*/
type FrontMiddleBackQueue struct {
	queue []int
}

func Constructor() FrontMiddleBackQueue {
	return FrontMiddleBackQueue{[]int{}}
}

func (this *FrontMiddleBackQueue) PushFront(val int) {
	this.queue = append([]int{val}, this.queue...)
}
func (this *FrontMiddleBackQueue) PushMiddle(val int) {
	mid := len(this.queue) / 2
	this.queue = append(this.queue[:mid], append([]int{val}, this.queue[mid:]...)...)
}

func (this *FrontMiddleBackQueue) PushBack(val int) {
	this.queue = append(this.queue, val)
}

func (this *FrontMiddleBackQueue) PopFront() int {
	if len(this.queue) == 0 {
		return -1
	}
	ans := (this.queue)[0]
	this.queue = (this.queue)[1:]
	return ans
}

func (this *FrontMiddleBackQueue) PopMiddle() int {
	if len(this.queue) == 0 {
		return -1
	}
	mid := (len(this.queue) - 1) / 2
	ans := this.queue[mid]
	this.queue = append((this.queue)[:mid], (this.queue)[mid+1:]...)
	return ans
}

func (this *FrontMiddleBackQueue) PopBack() int {
	if len(this.queue) == 0 {
		return -1
	}
	ans := (this.queue)[len(this.queue)-1]
	this.queue = (this.queue)[:len(this.queue)-1]
	return ans
}
func TestLeetCode1670() {
	q := Constructor()
	q.PushFront(1)
	q.PushBack(2)
	q.PushMiddle(3)
	q.PushMiddle(4)
	fmt.Println(q.PopFront())
	fmt.Println(q.PopMiddle())
	fmt.Println(q.PopMiddle())
	fmt.Println(q.PopBack())
	fmt.Println(q.PopFront())
}
