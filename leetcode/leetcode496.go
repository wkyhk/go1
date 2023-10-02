package leetcode

import "fmt"

/*
 496. 下一个更大元素 I
nums1 中数字 x 的 下一个更大元素 是指 x 在 nums2 中对应位置 右侧 的 第一个 比 x 大的元素。
给你两个 没有重复元素 的数组 nums1 和 nums2 ，下标从 0 开始计数，其中nums1 是 nums2 的子集。
对于每个 0 <= i < nums1.length ，找出满足 nums1[i] == nums2[j] 的下标 j ，并且在 nums2 确定 nums2[j] 的 下一个更大元素 。如果不存在下一个更大元素，那么本次查询的答案是 -1 。
返回一个长度为 nums1.length 的数组 ans 作为答案，满足 ans[i] 是如上所述的 下一个更大元素 。
链接：https://leetcode.cn/problems/next-greater-element-i
*/
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	m := make(map[int]int)
	ans := make([]int, 0, len(nums1))
	for i, v := range nums2 {
		for j := i + 1; j < len(nums2); j++ {
			if v < nums2[j] {
				m[v] = nums2[j]
				break
			} else if j == len(nums2)-1 && v > nums2[j] {
				m[v] = -1
			}
		}
	}
	m[nums2[len(nums2)-1]] = -1
	for _, v := range nums1 {

		ans = append(ans, m[v])
	}
	return ans
}
func Exmaple496() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
	fmt.Println(nextGreaterElement([]int{2, 4}, []int{1, 2, 3, 4}))
}
