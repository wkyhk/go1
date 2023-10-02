package leetcode

/*
930. 和相同的二元子数组
给你一个二元数组 nums ，和一个整数 goal ，请你统计并返回有多少个和为 goal 的 非空 子数组。
子数组 是数组的一段连续部分。
https://leetcode.cn/problems/binary-subarrays-with-sum/
*/
func numSubarraysWithSum(nums []int, goal int) int {
	ans := 0
	for i, _ := range nums {
		sum := 0
		for j := i; j < len(nums); j++ {
			sum += nums[j]
			if sum == goal {
				ans++
			} else if sum > goal {
				break
			}
		}
	}
	return ans
}
