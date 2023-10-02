package leetcode

import (
	"fmt"
	"strings"
)

/*
1592. 重新排列单词间的空格
给你一个字符串 text ，该字符串由若干被空格包围的单词组成。每个单词由一个或者多个小写英文字母组成，并且两个单词之间至少存在一个空格。题目测试用例保证 text 至少包含一个单词 。
请你重新排列空格，使每对相邻单词之间的空格数目都 相等 ，并尽可能 最大化 该数目。如果不能重新平均分配所有空格，请 将多余的空格放置在字符串末尾 ，这也意味着返回的字符串应当与原 text 字符串的长度相等。
返回 重新排列空格后的字符串 。
链接：https://leetcode.cn/problems/rearrange-spaces-between-words
*/
func reorderSpaces(s string) (ans string) {
	words := strings.Fields(s)
	space := strings.Count(s, " ")
	lw := len(words) - 1
	if lw == 0 {
		return words[0] + strings.Repeat(" ", space)
	}
	return strings.Join(words, strings.Repeat(" ", space/lw)) + strings.Repeat(" ", space%lw)
}

func Example1592() {
	fmt.Println(reorderSpaces("  this   is  a sentence "))
	fmt.Println(reorderSpaces(" practice   makes   perfect"))
	fmt.Println(reorderSpaces("hello   world"))
	fmt.Print(reorderSpaces("  walks  udp package   into  bar a"))
	fmt.Println(reorderSpaces("a"))
}
