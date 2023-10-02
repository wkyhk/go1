package leetcode

import "fmt"

/*
2351. 第一个出现两次的字母
给你一个由小写英文字母组成的字符串 s ，请你找出并返回第一个出现 两次 的字母。
注意：
如果 a 的 第二次 出现比 b 的 第二次 出现在字符串中的位置更靠前，则认为字母 a 在字母 b 之前出现两次。
s 包含至少一个出现两次的字母。
链接：https://leetcode.cn/problems/first-letter-to-appear-twice
*/
func repeatedCharacter(s string) byte {
	m:=make(map[byte]int,0)
	for _,v:=range []byte(s){
        if _,ok:=m[v];ok{
			return v
		}else {
			m[v]++
		}
	}
	return ' '
}
func TestLeetCode2351(){
	fmt.Println(repeatedCharacter("abccbaacz"))
	fmt.Println(repeatedCharacter("abcdd"))
}