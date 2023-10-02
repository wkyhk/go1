package leetcode

/*
215. 数组中的第K个最大元素
给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。
你必须设计并实现时间复杂度为 O(n) 的算法解决此问题。
链接：https://leetcode.cn/problems/kth-largest-element-in-an-array
*/
func findKthLargest(nums []int, k int) int {
	num := make([]int, 2*1e4+1)
	for _, v := range nums {
		num[v+1e4]++
	}
	for i := len(num) - 1; i > 0; i-- {
		if k-num[i] <= 0 {
			return i - 1e4
		} else {
			k -= num[i]
		}
	}
	return 0
}
