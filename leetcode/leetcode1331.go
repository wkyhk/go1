package leetcode

import "sort"

/*
1331. 数组序号转换
给你一个整数数组 arr ，请你将数组中的每个元素替换为它们排序后的序号。
序号代表了一个元素有多大。序号编号的规则如下：
序号从 1 开始编号。
一个元素越大，那么序号越大。如果两个元素相等，那么它们的序号相同。
每个数字的序号都应该尽可能地小。
 链接：https://leetcode.cn/problems/rank-transform-of-an-array
*/
func ArrayRankTransform(arr []int) []int {
	a := append([]int{}, arr...)
	sort.Ints(a)
	ranks := map[int]int{}
	for _, v := range a {
		if _, ok := ranks[v]; !ok {
			ranks[v] = len(ranks) + 1
		}
	}
	for i, v := range arr {
		arr[i] = ranks[v]
	}
	return arr
}
