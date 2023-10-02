package leetcode

import "strings"

/*
1324. 竖直打印单词
给你一个字符串 s。请你按照单词在 s 中的出现顺序将它们全部竖直返回。
单词应该以字符串列表的形式返回，必要时用空格补位，但输出尾部的空格需要删除（不允许尾随空格）。
每个单词只能放在一列上，每一列中也只能有一个单词。
链接：https://leetcode.cn/problems/print-words-vertically
*/
func printVertically(s string) []string {
	words := strings.Split(s, " ")
	rows := 0
	for _, word := range words {
		if len(word) > rows {
			rows = len(word)
		}
	}
	cols := len(words)
	res := make([][]byte, rows)
	for i := range res {
		res[i] = make([]byte, cols)
	}
	for j := 0; j < cols; j++ {
		for i := 0; i < len(words[j]); i++ {
			res[i][j] = words[j][i]
		}
		for i := len(words[j]); i < rows; i++ {
			res[i][j] = byte(' ')
		}
	}
	result := make([]string, rows)
	for i := range res {
		j := len(res[i]) - 1
		for res[i][j] == ' ' {
			j--
		}
		result[i] = string(res[i][:j+1])
	}
	return result
}
