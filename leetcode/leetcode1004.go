package leetcode

import (
	"fmt"
	"sort"
)

/*
1004. 最大连续1的个数 III
给定一个二进制数组 nums 和一个整数 k，如果可以翻转最多 k 个 0 ，则返回 数组中连续 1 的最大个数 。
链接：https://leetcode.cn/problems/max-consecutive-ones-iii/
*/
func longestOnes(nums []int, k int) (ans int) {
	n := len(nums)
	P := make([]int, n+1)
	for i, v := range nums {
		P[i+1] = P[i] + 1 - v
	}
	for right, v := range P {
		left := sort.SearchInts(P, v-k)
		ans = max(ans, right-left)
	}
	return
}




func TestLeetCode1004(){
	fmt.Println(longestOnes([]int{1,1,1,0,0,0,1,1,1,1,0},2))
	fmt.Println(longestOnes([]int{0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1},3))
}