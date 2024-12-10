package leetcode

/*
935. 骑士拨号器
象棋骑士有一个独特的移动方式，它可以垂直移动两个方格，水平移动一个方格，或者水平移动两个方格，垂直移动一个方格(两者都形成一个 L 的形状)。
象棋骑士可能的移动方式如下图所示:
*/
func knightDialer(n int) int {
	const mod = 1_000_000_007
	moves := [][]int{
		{4, 6},
		{6, 8},
		{7, 9},
		{4, 8},
		{3, 9, 0},
		{},
		{1, 7, 0},
		{2, 6},
		{1, 3},
		{2, 4},
	}
	d := [2][10]int{}
	for i := 0; i < 10; i++ {
		d[1][i] = 1
	}
	for i := 2; i <= n; i++ {
		x := i & 1
		for j := 0; j < 10; j++ {
			d[x][j] = 0
			for _, k := range moves[j] {
				d[x][j] = (d[x][j] + d[x^1][k]) % mod
			}
		}
	}
	res := 0
	for _, x := range d[n%2] {
		res = (res + x) % mod
	}
	return res
}
