package leetcode

import "fmt"

/*
704. 二分查找
给定一个n个元素有序的（升序）整型数组nums 和一个目标值target ，写一个函数搜索nums中的 target，如果目标值存在返回下标，否则返回 -1。
链接：https://leetcode.cn/problems/binary-search
*/
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		num := nums[mid]
		if num == target {
			return mid
		} else if num > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return -1
}

func Example704() {
	arr := []int{-1, 0, 3, 5, 9, 12}
	fmt.Println(search(arr, 2))
}
