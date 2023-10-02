package leetcode

/*
1800. 最大升序子数组和
给你一个正整数组成的数组 nums ，返回 nums 中一个 升序 子数组的最大可能元素和。
子数组是数组中的一个连续数字序列。
已知子数组 [numsl, numsl+1, ..., numsr-1, numsr] ，若对所有 i（l <= i < r），numsi < numsi+1 都成立，则称这一子数组为升序子数组。
注意，大小为 1 的子数组也视作 升序 子数组。
链接：https://leetcode.cn/problems/maximum-ascending-subarray-sum
*/


func maxAscendingSum(nums []int) (ans int) {
	for i, n := 0, len(nums); i < n; {
		sum := nums[i]
		for i++; i < n && nums[i] > nums[i-1]; i++ {
			sum += nums[i]
		}
		if sum > ans {
			ans = sum
		}
	}
	return ans
}