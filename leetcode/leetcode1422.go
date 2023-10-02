package leetcode

import "strings"

/*
1422.分割字符串最大得分
给你一个由若干 0 和 1 组成的字符串 s ，请你计算并返回将该字符串分割成两个 非空 子字符串（即 左 子字符串和 右 子字符串）所能获得的最大得分。
「分割字符串的得分」为 左 子字符串中 0 的数量加上 右 子字符串中 1 的数量。
链接：https://leetcode.cn/problems/maximum-score-after-splitting-a-string
*/
func MaxScore(s string) (ans int) {
	for i := 1; i < len(s); i++ {
		ans = max(ans, strings.Count(s[:i], "0")+strings.Count(s[i:], "1"))
	}
	return
}
