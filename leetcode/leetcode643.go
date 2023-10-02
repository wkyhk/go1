package leetcode

import "fmt"

/*
643. 子数组最大平均数 I
给你一个由 n 个元素组成的整数数组 nums 和一个整数 k 。
请你找出平均数最大且 长度为 k 的连续子数组，并输出该最大平均数。
任何误差小于 10-5 的答案都将被视为正确答案。
链接：https://leetcode.cn/problems/maximum-average-subarray-i
*/
func findMaxAverage(nums []int, k int) float64 {
	sum := 0
	for i := 0; i < k; i++ {
		sum += nums[i]
	}
	max := sum
	for i := k; i < len(nums); i++ {
		sum = sum - nums[i-k] + nums[i]

		if max < sum {
			max = sum
		}
	}
	return float64(max) / float64(k)

}
func Example643() {
	num := []int{1, 12, -5, -6, 50, 3}
	fmt.Println(findMaxAverage(num, 4))
	num1 := []int{5}
	fmt.Println(findMaxAverage(num1, 1))
}
