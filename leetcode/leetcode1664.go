package leetcode

import "fmt"

/*
1664. 生成平衡数组的方案数
给你一个整数数组 nums 。你需要选择 恰好 一个下标（下标从 0 开始）并删除对应的元素。请注意剩下元素的下标可能会因为删除操作而发生改变。
比方说，如果 nums = [6,1,7,4,1] ，那么：
选择删除下标 1 ，剩下的数组为 nums = [6,7,4,1] 。
选择删除下标 2 ，剩下的数组为 nums = [6,1,4,1] 。
选择删除下标 4 ，剩下的数组为 nums = [6,1,7,4] 。
如果一个数组满足奇数下标元素的和与偶数下标元素的和相等，该数组就是一个 平衡数组 。
请你返回删除操作后，剩下的数组 nums 是 平衡数组 的 方案数 。
链接：https://leetcode.cn/problems/ways-to-make-a-fair-array
*/
func waysToMakeFair(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		num := make([]int, len(nums))
		copy(num, nums)
		num = append(num[:i], num[i+1:]...)
		if fair(num) {
			count++
		}
	}
	return count
}

func fair(nums []int) bool {
	ans := 0
	for i := 0; i < len(nums); i++ {
		if i%2 == 0 {
			ans += nums[i]
		} else {
			ans -= nums[i]
		}
	}
	return ans == 0
}
func Example1664() {
	nums := []int{2, 1, 6, 4}
	nums1 := []int{1, 1, 1}
	nums2 := []int{1, 2, 3}
	fmt.Println(waysToMakeFair(nums))
	fmt.Println(waysToMakeFair(nums1))
	fmt.Println(waysToMakeFair(nums2))
}
