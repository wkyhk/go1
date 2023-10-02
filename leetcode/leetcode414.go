package leetcode

import (
	"sort"
)

/*
414. 第三大的数
给你一个非空数组，返回此数组中 第三大的数 。如果不存在，则返回数组中最大的数。
https://leetcode.cn/problems/third-maximum-number/
*/
func thirdMax(nums []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(nums)))
	for i, diff := 1, 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			diff++
			if diff == 3 { // 此时 nums[i] 就是第三大的数
				return nums[i]
			}
		}
	}
	return nums[0]
}
