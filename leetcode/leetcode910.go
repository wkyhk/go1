package leetcode

import (
	"fmt"
	"sort"
)

/*
910. 最小差值 II
给你一个整数数组 nums，和一个整数 k 。
对于每个下标 i（0 <= i < nums.length），将 nums[i] 变成 nums[i] + k 或 nums[i] - k 。
nums 的 分数 是 nums 中最大元素和最小元素的差值。
在更改每个下标对应的值之后，返回 nums 的最小 分数 。
链接：https://leetcode.cn/problems/smallest-range-ii
*/

func Max(x, y int) int {
	if x > y {
		return x
	} else {
		return y
	}
}

func Min(x, y int) int {
	if x < y {
		return x
	} else {
		return y
	}
}

func smallestRangeII(nums []int, k int) int {
	sort.Ints(nums)
	res := nums[len(nums)-1] - nums[0]
	var max, min int
	for i := 1; i < len(nums); i ++ {
		min = Min(nums[0] + k, nums[i] - k)
		max = Max(nums[len(nums)-1] - k, nums[i-1] + k)
		res = Min(res, max - min)
	}
	return res
}


func TestLeetCode910(){
	fmt.Println(smallestRangeII([]int{1},0))
	fmt.Println(smallestRangeII([]int{0,10},2))
	fmt.Println(smallestRangeII([]int{1,3,6},3))

}