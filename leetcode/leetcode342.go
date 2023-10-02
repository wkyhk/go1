package leetcode

/*
  342. 4的幂
给定一个整数，写一个函数来判断它是否是 4 的幂次方。如果是，返回 true ；否则，返回 false 。
整数 n 是 4 的幂次方需满足：存在整数 x 使得 n == 4^x
-2^31 <= n <= 2^31 - 1
*/
func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}
	for n%4 == 0 {
		n = n / 4
	}
	return n == 1
}
