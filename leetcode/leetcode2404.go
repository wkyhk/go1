package leetcode

import "fmt"

/*
2404. 出现最频繁的偶数元素给你一个整数数组 nums ，返回出现最频繁的偶数元素。
如果存在多个满足条件的元素，只需要返回 最小 的一个。如果不存在这样的元素，返回 -1 。
链接：https://leetcode.cn/problems/most-frequent-even-element/
*/
func mostFrequentEven(nums []int) int {
	m := make(map[int]int)
	ans := -1
	count := 0
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			m[nums[i]]++
		}
	}
	for i, v := range m {
		if v > count {
			ans = i
			count = v
		} else if v == count {
			ans = min(ans, i)
		}
	}
	return ans
}
func TestLeetCode2404() {
	fmt.Println(mostFrequentEven([]int{0, 1, 2, 2, 4, 4, 1}))
	fmt.Println(mostFrequentEven([]int{4, 4, 4, 9, 2, 4}))
	fmt.Println(mostFrequentEven([]int{29, 47, 21, 41, 13, 37, 25, 7}))
}
