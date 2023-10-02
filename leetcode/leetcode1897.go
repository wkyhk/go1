package leetcode

import "fmt"

/*
1897. 重新分配字符使所有字符串都相等
给你一个字符串数组 words（下标 从 0 开始 计数）。
在一步操作中，需先选出两个 不同 下标 i 和 j，其中 words[i] 是一个非空字符串，接着将 words[i] 中的 任一 字符移动到 words[j] 中的 任一 位置上。
如果执行任意步操作可以使 words 中的每个字符串都相等，返回 true ；否则，返回 false 。
链接：https://leetcode.cn/problems/redistribute-characters-to-make-all-strings-equal
*/
func makeEqual(words []string) bool {
	m := make(map[byte]int, 0)
	for _, word := range words {
		str := []byte(word)
		for _, val := range str {
			m[val]++
		}
	}
	for _, v := range m {
		if v%len(words) != 0 {
			return false
		}
	}
	return true
}
func Example1897() {
	fmt.Println(makeEqual([]string{"abc", "aabc", "bc"}))
	fmt.Println(makeEqual([]string{"ab", "a"}))
}
