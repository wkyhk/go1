package leetcode

import "fmt"

/*
2180. 统计各位数字之和为偶数的整数个数
给你一个正整数 num ，请你统计并返回 小于或等于 num 且各位数字之和为 偶数 的正整数的数目。
正整数的 各位数字之和 是其所有位上的对应数字相加的结果。
链接：https://leetcode.cn/problems/count-integers-with-even-digit-sum
*/
func countEven(num int) int {
	y, x := num/10, num%10
	ans := y * 5
	ySum := 0
	for ; y > 0; y /= 10 {
		ySum += y % 10
	}
	if ySum%2 == 0 {
		ans += x / 2
	} else {
		ans += (x+1)/2 - 1
	}
	return ans
}

func TestLeetCode2180() {
	fmt.Println(countEven(4))
}
