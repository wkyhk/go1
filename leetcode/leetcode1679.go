package leetcode

/*
1679. K 和数对的最大数目
给你一个整数数组 nums 和一个整数 k 。
每一步操作中，你需要从数组中选出和为 k 的两个整数，并将它们移出数组。
返回你可以对数组执行的最大操作数。
链接：https://leetcode.cn/problems/max-number-of-k-sum-pairs
*/
func maxOperations(nums []int, k int) int {
	m := make(map[int]int)
	res := 0
	for _, v := range nums {
		if m[k-v] > 0 {
			m[k-v]--
			res++
		} else {
			m[v]++
		}
	}
	return res
}
