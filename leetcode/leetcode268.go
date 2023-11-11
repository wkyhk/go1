package leetcode

/*
268. 丢失的数字
给定一个包含 [0, n] 中 n 个数的数组 nums ，找出 [0, n] 这个范围内没有出现在数组中的那个数。
链接：https://leetcode.cn/problems/missing-number/description/
*/
func missingNumber(nums []int) int {
	ans := 0
	for i := 0; i < len(nums); i++ {
		ans += i + 1 - nums[i]
	}
	return ans
}
