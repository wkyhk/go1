package leetcode

import "math"

/* 7. 整数反转
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
假设环境不允许存储 64 位整数（有符号或无符号）。
链接：https://leetcode.cn/problems/reverse-integer
*/

func reverse(x int) int {
	a := 1
	if x < 0 {
		a = -1
		x = -x
	}
	res := 0
	for x != 0 {
		res = res*10 + x%10
		x = x / 10
		if res >= math.MaxInt32 {
			return 0
		}
	}
	return a * res
}
