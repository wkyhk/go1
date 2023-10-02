package leetcode

import "fmt"

/*
2150. 找出数组中的所有孤独数字
给你一个整数数组 nums 。如果数字 x 在数组中仅出现 一次 ，且没有 相邻 数字（即，x + 1 和 x - 1）出现在数组中，则认为数字 x 是 孤独数字 。
返回 nums 中的 所有 孤独数字。你可以按 任何顺序 返回答案。
链接：https://leetcode.cn/problems/find-all-lonely-numbers-in-the-array
*/
func findLonely(nums []int) (ans []int) {
	nummap := make(map[int]int, 0)
	for _, v := range nums {
		nummap[v]++
	}
	for i, v := range nummap {
		if v == 1 {
			if nummap[i-1] == 0 && nummap[i+1] == 0 {
				ans = append(ans, i)
			}
		}
	}
	return
}
func Testleetcode2150() {
	fmt.Println(findLonely([]int{10, 6, 5, 8}))
	fmt.Println(findLonely([]int{1, 3, 5, 3}))
}
