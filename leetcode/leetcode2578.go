/*
 2578. 最小和分割

给你一个正整数 num ，请你将它分割成两个非负整数 num1 和 num2 ，满足：
num1 和 num2 直接连起来，得到 num 各数位的一个排列。
换句话说，num1 和 num2 中所有数字出现的次数之和等于 num 中所有数字出现的次数。
num1 和 num2 可以包含前导 0 。
请你返回 num1 和 num2 可以得到的和的 最小 值。
注意：
num 保证没有前导 0 。
num1 和 num2 中数位顺序可以与 num 中数位顺序不同。
链接：https://leetcode.cn/problems/split-with-minimum-sum/?envType=daily-question&envId=2023-10-09
*/
package leetcode

import "sort"

func splitNum(num int) int {
	m := make([]int, 0)
	for num > 0 {
		m = append(m, num%10)
		num /= 10
	}
	sort.Ints(m)
	num1, num2 := 0, 0
	for i := 0; i < len(m); i++ {
		if i%2 == 0 {
			num1 = num1*10 + m[i]
		} else {
			num2 = num2*10 + m[i]
		}
	}
	return num1 + num2
}
