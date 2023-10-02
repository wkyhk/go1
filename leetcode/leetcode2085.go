/*
2085. 统计出现过一次的公共字符串
给你两个字符串数组 words1 和 words2 ，请你返回在两个字符串数组中 都恰好出现一次 的字符串的数目。
https://leetcode.cn/problems/count-common-words-with-one-occurrence/
*/
package leetcode

func countWords(words1 []string, words2 []string) int {
	m := make(map[string]int, 0)
	m1 := make(map[string]int, 0)

	ans := 0
	for _, v := range words1 {
		m[v]++
	}
	for _, v := range words2 {
		m1[v]++
	}
	for i, v := range m {
		if v == 1 && m1[i] == 1 {
			ans++
		}
	}
	return ans
}
