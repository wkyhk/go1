package leetcode

import "fmt"

/*
2789. 合并后数组中的最大元素
给你一个下标从 0 开始、由正整数组成的数组 nums 。

你可以在数组上执行下述操作 任意 次：

选中一个同时满足 0 <= i < nums.length - 1 和 nums[i] <= nums[i + 1] 的整数 i 。将元素 nums[i + 1] 替换为 nums[i] + nums[i + 1] ，并从数组中删除元素 nums[i] 。
返回你可以从最终数组中获得的 最大 元素的值。
链接：https://leetcode.cn/problems/largest-element-in-an-array-after-merge-operations/description/
*/
func maxArrayValue(nums []int) int64 {
	ans := 0
	for i := len(nums) - 1; i >= 0; i-- {
		if ans >= nums[i] {
			ans = ans + nums[i]
		} else {
			ans = nums[i]
		}
	}
	return int64(ans)
}
func TestLeetcode2789() {
	nums1 := []int{2, 3, 7, 9, 3}
	fmt.Println(maxArrayValue(nums1))
	nums2 := []int{5, 3, 3}
	fmt.Println(maxArrayValue(nums2))
}
