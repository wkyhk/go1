package leetcode

/*
985. 查询后的偶数和
给出一个整数数组 A 和一个查询数组 queries。
对于第 i 次查询，有 val = queries[i][0], index = queries[i][1]，我们会把 val 加到 A[index] 上。然后，第 i 次查询的答案是 A 中偶数值的和。
（此处给定的 index = queries[i][1] 是从 0 开始的索引，每次查询都会永久修改数组 A。）
链接：https://leetcode.cn/problems/sum-of-even-numbers-after-queries
*/
func sumEvenAfterQueries(nums []int, queries [][]int) []int {
	sum := 0
	ans := make([]int, 0)
	for _, v := range nums {
		if v%2 == 0 {
			sum += v
		}
	}
	for _, v := range queries {
		if nums[v[1]]%2 == 0 {
			sum -= nums[v[1]]
		}
		nums[v[1]] += v[0]
		if nums[v[1]]%2 == 0 {
			sum += nums[v[1]]
		}
		ans = append(ans, sum)
	}
	return ans
}
