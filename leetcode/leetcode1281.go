package leetcode

/*
1281. 整数的各位积和之差
给你一个整数 n，请你帮忙计算并返回该整数「各位数字之积」与「各位数字之和」的差。
1-10^5
*/
func SubtractProductAndSum(n int) int {
	sum := 0
	mul := 1
	for ; n > 0; n = n / 10 {
		temp := n % 10
		sum += temp
		mul *= temp
	}
	return mul - sum
}
