package leetcode

import "fmt"

/*
2348. 全 0 子数组的数目
给你一个整数数组 nums ，返回全部为 0 的 子数组 数目。
子数组 是一个数组中一段连续非空元素组成的序列。
https://leetcode.cn/problems/number-of-zero-filled-subarrays/
*/
func zeroFilledSubarray(nums []int) (ans int64) {
	c := 0
	for _, num := range nums {
		if num == 0 {
			c++
			ans += int64(c)
		} else {
			c = 0
		}
	}
	return
}

func Example2348() {
	nums := []int{1, 3, 0, 0, 2, 0, 0, 4}
	nums1 := []int{0, 0, 0, 2, 0, 0}
	nums2 := []int{2, 10, 2019}
	fmt.Println(zeroFilledSubarray(nums))
	fmt.Println(zeroFilledSubarray(nums1))
	fmt.Println(zeroFilledSubarray(nums2))
}
