package leetcode

/*
1624. 两个相同字符之间的最长子字符串
给你一个字符串 s，请你返回 两个相同字符之间的最长子字符串的长度 ，计算长度时不含这两个字符。如果不存在这样的子字符串，返回 -1 。
子字符串 是字符串中的一个连续字符序列。
链接：https://leetcode.cn/problems/largest-substring-between-two-equal-characters
*/
func maxLengthBetweenEqualCharacters(s string) int {
	m := make(map[byte]int, 0)
	ans := -1
	b := []byte(s)
	for i := 0; i < len(b); i++ {
		m[b[i]] = i
	}
	for i := 0; i < len(b); i++ {
		ans = max(ans, m[b[i]]-i-1)
	}
	return ans
}

/* func max(a,b int)int{
    if a>b{
        return a
    }
    return b
} */
