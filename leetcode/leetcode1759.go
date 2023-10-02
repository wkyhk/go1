package leetcode

/*
1759. 统计同构子字符串的数目
给你一个字符串 s ，返回 s 中 同构子字符串 的数目。由于答案可能很大，只需返回对 109 + 7 取余 后的结果。
同构字符串 的定义为：如果一个字符串中的所有字符都相同，那么该字符串就是同构字符串。
子字符串 是字符串中的一个连续字符序列。
链接：https://leetcode.cn/problems/count-number-of-homogenous-substrings
*/
func CountHomogenous(s string) int {
	ans := 0
	const mod int = 1e9 + 7
	n := 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			n++
		} else {
			ans = (ans + (n+1)*n/2) % mod
			n = 1
		}
	}
	ans = (ans + (n+1)*n/2) % mod
	return ans

}
