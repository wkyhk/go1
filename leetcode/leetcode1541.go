package leetcode

import "fmt"

/*
1541. 平衡括号字符串的最少插入次数
给你一个括号字符串 s ，它只包含字符 '(' 和 ')' 。一个括号字符串被称为平衡的当它满足：
任何左括号 '(' 必须对应两个连续的右括号 '))' 。
左括号 '(' 必须在对应的连续两个右括号 '))' 之前。
比方说 "())"， "())(())))" 和 "(())())))" 都是平衡的， ")()"， "()))" 和 "(()))" 都是不平衡的。
你可以在任意位置插入字符 '(' 和 ')' 使字符串平衡。
请你返回让 s 平衡的最少插入次数。
链接：https://leetcode.cn/problems/minimum-insertions-to-balance-a-parentheses-string
*/
func minInsertions(s string) int {
	ans := 0  // 记录插入次数
	need := 0 //  记录右括号的需求量

	for _, c := range s {
		if string(c) == "(" {
			need += 2 // 对右括号的需求+2

			if need%2 == 1 { //左括号对右括号的需求必须为偶数
				ans++
				need--
			}
		}
		if string(c) == ")" {
			need-- // 对右括号的需求-1

			if need == -1 {
				ans++    // 需插入一个左括号
				need = 1 // 同时对右括号的需求变为1
			}
		}
	}

	return ans + need
}
func Example1541() {
	/* s := "(()))"
	s1 := "())"
	s2 := "))())("
	fmt.Println(minInsertions(s))
	fmt.Println(minInsertions(s1))
	fmt.Println(minInsertions(s2)) */
	s := "(()))(()))()())))"
	fmt.Println(minInsertions(s))
}
