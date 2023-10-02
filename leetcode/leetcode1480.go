package leetcode

/*
1480. 一维数组的动态和
给你一个数组 nums 。数组「动态和」的计算公式为：runningSum[i] = sum(nums[0]…nums[i]) 。
请返回 nums 的动态和。
链接：https://leetcode.cn/problems/running-sum-of-1d-array
*/
func RunningSum(nums []int) []int {
	sum := 0
	for i, v := range nums {
		nums[i] = sum + v
		sum = nums[i]
	}
	return nums
}
