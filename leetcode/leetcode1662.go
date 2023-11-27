package leetcode

import "strings"

/*
1662. 检查两个字符串数组是否相等
给你两个字符串数组 word1 和 word2 。如果两个数组表示的字符串相同，返回 true ；否则，返回 false 。

数组表示的字符串 是由数组中的所有元素 按顺序 连接形成的字符串。
链接：https://leetcode.cn/problems/check-if-two-string-arrays-are-equivalent/description/
*/
func arrayStringsAreEqual(word1 []string, word2 []string) bool {
	return strings.Join(word1, "") == strings.Join(word2, "")
}
