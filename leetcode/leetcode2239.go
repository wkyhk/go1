package leetcode

import "math"

/*
2239. 找到最接近 0 的数字
给你一个长度为 n 的整数数组 nums ，请你返回 nums 中最 接近 0 的数字。如果有多个答案，请你返回它们中的 最大值 。
*/
func FindClosestNumber(nums []int) int {
	index := 0
	max := math.MaxInt
	for i, v := range nums {
		if max > abs(v) {
			index = i
			max = abs(v)
		} else if max == abs(v) {
			if v > 0 {
				index = i
			}
		}
	}
	return nums[index]
}
func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}
