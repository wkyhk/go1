package leetcode

import "fmt"

/*
1750. 删除字符串两端相同字符后的最短长度
给你一个只包含字符 'a'，'b' 和 'c' 的字符串 s ，你可以执行下面这个操作（5 个步骤）任意次：
选择字符串 s 一个 非空 的前缀，这个前缀的所有字符都相同。
选择字符串 s 一个 非空 的后缀，这个后缀的所有字符都相同。
前缀和后缀在字符串中任意位置都不能有交集。
前缀和后缀包含的所有字符都要相同。
同时删除前缀和后缀。
请你返回对字符串 s 执行上面操作任意次以后（可能 0 次），能得到的 最短长度 。
链接：https://leetcode.cn/problems/minimum-length-of-string-after-deleting-similar-ends
*/
func minimumLength(s string) int {
    l,r:=0,len(s)-1
	for ;l<r;{
		if s[l]!=s[r]{
			break
		}else{
			c:=s[l]
			for ;(s[l]==c||s[r]==c)&&l<r;{
				if s[l]==c{
					l++
				}
				if s[r]==c{
					r--
				}
			}
		}
	}
	if l>r{
		 return 0
	}
	return r-l+1
}
func TestLeetCode1750(){
	fmt.Println(minimumLength("ca"))
	fmt.Println(minimumLength("cabaabac"))
	fmt.Println(minimumLength("aabccabba"))
}
