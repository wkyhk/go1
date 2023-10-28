/*
1726. 同积元组
给你一个由 不同 正整数组成的数组 nums ，请你返回满足 a * b = c * d 的元组 (a, b, c, d) 的数量。其中 a、b、c 和 d 都是 nums 中的元素，且 a != b != c != d 。
链接：https://leetcode.cn/problems/tuple-with-same-product/
*/
package leetcode

func tupleSameProduct(nums []int) int {
	if len(nums) < 4 {
		return 0
	}
	ans := 0
	m := make(map[int]int, 0)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; i < len(nums); j++ {
			m[nums[i]*nums[j]]++
		}
	}
	for _, v := range m {
		ans += v * (v - 1) * 4
	}
	return ans
}
