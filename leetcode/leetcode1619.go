package leetcode

import "sort"

/*
1619. 删除某些元素后的数组均值
给你一个整数数组arr，请你删除最小5%的数字和最大 5%的数字后，剩余数字的平均值。
与 标准答案误差在10-5的结果都被视为正确结果。
链接：https://leetcode.cn/problems/mean-of-array-after-removing-some-elements
*/
func trimMean(arr []int) float64 {
	n := len(arr) / 20
	ans := 0
	sort.Ints(arr)
	for i := n; i < len(arr)-n; i++ {
		ans += arr[i]
	}
	return float64(ans) / float64(18*n)
}
