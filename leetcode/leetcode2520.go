/*
2520. 统计能整除数字的位数
给你一个整数 num ，返回 num 中能整除 num 的数位的数目。
如果满足 nums % val == 0 ，则认为整数 val 可以整除 nums 。
链接：https://leetcode.cn/problems/count-the-digits-that-divide-a-number/?envType=daily-question&envId=2023-10-26
*/
package leetcode

func countDigits(num int) int {
	ans := 0
	n := num
	for num > 0 {
		if n%(num%10) == 0 {
			ans++
		}
		num /= 10
	}
	return ans
}
