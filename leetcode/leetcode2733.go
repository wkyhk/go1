package leetcode

import "sort"

/*
2733. 既不是最小值也不是最大值
给你一个整数数组 nums ，数组由 不同正整数 组成，请你找出并返回数组中 任一 既不是 最小值 也不是 最大值 的数字，如果不存在这样的数字，返回 -1 。

返回所选整数。
链接：https://leetcode.cn/problems/neither-minimum-nor-maximum/
*/
func findNonMinOrMax(nums []int) int {
	if len(nums) < 3 {
		return -1
	}
	sort.Ints(nums[:3]) // 只对前三个数排序
	return nums[1]
}
