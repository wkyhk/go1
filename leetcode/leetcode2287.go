package leetcode

import "math"

/*
2287. 重排字符形成目标字符串
给你两个下标从 0 开始的字符串 s 和 target 。你可以从 s 取出一些字符并将其重排，得到若干新的字符串。
从 s 中取出字符并重新排列，返回可以形成 target 的 最大 副本数。
链接：https://leetcode.cn/problems/rearrange-characters-to-make-target-string
*/
func rearrangeCharacters(s string, target string) int {
	m := make(map[byte]int, 0)
	m1 := make(map[byte]int, 0)
	ans := math.MaxInt32
	for _, v := range []byte(s) {
		m[v]++
	}
	for _, v := range []byte(target) {
		m1[v]++
	}
	for i, _ := range m1 {
		ans = min(ans, m[i]/m1[i])
	}
	return ans
}
