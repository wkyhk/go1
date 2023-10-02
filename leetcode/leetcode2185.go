package leetcode

import "strings"

/*
2185. 统计包含给定前缀的字符串
给你一个字符串数组 words 和一个字符串 pref 。
返回 words 中以 pref 作为 前缀 的字符串的数目。
字符串 s 的 前缀 就是s 的任一前导连续字符串。
链接：https://leetcode.cn/problems/counting-words-with-a-given-prefix
*/
func prefixCount(words []string, pref string) int {
	ans := 0
	for _, v := range words {
		if strings.HasPrefix(v, pref) {
			ans++
		}
	}
	return ans
}
