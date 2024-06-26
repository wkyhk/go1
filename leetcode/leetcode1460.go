package leetcode

import "sort"

/*
1460. 通过翻转子数组使两个数组相等
给你两个长度相同的整数数组 target 和 arr 。每一步中，你可以选择 arr 的任意 非空子数组 并将它翻转。你可以执行此过程任意次。
如果你能让 arr 变得与 target 相同，返回 True；否则，返回 False 。
链接：https://leetcode.cn/problems/make-two-arrays-equal-by-reversing-sub-arrays
*/
func canBeEqual(target, arr []int) bool {
	sort.Ints(target)
	sort.Ints(arr)
	for i, v := range target {
		if arr[i] != v {
			return false
		}
	}
	return true
}
