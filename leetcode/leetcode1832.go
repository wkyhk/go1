package leetcode

/*
1832. 判断句子是否为全字母句
全字母句 指包含英语字母表中每个字母至少一次的句子。
给你一个仅由小写英文字母组成的字符串 sentence ，请你判断sentence 是否为 全字母句 。
如果是，返回 true ；否则，返回 false 。
链接：https://leetcode.cn/problems/check-if-the-sentence-is-pangram
*/
func checkIfPangram(sentence string) bool {
	if len(sentence) < 26 {
		return false
	}
	m := make(map[byte]struct{})
	for _, v := range []byte(sentence) {
		m[v] = struct{}{}
		if len(m) == 26 {
			return true
		}
	}
	return false
}
