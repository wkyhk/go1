package leetcode

/*
2586. 统计范围内的元音字符串数
给你一个下标从 0 开始的字符串数组 words 和两个整数：left 和 right 。

如果字符串以元音字母开头并以元音字母结尾，那么该字符串就是一个 元音字符串 ，其中元音字母是 'a'、'e'、'i'、'o'、'u' 。

返回 words[i] 是元音字符串的数目，其中 i 在闭区间 [left, right] 内。
链接：https://leetcode.cn/problems/count-the-number-of-vowel-strings-in-range/
*/
func vowelStrings(words []string, left int, right int) int {
	m := map[byte]int{'a': 1, 'e': 1, 'i': 1, 'o': 1, 'u': 1}
	ans := 0
	for i := left; i <= right; i++ {
		if m[words[i][0]] == 1 && m[words[i][len(words[i])-1]] == 1 {
			ans++
		}
	}
	return ans
}
