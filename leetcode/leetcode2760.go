package leetcode

import "fmt"

/*
2760. 最长奇偶子数组
给你一个下标从 0 开始的整数数组 nums 和一个整数 threshold 。

请你从 nums 的子数组中找出以下标 l 开头、下标 r 结尾 (0 <= l <= r < nums.length) 且满足以下条件的 最长子数组 ：

nums[l] % 2 == 0
对于范围 [l, r - 1] 内的所有下标 i ，nums[i] % 2 != nums[i + 1] % 2
对于范围 [l, r] 内的所有下标 i ，nums[i] <= threshold
以整数形式返回满足题目要求的最长子数组的长度。

注意：子数组 是数组中的一个连续非空元素序列。
链接：https://leetcode.cn/problems/longest-even-odd-subarray-with-threshold/
*/
func longestAlternatingSubarray(nums []int, threshold int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 && nums[i] <= threshold {
			temp := 1
			for j := i + 1; j < len(nums); j++ {
				if nums[j]%2 == 1 && temp%2 == 1 && nums[j] <= threshold {
					temp++
				} else if nums[j]%2 == 0 && temp%2 == 0 && nums[j] <= threshold {
					temp++
				} else {
					i = j - 1
					break
				}
			}
			ans = max(ans, temp)
		}
	}
	return ans
}
func Testleetcode2760() {
	nums := []int{4, 10, 3}
	fmt.Println(longestAlternatingSubarray(nums, 10))
}
