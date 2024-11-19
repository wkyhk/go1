package leetcode
/*2207. 字符串中最多数目的子序列
给你一个下标从 0 开始的字符串 text 和另一个下标从 0 开始且长度为 2 的字符串 pattern ，两者都只包含小写英文字母。

你可以在 text 中任意位置插入 一个 字符，这个插入的字符必须是 pattern[0] 或者 pattern[1] 。注意，这个字符可以插入在 text 开头或者结尾的位置。

请你返回插入一个字符后，text 中最多包含多少个等于 pattern 的 子序列 。

子序列 指的是将一个字符串删除若干个字符后（也可以不删除），剩余字符保持原本顺序得到的字符串。
链接：https://leetcode.cn/problems/maximize-number-of-subsequences-in-a-string/description/
*/
func maximumSubsequenceCount(text string, pattern string) int64 {
    ans:=0
	cnt1, cnt2 := 0, 0
	for i := 0; i < len(text); i++ {
		if text[i] == pattern[1] {
			ans += cnt1
			cnt2++
		}
		if text[i] == pattern[0] {
			cnt1++
		}
	}
	ans=ans+max(cnt1, cnt2)
	return int64(ans)
}