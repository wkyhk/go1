package leetcode
/*
1658. 将 x 减到 0 的最小操作数
给你一个整数数组 nums 和一个整数 x 。每一次操作时，你应当移除数组 nums 最左边或最右边的元素，然后从 x 中减去该元素的值。请注意，需要修改
数组以供接下来的操作使用。
如果可以将 x恰好 减到0 ，返回 最小操作数 ；否则，返回 -1 。

链接：https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero

*/

func minOperations(nums []int, x int) int {
	s, n := 0, len(nums)
	right := n
	for right > 0 && s+nums[right-1] <= x { // 计算最长后缀
		right--
		s += nums[right]
	}
	if right == 0 && s < x { // 全部移除也无法满足要求
		return -1
	}
	ans := n + 1
	if s == x {
		ans = n - right
	}
	for left, num := range nums {
		s += num
		for ; right < n && s > x; right++ { // 缩小后缀长度
			s -= nums[right]
		}
		if s > x { // 缩小失败，说明前缀过长
			break
		}
		if s == x {
			ans = min(ans, left+1+n-right) // 前缀+后缀长度
		}
	}
	if ans > n {
		return -1
	}
	return ans
}



