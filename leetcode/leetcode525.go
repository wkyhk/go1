package leetcode

import "fmt"

/*
525. 连续数组
给定一个二进制数组 nums , 找到含有相同数量的 0 和 1 的最长连续子数组，并返回该子数组的长度。
链接：https://leetcode.cn/problems/contiguous-array/
*/
func findMaxLength(nums []int) (maxLength int) {
	mp := map[int]int{0: -1}
	counter := 0
	for i, num := range nums {
		if num == 1 {
			counter++
		} else {
			counter--
		}
		if prevIndex, has := mp[counter]; has {
			maxLength = max(maxLength, i-prevIndex)
		} else {
			mp[counter] = i
		}
	}
	return
}

func TestLeetCode525() {
	fmt.Println(findMaxLength([]int{0, 1}))
	fmt.Println(findMaxLength([]int{0, 1, 0}))
}
