package leetcode

/*
991. 坏了的计算器
在显示着数字 startValue 的坏计算器上，我们可以执行以下两种操作：
双倍（Double）：将显示屏上的数字乘 2；
递减（Decrement）：将显示屏上的数字减 1 。
给定两个整数 startValue 和 target 。返回显示数字 target 所需的最小操作数。
链接：https://leetcode.cn/problems/broken-calculator
*/

func BrokenCalc(startValue int, target int) int {
	ans := 0
	for startValue < target {
		if target%2 != 0 {
			target = target + 1
		} else {
			target = target / 2
		}
		ans++
	}
	ans = ans + startValue - target
	return ans
}
