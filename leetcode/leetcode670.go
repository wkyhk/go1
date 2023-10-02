package leetcode

import "strconv"

/*
670. 最大交换
给定一个非负整数，你至多可以交换一次数字中的任意两位。返回你能得到的最大值。
https://leetcode.cn/problems/maximum-swap/
*/
func maximumSwap(num int) int {
	ans := num
	s := []byte(strconv.Itoa(num))
	for i := range s {
		for j := range s[:i] {
			s[i], s[j] = s[j], s[i]
			v, _ := strconv.Atoi(string(s))
			ans = max(ans, v)
			s[i], s[j] = s[j], s[i]
		}
	}
	return ans
}

/*func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}
*/
