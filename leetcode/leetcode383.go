package leetcode

/*
383. 赎金信
给你两个字符串：ransomNote 和 magazine ，判断 ransomNote 能不能由 magazine 里面的字符构成。
如果可以，返回 true ；否则返回 false 。
magazine 中的每个字符只能在 ransomNote 中使用一次。
链接：https://leetcode.cn/problems/ransom-note
*/
func canConstruct(ransomNote string, magazine string) bool {
	m1 := make(map[byte]int, 0)
	m2 := make(map[byte]int, 0)
	for _, v := range []byte(ransomNote) {
		m1[v]++
	}
	for _, v := range []byte(magazine) {
		m2[v]++
	}
	for i, _ := range m1 {
		if m1[i] > m2[i] {
			return false
		}
	}
	return true
}
