package leetcode

import "fmt"

/*
2309. 兼具大小写的最好英文字母
给你一个由英文字母组成的字符串 s ，请你找出并返回 s 中的 最好 英文字母。返回的字母必须为大写形式。
如果不存在满足条件的字母，则返回一个空字符串。
最好 英文字母的大写和小写形式必须 都 在 s 中出现。
英文字母 b 比另一个英文字母a更好 的前提是：英文字母表中，b 在 a 之 后 出现。
链接：https://leetcode.cn/problems/greatest-english-letter-in-upper-and-lower-case
*/
func greatestLetter(s string) string {
	cntA:=[26]int{}
	cnta:=[26]int{}
	for _,v:=range []byte(s){
		if v-'a'<26{
			cnta[v-'a']=1
		}
		if v-'A'<26{
			cntA[v-'A']=1
		}
	}
	for i:=25;i>=0;i--{
		if cntA[i]==1&&cnta[i]==1{
			return string('A'+i)
		}
	}
	return ""
}
func TestLeetCode2309(){
	fmt.Println(greatestLetter("lEeTcOdE"))
	fmt.Println(greatestLetter("arRAzFif"))
	fmt.Println(greatestLetter("AbCdEfGhIjK"))
}