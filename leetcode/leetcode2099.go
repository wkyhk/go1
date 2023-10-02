package leetcode

import (
	"fmt"
	"sort"
)

/*
2099. 找到和最大的长度为 K 的子序列
给你一个整数数组 nums 和一个整数 k 。你需要找到 nums 中长度为 k 的 子序列 ，且这个子序列的 和最大 。
请你返回 任意 一个长度为 k 的整数子序列。
子序列 定义为从一个数组里删除一些元素后，不改变剩下元素的顺序得到的数组。
链接：https://leetcode.cn/problems/find-subsequence-of-length-k-with-the-largest-sum
*/

func maxSubsequence(nums []int, k int) []int {
	var n int = len(nums)

	var nums2 []int = make([]int, n)
	copy(nums2, nums)
	sort.Ints(nums2)

	var num_freq map[int]int = make(map[int]int)
	for i := n - 1; i > n-k-1; i-- {
		num_freq[nums2[i]]++
	}

	var res []int = make([]int, 0)
	for _, x := range nums {
		if num_freq[x] > 0 {
			res = append(res, x)
			num_freq[x]--
		}
	}
	return res
}

func Example2099() {
	nums := []int{2, 1, 3, 3}
	nums1 := []int{-1, -2, 3, 4}
	nums2 := []int{3, 4, 3, 3}
	fmt.Println(maxSubsequence(nums, 2))
	fmt.Println(maxSubsequence(nums1, 3))
	fmt.Println(maxSubsequence(nums2, 2))
}
