package leetcode

import "fmt"

/*
219. 存在重复元素 II
给你一个整数数组 nums 和一个整数 k ，判断数组中是否存在两个 不同的索引 i 和 j ，满足 nums[i] == nums[j] 且 abs(i - j) <= k 。如果存在，返回 true ；否则，返回 false 。
链接：https://leetcode.cn/problems/contains-duplicate-ii
*/
func containsNearbyDuplicate(nums []int, k int) bool {
	m := make(map[int]int, 0)
	for i, v := range nums {
		if i1, ok := m[v]; ok {
			if i-i1 <= k {
				return true
			}
			m[v] = i
		} else {
			m[v] = i
		}
	}
	return false
}
func TestLeetCode219() {
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1}, 3))
	fmt.Println(containsNearbyDuplicate([]int{1, 0, 1, 1}, 1))
	fmt.Println(containsNearbyDuplicate([]int{1, 2, 3, 1, 2, 3}, 2))
}
