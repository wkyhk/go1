package leetcode

/*
1876. 长度为三且各字符不同的子字符串
如果一个字符串不含有任何重复字符，我们称这个字符串为 好 字符串。
给你一个字符串 s ，请你返回 s 中长度为 3 的 好子字符串 的数量。
注意，如果相同的好子字符串出现多次，每一次都应该被记入答案之中。
子字符串 是一个字符串中连续的字符序列。
链接：https://leetcode.cn/problems/substrings-of-size-three-with-distinct-characters
*/
func CountGoodSubstrings(s string) int {
	count := 0
	for i := 0; i < len(s)-2; i++ {
		if isGoodSub([]byte(s[i : i+3])) {
			count++
		}
	}
	return count
}
func isGoodSub(s []byte) bool {
	m := make(map[byte]struct{}, 0)
	for _, v := range s {
		m[v] = struct{}{}
	}
	return len(s) == len(m)
}
