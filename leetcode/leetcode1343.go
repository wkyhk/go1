package leetcode

/*
1343. 大小为 K 且平均值大于等于阈值的子数组数目
给你一个整数数组 arr 和两个整数 k 和 threshold 。
请你返回长度为 k 且平均值大于等于 threshold 的子数组数目。
https://leetcode.cn/problems/number-of-sub-arrays-of-size-k-and-average-greater-than-or-equal-to-threshold/
*/
func numOfSubarrays(arr []int, k int, threshold int) int {
	ans := 0
	sum := 0
	for i := 0; i < k; i++ {
		sum += arr[i]
	}
	if sum >= k*threshold {
		ans++
	}
	for i := k; i < len(arr); i++ {
		sum = sum + arr[i] - arr[i-k]
		if sum >= k*threshold {
			ans++
		}
	}
	return ans
}
