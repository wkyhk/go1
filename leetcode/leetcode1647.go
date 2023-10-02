package leetcode

import (
	"fmt"
)

/*
1647. 字符频次唯一的最小删除次数
如果字符串 s 中 不存在 两个不同字符 频次 相同的情况，就称 s 是 优质字符串 。
给你一个字符串 s，返回使 s 成为 优质字符串 需要删除的 最小 字符数。
字符串中字符的 频次 是该字符在字符串中的出现次数。例如，在字符串 "aab" 中，'a' 的频次是 2，而 'b' 的频次是 1 。
链接：https://leetcode.cn/problems/minimum-deletions-to-make-character-frequencies-unique
*/
func minDeletions(s string) int {
	cntMap := make(map[byte]int)
	for i, _ := range s {
		cntMap[s[i]]++
	}
	byteMap := make(map[int]bool)
	res := 0
	for _, v := range cntMap {
		if !byteMap[v] {
			byteMap[v] = true
		} else {
			for v > 0 {
				v--
				res++
				if byteMap[v] == false {
					byteMap[v] = true
					break
				}
			}
		}
	}
	return res
}
func TestLeetCode1647() {
	fmt.Println(minDeletions("aab"))
	fmt.Println(minDeletions("aaabbbcc"))
	fmt.Println(minDeletions("ceabaacb"))
}
