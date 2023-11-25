package leetcode

/*
2824. 统计和小于目标的下标对数目
给你一个下标从 0 开始长度为 n 的整数数组 nums 和一个整数 target ，请你返回满足 0 <= i < j < n 且 nums[i] + nums[j] < target 的下标对 (i, j) 的数目。
链接：https://leetcode.cn/problems/count-pairs-whose-sum-is-less-than-target/
*/
func countPairs(nums []int, target int) int {
	ans := 0
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] < target {
				ans++
			}
		}
	}
	return ans
}
