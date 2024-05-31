package leetcode

import "fmt"

/*
2965. 找出缺失和重复的数字
给你一个下标从 0 开始的二维整数矩阵 grid，大小为 n * n ，其中的值在 [1, n2] 范围内。除了 a 出现 两次，b 缺失 之外，每个整数都 恰好出现一次 。
任务是找出重复的数字a 和缺失的数字 b 。
返回一个下标从 0 开始、长度为 2 的整数数组 ans ，其中 ans[0] 等于 a ，ans[1] 等于 b 。
链接：https://leetcode.cn/problems/find-missing-and-repeated-values/description/
*/
func findMissingAndRepeatedValues(grid [][]int) []int {
	ans := make([]int, 2)
	sum := 0
	n := len(grid)
	m := make(map[int]int)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			sum += grid[i][j]
			m[grid[i][j]]++
			if m[grid[i][j]] == 2 {
				ans[0] = grid[i][j]
			}
		}
	}
	ans[1] = n*n*(n*n+1)/2 - sum + ans[0]
	return ans
}
func Testleetcode2965() {
	nums1 := [][]int{{1, 3}, {2, 2}}
	nums2 := [][]int{{9, 1, 7}, {8, 9, 2}, {3, 4, 6}}
	fmt.Println(findMissingAndRepeatedValues(nums1))
	fmt.Println(findMissingAndRepeatedValues(nums2))
}
