package leetcode

import (
	"fmt"
	"sort"
)

/*
665. 非递减数列
给你一个长度为 n 的整数数组 nums ，请你判断在 最多 改变 1 个元素的情况下，该数组能否变成一个非递减数列。
我们是这样定义一个非递减数列的： 对于数组中任意的 i (0 <= i <= n-2)，总满足 nums[i] <= nums[i + 1]。
链接：https://leetcode.cn/problems/non-decreasing-array
*/
func checkPossibility(nums []int) bool {
	for i := 0; i < len(nums)-1; i++ {
		x, y := nums[i], nums[i+1]
		if x > y {
			nums[i] = y
			if sort.IntsAreSorted(nums) {
				return true
			}
			nums[i] = x // 复原
			nums[i+1] = x
			return sort.IntsAreSorted(nums)
		}
	}
	return true
}

func Example665() {
	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(checkPossibility(arr))
}
