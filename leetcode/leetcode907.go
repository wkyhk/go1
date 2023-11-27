/*
907. 子数组的最小值之和
给定一个整数数组 arr，找到 min(b) 的总和，其中 b 的范围为 arr 的每个（连续）子数组。
由于答案可能很大，因此 返回答案模 10^9 + 7 。
链接：https://leetcode.cn/problems/sum-of-subarray-minimums
*/
package leetcode

func sumSubarrayMins(A []int) int {
	// 单调栈+dp

	stack, dp, res, mod := []int{}, make([]int, len(A)+1), 0, int(1e9+7)
	stack = append(stack, -1)

	for i := 0; i < len(A); i++ {
		for stack[len(stack)-1] != -1 && A[i] <= A[stack[len(stack)-1]] {
			stack = stack[:len(stack)-1]
		}
		dp[i+1] = (dp[stack[len(stack)-1]+1] + (i-stack[len(stack)-1])*A[i]) % mod
		stack = append(stack, i)
		res += dp[i+1]
		res %= mod
	}
	return res
}
