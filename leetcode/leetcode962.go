package leetcode

/*
962. 最大宽度坡
给定一个整数数组A，坡是元组(i, j)，其中i < j且A[i] <= A[j]。这样的坡的宽度为j - i。
找出A中的坡的最大宽度，如果不存在，返回 0 。
链接：https://leetcode.cn/problems/maximum-width-ramp
*/

func maxWidthRamp(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		for j := len(nums) - 1; j > i && j-i > ans; j-- {

			if nums[j] >= nums[i] {
				ans = max(ans, j-i)
				break
			}
		}
	}
	return ans
}
