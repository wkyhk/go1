package leetcode

/*
1653. 使字符串平衡的最少删除次数
给你一个字符串 s ，它仅包含字符 'a' 和 'b'​​​​ 。
你可以删除 s 中任意数目的字符，使得 s 平衡 。我们称 s 平衡的 当不存在下标对 (i,j) 满足 i < j 且 s[i] = 'b' 同时 s[j]= 'a' 。
请你返回使 s 平衡 的 最少 删除次数。
链接：https://leetcode.cn/problems/minimum-deletions-to-make-string-balanced
*/
func MinimumDeletions(s string) int {
	// dp[i]为使字符串s[0:i]平衡的最少删除次数
	dp := make([]int, len(s))
	numb := 0
	if s[0] == 'b' {
		numb++
	}
	for i := 1; i < len(s); i++ {
		if s[i] == 'b' {
			dp[i] = dp[i-1] // 当出现b说明已经平衡
			numb++
		} else {
			//如果 为a 则需要判断是删除前面所有的b还是删除当前a
			dp[i] = min(dp[i-1]+1, numb)
		}
	}
	return dp[len(s)-1]
}
