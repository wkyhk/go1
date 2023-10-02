package leetcode

/*
1768. 交替合并字符串
给你两个字符串 word1 和 word2 。请你从 word1 开始，通过交替添加字母来合并字符串。如果一个字符串比另一个字符串长，就将多出来的字母追加到合并后字符串的末尾。
返回 合并后的字符串 。
链接：https://leetcode.cn/problems/merge-strings-alternately
*/
func MergeAlternately(word1 string, word2 string) string {
	n := min(len(word1), len(word2))
	ans := make([]byte, 2*n, len(word1)+len(word2))
	for i := 0; i < n; i++ {
		ans[2*i] = word1[i]
		ans[2*i+1] = word2[i]
	}
	if len(word1) < len(word2) {
		ans = append(ans, []byte(word2[len(word1):])...)
	} else {
		ans = append(ans, []byte(word1[len(word2):])...)
	}
	return string(ans)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
