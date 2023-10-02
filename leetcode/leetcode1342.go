/*
1342. 将数字变成 0 的操作次数
给你一个非负整数 num ，请你返回将它变成 0 所需要的步数。 如果当前数字是偶数，你需要把它除以 2 ；否则，减去 1 。
https://leetcode.cn/problems/number-of-steps-to-reduce-a-number-to-zero/
*/
package leetcode

func numberOfSteps(num int) int {
	ans := 0
	for num > 0 {
		if num%2 == 0 {
			num /= 2
		} else {
			num -= 1
		}
		ans++
	}
	return ans
}
