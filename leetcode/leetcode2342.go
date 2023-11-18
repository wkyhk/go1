package leetcode

import "fmt"

/*
2342. 数位和相等数对的最大和
给你一个下标从 0 开始的数组 nums ，数组中的元素都是 正 整数。请你选出两个下标 i 和 j（i != j），且 nums[i] 的数位和 与  nums[j] 的数位和相等。

请你找出所有满足条件的下标 i 和 j ，找出并返回 nums[i] + nums[j] 可以得到的 最大值 。
链接：https://leetcode.cn/problems/max-sum-of-a-pair-with-equal-sum-of-digits/description/
*/
func maximumSum(nums []int) int {
	ans := -1
	m := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		temp := imum(nums[i])
		if v, ok := m[temp]; ok {
			ans = max(ans, nums[i]+v)
			if nums[i] > v {
				m[temp] = nums[i]
			}
		} else {
			m[temp] = nums[i]
		}
	}
	return ans
}
func imum(num int) int {
	ans := 0
	for num > 0 {
		ans += num % 10
		num = num / 10
	}
	return ans
}
func TestLeetcode2342() {
	nums := []int{18, 43, 36, 13, 7}
	fmt.Println(maximumSum(nums))
}
