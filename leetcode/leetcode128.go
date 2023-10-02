package leetcode

import (
	"fmt"
)

/*
128. 最长连续序列
给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
链接：https://leetcode.cn/problems/longest-consecutive-sequence
*/
func longestConsecutive(nums []int) (res int) {
	set := map[int]struct{}{}

	for _, num := range nums {
		set[num] = struct{}{}
	}
	for _, num := range nums {
		_, ok := set[num-1]
		if !ok {
			cnt := 0
			_, ok = set[num]
			for ok {
				cnt++
				num++
				_, ok = set[num]
			}
			res = max(res, cnt)
		}
	}
	return
}

func Example128() {
	nums := []int{100, 4, 200, 1, 3, 2}
	nums1 := []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1}
	fmt.Println(longestConsecutive(nums))
	fmt.Println(longestConsecutive(nums1))
}
