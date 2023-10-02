package leetcode

import "fmt"

/*
1663. 具有给定数值的最小字符串
小写字符 的 数值 是它在字母表中的位置（从 1 开始），因此 a 的数值为 1 ，b 的数值为 2 ，c 的数值为 3 ，以此类推。
字符串由若干小写字符组成，字符串的数值 为各字符的数值之和。例如，字符串 "abe" 的数值等于 1 + 2 + 5 = 8 。
给你两个整数 n 和 k 。返回 长度 等于 n 且 数值 等于 k 的 字典序最小 的字符串。
注意，如果字符串 x 在字典排序中位于 y 之前，就认为 x 字典序比 y 小，有以下两种情况：
x 是 y 的一个前缀；
如果 i 是x[i] != y[i] 的第一个位置，且 x[i]在字母表中的位置比y[i]靠前。
链接：https://leetcode.cn/problems/smallest-string-with-a-given-numeric-value
*/
func getSmallestString(n int, k int) string {
     b:=make([]byte,n)
	 for i:=0;i<n;i++{
		 b[i]='a'
	 }
	 k=k-n
	 for i:=n-1;i>=0;i--{
		 if k>25{
			 b[i]+=byte(25)
			 k=k-25
		 }else{
			 b[i]+=byte(k)
			 break
		 }
	 }
	 return string(b)
}
func TestLeetCode1663(){
	fmt.Println(getSmallestString(3,27))
	fmt.Println(getSmallestString(5,73))
}