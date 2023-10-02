package leetcode

import "fmt"

/*
392. 判断子序列
给定字符串 s 和 t ，判断 s 是否为 t 的子序列。
字符串的一个子序列是原始字符串删除一些（也可以不删除）字符而不改变剩余字符相对位置形成的新字符串。（例如，"ace"是"abcde"的一个子序列，而"aec"不是）。
链接：https://leetcode.cn/problems/is-subsequence
*/
func isSubsequence(s string, t string) bool {
	if len(s) == 0 {
		return true
	}
	str := []byte(s)
	index := 0
	for _, v := range []byte(t) {
		if v == str[index] {
			index++
			if index == len(str) {
				return true
			}
		}
	}
	return false
}
func Example392() {
	fmt.Println(isSubsequence("abc", "ahbgdc"))
}
