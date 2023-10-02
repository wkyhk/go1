package leetcode

/*
494. 目标和
给你一个整数数组 nums 和一个整数 target 。
向数组中的每个整数前添加'+' 或 '-' ，然后串联起所有整数，可以构造一个 表达式 ：
例如，nums = [2, 1] ，可以在 2 之前添加 '+' ，在 1 之前添加 '-' ，然后串联起来得到表达式 "+2-1" 。
返回可以通过上述方法构造的、运算结果等于 target 的不同 表达式 的数目。
链接：https://leetcode.cn/problems/target-sum
*/
func findTargetSumWays(nums []int, target int) int {
	ans := 0
	var find func([]int, int)
	find = func(arr []int, n int) {
		if len(arr) == 0 {
			return
		}

		if len(arr) == 1 {
			if n+arr[0] == target {
				ans++
			}
			if n-arr[0] == target {
				ans++
			}
			return
		}
		find(arr[1:], n-arr[0])
		find(arr[1:], n+arr[0])
	}
	find(nums, ans)
	return ans
}
func Example494() {
	findTargetSumWays([]int{1, 0}, 1)
}
