package offer

import "fmt"

/*
剑指 Offer 21. 调整数组顺序使奇数位于偶数前面
输入一个整数数组，实现一个函数来调整该数组中数字的顺序，使得所有奇数在数组的前半部分，所有偶数在数组的后半部分。
https://leetcode.cn/problems/diao-zheng-shu-zu-shun-xu-shi-qi-shu-wei-yu-ou-shu-qian-mian-lcof/
*/
func exchange(nums []int) []int {
	l, r := 0, len(nums)-1
	for l < r {
		if !isOdd(nums[l]) && isOdd(nums[r]) {
			nums[l], nums[r] = nums[r], nums[l]
		}
		if isOdd(nums[l]) {
			l++
		}
		if !isOdd(nums[r]) {
			r--
		}
	}
	return nums
}

func isOdd(a int) bool {
	return a%2 == 1
}
func Example21() {
	fmt.Println(exchange([]int{1, 2, 3, 4}))
}
