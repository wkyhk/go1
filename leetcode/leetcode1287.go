package leetcode

/*
1287. 有序数组中出现次数超过25%的元素
给你一个非递减的 有序 整数数组，已知这个数组中恰好有一个整数，它的出现次数超过数组元素总数的 25%。

请你找到并返回这个整数
链接：https://leetcode.cn/problems/element-appearing-more-than-25-in-sorted-array/description/
*/
func findSpecialInteger(arr []int) int {
	m := make(map[int]int)
	for i := 0; i < len(arr); i++ {
		m[arr[i]]++
	}
	for k, v := range m {
		if v > len(arr)/4 {
			return k
		}
	}
	return 0
}
