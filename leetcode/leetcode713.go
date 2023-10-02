package leetcode

/*
713. 乘积小于 K 的子数组
给你一个整数数组 nums 和一个整数 k ，请你返回子数组内所有元素的乘积严格小于 k 的连续子数组的数目。
https://leetcode.cn/problems/subarray-product-less-than-k/
*/
func numSubarrayProductLessThanK(nums []int, k int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		temp := 1
		for j := i; j < len(nums); j++ {
			temp *= nums[j]
			if temp < k {
				ans++
			} else {
				break
			}
		}
	}
	return ans
}
