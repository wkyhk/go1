package leetcode

/*
509. 斐波那契数
斐波那契数 （通常用 F(n) 表示）形成的序列称为 斐波那契数列 。该数列由 0 和 1 开始，后面的每一项数字都是前面两项数字的和。也就是：
0 <= n <= 30
*/
func fib(n int) int {
	if n <= 1 {
		return n
	}
	f, s, t := 0, 1, 1
	for i := 1; i < n; i++ {
		f = s
		s = t
		t = f + s
	}
	return t
}
