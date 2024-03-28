package leetcode

import "math"

/*
2908. 元素和最小的山形三元组 I
给你一个下标从 0 开始的整数数组 nums 。

如果下标三元组 (i, j, k) 满足下述全部条件，则认为它是一个 山形三元组 ：

i < j < k
nums[i] < nums[j] 且 nums[k] < nums[j]
请你找出 nums 中 元素和最小 的山形三元组，并返回其 元素和 。如果不存在满足条件的三元组，返回 -1 。
链接：https://leetcode.cn/problems/minimum-sum-of-mountain-triplets-i/description/
*/
func minimumSum(nums []int) int {
	n := len(nums)
	suf := make([]int, n) // 后缀最小值
	suf[n-1] = nums[n-1]
	for i := n - 2; i > 1; i-- {
		suf[i] = min(suf[i+1], nums[i])
	}

	ans := math.MaxInt
	pre := nums[0] // 前缀最小值
	for j := 1; j < n-1; j++ {
		if pre < nums[j] && nums[j] > suf[j+1] { // 山形
			ans = min(ans, pre+nums[j]+suf[j+1]) // 更新答案
		}
		pre = min(pre, nums[j])
	}
	if ans == math.MaxInt {
		return -1
	}
	return ans
}
