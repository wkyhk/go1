package leetcode

import "sort"

/*
 1636. 按照频率将数组升序排序
给你一个整数数组 nums ，请你将数组按照每个值的频率 升序 排序。如果有多个值的频率相同，请你按照数值本身将它们 降序 排序。
请你返回排序后的数组。
https://leetcode.cn/problems/sort-array-by-increasing-frequency/
*/
func frequencySort(nums []int) []int {
	cnt := map[int]int{}
	for _, x := range nums {
		cnt[x]++
	}
	sort.Slice(nums, func(i, j int) bool {
		a, b := nums[i], nums[j]
		return cnt[a] < cnt[b] || cnt[a] == cnt[b] && a > b
	})
	return nums
}
