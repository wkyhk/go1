package leetcode

/*
3146. 两个字符串的排列差
给你两个字符串 s 和 t，每个字符串中的字符都不重复，且 t 是 s 的一个排列。

排列差 定义为 s 和 t 中每个字符在两个字符串中位置的绝对差值之和。

返回 s 和 t 之间的 排列差 。
链接：https://leetcode.cn/problems/permutation-difference-between-two-strings/description/
*/
func findPermutationDifference(s string, t string) int {
	sum := 0
	m := make(map[byte]int)
	t1 := []byte(t)
	for i, v := range t1 {
		m[v] = i
	}
	for i := 0; i < len(s); i++ {
		sum = abs(m[s[i]]-i) + sum
	}
	return sum
}
