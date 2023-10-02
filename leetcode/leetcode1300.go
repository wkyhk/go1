package leetcode

import "sort"

/*
1300. 转变数组后最接近目标值的数组和
给你一个整数数组arr 和一个目标值target ，请你返回一个整数value，使得将数组中所有大于value 的值变成value 后，数组的和最接近 target
（最接近表示两者之差的绝对值最小）。

如果有多种使得和最接近target的方案，请你返回这些整数中的最小值。

请注意，答案不一定是arr 中的数字。

链接：https://leetcode.cn/problems/sum-of-mutated-array-closest-to-target
*/
func findBestValue(arr []int, target int) int {
	sort.Ints(arr)
	n := len(arr)
	prefix := make([]int, n+1)
	for i := 1; i <= n; i++ {
		prefix[i] = prefix[i-1] + arr[i-1]
	}
	r := arr[n-1]
	ans, diff := 0, target
	for i := 1; i <= r; i++ {
		index := sort.SearchInts(arr, i)
		if index < 0 {
			index = -index - 1
		}
		cur := prefix[index] + (n-index)*i
		if abs(cur-target) < diff {
			ans = i
			diff = abs(cur - target)
		}
	}
	return ans
}

/*func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}*/
