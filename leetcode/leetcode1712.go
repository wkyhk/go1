package leetcode

import "sort"

/*
1712. 将数组分成三个子数组的方案数
我们称一个分割整数数组的方案是 好的 ，当它满足：
数组被分成三个 非空 连续子数组，从左至右分别命名为 left ， mid ， right 。
left 中元素和小于等于 mid 中元素和，mid 中元素和小于等于 right 中元素和。
给你一个 非负 整数数组 nums ，请你返回 好的 分割 nums 方案数目。由于答案可能会很大，请你将结果对 10^9 + 7 取余后返回。
链接：https://leetcode.cn/problems/ways-to-split-array-into-three-subarrays
*/
func WaysToSplit(nums []int) int {
	left := 0
	mid := 0
	sum := 0
	ans := 0
	for _, v := range nums {
		sum += v
	}
	for i, v := range nums {
		left += v
		mid = 0
		for j := i + 1; j < len(nums)-1; j++ {
			mid += nums[j]
			if left > mid {
				continue
			}
			if sum-left-mid < mid {
				break
			}
			ans = (ans + 1) % (1e9 + 7)
		}
	}
	return ans
}
func waysToSplit(a []int) (ans int) {
	n := len(a)
	sum := make([]int, n+1)
	for i, v := range a {
		sum[i+1] = sum[i] + v
	}
	for r := 2; r < n && 3*sum[r] <= 2*sum[n]; r++ {
		l1 := sort.SearchInts(sum[1:r], 2*sum[r]-sum[n]) + 1
		ans += sort.SearchInts(sum[l1:r], sum[r]/2+1)
	}
	return ans % (1e9 + 7)
}
