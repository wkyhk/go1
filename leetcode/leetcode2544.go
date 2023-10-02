package leetcode

import "fmt"

/*
2544. 交替数字和
给你一个正整数 n 。n 中的每一位数字都会按下述规则分配一个符号：
最高有效位 上的数字分配到 正 号。
剩余每位上数字的符号都与其相邻数字相反。
返回所有数字及其对应符号的和。
链接：https://leetcode.cn/problems/alternating-digit-sum
*/
func alternateDigitSum(n int) int {
	ans := 0
	count := 0
	for n > 0 {
		if count%2 == 0 {
			ans = ans + (n % 10)
		} else {
			ans = ans - (n % 10)
		}
		count++
		n = n / 10
	}
	if count%2 == 0 {
		return -ans
	}
	return ans
}
func TestLeetCode2544() {
	fmt.Println(alternateDigitSum(521))
	fmt.Println(alternateDigitSum(111))
	fmt.Println(alternateDigitSum(886996))
}
