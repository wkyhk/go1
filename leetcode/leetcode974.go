package leetcode

/*
974. 和可被 K 整除的子数组
给定一个整数数组 nums 和一个整数 k ，返回其中元素之和可被 k 整除的（连续、非空） 子数组 的数目。
子数组 是数组的 连续 部分。
https://leetcode.cn/problems/subarray-sums-divisible-by-k/
*/
func subarraysDivByK(nums []int, k int) int {
	ans := 0
	m := make(map[int]int, 0)
	sum := 0
	for _, v := range nums {
		sum += v
		value := (sum%k + k) % k
		ans += m[value]
		m[value]++
	}
	return ans
}
