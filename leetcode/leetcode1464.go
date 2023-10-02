package leetcode

import (
	"fmt"
	"sort"
)

/*
1464. 数组中两元素的最大乘积
给你一个整数数组 nums，请你选择数组的两个不同下标 i 和 j，使 (nums[i]-1)*(nums[j]-1) 取得最大值。
请你计算并返回该式的最大值。
链接：https://leetcode.cn/problems/maximum-product-of-two-elements-in-an-array
*/
func maxProduct(nums []int) int {
	max1, max2 := nums[0], nums[1]
	for i := 2; i < len(nums); i++ {
		max1, max2 = top2(max1, max2, nums[i])
	}
	return (max1 - 1) * (max2 - 1)
}
func top2(a, b, c int) (int, int) {
	if a < b {
		a, b = b, a
	}
	if a < c {
		a, c = c, a
	}
	if b < c {
		b, c = c, b
	}
	return a, b
}
func maxProduct1(nums []int) int {
	sort.Ints(nums)
	return (nums[len(nums)-1] - 1) * (nums[len(nums)-2] - 1)
}
func Example1464() {
	arr := []int{3, 4, 5, 2}
	arr1 := []int{1, 5, 4, 5}
	arr2 := []int{3, 7}
	fmt.Println(maxProduct(arr))
	fmt.Println(maxProduct(arr1))
	fmt.Println(maxProduct(arr2))
	fmt.Println(maxProduct1(arr))
	fmt.Println(maxProduct1(arr1))
	fmt.Println(maxProduct1(arr2))
}
