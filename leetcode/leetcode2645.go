package leetcode

/*
2645. 构造有效字符串的最少插入数给你一个字符串 word ，你可以向其中任何位置插入 "a"、"b" 或 "c" 任意次，返回使 word 有效 需要插入的最少字母数。

如果字符串可以由 "abc" 串联多次得到，则认为该字符串 有效 。
链接：https://leetcode.cn/problems/minimum-additions-to-make-valid-string/description/
*/
func addMinimum(word string) int {
	n := len(word)
	res := int(word[0] - word[n-1] + 2)
	for i := 1; i < n; i++ {
		res += int(word[i]-word[i-1]+2) % 3
	}
	return res
}
