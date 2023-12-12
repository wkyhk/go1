package leetcode

import "container/heap"

/*
2454. 下一个更大元素 IV
给你一个下标从 0 开始的非负整数数组 nums 。对于 nums 中每一个整数，你必须找到对应元素的 第二大 整数。

如果 nums[j] 满足以下条件，那么我们称它为 nums[i] 的 第二大 整数：

j > i
nums[j] > nums[i]
恰好存在 一个 k 满足 i < k < j 且 nums[k] > nums[i] 。
如果不存在 nums[j] ，那么第二大整数为 -1 。

比方说，数组 [1, 2, 4, 3] 中，1 的第二大整数是 4 ，2 的第二大整数是 3 ，3 和 4 的第二大整数是 -1 。
请你返回一个整数数组 answer ，其中 answer[i]是 nums[i] 的第二大整数。
链接:https://leetcode.cn/problems/next-greater-element-iv/
*/
type Pair struct {
	first  int
	second int
}

type hp []Pair

func (h hp) Len() int {
	return len(h)
}

func (h hp) Less(i, j int) bool {
	a, b := h[i], h[j]
	return a.first < b.first || a.first == b.first && a.second < b.second
}

func (h hp) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *hp) Push(v interface{}) {
	*h = append(*h, v.(Pair))
}

func (h *hp) Pop() interface{} {
	a := *h
	v := a[len(a)-1]
	*h = a[:len(a)-1]
	return v
}

func secondGreaterElement(nums []int) []int {
	res := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		res[i] = -1
	}
	st := []int{}
	q := hp{}
	for i := 0; i < len(nums); i++ {
		for len(q) > 0 && q[0].first < nums[i] {
			res[q[0].second] = nums[i]
			heap.Pop(&q)
		}
		for len(st) > 0 && nums[st[len(st)-1]] < nums[i] {
			heap.Push(&q, Pair{nums[st[len(st)-1]], st[len(st)-1]})
			st = st[:len(st)-1]
		}
		st = append(st, i)
	}
	return res
}
