package leetcode

import "fmt"

/*
2032. 至少在两个数组中出现的值
给你三个整数数组 nums1、nums2 和 nums3 ，请你构造并返回一个 元素各不相同的 数组，且由 至少 在 两个 数组中出现的所有值组成。数组中的元素可以按 任意 顺序排列。
链接：https://leetcode.cn/problems/two-out-of-three
*/
func twoOutOfThree(nums1, nums2, nums3 []int) (ans []int) {
	mask := map[int]int8{}
	for i, a := range [][]int{nums1, nums2, nums3} {
		for _, v := range a {
			mask[v] |= 1 << i // 标记 v 在哪些数组中出现过
		}
	}
	for v, m := range mask {
		if m > 2 && m != 4 { // 二进制包含至少两个 1，即 3 5 6 7。也可以用 bits.OnesCount8 来求
			ans = append(ans, v)
		}
	}
	return
}

func Example2032() {
	fmt.Println(twoOutOfThree([]int{1, 1, 3, 2}, []int{2, 3}, []int{3}))
}
