/*
2679. 矩阵中的和
给你一个下标从 0 开始的二维整数数组 nums 。一开始你的分数为 0 。你需要执行以下操作直到矩阵变为空：

矩阵中每一行选取最大的一个数，并删除它。如果一行中有多个最大的数，选择任意一个并删除。
在步骤 1 删除的所有数字中找到最大的一个数字，将它添加到你的 分数 中。
请你返回最后的 分数 。
链接：https://leetcode.cn/problems/sum-in-a-matrix/?envType=daily-question&envId=2023-10-15
*/
package leetcode

import (
	"fmt"
	"sort"
)

func matrixSum(nums [][]int) int {
	res := 0
	m := len(nums)
	n := len(nums[0])
	for i := 0; i < m; i++ {
		sort.Ints(nums[i])
	}
	for j := 0; j < n; j++ {
		maxVal := 0
		for i := 0; i < m; i++ {
			if nums[i][j] > maxVal {
				maxVal = nums[i][j]
			}
		}
		res += maxVal
	}
	return res
}
func TestmamatrixSum() {
	nums := [][]int{{7, 2, 1}, {6, 4, 2}, {6, 5, 3}, {3, 2, 1}}
	ans := matrixSum(nums)
	fmt.Println(ans)

}
